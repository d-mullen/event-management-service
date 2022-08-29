//go:build integration

package mongo_test

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	. "github.com/onsi/gomega/gstruct"
	"github.com/onsi/gomega/types"
	"github.com/spf13/viper"
	"github.com/zenoss/event-management-service/config"
	redisCursors "github.com/zenoss/event-management-service/pkg/adapters/datasources/cursors/redis"
	"github.com/zenoss/event-management-service/pkg/adapters/datasources/eventcontext/mongo"
	"github.com/zenoss/zenkit/v5"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
	"github.com/zenoss/event-management-service/pkg/adapters/framework/mongodb"
	"github.com/zenoss/event-management-service/pkg/models/event"
)

func toDocument(v any) (bson.D, error) {
	vBytes, err := bson.Marshal(v)
	if err != nil {
		return nil, err
	}
	result := bson.D{}
	err = bson.Unmarshal(vBytes, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func mkEventResult(
	numClosedOccurrences int,
	ID, tenant, entity string,
	dimensions map[string]any,
	now time.Time,
) *mongo.Event {
	result := &mongo.Event{
		ID:         ID,
		Tenant:     tenant,
		Entity:     entity,
		Dimensions: dimensions,
		CreatedAt:  time.Now().Add(-8 * 24 * time.Hour),
		UpdatedAt:  time.Now().Add(-24 * time.Hour),
	}
	occurrences := make([]*mongo.Occurrence, 0)
	if numClosedOccurrences > 0 {
		for i := 0; i < numClosedOccurrences; i++ {
			startOffset, _ := time.ParseDuration(fmt.Sprintf("-%dd", 8-i))
			endOffset, _ := time.ParseDuration(fmt.Sprintf("-%dd", 7-i))
			start := now.Add(startOffset)
			end := now.Add(endOffset)
			occurrences = append(occurrences, &mongo.Occurrence{
				ID:           fmt.Sprintf("%s:%d", ID, i+1),
				EventID:      ID,
				Tenant:       tenant,
				Status:       event.StatusClosed,
				Severity:     event.SeverityInfo,
				Acknowledged: new(bool),
				StartTime:    start.UnixMilli(),
				EndTime:      end.UnixMilli(),
				CurrentTime:  end.UnixMilli(),
				CreatedAt:    start.Add(10 * time.Second),
				UpdatedAt:    end.Add(10 * time.Second),
				Entity:       entity,
				Dimensions:   dimensions,
				LastSeen:     end.UnixMilli(),
			})
		}
	}
	startOffset, _ := time.ParseDuration(fmt.Sprintf("-%dd", 1))
	start := now.Add(startOffset)
	occurrences = append(occurrences, &mongo.Occurrence{
		ID:           fmt.Sprintf("%s:%d", ID, len(occurrences)+1),
		EventID:      ID,
		Tenant:       tenant,
		Status:       event.StatusOpen,
		Severity:     event.SeverityInfo,
		Acknowledged: new(bool),
		StartTime:    start.UnixMilli(),
		CurrentTime:  start.UnixMilli(),
		CreatedAt:    start.Add(10 * time.Second),
		UpdatedAt:    start.Add(10 * time.Second),
		Entity:       entity,
		Dimensions:   dimensions,
		LastSeen:     start.UnixMilli(),
	})
	result.Occurrences = occurrences
	result.OccurrenceCount = uint64(len(occurrences))
	result.LastActiveOccurrence = occurrences[len(occurrences)-1]
	return result
}

func matchFindResults(
	resp *event.Page,
	expectedEvents []any,
) {
	for i, eventAny := range expectedEvents {
		event, ok := eventAny.(*mongo.Event)
		Expect(ok).To(BeTrue())
		Expect(resp.Results[i]).To(PointTo(MatchFields(IgnoreExtras, Fields{
			"ID": Equal(event.ID),
		})))
		actual := resp.Results[i]
		matchers := make([]types.GomegaMatcher, 0)
		for _, expectedOcc := range event.Occurrences {
			matchers = append(matchers, PointTo(MatchFields(IgnoreExtras, Fields{
				"ID":        Equal(expectedOcc.ID),
				"EventID":   Equal(expectedOcc.EventID),
				"Status":    Equal(expectedOcc.Status),
				"Severity":  Equal(expectedOcc.Severity),
				"StartTime": Equal(expectedOcc.StartTime),
				"Entity":    Equal(expectedOcc.Entity),
			})))
		}
		Expect(actual.Occurrences).To(ConsistOf(matchers))
	}
}

var _ = Describe("MongoDB Integration Test", func() {
	var (
		testCtx    context.Context
		testCancel context.CancelFunc
	)
	BeforeEach(func() {
		zenkit.InitConfigForTests(GinkgoT(), "event-management-service")
		testCtx, testCancel = context.WithCancel(context.Background())
		log := logrus.NewEntry(logrus.StandardLogger())
		log.Logger.SetFormatter(&logrus.JSONFormatter{})
		log.Logger.SetLevel(logrus.DebugLevel)
		testCtx = ctxlogrus.ToContext(testCtx, log)
	})
	AfterEach(func() {
		testCancel()
	})
	It("should connect to the configured MongoDB deployment", func() {
		cfg := mongodb.Config{
			Address:    viper.GetString(config.MongoDBAddr),
			Username:   "zing-user",
			Password:   ".H1ddenPassw0rd.",
			DefaultTTL: 90 * 24 * time.Hour,
			DBName:     "event-context-svc",
		}
		addresses := make(map[string]string)
		for i, addr := range viper.GetStringSlice(zenkit.GCMemstoreAddressConfig) {
			addresses[fmt.Sprintf("%d", i+1)] = addr
		}
		Expect(addresses).ToNot(BeEmpty())
		redisClient := redis.NewRing(&redis.RingOptions{
			Addrs:       addresses,
			DialTimeout: 2 * time.Second,
		})
		cursorAdapter := redisCursors.NewAdapter(redisClient)
		db, err := mongodb.NewMongoDatabaseConnection(testCtx, cfg)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(db).ShouldNot(BeNil())
		adapter, err := mongo.NewAdapter(testCtx, cfg, db, cursorAdapter)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(adapter).ShouldNot(BeNil())
	})
	Context("mongodb.Adapter.Find", func() {
		var (
			cursorAdapter event.CursorRepository
			adapter       *mongo.Adapter
			cfg           mongodb.Config
		)
		BeforeEach(func() {
			var err error
			cfg = mongodb.Config{
				Address:    viper.GetString(config.MongoDBAddr),
				Username:   "zing-user",
				Password:   ".H1ddenPassw0rd.",
				DefaultTTL: 90 * 24 * time.Hour,
				DBName:     "event-context-svc",
			}

			addresses := make(map[string]string)
			for i, addr := range viper.GetStringSlice(zenkit.GCMemstoreAddressConfig) {
				addresses[fmt.Sprintf("%d", i+1)] = addr
			}
			Expect(addresses).ToNot(BeEmpty())
			redisClient := redis.NewRing(&redis.RingOptions{
				Addrs:       addresses,
				DialTimeout: 2 * time.Second,
			})
			cursorAdapter = redisCursors.NewAdapter(redisClient)

			db, err := mongodb.NewMongoDatabaseConnection(testCtx, cfg)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(db).ShouldNot(BeNil())
			adapter, err = mongo.NewAdapter(testCtx, cfg, db, cursorAdapter)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(adapter).ShouldNot(BeNil())
		})
		It("finds documents", func() {
			input := &event.Query{
				Tenant: "acme",
				TimeRange: event.TimeRange{
					Start: time.Now().Add(-1 * time.Hour).UnixMilli(),
					End:   time.Now().UnixMilli(),
				},
				Statuses:   []event.Status{event.StatusOpen},
				Severities: []event.Severity{event.SeverityError, event.SeverityCritical},
			}
			resp, err := adapter.Find(testCtx, input)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp).ShouldNot(BeNil())
		})
		Context("Pagination Integration Tests", func() {
			It("page through event query results", func() {
				var (
					numResults int = 15
					pageSize   int = 5
				)
				integrationTestTenant := fmt.Sprintf("acme_%s", zenkit.RandString(10))
				db, err := mongodb.NewMongoDatabaseConnection(testCtx, cfg)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(db).ShouldNot(BeNil())
				eventCollection := db.Collection(mongo.CollEvents)
				testEvents := make([]any, 0)
				occCollection := db.Collection(mongo.CollOccurrences)
				testOccurrences := make([]any, 0)
				now := time.Now()
				randID := zenkit.RandString(16)
				entityIds := make([]string, 0)
				for i := 0; i < numResults; i++ {
					eventResult := mkEventResult(
						0,
						fmt.Sprintf("event_%s_%02d", randID, i),
						integrationTestTenant,
						fmt.Sprintf("entity_%s", zenkit.RandString(6)),
						nil,
						now.Add(time.Duration(-1*i)*(time.Hour)),
					)
					testEvents = append(testEvents, eventResult)
					for _, occ := range eventResult.Occurrences {
						testOccurrences = append(testOccurrences, occ)
						entityIds = append(entityIds, occ.Entity)
					}
				}

				insertManyResult1, err := eventCollection.InsertMany(testCtx, testEvents)
				Expect(err).ToNot(HaveOccurred())
				Expect(insertManyResult1).ToNot(BeNil())

				insertManyResult, err := occCollection.InsertMany(testCtx, testOccurrences)
				Expect(err).ToNot(HaveOccurred())
				Expect(insertManyResult).ToNot(BeNil())

				cursor := ""
				for i := 0; i < len(testEvents); i += pageSize {
					input := &event.Query{
						Tenant: integrationTestTenant,
						TimeRange: event.TimeRange{
							Start: now.Add(-8 * 24 * time.Hour).UnixMilli(),
							End:   now.UnixMilli(),
						},
						Statuses:   []event.Status{event.StatusOpen},
						Severities: []event.Severity{event.SeverityInfo},
						Filter: &event.Filter{
							Op:    event.FilterOpAnd,
							Field: string(event.FilterOpAnd),
							Value: []any{
								&event.Filter{
									Field: "startTime",
									Op:    event.FilterOpGreaterThan,
									Value: 0,
								},
								&event.Filter{
									Field: "entity",
									Op:    event.FilterOpIn,
									Value: entityIds,
								},
							},
						},
						PageInput: &event.PageInput{
							Cursor: cursor,
							Limit:  uint64(pageSize),
							SortBy: []event.SortOpt{
								{
									Field:     "startTime",
									SortOrder: event.SortOrderDescending,
								},
							},
						},
					}
					resp, err := adapter.Find(testCtx, input)
					Expect(err).ShouldNot(HaveOccurred())
					Expect(resp.Results).To(HaveLen(pageSize))
					Expect(resp.Cursor).ToNot(BeEmpty())
					if i+len(resp.Results) < len(testEvents) {
						Expect(resp.HasNext).To(BeTrue())
					} else {
						Expect(resp.HasNext).To(BeFalse())
					}
					end := i + pageSize
					if end > len(testEvents) {
						end = len(testEvents)
					}
					matchFindResults(resp, testEvents[i:end])
					cursor = resp.Cursor

				}
				for i := len(testEvents); i > 0; i -= pageSize {
					input := &event.Query{
						Tenant: integrationTestTenant,
						TimeRange: event.TimeRange{
							Start: now.Add(-8 * 24 * time.Hour).UnixMilli(),
							End:   now.UnixMilli(),
						},
						Statuses:   []event.Status{event.StatusOpen},
						Severities: []event.Severity{event.SeverityInfo},
						Filter: &event.Filter{
							Op:    event.FilterOpAnd,
							Field: string(event.FilterOpAnd),
							Value: []any{
								&event.Filter{
									Field: "startTime",
									Op:    event.FilterOpGreaterThan,
									Value: 0,
								},
								&event.Filter{
									Field: "entity",
									Op:    event.FilterOpIn,
									Value: entityIds,
								},
							},
						},
						PageInput: &event.PageInput{
							Cursor:    cursor,
							Direction: event.PageDirectionBackward,
							Limit:     uint64(pageSize),
							SortBy: []event.SortOpt{
								{
									Field:     "startTime",
									SortOrder: event.SortOrderDescending,
								},
							},
						},
					}
					resp, err := adapter.Find(testCtx, input)
					Expect(err).ShouldNot(HaveOccurred())
					Expect(resp.Results).To(HaveLen(pageSize))
					Expect(resp.Cursor).ToNot(BeEmpty())
					// if i+len(resp.Results) < len(testEvents) {
					// 	Expect(resp.HasNext).To(BeTrue())
					// } else {
					// 	Expect(resp.HasNext).To(BeFalse())
					// }
					start := i - pageSize
					matchFindResults(resp, testEvents[start:i])
					cursor = resp.Cursor
				}
			})
		})
	})
})
