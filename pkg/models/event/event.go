package event

import "time"

type Status int

const (
	StatusDefault Status = iota
	StatusOpen
	StatusSuppressed
	StatusClosed
)

const StatusMax = StatusClosed

type Severity int

const (
	SeverityDefault Severity = iota
	SeverityDebug
	SeverityInfo
	SeverityWarning
	SeverityError
	SeverityCritical
)

const SeverityMax = SeverityCritical

// enum <-> string mappings copied over from the
// golang protobuf bindings to avoid this pkg depending
// on the protobufs
var (
	Severity_name = map[Severity]string{
		SeverityDefault:  "SEVERITY_DEFAULT",
		SeverityDebug:    "SEVERITY_DEBUG",
		SeverityInfo:     "SEVERITY_INFO",
		SeverityWarning:  "SEVERITY_WARNING",
		SeverityError:    "SEVERITY_ERROR",
		SeverityCritical: "SEVERITY_CRITICAL",
	}
	Severity_value = map[string]Severity{
		"SEVERITY_DEFAULT":  SeverityDefault,
		"SEVERITY_DEBUG":    SeverityDebug,
		"SEVERITY_INFO":     SeverityInfo,
		"SEVERITY_WARNING":  SeverityWarning,
		"SEVERITY_ERROR":    SeverityError,
		"SEVERITY_CRITICAL": SeverityCritical,
	}
	Status_name = map[Status]string{
		StatusDefault:    "STATUS_DEFAULT",
		StatusOpen:       "STATUS_OPEN",
		StatusSuppressed: "STATUS_SUPPRESSED",
		StatusClosed:     "STATUS_CLOSED",
	}
	Status_value = map[string]Status{
		"STATUS_DEFAULT":    StatusDefault,
		"STATUS_OPEN":       StatusOpen,
		"STATUS_SUPPRESSED": StatusSuppressed,
		"STATUS_CLOSED":     StatusClosed,
	}
)

type (
	// Event is a structured record of a notable change in state of a managed resource
	Event struct {
		ID                   string         `json:"id,omitempty" bson:"_id"`
		Tenant               string         `json:"tenant" bson:"tenantId"`
		Entity               string         `json:"entity,omitempty"`
		Name                 string         `json:"name,omitempty"`
		Dimensions           map[string]any `json:"dimensions,omitempty"`
		Occurrences          []*Occurrence  `json:"occurrences,omitempty"`
		LastActiveOccurrence *Occurrence    `json:"lastActiveOccurrence,omitempty"`
		CreatedAt            time.Time      `json:"createdAt,omitempty"`
		UpdatedAt            time.Time      `json:"updatedAt,omitempty"`
		OccurrenceCount      uint64         `json:"occurrenceCount,omitempty"`
	}

	// Occurrence represents a specific episode of the state changes represented by an event
	Occurrence struct {
		ID            string           `json:"id,omitempty" bson:"_id"`
		EventID       string           `json:"eventID,omitempty"  bson:"eventId"`
		Tenant        string           `json:"tenant" bson:"tenantId"`
		Summary       string           `json:"summary,omitempty"`
		Body          string           `json:"body,omitempty"`
		Type          string           `json:"type,omitempty"`
		Status        Status           `json:"status,omitempty"`
		Severity      Severity         `json:"severity,omitempty"`
		Acknowledged  *bool            `json:"acknowledged,omitempty"`
		StartTime     int64            `json:"startTime,omitempty"`
		EndTime       int64            `json:"endTime,omitempty"`
		CurrentTime   int64            `json:"currentTime,omitempty"`
		LastSeen      int64            `json:"lastSeen,omitempty"`
		Notes         []*Note          `json:"notes,omitempty"`
		CreatedAt     time.Time        `json:"createdAt,omitempty"`
		UpdatedAt     time.Time        `json:"updatedAt,omitempty"`
		InstanceCount int64            `json:"instanceCount,omitempty"`
		Entity        string           `json:"entity,omitempty"`
		Metadata      map[string][]any `json:"metadata,omitempty"`
	}
	// Note represents annotations that can be added to an event occurrence
	Note struct {
		ID        string    `json:"id,omitempty" bson:"_id"`
		OccID     string    `json:"occID,omitempty"  bson:"occid"`
		Content   string    `json:"content,omitempty"`
		CreatedBy string    `json:"createdBy,omitempty"`
		UpdatedBy string    `json:"updatedBy,omitempty"`
		CreatedAt time.Time `json:"createdAt,omitempty"`
		UpdatedAt time.Time `json:"updatedAt,omitempty"`
	}
)
