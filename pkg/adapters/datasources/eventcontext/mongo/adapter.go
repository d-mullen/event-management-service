package mongo

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/zenoss/event-management-service/internal/batchops"
	"github.com/zenoss/event-management-service/internal/instrumentation"
	"github.com/zenoss/event-management-service/pkg/adapters/framework/mongodb"
	"github.com/zenoss/event-management-service/pkg/models/event"
	"github.com/zenoss/zenkit/v5"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/exp/slices"
	"sync"
	"time"
)

const (
	CollEvents      = "event"
	CollOccurrences = "occurrence"
	CollNotes       = "note"
)

type Adapter struct {
	db          mongodb.Database
	collections map[string]mongodb.Collection
	ttlMap      map[string]time.Duration
}

var _ event.Repository = &Adapter{}

func NewAdapter(_ context.Context, cfg mongodb.Config, database mongodb.Database) (*Adapter, error) {
	if database == nil {
		return nil, errors.New("invalid argument: nil database")
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
		hasNext bool
		limit   int
		log     = zenkit.ContextLogger(ctx)
	)
	ctx, span := instrumentation.StartSpan(ctx, "mongodb.Adapter.Find")
	defer span.End()
	err := query.Validate()
	if err != nil {
		return nil, errors.Wrap(err, "failed to find events")
	}
	if pi := query.PageInput; pi != nil {
		limit = int(pi.Limit)
	}
	filters, findOpts, err := QueryToFindArguments(query)
	if err != nil {
		return nil, err
	}
	instrumentation.AnnotateSpan("occurrence.Find",
		fmt.Sprintf("executing %s.Find", CollOccurrences),
		span,
		map[string]any{
			"filter":   filters,
			"findOpts": findOpts,
		}, nil)
	cursor, err := db.collections[CollOccurrences].Find(ctx, filters, findOpts...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find events with aggregation")
	}
	defer func() {
		err := cursor.Close(ctx)
		if err != nil {
			log.Warnf("failure closing cursor: %q", err)
		}
	}()
	docs := make([]*bson.M, 0)
	for cursor.Next(ctx) {
		if err := cursor.Err(); err != nil {
			return nil, errors.Wrap(err, "failed to find events with cursor")
		}
		result := &bson.M{}
		if err := cursor.Decode(result); err != nil {
			log.Errorf("failed to decode result: %q", err)
			continue
		}
		docs = append(docs, result)
		if limit > 0 && len(docs) == limit {
			break
		}
	}
	hasNext = !(cursor.ID() == 0)
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

	if slices.Contains(query.Fields, "notes") {
		// TODO: only get notes if requested
		notesInFilter := bson.D{{Key: "occid", Value: bson.D{{Key: OpIn, Value: occurrenceIDs}}}}
		log.WithFields(logrus.Fields{
			"filters": notesInFilter,
		}).Debugf("executing %s.Find", "note")
		notesCursor, err := db.collections[CollNotes].Find(ctx, notesInFilter)
		if err != nil {
			return nil, errors.Wrap(err, "failed to find notes")
		}
		defer func() {
			err := notesCursor.Close(ctx)
			if err != nil {
				log.Warnf("failed to close notes cursor: %q", err)
			}
		}()
		notesDocs := make([]bson.M, 0)
		allNotes := make([]*Note, 0)
		err = notesCursor.All(ctx, &notesDocs)
		if err != nil {
			return nil, errors.Wrap(err, "failed to retrieve note results")
		}
		for _, doc := range notesDocs {
			newNote := &Note{}
			err = defaultDecodeFunc(&doc, newNote)
			if err != nil {
				return nil, errors.Wrap(err, "failed to umarshal notes bytes")
			}
			allNotes = append(allNotes, newNote)
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

	if slices.Contains(query.Fields, "dimensions") || slices.Contains(query.Fields, "entity") {
		// TODO: only get event docs if requested; otherwise we can construct partial
		// event docs from the occurrence docs
		eventMapMut := sync.Mutex{}
		eventMap := make(map[string]*EventDimensions)
		err := batchops.DoConcurrently(ctx, 500, 10, eventIDs,
			func(batch []string) (mongodb.Cursor, error) {
				eventInFilter := bson.D{{Key: OpIn, Value: batch}}
				log.WithFields(logrus.Fields{
					"filters": eventInFilter,
				}).Debugf("executing %s.Find", "event")
				eventCursor, err := db.collections[CollEvents].Find(ctx, bson.D{{Key: "_id", Value: eventInFilter}})
				if err != nil {
					return nil, errors.Wrap(err, "failed to find events")
				}
				return eventCursor, nil
			},
			func(eventCursor mongodb.Cursor) (bool, error) {
				defer func() {
					err := eventCursor.Close(ctx)
					if err != nil {
						log.Warnf("failed to close events cursor: %q", err)
					}
				}()
				for eventCursor.Next(ctx) {
					if err := eventCursor.Err(); err != nil {
						return false, errors.Wrap(err, "failed to find events with cursor")
					}
					result := &EventDimensions{}
					if err := eventCursor.Decode(result); err != nil {
						log.Errorf("failed to decode result: %q", err)
						continue
					}
					eventMapMut.Lock()
					eventMap[result.ID] = result
					eventMapMut.Unlock()
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
	}, nil
}

func (db *Adapter) Update(_ context.Context, _ *event.Event) (*event.Event, error) {
	panic("not implemented") // TODO: Implement
}
