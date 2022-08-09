//go:build integration

package mongodb_test

import (
	"context"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
	"github.com/zenoss/event-management-service/pkg/adapters/framework/mongodb"
)

var _ = Describe("MongoDB Integration Test", Ordered, func() {
	var (
		testCtx            context.Context
		testCancel         context.CancelFunc
		testOnce           = sync.Once{}
		testDBName         = "testDB"
		testCollectionName = "testCollection"
		mongoAddr          = "mongodb"
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
	Context("MongoDB Helper Classes", func() {
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
				_, err = testCollection.InsertOne(testCtx, bson.D{
					{Key: "title", Value: "Owning IT"},
					{Key: "author", Value: "Zenny Zebra"},
				})
				Expect(err).ToNot(HaveOccurred())
			})
		})
		It("should connect to the configured MongoDB and make a query", func() {
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

			cur, err := collection.Find(testCtx, bson.D{})
			Expect(err).ShouldNot(HaveOccurred())
			Expect(cur).ToNot(BeNil())
		})
	})
})
