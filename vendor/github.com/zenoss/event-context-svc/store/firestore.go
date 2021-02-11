package store

import (
	"context"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/zenoss/event-context-svc/config"
	"golang.org/x/sync/semaphore"

	"github.com/zenoss/event-context-svc/utils"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/zenoss/zingo/v4/protobufutils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.opencensus.io/trace"

	"cloud.google.com/go/firestore"
	"github.com/pkg/errors"
	"github.com/zenoss/event-context-svc/metrics"
	"github.com/zenoss/zenkit/v5"
	eventProto "github.com/zenoss/zing-proto/v11/go/event"
	"google.golang.org/api/iterator"

	"go.opencensus.io/stats"
)

const (
	tenantsCollectionName         = "EventContextTenants"
	allEventsByDateCollectionName = "AllEventsByDate"
	eventsCollectionName          = "Events"
	entitiesCollectionName        = "Entities"
	allActiveEventsCollectionName = "AllActiveEvents"
	activityDatesCollectionName   = "ActivityDates"
	activeEventsCollectionName    = "ActiveEvents"
	occurrencesCollectionName     = "Occurrences"
	notesCollectionName           = "Notes"
)

var (
	// ErrInvalidArgument occurs when a client makes a request with an invalid argument
	ErrInvalidArgument = errors.New("invalid argument")
	// Encoding is used for creating/reading a cursor for pagination
	Encoding = base64.URLEncoding
	// ErrRestrictedEvent occurs when attempting to hydrate a restricted event
	ErrRestrictedEvent = errors.New("restricted event")
)

// Client an interface implemented by cloud.google.com/go/firestore.Client
// generated with `interfacer -for cloud.google.com/go/firestore.Client -as "CFSClient"`
type Client interface {
	Batch() *firestore.WriteBatch
	Close() error
	Collection(string) *firestore.CollectionRef
	CollectionGroup(string) *firestore.CollectionGroupRef
	Collections(context.Context) *firestore.CollectionIterator
	Doc(string) *firestore.DocumentRef
	GetAll(context.Context, []*firestore.DocumentRef) ([]*firestore.DocumentSnapshot, error)
	RunTransaction(context.Context, func(context.Context, *firestore.Transaction) error, ...firestore.TransactionOption) error
}

// CFEventStore is a EventContextStore implementation backed by Cloud Firestore
type CFEventStore struct {
	Client   Client
	mu       sync.RWMutex
	shardMap ShardMap
	workers  *semaphore.Weighted
}

// ActiveEvents is the return value for GetActiveEvents/GetAllActiveEvents
type ActiveEvents struct {
	ActiveEvents []ActiveEvent
	PageInfo     PageInfo
}

// SearchEvents is the return value for Search
type SearchEvents struct {
	Events   []*Event
	PageInfo PageInfo
}

//ActiveEventsCursor contains document information for pagination
type ActiveEventsCursor struct {
	EntityID string
	Day      string
	EventID  string
}

func getValidCursor(c string) (*ActiveEventsCursor, error) {
	if c == "" {
		return nil, errors.New("empty cursor")
	}
	aec := &ActiveEventsCursor{
		EntityID: "",
		Day:      "",
		EventID:  "",
	}
	e := DecodeCursor(c, aec)
	if e != nil {
		return nil, errors.New("unable to decode cursor")
	}
	return aec, nil
}

func stringSliceIndex(a string, list []string) int {
	for i, b := range list {
		if b == a {
			return i
		}
	}
	return -1
}

func reverseSlice(list []string) {
	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}
}

// Transaction is an interface generated for "cloud.google.com/go/firestore.Transaction".
type Transaction interface {
	Create(*firestore.DocumentRef, interface{}) error
	Delete(*firestore.DocumentRef, ...firestore.Precondition) error
	DocumentRefs(*firestore.CollectionRef) *firestore.DocumentRefIterator
	Documents(firestore.Queryer) *firestore.DocumentIterator
	Get(*firestore.DocumentRef) (*firestore.DocumentSnapshot, error)
	GetAll([]*firestore.DocumentRef) ([]*firestore.DocumentSnapshot, error)
	Set(*firestore.DocumentRef, interface{}, ...firestore.SetOption) error
	Update(*firestore.DocumentRef, []firestore.Update, ...firestore.Precondition) error
}

// NewCloudFirestoreEventsStore returns a event store for the client
func NewCloudFirestoreEventsStore(client Client) EventContextStore {
	return &CFEventStore{Client: client, shardMap: NewShardMap(), workers: semaphore.NewWeighted(viper.GetInt64(config.EventContextStoreNumWorkersConfig))}
}

// Path returns the path to receiver's document in Cloud Firestore
func (e *Event) Path(tenant string) string {
	return fmt.Sprintf("%s/%s/%s/%s", tenantsCollectionName, tenant, eventsCollectionName, e.ID)
}

// Path returns the path to receiver's document in Cloud Firestore
func (o *Occurrence) Path(tenant, eventID string) string {
	return fmt.Sprintf("%s/%s/%s/%s/%s/%s", tenantsCollectionName, tenant, eventsCollectionName, eventID, occurrencesCollectionName, o.ID)
}

// Path return the path the receiver's document in the Cloud Firestore
func (n *Note) Path(tenant, eventID, occurrenceID string) string {
	return fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s/%s", tenantsCollectionName, tenant, eventsCollectionName, eventID, occurrencesCollectionName, occurrenceID, notesCollectionName, n.ID)
}

// To add a new note to batch
func addNoteToBatch(note *Note, notesCollectionRef *firestore.CollectionRef, batch *firestore.WriteBatch) {
	newNoteDocRef := notesCollectionRef.NewDoc()
	// store doc ref id as ID so that we can get it for updates.
	note.ID = newNoteDocRef.ID
	batch.Create(newNoteDocRef, note)
}

func (eventsStore *CFEventStore) PutInTransaction(ctx context.Context, eventRef *firestore.DocumentRef, event *Event) error {
	ctx, span := trace.StartSpan(ctx, "CFEventStore.PutInTransaction")
	defer span.End()
	return eventsStore.Client.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {

		return eventsStore.CreateEventWithTransaction(ctx, tx, eventRef, event)
	})
}

type writeOp struct {
	dr     *firestore.DocumentRef
	data   interface{}
	errMsg string
}

type WriteOp interface {
	isWriteOp()
	errMsg() string
}

type opCreate struct {
	dr   *firestore.DocumentRef
	data interface{}
	msg  string
}

func (o *opCreate) isWriteOp() {}
func (o *opCreate) errMsg() string {
	return o.msg
}
func (o *opCreate) String() string {
	return fmt.Sprintf("opCreate{%s}", o.dr.Path)
}

var _ WriteOp = &opCreate{}

type opUpdate struct {
	dr   *firestore.DocumentRef
	data []firestore.Update
	opts []firestore.Precondition
	msg  string
}

func (o *opUpdate) isWriteOp() {}
func (o *opUpdate) errMsg() string {
	return o.msg
}
func (o *opUpdate) String() string {
	return fmt.Sprintf("opUpdate{%s}", o.dr.Path)
}

var _ WriteOp = &opUpdate{}

type opDelete struct {
	dr   *firestore.DocumentRef
	opts []firestore.Precondition
	msg  string
}

func (o *opDelete) isWriteOp() {}
func (o *opDelete) errMsg() string {
	return o.msg
}
func (o *opDelete) String() string {
	return fmt.Sprintf("opUpdate{%s}", o.dr.Path)
}

var _ WriteOp = &opDelete{}

type opSet struct {
	dr   *firestore.DocumentRef
	data interface{}
	opts []firestore.SetOption
	msg  string
}

func (o *opSet) isWriteOp() {}
func (o *opSet) errMsg() string {
	return o.msg
}
func (o *opSet) String() string {
	return fmt.Sprintf("opSet{%s}", o.dr.Path)
}

var _ WriteOp = &opSet{}

func DoWriteWithTransaction(ctx context.Context, tx Transaction, op WriteOp) error {
	log := zenkit.ContextLogger(ctx)
	switch o := op.(type) {
	case *opCreate:
		log.WithFields(logrus.Fields{
			"path":    o.dr.Path,
			"updates": o.data,
		}).Debug("performing Create operation")
		return tx.Create(o.dr, o.data)
	case *opUpdate:
		log.WithFields(logrus.Fields{
			"path":    o.dr.Path,
			"updates": o.data,
		}).Debug("performing Update operation")
		return tx.Update(o.dr, o.data, o.opts...)
	case *opSet:
		log.WithFields(logrus.Fields{
			"path":    o.dr.Path,
			"updates": o.data,
		}).Debug("performing Set operation")
		return tx.Set(o.dr, o.data, o.opts...)
	case *opDelete:
		log.WithField("path", o.dr.Path).Debug("performing Delete operation")
		return tx.Delete(o.dr, o.opts...)
	default:
		return errors.New("unsupported operation type")
	}
}

func (eventsStore *CFEventStore) CreateEventWithTransaction(ctx context.Context, tx *firestore.Transaction, eventRef *firestore.DocumentRef, event *Event) error {
	ctx, span := trace.StartSpan(ctx, "CFEventStore.CreateEventWithTransaction")
	defer span.End()
	writeOps := make([]*writeOp, 0, 1+len(event.Occurrences))
	writeOps = append(writeOps, &writeOp{dr: eventRef, data: event})
	for _, occ := range event.Occurrences {
		occurrenceDocRef := eventsStore.Client.Doc(occ.Path(event.Tenant, event.ID))
		occ.CreatedAt = time.Now()
		writeOps = append(writeOps, &writeOp{dr: occurrenceDocRef, data: occ})
		notesCollectionRef := occurrenceDocRef.Collection(notesCollectionName)
		for _, note := range occ.Notes {
			newNoteDocRef := notesCollectionRef.NewDoc()
			note.ID = newNoteDocRef.ID
			note.CreatedAt = time.Now()
			writeOps = append(writeOps, &writeOp{
				dr:   newNoteDocRef,
				data: note,
			})
		}
	}
	var err error
	for _, op := range writeOps {
		err = tx.Create(op.dr, op.data)
		if err != nil {
			zenkit.ContextLogger(ctx).WithError(err).Errorf("failed to write doc: %s", op.dr.Path)
			return err
		}
	}
	return nil
}

// PutEvent puts an event in the Events collection in Cloud Firestore
func (eventsStore *CFEventStore) PutEvent(ctx context.Context, tenant string, event *Event) error {
	ctx, span := trace.StartSpan(ctx, "CFEventStore.PutEvent")
	defer span.End()
	log := zenkit.ContextLogger(ctx)
	// grab the tenant
	if strings.TrimSpace(tenant) == "" {
		if strings.TrimSpace(event.Tenant) == "" {
			log.Error("invalid argument: no tenant")
			return utils.ErrInvalidIdentity
		}
		tenant = event.Tenant
	}
	evtRef := eventsStore.Client.Doc(event.Path(tenant))
	return eventsStore.PutInTransaction(ctx, evtRef, event)
}

type IDSnapshotPair struct {
	ID       string
	Snapshot *firestore.DocumentSnapshot
}
type CreateOrUpdateCallbackConfig struct {
	IdSnapshotPairs []*IDSnapshotPair
}

type CreateOrUpdateOperations struct {
	DocRef           *firestore.DocumentRef
	EventToBeCreated *Event
	EventToBeUpdated *UpdateEventPayload
}

func (opts *CreateOrUpdateOperations) Validate() error {
	if opts.EventToBeCreated != nil && opts.EventToBeUpdated != nil {
		return errors.New("cannot create and update in same operation")
	}
	return nil
}

// getFirestoreUpdates uses reflection to iterate over the UpdateOccurrencePayload
// field values and returns an array of firestore.Updates.
// Please update this method if you wish to exclude any fields from the updates returned
// by it, which should be unnecessary given that UpdateOccurrencePayload should only
// fields that should be updated.
func (occ *UpdateOccurrencePayload) getFirestoreUpdates() []firestore.Update {
	updates := make([]firestore.Update, 0)
	iVal := reflect.ValueOf(occ).Elem()
	for i := 0; i < iVal.NumField(); i++ {
		field := iVal.Field(i)
		typ := iVal.Type().Field(i)
		fieldName := typ.Name
		if fieldName == "ID" || fieldName == "NotesToBeAdded" || fieldName == "NotesToBeUpdated" {
			continue
		}
		if !field.IsZero() {
			if field.Type().Name() != "Acknowledged" && field.Kind() == reflect.Ptr {
				updates = append(updates, firestore.Update{Path: typ.Name, Value: field.Elem().Interface()})
			} else {
				updates = append(updates, firestore.Update{Path: typ.Name, Value: field.Interface()})
			}
		}
	}
	return updates
}

// UpdateEvent updates an event document in Cloud Firestore
func (eventsStore *CFEventStore) UpdateEvent(ctx context.Context, tenant string, payload *UpdateEventPayload) error {
	ctx, span := trace.StartSpan(ctx, "CFEventStore.UpdateEvent")
	defer span.End()
	log := zenkit.ContextLogger(ctx)
	// grab the tenant
	if strings.TrimSpace(payload.Tenant) == "" {
		log.Error("invalid argument: no tenant")
		return utils.ErrInvalidIdentity
	}
	shard := eventsStore.shardMap.GetShard(payload.ID)
	shard.Lock()
	defer shard.Unlock()
	err := eventsStore.Client.RunTransaction(ctx, func(txCtx context.Context, tx *firestore.Transaction) error {
		writeOps := make([]*writeOp, 0)
		// update existing occurrences
		for _, otbu := range payload.OccurrencesToBeUpdated {
			evOcc := eventsStore.Client.Doc(fmt.Sprintf("%s/%s/%s/%s/%s/%s",
				tenantsCollectionName, payload.Tenant, eventsCollectionName, payload.ID, occurrencesCollectionName, otbu.ID))

			occUpdates := otbu.getFirestoreUpdates()
			if len(occUpdates) > 0 {
				writeOps = append(writeOps, &writeOp{
					dr:     evOcc,
					data:   occUpdates,
					errMsg: "failed to update occurrence",
				})
			}

			// add new notes
			notesCollectionRef := evOcc.Collection(notesCollectionName)
			for _, note := range otbu.NotesToBeAdded {
				newNoteDocRef := notesCollectionRef.NewDoc()
				// store doc ref id as ID so that we can get it for updates.
				note.ID = newNoteDocRef.ID
				writeOps = append(writeOps, &writeOp{
					dr:     newNoteDocRef,
					data:   note,
					errMsg: "failed to create new note",
				})
			}

			// update notes
			for _, note := range otbu.NotesToBeUpdated {
				// get the note from the collection to update
				noteRef := notesCollectionRef.Doc(note.ID)
				if noteRef != nil {
					writeOps = append(writeOps, &writeOp{
						dr:     noteRef,
						data:   []firestore.Update{{Path: "content", Value: note.Content}},
						errMsg: "failed to update note",
					})
				} else {
					err := errors.New("Cannot find note " + note.ID + " to update for occurrence" + otbu.ID + " for event " + payload.ID)
					log.WithError(err).Error("failed to update note")
					return err
				}
			}
		}
		// add new occurrences
		for _, occ := range payload.OccurrencesToBeAdded {
			occurrenceDocRef := eventsStore.Client.Doc(occ.Path(payload.Tenant, payload.ID))
			if occurrenceDocRef == nil {
				log.WithField("currentOccurrence", occ).Error("invalid path for new occurrence")
				return ErrInvalidArgument
			}
			writeOps = append(writeOps, &writeOp{
				dr:     occurrenceDocRef,
				data:   occ,
				errMsg: "failed to create new occurrence",
			})
			notesCollectionRef := occurrenceDocRef.Collection(notesCollectionName)
			for _, note := range occ.Notes {
				newNoteDocRef := notesCollectionRef.NewDoc()
				// store doc ref id as ID so that we can get it for updates.
				note.ID = newNoteDocRef.ID
				writeOps = append(writeOps, &writeOp{
					dr:     newNoteDocRef,
					data:   note,
					errMsg: "failed to create new note",
				})
			}
		}
		for _, op := range writeOps {
			if updates, ok := op.data.([]firestore.Update); ok {
				err := tx.Update(op.dr, updates)
				if err != nil {
					return errors.Wrap(err, op.errMsg)
				}
			} else {
				err := tx.Create(op.dr, op.data)
				if err != nil {
					return errors.Wrap(err, op.errMsg)
				}
			}
		}
		return nil
	})
	return err
}

func (eventsStore *CFEventStore) getEventsCollection(ctx context.Context, tenant string) *firestore.CollectionRef {
	return eventsStore.Client.Collection(tenantsCollectionName).Doc(tenant).Collection(eventsCollectionName)
}

// GetEvent retrieve the event for the given ID if it founds in the collection, or an error otherwise
func (eventsStore *CFEventStore) GetEvent(ctx context.Context, tenant string, ID string) (*Event, error) {
	ctx, span := trace.StartSpan(ctx, "CFEventStore.GetEvent")
	defer span.End()
	log := zenkit.ContextLogger(ctx)
	// grab the tenant
	if strings.TrimSpace(tenant) == "" {
		log.Error("invalid argument: no tenant")
		return nil, utils.ErrInvalidIdentity
	}
	docRef := eventsStore.getEventsCollection(ctx, tenant).Doc(ID)
	var result *Event
	err := eventsStore.Client.RunTransaction(ctx, func(ctx context.Context, transaction *firestore.Transaction) error {
		snapshot, err := docRef.Get(ctx)
		if err != nil {
			log.WithError(err).Error("failed to get event")
			return err
		}
		if !snapshot.Exists() {
			return errors.New("event not found")
		}
		result, err = hydrateEventDocument(ctx, true, -1, -1, snapshot, nil)
		if err != nil {
			return err
		}
		return nil
	}, firestore.ReadOnly)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func sinceInMilliseconds(startTime time.Time) float64 {
	return float64(time.Since(startTime).Nanoseconds()) / 1e6
}

type EventResult struct {
	event *Event
	err   error
}

func applyOccurrenceCriteria(ctx context.Context, eventSnapshot *firestore.DocumentSnapshot, query *Query, workers *semaphore.Weighted) chan EventResult {
	ctx, span := trace.StartSpan(ctx, "applyOccurrenceCriteria")
	defer span.End()
	log := zenkit.ContextLogger(ctx)
	out := make(chan EventResult)
	go func() {
		defer close(out)
		err := workers.Acquire(ctx, 1)
		if err != nil {
			log.WithError(err).Error("failed to hydrate event document")
			return
		}
		defer workers.Release(1)
		event, err := hydrateEventDocument(ctx, false, -1, -1, eventSnapshot, query)
		if err != nil {
			log.WithError(err).Error("failed to hydrate event document")
			out <- EventResult{
				err: err,
			}
			return
		}
		if len(event.Occurrences) > 0 {
			out <- EventResult{
				event: event,
			}
		}
		return
	}()
	return out
}

func mergeEventResultChannels(ctx context.Context, cs ...<-chan EventResult) <-chan EventResult {
	out := make(chan EventResult)
	var wg sync.WaitGroup
	wg.Add(len(cs))
	for _, c := range cs {
		go func(c <-chan EventResult) {
			defer wg.Done()
			for r := range c {
				select {
				case <-ctx.Done():
					return
				case out <- r:
				}
			}
		}(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// Search retrieves an array of events from the Events collection in Cloud Firestore
func (eventsStore *CFEventStore) Search(ctx context.Context, query *Query) (*SearchEvents, error) {
	var (
		qCursor        = new(QueryCursor)
		hasNext        = false
		hasPrev        = false
		lastOccurrence *Occurrence
		cursor         = make([]byte, 0)
		cursErr        error
	)
	ctx, span := trace.StartSpan(ctx, "CFEventStore.Search")
	log := zenkit.ContextLogger(ctx)
	log.WithField("query", query).Debug("performing search")
	// grab the tenant

	eventResults := make([]*Event, 0)

	// setup metrics
	mTime := time.Now()
	defer func() {
		span.End()
		stats.Record(ctx, metrics.MSearchTimeMs.M(sinceInMilliseconds(mTime)),
			metrics.MSearchCount.M(int64(len(eventResults))))
	}()

	if len(query.EventFilters)+len(query.DimensionFilters) == 0 {
		currQuery := &Query{
			Tenant:            query.Tenant,
			TimeRange:         query.TimeRange,
			Severities:        query.Severities,
			Statuses:          query.Statuses,
			EventFilters:      query.EventFilters,
			DimensionFilters:  query.DimensionFilters,
			OccurrenceFilters: query.OccurrenceFilters,
			PageInput:         query.PageInput,
		}
		var (
			origPageInput   *PageInput = query.PageInput
			pageInfo        *PageInfo
			eventsSnapshots []*firestore.DocumentSnapshot
			err             error
		)
	activeEventQueryLoop:
		for {
			eventsSnapshots, pageInfo, err = ResolveOccurrencesViaActiveEvents(ctx, eventsStore.Client, currQuery, eventsStore)
			if err != nil {
				log.WithError(err).Error("failed to get active event snapshots")
				return nil, err
			}
			eventIDSet := utils.NewStringSet()
			eventChannels := make([]<-chan EventResult, 0, len(eventsSnapshots))
			for _, snap := range eventsSnapshots {
				idV, err := snap.DataAt("id")
				if err != nil {
					log.WithError(err).Warn("got an event without an ID, skipping")
					continue
				}
				// drop duplicate eventIDs
				if id, ok := idV.(string); !ok || !eventIDSet.Add(id) {
					continue
				}
				eventChannels = append(eventChannels, applyOccurrenceCriteria(ctx, snap, currQuery, eventsStore.workers))
				//event, err := hydrateEventDocument(ctx, false, -1, -1, snap, currQuery)
				//if err != nil {
				//    log.WithError(err).Error("failed to hydrate event document")
				//    return nil, err
				//}
				//if len(event.Occurrences) > 0 {
				//    eventResults = append(eventResults, event)
				//}
			}
			for result := range mergeEventResultChannels(ctx, eventChannels...) {
				if result.err != nil {
					return nil, err
				}
				if result.event == nil {
					continue
				}
				eventResults = append(eventResults, result.event)
			}
			// retrieve more active events unless the client either
			//   - did not set a limit on results
			//   - or there were no results in the last attempt
			//   - or the count of final results matching query filters equal the request limit
			if origPageInput == nil || pageInfo.Count == 0 || !pageInfo.HasNext ||
				len(eventResults) >= origPageInput.Limit {
				break activeEventQueryLoop
			}
			currQuery.PageInput.Limit = origPageInput.Limit - len(eventResults)
			currQuery.PageInput.Cursor = pageInfo.Cursor
		}

		return &SearchEvents{
			Events: eventResults,
			PageInfo: PageInfo{
				Cursor:  pageInfo.Cursor,
				Count:   uint64(len(eventResults)),
				HasNext: pageInfo.HasNext,
				HasPrev: pageInfo.HasPrev,
			},
		}, nil
		//return eventsStore.searchViaOccurrenceCollectionGroupQuery(ctx, query, log, lastOccurrence, eventResults, qCursor, hasNext, cursor, cursErr, hasPrev)
	}

	eventSnapshots, err := ResolveEventDocuments(ctx, query, eventsStore.Client)
	if err != nil {
		return nil, err
	}
	eventIDs := make([]string, len(eventSnapshots))
	eventResult := make([]*Event, len(eventSnapshots))
	for i, snap := range eventSnapshots {

		e := &Event{}
		err = snap.DataTo(e)
		if err != nil {
			return nil, err
		}
		eventIDs[i] = e.ID
		eventResult[i] = e
	}
	if len(eventSnapshots)-1 != -1 {
		qCursor.EventID = eventIDs[len(eventSnapshots)-1]
	}
	byEventID, err := ResolveOccurrencesByEventIds(ctx, eventsStore.Client, query, eventIDs)
	if err != nil {
		return nil, err
	}

	for _, e := range eventResult {
		occs := make([]*Occurrence, 0)
		for _, occSnap := range byEventID[e.ID] {
			occ := &Occurrence{}
			err = occSnap.DataTo(occ)
			if err != nil {
				return nil, err
			}
			occs = append(occs, occ)
			lastOccurrence = occ
		}
		if len(occs) > 0 {
			// add only if there are occurrences for the event that match
			e.Occurrences = occs
			eventResults = append(eventResults, e)
		}
	}
	if query.PageInput != nil {
		if lastOccurrence != nil {
			qCursor.StartTime = lastOccurrence.StartTime
		}
		if len(eventResults) == query.PageInput.Limit || len(eventResults) < len(eventIDs) {
			hasNext = true
		}
		if len(eventResults) > 0 {
			if eventResult[len(eventSnapshots)-1].OccurrenceCount > int64(len(eventResults[len(eventResults)-1].Occurrences)) {
				hasNext = true
				qCursor.HasNextOccurrence = true
			}
		}
		if query.PageInput != nil && query.PageInput.Cursor != "" {
			hasPrev = true
		}
		cursor, cursErr = json.Marshal(qCursor)
		if cursErr != nil {
			log.WithError(cursErr).Error("error creating cursor while searching events")
			cursor, _ = json.Marshal("")
		}
	}
	results := new(SearchEvents)
	results.Events = eventResults
	results.PageInfo = PageInfo{
		Cursor:  Encoding.EncodeToString([]byte(cursor)),
		Count:   uint64(len(eventResults)),
		HasPrev: hasPrev,
		HasNext: hasNext,
	}

	return results, nil
}

func (eventsStore *CFEventStore) searchViaOccurrenceCollectionGroupQuery(ctx context.Context, query *Query, log *logrus.Entry, lastOccurrence *Occurrence, result []*Event, qCursor *QueryCursor, hasNext bool, cursor []byte, cursErr error, hasPrev bool) (*SearchEvents, error) {
	occurrenceSnapshots, err := ResolveOccurrencesCollectionGroupQuery(ctx, eventsStore.Client, query)
	if err != nil {
		log.WithError(err).Error("failed getting occurrences from collection group")
		return nil, err
	}
	byEventID := make(map[string][]*Occurrence)

	totalOccurrences := 0
	for _, snap := range occurrenceSnapshots {
		occur := &Occurrence{}
		if !snap.Exists() {
			continue
		}
		err = snap.DataTo(occur)
		if err != nil {
			return nil, err
		}
		if byEventID[occur.EventID] == nil {
			byEventID[occur.EventID] = make([]*Occurrence, 0)
		}
		byEventID[occur.EventID] = append(byEventID[occur.EventID], occur)
		totalOccurrences++
	}
	for eventID, occurs := range byEventID {
		snap, err := eventsStore.Client.Doc(fmt.Sprintf("%s/%s/%s/%s", tenantsCollectionName, query.Tenant, eventsCollectionName, eventID)).Get(ctx)
		if err != nil {
			if status.Code(err) == codes.NotFound {
				log.WithError(err).Warn("firestore returned non-existent document")
				continue
			}
			log.WithError(err).Error("failed to get occurrence's parent event")
			return nil, err
		}
		event := &Event{}
		err = snap.DataTo(event)
		if err != nil {
			if status.Code(err) == codes.NotFound {
				log.WithError(err).Warn("firestore returned non-existent document")
				continue
			}
			log.WithError(err).Error("failed to unmarshal occurrence's parent event")
			return nil, err
		}
		event.Occurrences = occurs
		occursCount := len(occurs)

		lastOccurrence = occurs[occursCount-1]
		result = append(result, event)
	}
	if query.PageInput != nil {
		if lastOccurrence != nil {
			qCursor.StartTime = lastOccurrence.StartTime
			qCursor.EventID = lastOccurrence.EventID
		}
		limit := query.PageInput.OccurrenceLimit
		if limit == 0 {
			limit = query.PageInput.Limit
		}
		if query.PageInput != nil && totalOccurrences == limit {
			hasNext = true
		}
		if resultLen := len(result); resultLen > 0 {
			if result[resultLen-1].OccurrenceCount > int64(len(result[resultLen-1].Occurrences)) {
				hasNext = true
				qCursor.HasNextOccurrence = true
			}
		}
		cursor, cursErr = json.Marshal(qCursor)
		if cursErr != nil {
			log.WithError(cursErr).Error("error creating cursor while searching events")
			cursor, _ = json.Marshal("")
		}
	}
	results := new(SearchEvents)
	results.Events = result
	results.PageInfo = PageInfo{
		Cursor:  Encoding.EncodeToString([]byte(cursor)),
		Count:   uint64(len(result)),
		HasPrev: hasPrev,
		HasNext: hasNext,
	}

	return results, nil
}

const (
	EntityKey = "_zen_direct_entity_id"
)

func OccurrenceFirestoreUpdatesFromProto(event *eventProto.Event, shouldUpdateStatus bool) []firestore.Update {
	results := []firestore.Update{
		{
			Path:  "severity",
			Value: SeverityFromProto(event.Severity),
		},
		{
			Path:  "type",
			Value: event.Type,
		},
		{
			Path:  "body",
			Value: event.Body,
		},
		{
			Path:  "summary",
			Value: event.Summary,
		},
	}
	if event.Acknowledged != nil {
		results = append(results, firestore.Update{
			Path:  "acknowledged",
			Value: &event.Acknowledged.Value,
		})
	}
	if shouldUpdateStatus {
		results = append(results, firestore.Update{
			Path:  "status",
			Value: StatusFromProto(event.Status),
		})
	}
	return results
}

func createNewOccurrence(txCtx context.Context, tx Transaction, eventMsg *eventProto.Event, eventDocRef *firestore.DocumentRef, currEvent *Event, occursColRef *firestore.CollectionRef) ([]WriteOp, error) {
	writeOps := make([]WriteOp, 0)
	occurrenceCount := len(currEvent.Occurrences)
	occID := fmt.Sprintf("%s:%d", eventMsg.Id, occurrenceCount+1)
	newOccurrence := Occurrence{
		ID:        occID,
		EventID:   currEvent.ID,
		Tenant:    currEvent.Tenant,
		Status:    StatusFromProto(eventMsg.Status),
		Severity:  SeverityFromProto(eventMsg.Severity),
		StartTime: eventMsg.Timestamp,
		Body:      eventMsg.Body,
		Summary:   eventMsg.Summary,
		Type:      eventMsg.Type,
		CreatedAt: time.Now(),
	}
	if eventMsg.Status == eventProto.Status_STATUS_CLOSED {
		newOccurrence.EndTime = eventMsg.Timestamp
	}
	if eventMsg.Acknowledged != nil {
		newOccurrence.Acknowledged = &eventMsg.Acknowledged.Value
	}
	writeOps = append(writeOps, &opCreate{
		dr:   occursColRef.Doc(occID),
		data: &newOccurrence,
		msg:  "create new occurrence",
	})
	writeOps = append(writeOps, &opUpdate{
		dr: eventDocRef,
		data: []firestore.Update{
			{
				Path:  "lastActiveOccurrenceID",
				Value: occID,
			},
			{
				Path:  "occurrenceCount",
				Value: occurrenceCount + 1,
			},
		},
		msg: "update event: last active occurrence",
	})
	return writeOps, nil
}

func nextOccSplitID(occID string) (string, error) {
	// This regex will recognize all components of a occurrenceID:
	// `([-a-zA-Z0-9=]+):(?P<occNum>\d+)(_(?P<splitID>\w))?`;
	// however, given a ID for the next split occurrence, we only need the first two
	// components to construct a split occurrence ID
	re, err := regexp.Compile(`(?P<eventID>[-a-zA-Z0-9=_]+):(?P<occNum>\d+)(_(?P<splitID>\w))?`)
	if err != nil {
		return "", err
	}
	buf := make([]byte, binary.MaxVarintLen64)
	binary.LittleEndian.PutUint64(buf, uint64(time.Now().UnixNano()))

	// splitID is a base64 string encoding of the current time
	splitID := base64.StdEncoding.EncodeToString(buf)
	match := re.FindStringSubmatch(occID)
	switch len(match) {
	case 3:
		return fmt.Sprintf("%s_%s", match[0], splitID), nil
	case 5:
		eventID, occNum := match[1], match[2]
		return fmt.Sprintf("%s:%s_%s", eventID, occNum, splitID), nil
	default:
		return "", errors.New("failed to recognize occurrence ID")
	}
}

func splitOccurrences(txCtx context.Context, tx Transaction, lastActiveOccurrence *Occurrence, event *eventProto.Event, docRef *firestore.DocumentRef, occursColRef *firestore.CollectionRef) ([]WriteOp, error) {
	txCtx, span := trace.StartSpan(txCtx, "splitOccurrences")
	defer span.End()
	log := zenkit.ContextLogger(txCtx)
	writeOps := make([]WriteOp, 0)
	var occToBeSplit Occurrence
	var targetOccRef *firestore.DocumentRef
	if lastActiveOccurrence.StartTime < event.Timestamp && event.Timestamp < lastActiveOccurrence.EndTime {
		targetOccRef = occursColRef.Doc(lastActiveOccurrence.ID)
		occToBeSplit = *lastActiveOccurrence
	}
	if targetOccRef == nil {
		// only update occurrences within the last 24 hours
		occSnapshots, err := occursColRef.
			Where("createdAt", ">=", time.Now().Add(-24*time.Hour)).
			Documents(txCtx).GetAll()
		if err != nil {
			log.WithError(err).Error("failed to get occurrences")
			return nil, errors.Wrap(err, "failed to get occurrences")
		}
		for _, snap := range occSnapshots {
			if snap.Exists() {
				err = snap.DataTo(&occToBeSplit)
				if err != nil {
					return nil, errors.Wrap(err, "failed to unmarshal occurrence")
				}
				if occToBeSplit.StartTime < event.Timestamp && event.Timestamp < occToBeSplit.EndTime {
					targetOccRef = snap.Ref
					break
				}
			}
		}
	}
	if targetOccRef != nil {
		writeOps = append(writeOps, &opUpdate{
			dr: targetOccRef,
			data: []firestore.Update{
				{Path: "endTime", Value: event.Timestamp},
				{Path: "currentTime", Value: event.Timestamp},
				{Path: "splitCount", Value: firestore.Increment(1)},
			},
			msg: "update occurrence to be split",
		})
		newOccID, err := nextOccSplitID(occToBeSplit.ID)
		if err != nil {
			return nil, errors.Wrap(err, "failed to generate split occurrence ID")
		}
		newOccDocRef := docRef.Collection(occurrencesCollectionName).Doc(newOccID)
		writeOps = append(writeOps, &opCreate{
			dr: newOccDocRef,
			data: &Occurrence{
				ID:           newOccID,
				EventID:      occToBeSplit.EventID,
				Tenant:       occToBeSplit.Tenant,
				Status:       occToBeSplit.Status,
				Severity:     occToBeSplit.Severity,
				Acknowledged: occToBeSplit.Acknowledged,
				StartTime:    occToBeSplit.EndTime - 1, // TODO: use some kind of temporary accounting to correctly set this
				EndTime:      occToBeSplit.EndTime,
				CurrentTime:  occToBeSplit.CurrentTime,
				Body:         occToBeSplit.Body,
				Summary:      occToBeSplit.Summary,
				Type:         occToBeSplit.Type,
				CreatedAt:    time.Now(),
			},
			msg: "create split occurrence",
		})
		writeOps = append(writeOps, &opUpdate{
			dr: docRef,
			data: []firestore.Update{
				{Path: "occurrenceCount", Value: firestore.Increment(1)},
			},
			msg: "update event: occurrenceCount",
		})
		if occToBeSplit.ID == lastActiveOccurrence.ID {
			writeOps = append(writeOps, &opUpdate{
				dr: docRef,
				data: []firestore.Update{
					{Path: "lastActiveOccurrenceID", Value: newOccID},
				},
				msg: "update event: last active occurrence",
			})
		}
	}
	return writeOps, nil
}

func (eventsStore *CFEventStore) createOrGetEventOccurrence(ctx context.Context, tx Transaction, eventDocRef *firestore.DocumentRef, eventMsg *eventProto.Event) (event *Event, occurrence *Occurrence, err error, isNew bool, writeOps []WriteOp) {
	ctx, span := trace.StartSpan(ctx, "CFEventStore.createOrGetEventOccurrence")
	defer span.End()
	log := zenkit.ContextLogger(ctx)
	writeOps = make([]WriteOp, 0)
	var docSnapshot *firestore.DocumentSnapshot
	docSnapshot, err = tx.Get(eventDocRef)
	if err != nil && status.Code(err) != codes.NotFound {
		log.WithError(err).Error("failed to get event document")
		return
	}
	if !docSnapshot.Exists() {
		var entity string
		var eventMetadata []interface{}
		eventMetadata, err = protobufutils.FromScalarArray(eventMsg.Metadata[EntityKey])
		if err != nil {
			log.WithError(err).Error("failed to get entity on event")
			return
		}
		if len(eventMetadata) > 0 {
			entity, _ = eventMetadata[0].(string)
		}
		var dim map[string]interface{}
		dim, err = protobufutils.FromScalarMap(eventMsg.Dimensions)
		if err != nil {
			log.WithError(err).Error("failed to get event dimensions")
			return
		}
		newOccurrence := &Occurrence{
			ID:          fmt.Sprintf("%s:1", eventMsg.Id),
			EventID:     eventMsg.Id,
			Tenant:      eventMsg.Tenant,
			StartTime:   eventMsg.Timestamp,
			CurrentTime: eventMsg.Timestamp,
			Status:      StatusFromProto(eventMsg.Status),
			Severity:    SeverityFromProto(eventMsg.Severity),
			Body:        eventMsg.Body,
			Summary:     eventMsg.Summary,
			Type:        eventMsg.Type,
		}
		if eventMsg.Acknowledged != nil {
			newOccurrence.Acknowledged = &eventMsg.Acknowledged.Value
		}
		if eventMsg.Status == eventProto.Status_STATUS_CLOSED {
			newOccurrence.EndTime = eventMsg.Timestamp
		}

		if len(eventMsg.Type) != 0 {
			newOccurrence.Type = eventMsg.Type
		}
		if len(eventMsg.Body) != 0 {
			newOccurrence.Body = eventMsg.Body
		}
		if len(eventMsg.Summary) != 0 {
			newOccurrence.Summary = eventMsg.Summary
		}

		newOccurrenceDocRef := eventDocRef.Collection(occurrencesCollectionName).Doc(newOccurrence.ID)
		newEvent := &Event{
			ID:                     eventMsg.Id,
			Tenant:                 eventMsg.Tenant,
			Entity:                 entity,
			Dimensions:             dim,
			LastActiveOccurrenceID: newOccurrence.ID,
			CreatedAt:              time.Now(),
			OccurrenceCount:        1,
		}
		writeOps = append(writeOps, &opCreate{
			dr:   eventDocRef,
			data: newEvent,
			msg:  fmt.Sprintf("create event: %s", eventDocRef.Path),
		})
		writeOps = append(writeOps, &opCreate{
			dr:   newOccurrenceDocRef,
			data: newOccurrence,
			msg:  fmt.Sprintf("create new occurrence: %s", newOccurrenceDocRef.Path),
		})
		newEvent.LastActiveOccurrence = newOccurrence
		event, occurrence, isNew = newEvent, newOccurrence, true
		return
	}
	var currEvent *Event
	currEvent, err = hydrateEventDocument(ctx, true, -1, -1, docSnapshot, nil)
	if err != nil {
		log.WithError(err).Error("failed to unmarshal event document")
		return
	}
	if currEvent.LastActiveOccurrenceID == "" {
		err = errors.New("failed to get active occurrence")
		return
	}
	occursColRef := eventDocRef.Collection(occurrencesCollectionName)
	activeOccDocRef := occursColRef.Doc(currEvent.LastActiveOccurrenceID)
	var activeOccSnapshot *firestore.DocumentSnapshot
	activeOccSnapshot, err = activeOccDocRef.Get(ctx)
	if err != nil {
		log.WithError(err).Error("failed to get active occurrence snapshot")
		return
	}
	var activeOccurrence Occurrence
	err = activeOccSnapshot.DataTo(&activeOccurrence)
	if err != nil {
		log.WithError(err).Error("failed to unmarshal occurrence document")
	}
	event, occurrence = currEvent, &activeOccurrence
	return
}

func (eventsStore *CFEventStore) updateOccurrences(ctx context.Context, tx Transaction, docRef *firestore.DocumentRef, event *eventProto.Event, currEvent *Event, activeOccurrence *Occurrence) ([]WriteOp, error) {
	ctx, span := trace.StartSpan(ctx, "CFEventStore.updateOccurrences")
	defer span.End()
	log := zenkit.ContextLogger(ctx)
	writeOps := make([]WriteOp, 0)
	occursColRef := docRef.Collection(occurrencesCollectionName)
	if occursColRef == nil {
		return nil, errors.New("invalid doc")
	}
	activeOccDocRef := occursColRef.Doc(activeOccurrence.ID)
	if activeOccDocRef == nil {
		return nil, errors.New("invalid doc")
	}

	if activeOccurrence.Status != StatusClosed {
		if event.Status != eventProto.Status_STATUS_CLOSED {
			if event.Timestamp < activeOccurrence.StartTime {
				writeOps = append(writeOps, &opUpdate{
					dr: activeOccDocRef,
					data: []firestore.Update{
						{Path: "startTime", Value: firestore.FieldTransformMinimum(event.Timestamp)},
					},
					msg: "update occurrence: startTime",
				})
			}
		} else {
			writeOps = append(writeOps, &opUpdate{
				dr: activeOccDocRef,
				data: []firestore.Update{
					{Path: "status", Value: StatusClosed},
					{Path: "endTime", Value: firestore.FieldTransformMaximum(event.Timestamp)},
				},
				msg: "update occurrence: status, endTime",
			})
		}
		if event.Timestamp > activeOccurrence.CurrentTime {
			writeOps = append(writeOps, &opUpdate{
				dr:   activeOccDocRef,
				data: OccurrenceFirestoreUpdatesFromProto(event, true),
				msg:  "update occurrence",
			})
		}
	} else {
		if event.Status != eventProto.Status_STATUS_CLOSED {
			if event.Timestamp < activeOccurrence.StartTime {
				writeOps = append(writeOps, &opUpdate{
					dr: activeOccDocRef,
					data: []firestore.Update{
						{Path: "startTime", Value: firestore.FieldTransformMinimum(event.Timestamp)},
					},
					msg: "update occurrence: startTime",
				})
			}
			if event.Timestamp > activeOccurrence.EndTime {
				ops, err := createNewOccurrence(ctx, tx, event, docRef, currEvent, occursColRef)
				if err != nil {
					log.WithError(err).Error("failed to create new occurrence")
					return nil, errors.Wrap(err, "failed to create new occurrence")
				}
				writeOps = append(writeOps, ops...)
			} else if event.Timestamp > activeOccurrence.CurrentTime {
				writeOps = append(writeOps, &opUpdate{
					dr:   activeOccDocRef,
					data: OccurrenceFirestoreUpdatesFromProto(event, true),
					msg:  "update occurrence",
				})
			}
		} else {
			if event.Timestamp > activeOccurrence.EndTime {
				ops, err := createNewOccurrence(ctx, tx, event, docRef, currEvent, occursColRef)
				if err != nil {
					log.WithError(err).Error("failed to create new occurrence")
					return nil, errors.Wrap(err, "failed to create new occurrence")
				}
				writeOps = append(writeOps, ops...)
			}
			if event.Timestamp < activeOccurrence.EndTime {
				ops, err := splitOccurrences(ctx, tx, activeOccurrence, event, docRef, occursColRef)
				if err != nil {
					log.WithError(err).Error("failed to update occurrences: split occurrences")
					return nil, errors.Wrap(err, "failed to split occurrence")
				}
				writeOps = append(writeOps, ops...)
			}
		}
	}
	writeOps = append(writeOps, &opUpdate{
		dr: activeOccDocRef,
		data: []firestore.Update{
			{Path: "currentTime", Value: firestore.FieldTransformMaximum(event.Timestamp)},
		},
		msg: "update occurrence: currentTime",
	})
	return writeOps, nil
}

func (eventsStore *CFEventStore) createOrUpdateActiveEvent(ctx context.Context, tx Transaction, timestamp int64, eventID, entityID, tenant string) ([]WriteOp, error) {
	ctx, span := trace.StartSpan(ctx, "CFEventStore.createOrUpdateActiveEvent")
	defer span.End()
	writeOps := make([]WriteOp, 0)
	tsAsTm := time.Unix(timestamp/1e3, timestamp%1e3*1e6)
	activityDate := fmt.Sprintf("%d%02d%02d", tsAsTm.Year(), tsAsTm.Month(), tsAsTm.Day())
	activeEventDocPath := fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s/%s", tenantsCollectionName, tenant, entitiesCollectionName, entityID, activityDatesCollectionName, activityDate, activeEventsCollectionName, eventID)
	activeEventDocRef := eventsStore.Client.Doc(activeEventDocPath)
	allEventDocPath := fmt.Sprintf("%s/%s/%s/%s/%s/%s", tenantsCollectionName, tenant, allEventsByDateCollectionName, activityDate, allActiveEventsCollectionName, eventID)
	allEventDocRef := eventsStore.Client.Doc(allEventDocPath)

	writeOps = append(writeOps, &opSet{
		dr: activeEventDocRef,
		data: &ActiveEvent{
			ID:        eventID,
			Entity:    entityID,
			Tenant:    tenant,
			Timestamp: timestamp,
		},
		msg: fmt.Sprintf("set active event: %s", activeEventDocRef.Path),
	})

	// write event in the second path tenant>>active events by date
	writeOps = append(writeOps, &opSet{
		dr: allEventDocRef,
		data: &ActiveEvent{
			ID:        eventID,
			Entity:    entityID,
			Tenant:    tenant,
			Timestamp: timestamp,
		},
		msg: fmt.Sprintf("set active event: %s", activeEventDocRef.Path),
	})
	return writeOps, nil
}

// CreateOrUpdateEventTimeseries ...
func (eventsStore *CFEventStore) CreateOrUpdateEventTimeseries(ctx context.Context, event *eventProto.Event) error {
	ctx, span := trace.StartSpan(ctx, "CFEventStore.CreateOrUpdateEventTimeseries")
	defer span.End()
	shard := eventsStore.shardMap.GetShard(event.Id)
	shard.Lock()
	defer shard.Unlock()
	log := zenkit.ContextLogger(ctx).
		WithFields(logrus.Fields{
			"eventId":   event.Id,
			"timestamp": event.Timestamp,
		})
	eventDocPath := fmt.Sprintf("%s/%s/%s/%s", tenantsCollectionName, event.Tenant, eventsCollectionName, event.Id)
	eventDocRef := eventsStore.Client.Doc(eventDocPath)
	if eventDocRef == nil {
		log.WithField("docPath", eventDocPath).Error("invalid document path")
		return errors.New("invalid document path")
	}
	writeOps := make([]WriteOp, 0)
	var (
		currEventDoc     *Event
		activeOccurrence *Occurrence
		isNew            bool
		err              error
	)
	err = eventsStore.Client.RunTransaction(ctx, func(txCtx context.Context, tx *firestore.Transaction) error {
		log := zenkit.ContextLogger(txCtx)
		var ops []WriteOp
		currEventDoc, activeOccurrence, err, isNew, ops = eventsStore.createOrGetEventOccurrence(txCtx, tx, eventDocRef, event)
		if err != nil {
			log.WithError(err).Error("failed to put event")
			return errors.Wrap(err, "failed to put event: createOrGetEventOccurrences")
		}
		for _, op := range ops {
			err = DoWriteWithTransaction(txCtx, tx, op)
			if err != nil {
				log.WithError(err).Errorf("failed to do write operation: %s", op.errMsg())
				return errors.Wrap(err, fmt.Sprintf("failed operation: %s", op.errMsg()))
			}
		}
		return nil
	})
	if err != nil {
		log.WithError(err).Error("failed to create or update event")
		retErr := status.Error(status.Code(err), "failed to create or update event")
		if stat, ok := status.FromError(err); ok {
			retErr = status.Error(stat.Code(), stat.Message())
		}
		return retErr
	}
	err = eventsStore.Client.RunTransaction(ctx, func(txCtx context.Context, tx *firestore.Transaction) error {
		var ops []WriteOp
		if !isNew {
			ops, err = eventsStore.updateOccurrences(txCtx, tx, eventDocRef, event, currEventDoc, activeOccurrence)
			if err != nil {
				log.WithError(err).Error("failed to put event")
				return errors.Wrap(err, "failed to put event: updateOccurrences")
			}
			writeOps = append(writeOps, ops...)
		}
		if currEventDoc.Entity != "" {
			ops, err := eventsStore.createOrUpdateActiveEvent(txCtx, tx, event.Timestamp, currEventDoc.ID, currEventDoc.Entity, currEventDoc.Tenant)
			if err != nil {
				log.WithError(err).Error("failed to put event")
				return errors.Wrap(err, "failed to put event: createOrUpdateActiveEvent")
			}
			writeOps = append(writeOps, ops...)
		}
		for _, op := range writeOps {
			err = DoWriteWithTransaction(txCtx, tx, op)
			if err != nil {
				log.WithError(err).Errorf("failed to do write operation: %s", op.errMsg())
				return errors.Wrap(err, fmt.Sprintf("failed operation: %s", op.errMsg()))
			}
		}
		return nil
	})
	if err != nil {
		log.WithError(err).Error("failed to create or update event")
		retErr := status.Error(status.Code(err), "failed to create or update event")
		if stat, ok := status.FromError(err); ok {
			retErr = status.Error(stat.Code(), stat.Message())
		}
		return retErr
	}
	return nil
}

// GetBulkEvents retrieves an array of events from the Events collection in Cloud Firestore
func (eventsStore *CFEventStore) GetBulkEvents(ctx context.Context, tenant string, IDs []string, startTime int64, endTime int64, fields string) ([]*Event, error) {
	var query = &Query{Tenant: tenant, TimeRange: TimeRange{Start: startTime, End: endTime}}
	log := zenkit.ContextLogger(ctx)
	if len(tenant) == 0 {
		return nil, status.Error(codes.InvalidArgument, "empty tenant")
	}
	result := make([]*Event, 0, len(IDs))

	isNotesNeeded := false
	if len(fields) == 0 || strings.Contains(fields, "Notes") {
		isNotesNeeded = true
	}
	// setup metrics
	mTime := time.Now()
	defer func() {
		stats.Record(ctx, metrics.MGetBulkEventsTimeMs.M(sinceInMilliseconds(mTime)),
			metrics.MGetBulkEventsCount.M(int64(len(result))))
	}()

	path := fmt.Sprintf("%s/%s/%s", tenantsCollectionName, tenant, eventsCollectionName)
	docRefs := make([]*firestore.DocumentRef, len(IDs))
	for i, id := range IDs {
		fullPath := fmt.Sprintf("%s/%s", path, id)
		docRefs[i] = eventsStore.Client.Doc(fullPath)
	}

	sources := utils.GetRestrictedSources(ctx)
	if len(sources) > 1 {
		dimensionFilters := make([]*Filter, 0)
		dimensionFilters = append(dimensionFilters, &Filter{
			Field: "source",
			Op:    "in",
			Value: sources,
		})
		query.DimensionFilters = dimensionFilters
	} else if len(sources) == 1 {
		dimensionFilters := make([]*Filter, 0)
		dimensionFilters = append(dimensionFilters, &Filter{
			Field: "source",
			Op:    "==",
			Value: sources[0],
		})
		query.DimensionFilters = dimensionFilters
	} else {
		query = nil
	}
	//process it
	err := eventsStore.Client.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		docs, err := tx.GetAll(docRefs)
		if err == nil {
			for _, doc := range docs {
				evt, err := hydrateEventDocument(ctx, isNotesNeeded, startTime, endTime, doc, query)
				if err == nil {
					result = append(result, evt)
				} else if err == ErrRestrictedEvent {
					log.WithError(err).Debug("attempted to get restricted event")
				} else {
					log.WithError(err).Error("Failed hydrating in get bulk. aborting...")
					return err
				}
			}
		} else {
			log.WithError(err).Error("Failed getting all docs in get bulk. aborting...")
			return err
		}
		return nil // always for function
	})
	if err != nil {
		log.WithError(err).Error("Failed transaction in get bulk. aborting...")
		return nil, err
	}
	return result, nil
}

// GetActiveEvents retrieves an array of active event ids for the passed in days for the tenant
func (eventsStore *CFEventStore) GetActiveEvents(ctx context.Context, tenant string, entityIDs []string, days []string, pi *PageInput) (*ActiveEvents, error) {
	var (
		limit        = 0
		eventID      = ""
		piCursor     = new(ActiveEventsCursor)
		entityIndex  = 0
		dayIndex     = 0
		dir          = firestore.Asc
		hasNext      = false
		hasPrev      = false
		lastDay      = ""
		lastEntityID = ""
		lastEventID  = ""
		getAll       = false
		allCollRef   *firestore.CollectionRef
		err          error
	)
	log := zenkit.ContextLogger(ctx)
	// grab the tenant

	if entityIDs[0] == "GET_ALL_EVENTS" {
		getAll = true
	}
	// setup metrics
	mTime := time.Now()
	if getAll {
		defer func() {
			stats.Record(ctx, metrics.MGetAllActiveEventsTimeMs.M(sinceInMilliseconds(mTime)))
		}()
	} else {
		defer func() {
			stats.Record(ctx, metrics.MGetActiveEventsTimeMs.M(sinceInMilliseconds(mTime)))
		}()
	}

	iters := make([]*firestore.DocumentIterator, 0)
	defer func() {
		for _, iter := range iters {
			iter.Stop()
		}
	}()

	if pi != nil {
		if pi.Cursor != "" {
			piCursor, err = getValidCursor(pi.Cursor)
			if err != nil {
				log.WithError(err).Error("unable to validate cursor")
				return nil, err
			}
			eventID = piCursor.EventID
			entityIndex = stringSliceIndex(piCursor.EntityID, entityIDs)
			if entityIndex == -1 {
				entityIndex = 0
			}
			dayIndex = stringSliceIndex(piCursor.Day, days)
			if dayIndex == -1 {
				dayIndex = 0
			}
		}
		if pi.Limit != 0 {
			limit = pi.Limit
		}
		if pi.Direction == 1 {
			dir = firestore.Desc
		}
	}

	aevCursor := new(ActiveEventsCursor)
	aevs := make([]ActiveEvent, 0)
	daySlice := days
	if dir == firestore.Desc {
		reverseSlice(daySlice)
	}
	if getAll {
		allPath := fmt.Sprintf("%s/%s/%s", tenantsCollectionName, tenant, allEventsByDateCollectionName)
		allCollRef = eventsStore.Client.Collection(allPath)
	}
entityLoop:
	for _, eID := range entityIDs[entityIndex:] {
		for _, day := range daySlice[dayIndex:] {
			dayIndex = 0
			log.Debug("finding event ids for: " + day)
			var query firestore.Query
			if getAll {
				docRef := allCollRef.Doc(day)
				if limit > 0 {
					query = docRef.Collection(allActiveEventsCollectionName).Select().Limit(limit).OrderBy("id", dir)
					if eventID != "" {
						query = query.StartAfter(eventID)
						hasPrev = true
					}
				} else {
					query = docRef.Collection(allActiveEventsCollectionName).Select()
				}
			} else {
				path := fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s", tenantsCollectionName, tenant, entitiesCollectionName, eID, activityDatesCollectionName, day, activeEventsCollectionName)
				query = eventsStore.Client.Collection(path).Select()
				if limit > 0 {
					query = query.Limit(limit + 1)
					if eventID != "" {
						query = query.OrderBy("id", dir).StartAfter(eventID)
						eventID = ""
					}
				}
			}
			iter := query.Documents(ctx)
			iters = append(iters, iter)
		iterLoop:
			for {
				doc, err := iter.Next()
				if err == iterator.Done {
					lastDay = day
					lastEntityID = eID
					if len(aevs) > 0 {
						lastEventID = aevs[len(aevs)-1].ID
					}
					break iterLoop
				}
				if err != nil {
					log.WithError(err).Error("error iterating thru documents aborting...")
					return nil, err
				}
				if limit > 0 && len(aevs) == limit {
					hasNext = true
					lastDay = day
					lastEntityID = eID
					lastEventID = aevs[len(aevs)-1].ID
					break entityLoop
				}
				aevs = append(aevs, ActiveEvent{
					ID:     doc.Ref.ID,
					Entity: eID,
					Tenant: tenant,
					// Timestamp
				})
			}
		}
	}
	// save page info
	aevCursor.Day = lastDay
	aevCursor.EventID = lastEventID
	aevCursor.EntityID = lastEntityID

	cursor, cursErr := json.Marshal(aevCursor)
	if cursErr != nil {
		log.WithError(cursErr).Error("error creating cursor while getting all active events")
		cursor, _ = json.Marshal("")
	}
	results := new(ActiveEvents)

	results.ActiveEvents = aevs
	results.PageInfo = PageInfo{
		Cursor:  Encoding.EncodeToString([]byte(cursor)),
		Count:   uint64(len(aevs)),
		HasPrev: hasPrev,
		HasNext: hasNext,
	}
	return results, nil
}

func hydrateEventDocument(ctx context.Context, includeNotes bool, startTime int64, endTime int64, eventSnapshot *firestore.DocumentSnapshot, eventQuery *Query) (*Event, error) {
	ctx, span := trace.StartSpan(ctx, "hydrateEventDocument")
	defer span.End()
	log := zenkit.ContextLogger(ctx)
	var result Event
	err := eventSnapshot.DataTo(&result)
	if err != nil {
		log.WithError(err).Errorf("failed to unmarshal event(%s)", eventSnapshot.Ref.Path)
		return nil, err
	}
	if eventQuery != nil {
		dimensions := StringIfaceMap(result.Dimensions)
		if !dimensions.ApplyFilters(ctx, eventQuery.DimensionFilters) {
			return nil, ErrRestrictedEvent
		}
	}
	occurrenceCollRef := eventSnapshot.Ref.Collection(occurrencesCollectionName)
	var occurrenceSnaps []*firestore.DocumentSnapshot
	if startTime >= 0 && endTime > 0 {
		occurrenceSnaps, err = occurrenceCollRef.Where("startTime", "<=", endTime).Documents(ctx).GetAll()
	} else {
		occurrenceSnaps, err = occurrenceCollRef.Documents(ctx).GetAll()
	}

	if err != nil {
		err = errors.Wrap(err, "failed to get event occurrences")
		log.WithError(err).Error("failed to get event occurrences")
		return nil, err
	}
	if len(occurrenceSnaps) > 0 {
		result.Occurrences = make([]*Occurrence, 0, len(occurrenceSnaps))
		for _, occ := range occurrenceSnaps {
			if occ.Exists() {
				if eventQuery != nil {
					data := StringIfaceMap(occ.Data())
					lteFilters := []*Filter{
						&Filter{
							Field: "startTime",
							Op:    FilterOpLessThanOrEqualTo,
							Value: eventQuery.TimeRange.Start,
						},
						&Filter{
							Field: "startTime",
							Op:    FilterOpLessThanOrEqualTo,
							Value: eventQuery.TimeRange.End,
						},
						&Filter{
							Field: "endTime",
							Op:    FilterOpLessThanOrEqualTo,
							Value: eventQuery.TimeRange.Start,
						},
						&Filter{
							Field: "endTime",
							Op:    FilterOpLessThanOrEqualTo,
							Value: eventQuery.TimeRange.End,
						},
					}
					// If the occurrence interval does not intersect with the query interval, in other words,
					// !(timeRange.Start <= occ.startTime || occ.startTime <= timeRange.End
					//    || timeRange.Start <= occ.endTime || occ.endTime << timeRange.End)
					// then omit it
					if !lteFilters[0].Apply(ctx, data["startTime"]) &&
						!lteFilters[1].Apply(ctx, data["startTime"]) &&
						!lteFilters[2].Apply(ctx, data["endTime"]) &&
						!lteFilters[3].Apply(ctx, data["endTime"]) {
						continue
					}
					if !data.ApplyFilters(ctx, eventQuery.OccurrenceFilters) {
						continue
					}
				}
				o := &Occurrence{}
				err := occ.DataTo(o)
				if err != nil {
					err = errors.Wrap(err, "failed to get event")
					log.WithError(err).Errorf("failed to unmarshal event occurrence(%s)", occ.Ref.Path)
					return nil, err
				}
				// if we are not looking at endtime or (o has endtime and is between start and end time) then add
				if endTime <= 0 || (o.EndTime > 0 && o.EndTime > startTime && o.EndTime < endTime) {
					if includeNotes {
						noteRefs, err := occ.Ref.Collection(notesCollectionName).Documents(ctx).GetAll()
						if err != nil {
							err = errors.Wrap(err, "failed to get event")
							log.WithError(err).Errorf("failed to to get notes on occurrence(%s)", occ.Ref.Path)
							return nil, err
						}
						if len(noteRefs) > 0 {
							o.Notes = make([]*Note, 0, len(noteRefs))
							for _, noteRef := range noteRefs {
								if noteRef.Exists() {
									n := &Note{}
									err := noteRef.DataTo(n)
									if err != nil {
										err = errors.Wrap(err, "failed to get event")
										log.WithError(err).Errorf("failed to unmarshal note (%s)", noteRef.Ref.Path)
										return nil, err
									}
									o.Notes = append(o.Notes, n)
								}
							}
						}
					}
					if o.ID == result.LastActiveOccurrenceID {
						result.LastActiveOccurrence = o
					}
					result.Occurrences = append(result.Occurrences, o)
				}
			}
		}
	}
	return &result, nil
}
