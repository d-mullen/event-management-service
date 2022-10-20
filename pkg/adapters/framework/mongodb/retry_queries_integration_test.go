//go:build integration

package mongodb_test

import (
	"context"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
	"github.com/zenoss/event-management-service/pkg/adapters/framework/mongodb"
	"github.com/zenoss/zenkit/v5"
)

type (
	networkErrState struct {
		count, errCount int
		every, maxFails int
	}
	cursorFake struct {
		mongodb.Cursor
		err      error
		errState *networkErrState
	}
	collectionWrapper struct {
		mongodb.Collection
		every, maxFails int
		errState        *networkErrState
	}
)

func (s *networkErrState) IncrementCount() {
	s.count += 1
}

func (s *networkErrState) Reset() {
	s.count = 0
	s.errCount = 0
}

func (s *networkErrState) ShouldErr() bool {
	if (s.count%s.every) == 0 &&
		(s.maxFails == 0 || s.errCount < s.maxFails) {
		// if s.count%s.every == 0 {
		s.errCount++
		return true
	}
	return false
}

func (cur *cursorFake) Next(ctx context.Context) bool {
	cur.errState.IncrementCount()
	if cur.errState.ShouldErr() {
		cur.err = mongo.CommandError{Code: mongodb.CursorNotFound, Message: "test error"}
		return false
	}
	return cur.Cursor.Next(ctx)
}

func (cur *cursorFake) Err() error {
	if cur.err != nil {
		return cur.err
	}
	return cur.Cursor.Err()
}

func (coll *collectionWrapper) Find(ctx context.Context, filter any, opts ...*options.FindOptions) (mongodb.Cursor, error) {
	cur, err := coll.Collection.Find(ctx, filter, opts...)
	if err != nil {
		return nil, err
	}
	return &cursorFake{
		Cursor:   cur,
		errState: coll.errState,
	}, nil
}

var _ = Describe("Query Retry Integration Test", Ordered, func() {
	var (
		testCtx            context.Context
		testCancel         context.CancelFunc
		testOnce           = sync.Once{}
		testDBName         = "testDB"
		testCollectionName = "testCollection_" + zenkit.RandString(8)
		mongoAddr          = "mongodb"
		testResults        *mongo.InsertManyResult
		expectedIDs        = make([]string, 0)
	)
	BeforeAll(func() {
		if addr := os.Getenv("MONGO_ADDRESS"); len(addr) > 0 {
			mongoAddr = addr
		}
	})

	BeforeEach(func() {
		testCtx, testCancel = context.WithCancel(context.Background())
		log := logrus.NewEntry(logrus.StandardLogger())
		log.Logger.SetFormatter(&logrus.JSONFormatter{})
		log.Logger.SetLevel(logrus.DebugLevel)
		testCtx = ctxlogrus.ToContext(testCtx, log)
	})
	AfterEach(func() {
		testCancel()
	})
	Context("FindWithRetry", func() {
		BeforeEach(func() {
			testOnce.Do(func() {
				cfg := mongodb.Config{
					Address:    mongoAddr,
					Username:   "zing-user",
					Password:   ".H1ddenPassw0rd.",
					DefaultTTL: 90 * 24 * time.Hour,
					DBName:     testDBName,
				}
				clientOpts := options.Client().ApplyURI(cfg.URI())
				client, err := mongo.Connect(testCtx, clientOpts)
				Expect(err).ToNot(HaveOccurred())

				testDb := client.Database(testDBName)
				testCollection := testDb.Collection(testCollectionName)
				docs := make([]any, 500)
				for i := 0; i < 500; i++ {
					docs[i] = bson.D{
						{"_id", zenkit.RandString(16)},
						{Key: "name", Value: zenkit.RandString(8)},
					}
				}
				testResults, err = testCollection.InsertMany(testCtx, docs)
				Expect(err).ToNot(HaveOccurred())
				Expect(testResults).ToNot(BeNil())
				for _, r := range testResults.InsertedIDs {
					id, ok := r.(primitive.ObjectID)
					if ok {
						expectedIDs = append(expectedIDs, id.Hex())
					} else {
						idStr, ok2 := r.(string)
						Expect(ok2).To(BeTrue())
						expectedIDs = append(expectedIDs, idStr)
					}
				}
			})
		})
		It("should connect to MongoDB and make a network-error resilient query", func() {
			cfg := mongodb.Config{
				Address:    mongoAddr,
				Username:   "zing-user",
				Password:   ".H1ddenPassw0rd.",
				DefaultTTL: 90 * 24 * time.Hour,
				DBName:     testDBName,
			}
			db, err := mongodb.NewMongoDatabaseConnection(testCtx, cfg)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(db).ShouldNot(BeNil())

			collection := db.Collection(testCollectionName)
			Expect(collection).ToNot(BeNil())
			errState := &networkErrState{
				count:    0,
				every:    13,
				maxFails: 10,
			}
			collection = &collectionWrapper{
				Collection: collection,
				errState:   errState,
			}

			results := make([]*doc, 0)
			actualIDs := make([]string, 0)
			err = mongodb.FindWithRetry(
				testCtx,
				bson.D{},
				[]*options.FindOptions{options.Find().SetSort(bson.D{{"_id", 1}})},
				collection.Find,
				func(result doc) (bool, string, error) {
					results = append(results, &result)
					actualIDs = append(actualIDs, result.ID)
					return true, result.ID, nil
				},
			)
			Expect(err).ToNot(HaveOccurred())
			Expect(actualIDs).To(ConsistOf(expectedIDs))
		})
	})
})
