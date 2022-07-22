package mongo

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/zenoss/event-management-service/internal/batchops"
	"github.com/zenoss/event-management-service/internal/instrumentation"
	"github.com/zenoss/event-management-service/pkg/adapters/framework/mongodb"
	"github.com/zenoss/event-management-service/pkg/models/event"
	"github.com/zenoss/zenkit/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/exp/slices"
)

const (
	CollEvents      = "event"
	CollOccurrences = "occurrence"
	CollNotes       = "note"
)

const (
	CursorConfigKeyNextKey = "nextKey"
)

type Adapter struct {
	db           mongodb.Database
	collections  map[string]mongodb.Collection
	ttlMap       map[string]time.Duration
	queryCursors event.CursorRepository
}

var _ event.Repository = &Adapter{}

func NewAdapter(_ context.Context,
	cfg mongodb.Config,
	database mongodb.Database,
	queryCursors event.CursorRepository) (*Adapter, error) {
	if database == nil {
		return nil, errors.New("invalid argument: nil database")
	}
	if queryCursors == nil {
		return nil, errors.New("invalid argument: nil query cursor adapter")
	}
	// since collection names won't change while service is up and running
	// it's reasonable to store the handles
	// and avoid annoying call eventStore.Db.Collection(cname) to get it
	collections := map[string]mongodb.Collection{
		CollEvents:      database.Collection(CollEvents),
		CollOccurrences: database.Collection(CollOccurrences),
		CollNotes:       database.Collection(CollNotes),
	}

	return &Adapter{
		db:          database,
		collections: collections,
		ttlMap: map[string]time.Duration{
			"zing_ttl_default": cfg.DefaultTTL,
		},
		queryCursors: queryCursors,
	}, nil
}

func (db *Adapter) Create(_ context.Context, _ *event.Event) (*event.Event, error) {
	panic("not implemented") // TODO: Implement
}

func (db *Adapter) Get(_ context.Context, _ *event.GetRequest) ([]*event.Event, error) {
	panic("not implemented") // TODO: Implement
}

func (db *Adapter) Find(ctx context.Context, query *event.Query) (*event.Page, error) {
	var (
		hasNext     bool
		limit       int
		log         = zenkit.ContextLogger(ctx)
		queryCursor *event.Cursor
		sortField   string
		filters     primitive.D
		prevNextKey *primitive.D
		sortOpt     *primitive.D
		findOpt     *options.FindOptions
	)
	ctx, span := instrumentation.StartSpan(ctx, "mongodb.Adapter.Find")
	defer span.End()
	err := query.Validate()
	if err != nil {
		return nil, errors.Wrap(err, "failed to find events")
	}
	if pi := query.PageInput; pi != nil && pi.Limit > 0 {
		limit = int(pi.Limit) + 1
	}

	if pi := query.PageInput; pi != nil && len(pi.Cursor) > 0 {
		queryCursor, err = db.queryCursors.Get(ctx, pi.Cursor)
		if err != nil {
			return nil, errors.Wrap(err, "failed to process cursor")
		}
		log.WithField("queryCursor", queryCursor).Debug("got query cursor from store")
		filters, findOpt, err = QueryToFindArguments(&queryCursor.Query)
		if err != nil {
			return nil, err
		}
		if nextKeyAny, ok := queryCursor.Config[CursorConfigKeyNextKey]; ok {
			if pNK, ok2 := nextKeyAny.(*primitive.D); ok2 {
				prevNextKey = pNK
			}
		}
	} else {
		filters, findOpt, err = QueryToFindArguments(query)
		if err != nil {
			return nil, err
		}
	}
	// generate the key for the next page
	if sortAny := findOpt.Sort; sortAny != nil {
		if sortD, ok := sortAny.(*bson.D); ok {
			sortOpt = sortD
			data, _ := bson.Marshal(sortD)
			sortRaw := bson.Raw(data)
			elms, err := sortRaw.Elements()
			if err == nil && len(elms) > 0 {
				sortField = elms[0].Key()
			}
		}
	}
	// generate pagination query
	paginatedFilter := GeneratePaginationQuery(&filters, sortOpt, prevNextKey)
	log.WithField("paginationFilter", paginatedFilter).Debug("got pagination filter")
	instrumentation.AnnotateSpan("occurrence.Find",
		fmt.Sprintf("executing %s.Find", CollOccurrences),
		span,
		map[string]any{
			"filter":  paginatedFilter,
			"findOpt": findOpt,
			"nextKey": prevNextKey,
		}, nil)
	docs := make([]*bson.M, 0)
	err = mongodb.FindWithRetry(
		ctx,
		*paginatedFilter,
		[]*options.FindOptions{findOpt},
		db.collections[CollOccurrences].Find,
		func(result *bson.M) (bool, string, error) {
			if limit > 0 && len(docs) == limit {
				// If limit has been set; then it will equal requestedLimit+1
				// So if we have scanned the the extra result, there is a next result
				hasNext = true
				return false, "", nil
			}
			docs = append(docs, result)
			id, ok := (*result)["_id"].(string)
			return ok, id, nil
		},
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find occurrences with retry")
	}
	instrumentation.AnnotateSpan("occurrenceResults",
		"got occurrence query results",
		span,
		map[string]any{
			"count": len(docs),
		}, nil)

	results := make([]*Event, 0)
	occMap := make(map[string]*Occurrence)
	eventIDs := make([]string, 0)
	occurrenceIDs := make([]string, 0, len(docs))
	for _, doc := range docs {
		occurrence := &Occurrence{}
		err := defaultDecodeFunc(doc, occurrence)
		if err != nil {
			log.Errorf("failed to unmarshal document: %q", err)
		} else {
			occurrenceIDs = append(occurrenceIDs, occurrence.ID)
			occMap[occurrence.ID] = occurrence
			newResult := &Event{
				ID:          occurrence.EventID,
				Tenant:      occurrence.Tenant,
				Entity:      occurrence.Entity,
				Occurrences: []*Occurrence{occurrence},
			}
			if occurrence.Status != event.StatusClosed {
				newResult.LastActiveOccurrence = occurrence
			}
			results = append(results, newResult)
			eventIDs = append(eventIDs, newResult.ID)
		}
	}
	nextKey := NextKey(sortField, docs)
	cursorInput := &event.Cursor{
		Query: *query,
		ID:    uuid.New().String(),
		Config: map[string]any{
			CursorConfigKeyNextKey: nextKey,
		},
	}
	resultCursorString, err := db.queryCursors.New(ctx, cursorInput)
	if err != nil {
		log.WithField(logrus.ErrorKey, err).Error("failed to create query cursor")
	}
	instrumentation.AnnotateSpan("newCursor",
		"created new cursor",
		span,
		map[string]any{
			"cursorInput": cursorInput,
		}, nil)
	if slices.Contains(query.Fields, "notes") {
		noteMut := sync.Mutex{}
		allNotes := make([]*Note, 0)
		err := batchops.DoConcurrently(ctx, 500, 10, occurrenceIDs,
			func(batch []string) ([]*Note, error) {
				notesInFilter := bson.D{{Key: "occid", Value: bson.D{{Key: OpIn, Value: batch}}}}
				log.WithFields(logrus.Fields{
					"filters": notesInFilter,
				}).Debugf("executing %s.Find", "note")
				notes := make([]*Note, 0)
				err := mongodb.FindWithRetry[*Note](
					ctx,
					notesInFilter,
					[]*options.FindOptions{},
					db.collections[CollNotes].Find,
					func(note *Note) (bool, string, error) {
						notes = append(notes, note)
						return true, note.ID, nil
					},
				)
				if err != nil {
					return nil, err
				}
				return notes, nil
			},
			func(notes []*Note) (bool, error) {
				noteMut.Lock()
				defer noteMut.Unlock()
				allNotes = append(allNotes, notes...)
				return true, nil
			},
		)
		if err != nil {
			log.WithField(logrus.ErrorKey, err).Error("failed to find notes")
			return nil, err
		}
		for _, note := range allNotes {
			occ, ok := occMap[note.OccID]
			if !ok {
				continue
			}
			if len(occ.Notes) == 0 {
				occ.Notes = make([]*Note, 0)
			}
			occ.Notes = append(occ.Notes, note)
		}
	}

	if slices.Contains(query.Fields, "dimensions") {
		eventMapMut := sync.Mutex{}
		eventMap := make(map[string]*EventDimensions)
		err := batchops.DoConcurrently(ctx, 500, 10, eventIDs,
			func(batch []string) ([]*EventDimensions, error) {
				eventInFilter := bson.D{{Key: OpIn, Value: batch}}
				log.WithFields(logrus.Fields{
					"filters": eventInFilter,
				}).Debugf("executing %s.Find", "event")
				eventDocs := make([]*EventDimensions, 0)
				err := mongodb.FindWithRetry[*EventDimensions](
					ctx,
					bson.D{{Key: "_id", Value: eventInFilter}},
					[]*options.FindOptions{},
					db.collections[CollEvents].Find,
					func(r *EventDimensions) (bool, string, error) {
						eventDocs = append(eventDocs, r)
						return true, r.ID, nil
					},
				)
				if err != nil {
					return nil, errors.Wrap(err, "failed to find events")
				}
				return eventDocs, nil
			},
			func(eventDocs []*EventDimensions) (bool, error) {
				eventMapMut.Lock()
				defer eventMapMut.Unlock()
				for _, doc := range eventDocs {
					eventMap[doc.ID] = doc
				}
				return true, nil
			})
		if err != nil {
			return nil, errors.Wrap(err, "failed to find events during batch operation")
		}
		for _, result := range results {
			if other, ok := eventMap[result.ID]; ok {
				result.Dimensions = other.Dimensions
				result.Entity = other.Entity
			} else {
				log.
					WithField("currResult", result).
					Warn("got event-occurrence document without a corresponding event document")
			}
		}
	}

	modelResults := make([]*event.Event, 0)
	bytes, err := json.Marshal(results)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &modelResults)
	if err != nil {
		return nil, err
	}

	return &event.Page{
		Results: modelResults,
		HasNext: hasNext,
		Cursor:  resultCursorString,
	}, nil
}

func (db *Adapter) Update(_ context.Context, _ *event.Event) (*event.Event, error) {
	panic("not implemented") // TODO: Implement
}
