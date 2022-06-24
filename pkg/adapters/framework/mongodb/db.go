package mongodb

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/zenoss/event-management-service/pkg/domain/event"
	"github.com/zenoss/zenkit/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.opencensus.io/trace"
	"golang.org/x/exp/slices"
)

const (
	CollEvents      = "event"
	CollOccurrences = "occurrence"
	CollNotes       = "note"
)

type (
	Adapter struct {
		db          *mongo.Database
		collections map[string]*mongo.Collection
		ttlMap      map[string]time.Duration
	}

	Config struct {
		Address    string
		DBName     string
		Username   string
		Password   string
		CACert     string
		ClientCert string
		DefaultTTL time.Duration
	}
)

func (cfg Config) URI() string {
	switch {
	case len(cfg.Username) > 0 && len(cfg.Password) > 0:
		return fmt.Sprintf("mongodb://%s:%s@%s", cfg.Username, cfg.Password, cfg.Address)
	case len(cfg.CACert) > 0 && len(cfg.ClientCert) > 0:
		return fmt.Sprintf(
			"mongodb://%s/?authMechanism=MONGODB-X509&tlsCAFile=%s&tlsCertificateKeyFile=%s",
			cfg.Address, cfg.CACert, cfg.ClientCert)
	default:
		return fmt.Sprintf("mongodb://%s", cfg.Address)
	}
}

var _ event.Repository = &Adapter{}

func NewAdapter(ctx context.Context, cfg Config) (*Adapter, error) {
	log := zenkit.ContextLogger(ctx)
	log.Info("Connecting to mongo db....")
	clientOpts := options.Client().ApplyURI(cfg.URI())
	if len(cfg.CACert) > 0 && len(cfg.ClientCert) > 0 {
		credential := options.Credential{
			AuthMechanism: "MONGODB-X509",
		}
		clientOpts = clientOpts.SetAuth(credential)
	}
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Errorf("failed to connect to MongoDB(%s): %q", cfg.URI(), err)
		return nil, errors.Wrap(err, fmt.Sprintf("failed to get MongoDB client at %v", cfg.URI()))
	}
	log.Debugf("Connected to mongo db: %s", cfg.Address)
	// Ping the primary
	log.Debug("Pinging the primary...")
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Error("failed to connect to MongoDB")
		return nil, err
	}
	db := client.Database(cfg.DBName)

	// since collection names won't change while service is up and running
	// it's reasonable to store the handles
	// and avoid annoying call eventStore.Db.Collection(cname) to get it
	collections := map[string]*mongo.Collection{
		CollEvents:      db.Collection(CollEvents),
		CollOccurrences: db.Collection(CollOccurrences),
		CollNotes:       db.Collection(CollNotes),
	}

	return &Adapter{
		db:          db,
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
	ctx, span := trace.StartSpan(ctx, "mongodb.Adapter.Find")
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
	log.WithFields(logrus.Fields{
		"filters":  filters,
		"findOpts": findOpts,
	}).Debugf("executing %s.Find", CollOccurrences)
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
		notesInFilter := bson.D{{Key: "occid", Value: bson.D{{Key: MongoOpIn, Value: occurrenceIDs}}}}
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
		err := ConcurrentBatcher(ctx, 200, 3, eventIDs,
			func(batch []string) (*mongo.Cursor, error) {
				eventInFilter := bson.D{{Key: MongoOpIn, Value: batch}}
				log.WithFields(logrus.Fields{
					"filters": eventInFilter,
				}).Debugf("executing %s.Find", "event")
				eventCursor, err := db.collections[CollEvents].Find(ctx, bson.D{{Key: "_id", Value: eventInFilter}})
				if err != nil {
					return nil, errors.Wrap(err, "failed to find events")
				}
				return eventCursor, nil
			},
			func(eventCursor *mongo.Cursor) (bool, error) {
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
