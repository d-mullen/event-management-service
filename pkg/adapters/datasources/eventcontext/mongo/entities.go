package mongo

import (
	"time"

	"github.com/zenoss/event-management-service/pkg/models/event"
)

type (
	// Event is a structured record of a notable change in state of a managed resource
	Event struct {
		ID                   string         `json:"id"  bson:"_id"`
		Tenant               string         `json:"tenant"  bson:"tenantId,omitempty"`
		Entity               string         `json:"entity,omitempty"  bson:"entity,omitempty"`
		Dimensions           map[string]any `json:"dimensions,omitempty"  bson:"dimensions,omitempty"`
		Occurrences          []*Occurrence  `json:"occurrences,omitempty"  bson:"-"`
		LastActiveOccurrence *Occurrence    `json:"lastActiveOccurrence,omitempty"  bson:"-"`
		CreatedAt            time.Time      `json:"-"  bson:"createdAt,omitempty"`
		UpdatedAt            time.Time      `json:"-"  bson:"updatedAt,omitempty"`
		OccurrenceCount      uint64         `json:"occurrenceCount,omitempty"  bson:"occurrenceCount,omitempty"`
	}

	EventDimensions struct {
		ID         string         `json:"id"  bson:"_id"`
		Dimensions map[string]any `json:"dimensions,omitempty"  bson:"dimensions,omitempty"`
		Entity     string         `json:"entity,omitempty"  bson:"entity,omitempty"`
	}

	// Occurrence represents a specific episode of the state changes represented by an event
	Occurrence struct {
		ID            string         `json:"id,omitempty"  bson:"_id"`
		EventID       string         `json:"eventID,omitempty"  bson:"eventId"`
		Tenant        string         `json:"tenant,omitempty"  bson:"tenantId"`
		Summary       string         `json:"summary,omitempty"  bson:"summary"`
		Body          string         `json:"body,omitempty"  bson:"body"`
		Type          string         `json:"type,omitempty"  bson:"type"`
		Status        event.Status   `json:"status,omitempty"  bson:"status"`
		Severity      event.Severity `json:"severity,omitempty"  bson:"severity"`
		Acknowledged  *bool          `json:"acknowledged,omitempty"  bson:"acknowledged"`
		StartTime     int64          `json:"startTime,omitempty"  bson:"startTime"`
		EndTime       int64          `json:"endTime,omitempty"  bson:"endTime"`
		CurrentTime   int64          `json:"currentTime,omitempty"  bson:"currentTime"`
		Notes         []*Note        `json:"notes,omitempty"  bson:"notes"`
		CreatedAt     time.Time      `json:"createdAt,omitempty"  bson:"createdAt"`
		UpdatedAt     time.Time      `json:"updatedAt,omitempty"  bson:"updatedAt"`
		InstanceCount int64          `json:"instanceCount,omitempty"  bson:"instanceCount"`
		Entity        string         `json:"entity,omitempty"  bson:"entity"`
		Dimensions    map[string]any `json:"dimensions,omitempty"  bson:"dimensions"`
	}
	// Note represents annotations that can be added to an event occurrence
	Note struct {
		ID        string    `json:"id,omitempty"  bson:"_id"`
		OccID     string    `json:"occID,omitempty"  bson:"occid"`
		Content   string    `json:"content,omitempty"  bson:"content"`
		CreatedBy string    `json:"createdBy,omitempty"  bson:"createdBy"`
		UpdatedBy string    `json:"updatedBy,omitempty"  bson:"updatedBy"`
		CreatedAt time.Time `json:"createdAt,omitempty"  bson:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt,omitempty"  bson:"updatedAt"`
	}
)
