package mongodb_test

import (
	"context"
	"fmt"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
	"github.com/zenoss/event-management-service/pkg/adapters/framework/mongodb"
)

var _ = Describe("MongoDB Unit Tests", func() {
	var (
		testCtx    context.Context
		testCancel context.CancelFunc
		descrFunc  = func(cfg mongodb.Config, expected string) string {
			return fmt.Sprintf("%v.URI() == %s", cfg, expected)
		}
	)
	BeforeEach(func() {
		testCtx, testCancel = context.WithCancel(context.Background())
		log := logrus.NewEntry(logrus.StandardLogger())
		testCtx = ctxlogrus.ToContext(testCtx, log)
	})
	AfterEach(func() {
		testCancel()
	})
	DescribeTable("mongodb.Config Table-driven Tests", func(cfg mongodb.Config, expected string) {
		Ω(cfg.URI()).Should(Equal(expected))
	},
		Entry(descrFunc, mongodb.Config{Address: "mongodb"}, "mongodb://mongodb"),
		Entry(descrFunc, mongodb.Config{Address: "mongodb", Username: "testUser", Password: "testPasswd"}, "mongodb://testUser:testPasswd@mongodb"),
		Entry(descrFunc, mongodb.Config{Address: "mongodb", CACert: "/testCA.crt", ClientCert: "/testClient.crt"}, "mongodb://mongodb/?authMechanism=MONGODB-X509&tlsCAFile=/testCA.crt&tlsCertificateKeyFile=/testClient.crt"),
	)
})
