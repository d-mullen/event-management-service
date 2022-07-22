package event

import "time"

type Status int

const (
	StatusDefault Status = iota
	StatusOpen
	StatusSuppressed
	StatusClosed
)

type Severity int

const (
	SeverityDefault Severity = iota
	SeverityDebug
	SeverityInfo
	SeverityWarning
	SeverityError
	SeverityCritical
)

type (
	// Event is a structured record of a notable change in state of a managed resource
	Event struct {
		ID                   string         `json:"id,omitempty"`
		Tenant               string         `json:"tenant,omitempty"`
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
		ID            string           `json:"id,omitempty"`
		EventID       string           `json:"eventID,omitempty"`
		Tenant        string           `json:"tenant,omitempty"`
		Summary       string           `json:"summary,omitempty"`
		Body          string           `json:"body,omitempty"`
		Type          string           `json:"type,omitempty"`
		Status        Status           `json:"status,omitempty"`
		Severity      Severity         `json:"severity,omitempty"`
		Acknowledged  *bool            `json:"acknowledged,omitempty"`
		StartTime     int64            `json:"startTime,omitempty"`
		EndTime       int64            `json:"endTime,omitempty"`
		CurrentTime   int64            `json:"currentTime,omitempty"`
		Notes         []*Note          `json:"notes,omitempty"`
		CreatedAt     time.Time        `json:"createdAt,omitempty"`
		UpdatedAt     time.Time        `json:"updatedAt,omitempty"`
		InstanceCount int64            `json:"instanceCount,omitempty"`
		Entity        string           `json:"entity,omitempty"`
		Metadata      map[string][]any `json:"metadata,omitempty"`
	}
	// Note represents annotations that can be added to an event occurrence
	Note struct {
		ID        string    `json:"id,omitempty"`
		OccID     string    `json:"occID,omitempty"`
		Content   string    `json:"content,omitempty"`
		CreatedBy string    `json:"createdBy,omitempty"`
		UpdatedBy string    `json:"updatedBy,omitempty"`
		CreatedAt time.Time `json:"createdAt,omitempty"`
		UpdatedAt time.Time `json:"updatedAt,omitempty"`
	}
)
