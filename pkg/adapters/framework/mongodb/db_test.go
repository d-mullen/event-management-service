package mongodb_test

import (
	"context"
	"fmt"
	"net/url"

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
		urlParsed, err := url.Parse(cfg.URI())
		Ω(err).ShouldNot(HaveOccurred())
		expectedURL, _ := url.Parse(expected)
		expectedOpts, _ := url.ParseQuery(expectedURL.RawQuery)
		opts, err := url.ParseQuery(urlParsed.RawQuery)
		Ω(err).ShouldNot(HaveOccurred())

		Expect(urlParsed.Scheme).Should(BeEquivalentTo(expectedURL.Scheme))
		Expect(urlParsed.Host).Should(BeEquivalentTo(expectedURL.Host))
		Expect(urlParsed.User).Should(BeEquivalentTo(expectedURL.User))
		Expect(opts).Should(BeEquivalentTo(expectedOpts))

	},
		Entry(descrFunc, mongodb.Config{Address: "mongodb"}, "mongodb://mongodb/"),
		Entry(descrFunc, mongodb.Config{Address: "mongodb", Username: "testUser", Password: "testPasswd"}, "mongodb://testUser:testPasswd@mongodb/?authMechanism=SCRAM-SHA-256"),
		Entry(descrFunc, mongodb.Config{Address: "mongodb", CACert: "/testCA.crt", ClientCert: "/testClient.crt"}, "mongodb://mongodb/?authMechanism=MONGODB-X509&readPreference=secondary&tlsCAFile=/testCA.crt&tlsCertificateKeyFile=/testClient.crt"),
		Entry(descrFunc, mongodb.Config{Address: "mongodb",
			CACert:     "/testCA.crt",
			ClientCert: "/testClient.crt",
			Options: map[string]string{"tlsInsecure": "true",
				"connectTimeoutMS": "30000",
				"socketTimeoutMS":  "22222"},
		},
			"mongodb://mongodb?authMechanism=MONGODB-X509&connectTimeoutMS=30000&readPreference=secondary&socketTimeoutMS=22222&tlsCAFile=%2FtestCA.crt&tlsCertificateKeyFile=%2FtestClient.crt&tlsInsecure=true"),
	)
})
