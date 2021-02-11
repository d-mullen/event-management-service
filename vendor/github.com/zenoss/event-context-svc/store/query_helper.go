package store

import (
	"context"
	"fmt"

	"github.com/zenoss/event-context-svc/utils"
	"github.com/zenoss/zingo/v4/orderedbytes"

	"cloud.google.com/go/firestore"
	"github.com/zenoss/zenkit/v5"
	proto "github.com/zenoss/zing-proto/v11/go/cloud/event_context"
	"github.com/zenoss/zingo/v4/protobufutils"
	"go.opencensus.io/trace"

	"github.com/pkg/errors"
)

type FilterOp string

// Filter operations
const (
	FilterOpLessThan             = "<"
	FilterOpLessThanOrEqualTo    = "<="
	FilterOpGreaterThan          = ">"
	FilterOpGreaterThanOrEqualTo = ">="
	FilterOpEqualTo              = "=="
	FilterOpNotEqualTo           = "!="
	FilterOpIn                   = "in"
	FilterOpNotIn                = "not-in"
)

type Filter struct {
	Field string
	Op    FilterOp
	Value interface{}
}

// Apply applies the filter operation f against field value a using the filter value b
func (filter Filter) Apply(ctx context.Context, value interface{}) bool {
	log := zenkit.ContextLogger(ctx)
	valueBytes, err := orderedbytes.Encode(value, orderedbytes.Ascending)
	if err != nil {
		log.WithError(err).Error("failed to encode value")
		return false
	}
	var filterBytes []byte
	if filter.Op != FilterOpIn && filter.Op != FilterOpNotIn {
		filterBytes, err = orderedbytes.Encode(filter.Value, orderedbytes.Ascending)
		if err != nil {
			log.WithError(err).Error("failed to encode filter value")
			return false
		}
	}
	switch filter.Op {
	case FilterOpLessThan:
		return string(valueBytes) < string(filterBytes)
	case FilterOpLessThanOrEqualTo:
		return string(valueBytes) <= string(filterBytes)
	case FilterOpGreaterThan:
		return string(valueBytes) > string(filterBytes)
	case FilterOpGreaterThanOrEqualTo:
		return string(valueBytes) <= string(filterBytes)
	case FilterOpEqualTo:
		return string(valueBytes) == string(filterBytes)
	case FilterOpNotEqualTo:
		return !(string(valueBytes) == string(filterBytes))
	case FilterOpIn:
		if filterAsSlice, ok := filter.Value.([]interface{}); ok {
			for _, elem := range filterAsSlice {
				filterBytes, err := orderedbytes.Encode(elem, orderedbytes.Ascending)
				if err != nil {
					log.WithError(err).Error("failed to encode slice element")
					return false
				}
				if string(filterBytes) == string(valueBytes) {
					return true
				}
			}
			return false
		}
	case FilterOpNotIn:
		notin := true
		if filterAsSlice, ok := filter.Value.([]interface{}); ok {
			for _, elem := range filterAsSlice {
				filterBytes, err := orderedbytes.Encode(elem, orderedbytes.Ascending)
				if err != nil {
					log.WithError(err).Error("failed to encode slice element")
					return false
				}
				if string(filterBytes) == string(valueBytes) {
					notin = false
				}
			}
		}
		return notin
	}
	return false
}

type TimeRange struct {
	Start int64
	End   int64
}

func (t *TimeRange) StartTS() int64 {
	return t.Start
}

func (t *TimeRange) EndTS() int64 {
	return t.End
}

// Query used by search
type Query struct {
	Tenant            string
	TimeRange         TimeRange
	Severities        []Severity
	Statuses          []Status
	EventFilters      []*Filter
	DimensionFilters  []*Filter
	OccurrenceFilters []*Filter
	PageInput         *PageInput
}

var (
	eventFieldMap = map[string]bool{
		"entity": true,
		"tenant": true,
	}
	occurrenceFieldMap = map[string]bool{
		"tenant":       true,
		"type":         true,
		"acknowledged": true,
		"summary":      true,
		"body":         true,
	}
)

// QueryFromProto creates a store.Query from incoming search request
func QueryFromProto(ctx context.Context, tenant string, req *proto.ECSearchRequest) (*Query, error) {
	log := zenkit.ContextLogger(ctx)
	var pageInput *PageInput
	if len(req.Tenants) == 0 {
		log.Error("invalid argument: missing tenant")
		return nil, ErrInvalidArgument
	}
	severities := make([]Severity, len(req.Severities))
	for i, s := range req.Severities {
		severities[i] = SeverityFromProto(s)
	}
	statuses := make([]Status, len(req.Statuses))
	for i, s := range req.Statuses {
		statuses[i] = StatusFromProto(s)
	}
	eventFilters := make([]*Filter, 0)
	dimensionFilters := make([]*Filter, 0)
	occurrenceFilters := make([]*Filter, 0)
	for _, filterPb := range req.Filters {
		f, err := FilterFromProto(filterPb)
		if err != nil {
			log.WithError(err).Error("failed to convert filter")
			return nil, err
		}
		if eventFieldMap[f.Field] {
			eventFilters = append(eventFilters, f)
		}
		if occurrenceFieldMap[f.Field] {
			occurrenceFilters = append(occurrenceFilters, f)
		}
		if !eventFieldMap[f.Field] && !occurrenceFieldMap[f.Field] {
			dimensionFilters = append(dimensionFilters, f)
		}
	}
	// add source restriction clause, if any
	sources := utils.GetRestrictedSources(ctx)
	if len(sources) > 1 {
		dimensionFilters = append(dimensionFilters, &Filter{
			Field: "source",
			Op:    "in",
			Value: sources,
		})
	} else if len(sources) == 1 {
		dimensionFilters = append(dimensionFilters, &Filter{
			Field: "source",
			Op:    "==",
			Value: sources[0],
		})
	}
	if pi := req.GetPageInput(); pi != nil {
		pageInput = &PageInput{
			Cursor:          pi.GetCursor(),
			Direction:       int(pi.GetDirection()),
			Limit:           int(pi.GetLimit()),
			OccurrenceLimit: int(pi.GetOccurrenceLimit()),
		}
	}

	return &Query{
		Tenant:            tenant,
		TimeRange:         TimeRange{Start: req.TimeRange.Start, End: req.TimeRange.End},
		Severities:        severities,
		Statuses:          statuses,
		EventFilters:      eventFilters,
		DimensionFilters:  dimensionFilters,
		OccurrenceFilters: occurrenceFilters,
		PageInput:         pageInput,
	}, nil
}

func validateQuery(filters []*Filter, isInUsed bool, isRangeUsed bool) error {
	// not-in and !=
	// range(<,> =)  and !=

	// one of in , not-in
	// range on only one field

	// in, not-in upto 10 val
	//
	isInThere := false
	isInThereField := ""
	isNotInThere := false
	isNotInThereField := ""
	isNotEqual := false
	isNotEqualField := ""
	rangeField := ""
	valCount := 0
	for _, f := range filters {
		switch f.Op {
		case FilterOpIn:
			{
				if isInUsed {
					return errors.New("In filter cannot be used by " + f.Field)
				}
				isInThere = true
				isInThereField = f.Field
				arr, ok := f.Value.([]interface{})
				if !ok {
					return errors.New("In filter has invalid values for " + isInThereField)
				}
				valCount = len(arr)
			}
		case FilterOpNotIn:
			{
				isNotInThere = true
				isNotInThereField = f.Field
				arr, ok := f.Value.([]interface{})
				if !ok {
					return errors.New("NotIn filter has invalid values for " + isNotInThereField)
				}
				valCount = len(arr)
			}
		case FilterOpNotEqualTo:
			{
				isNotEqual = true
				isNotEqualField = f.Field
			}
		case FilterOpLessThanOrEqualTo, FilterOpLessThan, FilterOpGreaterThanOrEqualTo, FilterOpGreaterThan:
			{
				if rangeField == "" {
					rangeField = f.Field
				}
				if rangeField != f.Field {
					return errors.New("Range(<,>,<=,>=) filters should be on one field only." + rangeField + "," + f.Field)
				}
			}
		}

		if (f.Op == FilterOpIn || f.Op == FilterOpNotIn) && valCount > 10 {
			return errors.New("Cannot have more than 10 values to filter on." + f.Field)
		}

		if isNotInThere && isInThere {
			return errors.New("Cannot have not-in and in filter in a query." + isNotInThereField + "," + isInThereField)
		}

		if isNotInThere && isNotEqual {
			return errors.New("Cannot have not-in and != filter in a query." + isNotInThereField + "," + isNotEqualField)
		}
		if isNotEqualField != "" && rangeField != "" && isNotEqualField != rangeField {
			return errors.New("Range(<,>,<=,>=) filters and != should be on one field only." + rangeField + "," + f.Field)
		}
	}
	// all is fine
	return nil
}

// FilterFromProto creates a Filter from incoming proto filter
func FilterFromProto(filterPb *proto.Filter) (*Filter, error) {
	field := filterPb.Field
	if field == "" {
		return nil, ErrInvalidArgument
	}
	switch op := filterPb.Op.(type) {
	case *proto.Filter_In_:
		arr, err := protobufutils.FromScalarArray(op.In.Values)
		if err != nil {
			return nil, err
		}
		if len(arr) == 1 {
			return &Filter{
				Field: field,
				Op:    "==",
				Value: arr[0],
			}, nil
		}
		return &Filter{
			Field: filterPb.Field,
			Op:    "in",
			Value: arr,
		}, nil
	case *proto.Filter_LessThan_:
		v, err := protobufutils.FromScalar(op.LessThan.Value)
		if err != nil {
			return nil, err
		}
		return &Filter{
			Field: field,
			Op:    "<",
			Value: v,
		}, nil
	case *proto.Filter_LessThanOrEquals_:
		v, err := protobufutils.FromScalar(op.LessThanOrEquals.Value)
		if err != nil {
			return nil, err
		}
		return &Filter{
			Field: field,
			Op:    "<=",
			Value: v,
		}, nil
	case *proto.Filter_GreaterThan_:
		v, err := protobufutils.FromScalar(op.GreaterThan.Value)
		if err != nil {
			return nil, err
		}
		return &Filter{
			Field: field,
			Op:    ">",
			Value: v,
		}, nil
	case *proto.Filter_GreaterThanOrEquals_:
		v, err := protobufutils.FromScalar(op.GreaterThanOrEquals.Value)
		if err != nil {
			return nil, err
		}
		return &Filter{
			Field: field,
			Op:    ">=",
			Value: v,
		}, nil
	case *proto.Filter_Equals_:
		v, err := protobufutils.FromScalar(op.Equals.Value)
		if err != nil {
			return nil, err
		}
		return &Filter{
			Field: field,
			Op:    "==",
			Value: v,
		}, nil
	case *proto.Filter_NotEquals_:
		v, err := protobufutils.FromScalar(op.NotEquals.Value)
		if err != nil {
			return nil, err
		}
		return &Filter{
			Field: field,
			Op:    "!=",
			Value: v,
		}, nil
	case *proto.Filter_NotIn_:
		arr, err := protobufutils.FromScalarArray(op.NotIn.Values)
		if err != nil {
			return nil, err
		}
		if len(arr) == 1 {
			return &Filter{
				Field: field,
				Op:    "!=",
				Value: arr[0],
			}, nil
		}
		return &Filter{
			Field: field,
			Op:    "not-in",
			Value: arr,
		}, nil
	default:
		return nil, ErrInvalidArgument
	}
}

type QueryPlan struct {
	EventsQ      *firestore.Query
	OccurrencesQ *firestore.Query
}

//QueryCursor contains cursor information for event search queries
type QueryCursor struct {
	// id of event
	EventID string
	// last start time of occurrence
	StartTime int64
	// are there more occurrences in the series?
	HasNextOccurrence bool
}

// ResolveOccurrencesCollectionGroupQuery use Occurrences collection group to find all events with given filters
func ResolveOccurrencesCollectionGroupQuery(ctx context.Context, client Client, query *Query) ([]*firestore.DocumentSnapshot, error) {
	ctx, span := trace.StartSpan(ctx, "ResolveOccurrencesCollectionGroupQuery")
	defer span.End()
	log := zenkit.ContextLogger(ctx)
	log.WithField("query", query).Debug("ResolveOccurrenceCollectionGroupQuery")
	occColl := client.CollectionGroup(fmt.Sprintf(occurrencesCollectionName))
	q := occColl.Where("tenant", "==", query.Tenant).
		Where("startTime", "<=", query.TimeRange.End)

	isInUsed := false
	if len(query.Severities) > 0 {
		if len(query.Severities) == 1 {
			q = q.Where("severity", "==", query.Severities[0])
		} else {
			isInUsed = true
			q = q.Where("severity", "in", query.Severities)
		}
	}
	if len(query.Statuses) > 0 {
		if len(query.Statuses) == 1 {
			q = q.Where("status", "==", query.Statuses[0])
		} else {
			isInUsed = true
			q = q.Where("status", "in", query.Statuses)
		}
	}

	//validate
	if len(query.Severities) > 1 && len(query.Statuses) > 1 {
		return nil, errors.New("Cannot have two in filters with multiple values, for severity and status in a query.")
	}
	err := validateQuery(query.OccurrenceFilters, isInUsed, true)
	if err != nil {
		return nil, err
	}

	for _, f := range query.OccurrenceFilters {
		q = q.Where(f.Field, string(f.Op), f.Value)
	}
	if query.PageInput != nil && (query.PageInput.Limit > 0 || query.PageInput.OccurrenceLimit > 0) {
		piCursor := new(QueryCursor)
		if query.PageInput.Cursor != "" {
			piCursor = new(QueryCursor)
			err := DecodeCursor(query.PageInput.Cursor, piCursor)
			if err != nil {
				zenkit.ContextLogger(ctx).WithError(err).Error("unable to decode cursor")
			} else if piCursor.StartTime != 0 {
				q = q.StartAfter(piCursor.StartTime)
			}
			// reset cursor
			query.PageInput.Cursor = ""
		}
		dir := firestore.Asc
		if query.PageInput.Direction != int(proto.PageInput_FORWARD) {
			dir = firestore.Desc
		}
		if query.PageInput.OccurrenceLimit != 0 {
			q = q.OrderBy("startTime", dir).Limit(query.PageInput.OccurrenceLimit)
		} else if query.PageInput.Limit != 0 {
			q = q.OrderBy("startTime", dir).Limit(query.PageInput.Limit)
		}
	}

	occSnaps, err := q.Documents(ctx).GetAll()
	if err != nil {
		log.WithError(err).Error("failed to query occurrences collection group")
		return nil, err
	}
	// TODO filter occurrences that ended before start time

	return occSnaps, err
}

type StringIfaceMap map[string]interface{}

func (s StringIfaceMap) ApplyFilters(ctx context.Context, filters []*Filter) bool {
	for _, filter := range filters {
		if ok := filter.Apply(ctx, s[filter.Field]); !ok {
			return false
		}
	}
	return true
}

func ResolveOccurrencesViaActiveEvents(ctx context.Context, client Client, query *Query, store EventContextStore) ([]*firestore.DocumentSnapshot, *PageInfo, error) {
	var (
		err error
	)
	ctx, span := trace.StartSpan(ctx, "ResolveOccurrencesViaActiveEvents")
	defer span.End()
	log := zenkit.ContextLogger(ctx)
	log.WithField("query", query).Debug("ResolveOccurrencesViaActiveEvents")

	days := utils.EnumerateDays(&query.TimeRange)
	eventsStore := store
	if eventsStore == nil {
		eventsStore, err = DefaultEventsStore(ctx)
		if err != nil {
			log.WithError(err).Error("failed to get eventsStore")
			return nil, nil, err
		}
	}
	var activeEventsResponse *ActiveEvents
	activeEventsResponse, err = eventsStore.GetActiveEvents(ctx, query.Tenant, []string{"GET_ALL_EVENTS"}, days, query.PageInput)
	if err != nil {
		log.WithError(err).Error("failed to get active events")
		return nil, nil, err
	}
	log.Debug("got active events refs")
	eventDocRefs := make([]*firestore.DocumentRef, len(activeEventsResponse.ActiveEvents))
	for idx, ae := range activeEventsResponse.ActiveEvents {
		eventDocRefs[idx] = client.Doc(fmt.Sprintf("%s/%s/%s/%s", tenantsCollectionName, query.Tenant, eventsCollectionName, ae.ID))
	}
	var snapshots []*firestore.DocumentSnapshot
	err = client.RunTransaction(ctx, func(txCtx context.Context, tx *firestore.Transaction) error {
		var err error
		snapshots, err = tx.GetAll(eventDocRefs)
		if err != nil {
			log.WithError(err).Error("failed to get event snapshots")
			return err
		}
		return nil
	}, firestore.ReadOnly)
	if err != nil {
		log.WithError(err).Error("failed to get active event snapshots")
		return nil, nil, err
	}
	log.Debug("got active events snaps")
	return snapshots, &activeEventsResponse.PageInfo, nil
}

// ResolveOccurrencesByEventIds given a list of event ids, get occurrences
func ResolveOccurrencesByEventIds(ctx context.Context, client Client, query *Query, eventIDs []string) (map[string][]*firestore.DocumentSnapshot, error) {
	ctx, span := trace.StartSpan(ctx, "ResolveOccurrencesByEventIds")
	defer span.End()
	log := zenkit.ContextLogger(ctx)
	log.WithField("query", query).Debug("ResolveOccurrencesByEventIds")
	results := make(map[string][]*firestore.DocumentSnapshot)
	for _, eventID := range eventIDs {
		// TODO: can be run concurrently
		collref := client.Collection(fmt.Sprintf("%s/%s/%s/%s/%s", tenantsCollectionName, query.Tenant, eventsCollectionName, eventID, occurrencesCollectionName))
		q := collref.
			Where("startTime", "<=", query.TimeRange.End)

		var hasAppliedInOp bool
		if len(query.Severities) > 0 {
			if len(query.Severities) > 1 {
				q = q.Where("severity", "in", query.Severities)
				hasAppliedInOp = true
			} else {
				q = q.Where("severity", "==", query.Severities[0])
			}
		}
		if len(query.Statuses) > 0 {
			if len(query.Statuses) > 1 && !hasAppliedInOp {
				q = q.Where("status", "in", query.Statuses)
			} else {
				q = q.Where("status", "==", query.Statuses[0])
			}
			//TODO: handle case where both severity and status filters have more than one value
		}

		// validate
		//if len(query.Severities) > 1 && len(query.Statuses) > 1 {
		//	return nil, errors.New("Cannot have two in filters with multiple values, for severity and status in a query.")
		//}
		err := validateQuery(query.OccurrenceFilters, hasAppliedInOp, true)
		if err != nil {
			return nil, err
		}

		for _, f := range query.OccurrenceFilters {
			q = q.Where(f.Field, string(f.Op), f.Value)
		}
		if query.PageInput != nil && (query.PageInput.Limit > 0 || query.PageInput.OccurrenceLimit > 0) {
			piCursor := new(QueryCursor)
			if query.PageInput.Cursor != "" {
				piCursor = new(QueryCursor)
				err := DecodeCursor(query.PageInput.Cursor, piCursor)
				if err != nil {
					zenkit.ContextLogger(ctx).WithError(err).Error("unable to decode cursor")
				} else if piCursor.StartTime != 0 {
					q = q.StartAfter(piCursor.StartTime)
				}
				// reset cursor
				query.PageInput.Cursor = ""
			}
			dir := firestore.Asc
			if query.PageInput.Direction != int(proto.PageInput_FORWARD) {
				dir = firestore.Desc
			}
			if query.PageInput.OccurrenceLimit != 0 {
				q = q.OrderBy("startTime", dir).Limit(query.PageInput.OccurrenceLimit)
			}
		}

		occSnaps, err := q.Documents(ctx).GetAll()
		if err != nil {
			log.WithError(err).Error("failed to query occurrence")
			return nil, err
		}

		//TODO filter out occurrences that ended before start time

		results[eventID] = occSnaps
	}
	return results, nil
}

// ResolveEventDocuments finds events matching query
func ResolveEventDocuments(ctx context.Context, query *Query, client Client) ([]*firestore.DocumentSnapshot, error) {
	ctx, span := trace.StartSpan(ctx, "ExecuteQuery")
	defer span.End()
	zenkit.ContextLogger(ctx).WithField("query", query).Debug("running ResolveEventDocuments")

	//validate filters
	var combinedFilters []*Filter
	combinedFilters = append(combinedFilters, query.EventFilters...)
	combinedFilters = append(combinedFilters, query.DimensionFilters...)
	err := validateQuery(combinedFilters, false, false)
	if err != nil {
		return nil, err
	}

	eventsColl := client.Collection(fmt.Sprintf("%s/%s/%s", tenantsCollectionName, query.Tenant, eventsCollectionName))
	if eventsColl == nil {
		return nil, ErrInvalidArgument
	}
	eventLevelQ := eventsColl.Query
	for _, f := range query.EventFilters {
		eventLevelQ = eventLevelQ.Where(f.Field, string(f.Op), f.Value)
	}
	for _, f := range query.DimensionFilters {
		eventLevelQ = eventLevelQ.WherePath(firestore.FieldPath{"dimensions", f.Field}, string(f.Op), f.Value)
	}
	if query.PageInput != nil && query.PageInput.Limit > 0 {
		piCursor := new(QueryCursor)
		if query.PageInput.Cursor != "" {
			piCursor = new(QueryCursor)
			err := DecodeCursor(query.PageInput.Cursor, piCursor)
			if err != nil {
				zenkit.ContextLogger(ctx).WithError(err).Error("unable to decode cursor")
			} else if piCursor.EventID != "" {
				if piCursor.HasNextOccurrence {
					eventLevelQ = eventLevelQ.StartAt(piCursor.EventID)
				} else {
					eventLevelQ = eventLevelQ.StartAfter(piCursor.EventID)
					// reset cursor so we do not try to use an invalid occurrence start time
					query.PageInput.Cursor = ""
				}
			}
		}
		dir := firestore.Asc
		if query.PageInput.Direction != int(proto.PageInput_FORWARD) {
			dir = firestore.Desc
		}
		eventLevelQ = eventLevelQ.OrderBy("id", dir).Limit(query.PageInput.Limit)
	}

	return eventLevelQ.Documents(ctx).GetAll()
}
