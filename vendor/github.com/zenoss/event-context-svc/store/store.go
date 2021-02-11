package store

import (
	"context"
	"encoding/json"
	"time"

	"github.com/zenoss/zenkit/v5"

	"cloud.google.com/go/firestore"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	evtProto "github.com/zenoss/zing-proto/v11/go/event"
)

// Event is a structured record of a notable change in state of a managed resource
type Event struct {
	ID          string                 `json:"id,omitempty" firestore:"id,omitempty"`
	Tenant      string                 `json:"tenant,omitempty" firestore:"tenant,omitempty"`
	Entity      string                 `json:"entity,omitempty" firestore:"entity,omitempty"`
	Dimensions  map[string]interface{} `json:"dimensions,omitempty" firestore:"dimensions,omitempty"`
	Occurrences []*Occurrence          `json:"occurrences,omitempty" firestore:"-"`
	// do not store active occurrence in Firestore, only ID
	LastActiveOccurrenceID string      `json:"lastActiveOccurrenceID,omitempty" firestore:"lastActiveOccurrenceID,omitempty"`
	LastActiveOccurrence   *Occurrence `json:"lastActiveOccurrence,omitempty" firestore:"-"`
	CreatedAt              time.Time   `json:"-" firestore:"createdAt,omitempty"`
	UpdatedAt              time.Time   `json:"-" firestore:"updatedAt,serverTimestamp"`
	OccurrenceCount        int64       `json:"occurrenceCount,omitempty" firestore:"occurrenceCount,omitempty"`
}

type Status int

const (
	StatusDefault Status = iota
	StatusOpen
	StatusSuppressed
	StatusClosed
)

// StatusFromProto takes a github.com/zenoss/zing-proto/go/event.Status and returns a Status value.
func StatusFromProto(s evtProto.Status) Status {
	switch s {
	case evtProto.Status_STATUS_OPEN:
		return StatusOpen
	case evtProto.Status_STATUS_SUPPRESSED:
		return StatusSuppressed
	case evtProto.Status_STATUS_CLOSED:
		return StatusClosed
	default:
		return StatusDefault
	}
}

type Severity int

const (
	SeverityDefault Severity = iota
	SeverityDebug
	SeverityInfo
	SeverityWarning
	SeverityError
	SeverityCritical
)

func SeverityFromProto(s evtProto.Severity) Severity {
	switch s {
	case evtProto.Severity_SEVERITY_DEBUG:
		return SeverityDebug
	case evtProto.Severity_SEVERITY_INFO:
		return SeverityInfo
	case evtProto.Severity_SEVERITY_WARNING:
		return SeverityWarning
	case evtProto.Severity_SEVERITY_ERROR:
		return SeverityError
	case evtProto.Severity_SEVERITY_CRITICAL:
		return SeverityCritical
	default:
		return SeverityDefault
	}
}

// SeverityToProto from event to proto values
func SeverityToProto(s Severity) evtProto.Severity {
	switch s {
	case SeverityDebug:
		return evtProto.Severity_SEVERITY_DEBUG
	case SeverityInfo:
		return evtProto.Severity_SEVERITY_INFO
	case SeverityWarning:
		return evtProto.Severity_SEVERITY_WARNING
	case SeverityError:
		return evtProto.Severity_SEVERITY_ERROR
	case SeverityCritical:
		return evtProto.Severity_SEVERITY_CRITICAL
	default:
		return evtProto.Severity_SEVERITY_DEFAULT
	}
}

// AcknowledgedFromProto ...
func AcknowledgedFromProto(w *wrappers.BoolValue) *bool {
	if w == nil {
		return nil
	}
	return &w.Value
}

// Occurrence represents a specific episode of the state changes represented by an event
type Occurrence struct {
	ID           string    `json:"id,omitempty" firestore:"id,omitempty"`
	EventID      string    `json:"eventId,omitempty" firestore:"eventId,omitempty"`
	Tenant       string    `json:"tenant,omitempty" firestore:"tenant,omitempty"`
	Summary      string    `json:"summary,omitempty" firestore:"summary,omitempty"`
	Body         string    `json:"body,omitempty" firestore:"body,omitempty"`
	Type         string    `json:"type,omitempty" firestore:"type,omitempty"`
	Status       Status    `json:"status,omitempty" firestore:"status,omitempty"`
	Severity     Severity  `json:"severity,omitempty" firestore:"severity,omitempty"`
	Acknowledged *bool     `json:"acknowledged,omitempty" firestore:"acknowledged,omitempty"`
	StartTime    int64     `json:"startTime,omitempty" firestore:"startTime,omitempty"`
	EndTime      int64     `json:"endTime,omitempty" firestore:"endTime,omitempty"`
	CurrentTime  int64     `json:"currentTime,omitempty" firestore:"currentTime,omitempty"`
	Notes        []*Note   `json:"notes,omitempty" firestore:"-"`
	CreatedAt    time.Time `json:"-" firestore:"createdAt,omitempty"`
	UpdatedAt    time.Time `json:"-" firestore:"updatedAt,serverTimestamp"`
}

// Note represents annotations that can be added to an event occurrence
type Note struct {
	ID        string    `json:"-" firestore:"id,omitempty"`
	Content   string    `json:"content,omitempty" firestore:"content,omitempty"`
	CreatedBy string    `json:"createdBy,omitempty" firestore:"context,omitempty"`
	UpdatedBy string    `json:"updatedBy,omitempty" firestore:"updatedBy,omitempty"`
	CreatedAt time.Time `json:"-" firestore:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"-" firestore:"updatedAt,serverTimestamp"`
}

// UpdateEventPayload ...
type UpdateEventPayload struct {
	ID                     string
	Tenant                 string
	OccurrencesToBeAdded   []*Occurrence
	OccurrencesToBeUpdated []*UpdateOccurrencePayload
	ActiveOccurrence       *Occurrence
}

// UpdateOccurrencePayload ...
type UpdateOccurrencePayload struct {
	ID               string
	Summary          string
	Body             string
	Type             string
	Status           *Status
	Severity         *Severity
	Acknowledged     *bool
	NotesToBeAdded   []*Note
	NotesToBeUpdated []*Note
	StartTime        *int64
	EndTime          *int64
}

type EntityEventKey interface {
	Tenant() string
	EntityID() string
	DateTime() int64
	EventID() string
}

// EventContextStore enables storage and retrieval of metadata on event time-series data
type EventContextStore interface {
	PutEvent(ctx context.Context, tenant string, event *Event) error
	UpdateEvent(ctx context.Context, tenant string, payload *UpdateEventPayload) error
	GetEvent(ctx context.Context, tenant string, ID string) (*Event, error)
	Search(ctx context.Context, query *Query) (*SearchEvents, error)
	CreateOrUpdateEventTimeseries(ctx context.Context, event *evtProto.Event) error
	GetBulkEvents(ctx context.Context, tenant string, IDs []string, startTime int64, endTime int64, fields string) ([]*Event, error)
	GetActiveEvents(ctx context.Context, tenant string, entityIDs []string, days []string, pi *PageInput) (*ActiveEvents, error)
}

// PageInfo contains paging information from queries
type PageInfo struct {
	Cursor  string
	Count   uint64
	HasNext bool
	HasPrev bool
}

// PageInput contains paging input from the request
type PageInput struct {
	Cursor          string
	Direction       int
	Limit           int
	OccurrenceLimit int
}

// DecodeCursor is used to decode the page input cursor
func DecodeCursor(cursor string, piCursor interface{}) error {
	b, err := Encoding.DecodeString(cursor)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, &piCursor)
	if err != nil {
		return err
	}
	return nil
}

// ActiveEvent contains information about an active event
type ActiveEvent struct {
	ID        string `json:"id,omitempty" firestore:"id,omitempty"`
	Tenant    string `json:"tenant,omitempty" firestore:"tenant,omitempty"`
	Entity    string `json:"entity,omitempty" firestore:"entity,omitempty"`
	Timestamp int64  `json:"timestamp,omitempty" firestore:"timestamp,omitempty"`
}

// FilterOccurrences take a Occurrence slice and predicate function and returns a slice of Occurrences for
// which the predicate return a true value
func FilterOccurrences(slice []*Occurrence, f func(*Occurrence) bool) []*Occurrence {
	result := make([]*Occurrence, 0)
	for _, o := range slice {
		if f(o) {
			result = append(result, o)
		}
	}
	return result
}

func DefaultEventsStore(ctx context.Context) (EventContextStore, error) {
	client, err := DefaultFSClient(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get Firestore client")
	}
	return NewCloudFirestoreEventsStore(client), nil
}

func DefaultFSClient(ctx context.Context) (*firestore.Client, error) {
	var firestoreProjMap = map[string]string{
		"zing-dev-197522":     "zing-dev-adjunct",
		"zing-perf":           "zing-perf-adjunct",
		"zing-testing-200615": "zing-testing-adjunct",
		"zing-preview":        "zing-preview-adjunct",
		"zcloud-prod":         "zcloud-prod-adjunct",
		"zcloud-emea":         "zcloud-emea-adjunct",
		"zenoss-zing":         "zenoss-zing",
	}
	gcpProjectID := viper.GetString(zenkit.GCProjectIDConfig)
	if gcpProjectID == "" {
		return nil, errors.New("GCP_PROJECT_ID not set")
	}
	projectID := firestoreProjMap[gcpProjectID]
	if projectID == "" {
		return nil, errors.New("FIRESTORE_PROJECT_ID is not set")
	}
	// TODO: cleanup Firestore client connections
	return firestore.NewClient(ctx, projectID)
}
