package mongo

import (
	"context"
	"math"
	"sync"
	"time"

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
	"golang.org/x/exp/constraints"
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
	db          mongodb.Database
	collections map[string]mongodb.Collection
	ttlMap      map[string]time.Duration
	cursorRepo  event.CursorRepository
	pager       Pager
}

var _ event.Repository = &Adapter{}

func NewAdapter(_ context.Context,
	cfg mongodb.Config,
	database mongodb.Database,
	queryCursors event.CursorRepository,
) (*Adapter, error) {
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
		cursorRepo: queryCursors,
		pager:      NewSkipLimitPager(),
	}, nil
}

func (db *Adapter) Create(_ context.Context, _ *event.Event) (*event.Event, error) {
	panic("not implemented") // TODO: Implement
}

func (db *Adapter) Get(_ context.Context, _ *event.GetRequest) ([]*event.Event, error) {
	panic("not implemented") // TODO: Implement
}

func min[N constraints.Ordered](a N, rest ...N) N {
	result := a
	for _, b := range rest {
		if b < result {
			result = b
		}
	}
	return result
}

func max[N constraints.Ordered](a N, rest ...N) N {
	result := a
	for _, b := range rest {
		if b > result {
			result = b
		}
	}
	return result
}

func defaultFindOpts(opts ...*options.FindOptions) *options.FindOptions {
	opt := options.MergeFindOptions(opts...)
	sortDoc := bson.D{}
	if sortAny := opt.Sort; sortAny != nil {
		if doc, ok := sortAny.(bson.D); ok {
			sortDoc = append(sortDoc, doc...)
		}
	}
	keyFound := false
	for _, e := range sortDoc {
		if e.Key == "_id" {
			keyFound = true
		}
	}
	if !keyFound {
		sortDoc = append(sortDoc, bson.E{"_id", event.SortOrderAscending})
	}
	if len(sortDoc) > 0 {
		opt.SetSort(sortDoc)
	}
	return opt
}

func (db *Adapter) Find(ctx context.Context, query *event.Query, opts ...*event.FindOption) (*event.Page, error) {
	var (
		hasNext bool
		limit   int
		log     = zenkit.ContextLogger(ctx)
		filters primitive.D
		findOpt *options.FindOptions
	)
	ctx, span := instrumentation.StartSpan(ctx, "mongodb.Adapter.Find")
	defer span.End()
	err := query.Validate()
	if err != nil {
		return nil, errors.Wrap(err, "failed to find events: invalid query")
	}

	filters, findOpt, err = db.pager.GetPaginationQuery(ctx, query, db.cursorRepo)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get pagination parameters")
	}

	if pi := query.PageInput; pi != nil && pi.Limit > 0 {
		limit = int(pi.Limit) + 1
	}
	if findOpt.Limit != nil && *findOpt.Limit > 0 {
		limit = int(*findOpt.Limit)
	}

	defaultOpts := defaultFindOpts(findOpt)
	docs := make([]*bson.M, 0)
	err = mongodb.FindWithRetry(
		ctx,
		filters,
		[]*options.FindOptions{defaultOpts},
		db.collections[CollOccurrences].Find,
		func(result *bson.M) (bool, string, error) {
			if limit > 0 && len(docs)+1 == limit {
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

	opt := event.MergeFindOptions(opts...)
	filteredOccurrences := make([]*event.Occurrence, 0)
	filteredOccurrences, err = convertBulk[*bson.M, *event.Occurrence](docs)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert documents")
	}
	filteredOccurrences, err = opt.ApplyOccurrenceProcessors(ctx, filteredOccurrences)
	if err != nil {
		return nil, errors.Wrap(err, "failed to apply occurrence post-processors")
	}

	if limit > 0 && hasNext {
		numRemoved := len(docs) - len(filteredOccurrences)
		retryTicker := time.NewTimer(30 * time.Second)
		defer retryTicker.Stop()
		offset := len(docs)
		retries := 0
		for {
			if !hasNext {
				log.Trace("exiting: fetch more loop: hasNext == false")
				break
			}
			// only guard on numRemoved the first time
			if retries == 0 && numRemoved == 0 {
				log.Trace("exiting: fetch more loop: numRemoved == 0")
				break
			}
			if !(len(filteredOccurrences) < limit) {
				log.Trace("exiting: fetch more loop: len(filteredOccurrences) >= limit")
				break
			}
			if retries > 100 {
				break
			}
			// Limit how much we over-fetch docs from Mongo on retries
			newLimit := min(numRemoved, limit, 100)
			hasNext = false
			currDocs := make([]*bson.M, 0)
			err = mongodb.FindWithRetry(
				ctx,
				filters,
				[]*options.FindOptions{
					findOpt.
						SetSkip(int64(offset)).
						SetLimit(int64(newLimit)),
				},
				db.collections[CollOccurrences].Find,
				func(result *bson.M) (bool, string, error) {
					if newLimit > 0 && len(currDocs)+1 == newLimit {
						hasNext = true
						return false, "", nil
					}
					currDocs = append(currDocs, result)
					id, ok := (*result)["_id"].(string)
					return ok, id, nil
				},
			)
			if err != nil {
				return nil, errors.Wrap(err, "failed to find occurrences with retry")
			}
			var currOccurrences []*event.Occurrence
			currOccurrences, err = convertBulk[*bson.M, *event.Occurrence](currDocs)
			if err != nil {
				return nil, errors.Wrap(err, "failed to convert occurrence docs")
			}
			currOccurrences, err = opt.ApplyOccurrenceProcessors(ctx, currOccurrences)
			if err != nil {
				return nil, errors.Wrap(err, "failed to apply occurrence post-processors")
			}
			docs = append(docs, currDocs...)
			offset = len(docs)
			// If we did over-fetch, then only keep the amount needed
			// to fill out page
			j := min(len(currOccurrences), int(math.Abs(float64(limit-len(currOccurrences)))))
			filteredOccurrences = append(filteredOccurrences, currOccurrences[:j]...)
			numRemoved = len(currDocs) - len(currOccurrences)
			retries++
			select {
			case <-retryTicker.C:
				break
			default:
				// no op
			}
		}
	}

	if log.Logger.IsLevelEnabled(logrus.TraceLevel) {
		instrumentation.AnnotateSpan("occurrenceResults",
			"got occurrence query results",
			span,
			map[string]any{
				"rawOccurrenceCount":      len(docs),
				"filteredOccurrenceCount": len(filteredOccurrences),
			}, nil)
	}
	if log.Logger.IsLevelEnabled(logrus.DebugLevel) {
		log.WithFields(logrus.Fields{
			"rawOccurrenceCount":      len(docs),
			"filteredOccurrenceCount": len(filteredOccurrences),
		}).Debug("table:rawOccurrenceCounts")
	}

	results := make([]*event.Event, 0)
	occMap := make(map[string]*event.Occurrence)
	eventIDs := make([]string, 0)
	occurrenceIDs := make([]string, 0, len(filteredOccurrences))
	sliceStop := len(filteredOccurrences)
	if len(filteredOccurrences) > 0 && limit > 0 {
		sliceStop = min(sliceStop, limit-1)
	}
	for _, occurrence := range filteredOccurrences[:sliceStop] {
		occurrenceIDs = append(occurrenceIDs, occurrence.ID)
		occMap[occurrence.ID] = occurrence
		newResult := &event.Event{
			ID:          occurrence.EventID,
			Tenant:      occurrence.Tenant,
			Entity:      occurrence.Entity,
			Occurrences: []*event.Occurrence{occurrence},
		}
		if occurrence.Status != event.StatusClosed {
			newResult.LastActiveOccurrence = occurrence
		}
		results = append(results, newResult)
		eventIDs = append(eventIDs, newResult.ID)
	}

	if len(occurrenceIDs) > 0 && slices.Contains(query.Fields, "notes") {
		noteMut := sync.Mutex{}
		allNotes := make([]*Note, 0)
		err := batchops.DoConcurrently(ctx, 500, 10, occurrenceIDs,
			func(ctx context.Context, batch []string) ([]*Note, error) {
				notesInFilter := bson.D{{Key: "occid", Value: bson.D{{Key: OpIn, Value: batch}}}}
				log.WithFields(logrus.Fields{
					"filters": notesInFilter,
				}).Debugf("executing %s.Find", "note")
				notes := make([]*Note, 0)
				err := mongodb.FindWithRetry[*Note](
					ctx,
					notesInFilter,
					[]*options.FindOptions{defaultFindOpts(options.Find().SetSort(bson.D{{"updatedAt", -1}}))},
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
			func(ctx context.Context, notes []*Note) (bool, error) {
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
		var allNotes2 []*event.Note
		allNotes2, err = convertBulk[*Note, *event.Note](allNotes)
		if err != nil {
			return nil, errors.Wrap(err, "failed to convert note struct")
		}
		for _, note := range allNotes2 {
			occ, ok := occMap[note.OccID]
			if !ok {
				continue
			}
			if len(occ.Notes) == 0 {
				occ.Notes = make([]*event.Note, 0)
			}
			occ.Notes = append(occ.Notes, note)
		}
	}

	if len(eventIDs) > 0 && slices.Contains(query.Fields, "dimensions") {
		eventMapMut := sync.Mutex{}
		eventMap := make(map[string]*EventDimensions)
		err := batchops.DoConcurrently(ctx, 500, 10, eventIDs,
			func(ctx context.Context, batch []string) ([]*EventDimensions, error) {
				eventInFilter := bson.D{{Key: OpIn, Value: batch}}
				log.WithFields(logrus.Fields{
					"filters": eventInFilter,
				}).Debugf("executing %s.Find", "event")
				eventDocs := make([]*EventDimensions, 0)
				err := mongodb.FindWithRetry[*EventDimensions](
					ctx,
					bson.D{{Key: "_id", Value: eventInFilter}},
					[]*options.FindOptions{defaultFindOpts(options.Find())},
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
			func(ctx context.Context, eventDocs []*EventDimensions) (bool, error) {
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
				if len(other.Entity) > 0 {
					result.Entity = other.Entity
				}
			} else {
				log.
					WithField("currResult", result).
					Warn("got event-occurrence document without a corresponding event document")
			}
		}
	}
	resultsCursor := ""
	if len(docs) > 0 {
		// Find the index in the slice of all occurrence docs fetched
		// so the paginator can capture the correct offset for the next page.
		lastResultIdx := len(docs)
		if len(filteredOccurrences) > 0 {
			lastOccID := filteredOccurrences[len(filteredOccurrences)-1].ID
			for i, doc := range docs {
				if (*doc)["_id"] == lastOccID {
					lastResultIdx = i + 1
					break
				}
			}
		}
		resultsCursor, err = UpsertCursor(ctx, db.cursorRepo, db.pager, query, docs[:lastResultIdx])
		if err != nil {
			return nil, errors.Wrap(err, "failed to upsert cursor")
		}
	}
	hasPrev := false
	startCursor := ""
	if query.PageInput != nil {
		if len(query.PageInput.Cursor) > 0 {
			startCursor = query.PageInput.Cursor
			if query.PageInput.Direction == event.PageDirectionForward {
				hasPrev = true
			}
		}
	}
	return &event.Page{
		Results:     results,
		HasNext:     hasNext,
		HasPrev:     hasPrev,
		StartCursor: startCursor,
		EndCursor:   resultsCursor,
	}, nil
}

func (db *Adapter) Update(_ context.Context, _ *event.Event) (*event.Event, error) {
	panic("not implemented") // TODO: Implement
}
