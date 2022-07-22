package config

import (
	"time"

	"github.com/spf13/viper"
)

const (
	DefaultBackendDatabase = "mongo"
	// same defaults in deployment.yaml
	DefaultMongoDBAddr                    = "zing-mongodb"
	DefaultMongoDBCACertificate           = "/mongo-user/ca.crt"
	DefaultMongoDBClientCertificate       = "/mongo-user/user.pem"
	DefaultMongoDBName                    = "event-context-svc"
	DefaultActiveEntityStoreBucketSize    = 30 * 24 * time.Hour
	DefaultActiveEntityStoreMinBucketSize = time.Hour
	DefaultActiveEntityStoreTTL           = 24 * time.Hour

	// EventManagementEnabledConfig determines if the API endpoint for event management is enabled for this server
	EventManagementEnabledConfig = "eventmanagement.enabled"
	EventQueryEnabled            = "event.query.enabled"
	MongoDBAddr                  = "mongo.address"
	MongoDBName                  = "mongo.db.name"

	// TraceRateLimitedSamplingEnabled - specifies whether rate-limited tracing is enabled
	TraceRateLimitedSamplingEnabled = "trace.rate.limited.sampling.enabled"

	// The name of the CA certificate file: For Example: ca.crt
	MongoDBCACertificate = "mongo.certificate.ca.path"
	// The name of the client certificate and private key file.
	MongoDBClientCertificate = "mongo.certificate.client.path"

	// The name of the privileged CA certificate file: For Example: ca.crt
	MongoAdmDBCACertificate = "mongo.adm.certificateCAPath"
	// The name of the privileged client certificate and private key file.
	MongoAdmDBClientCertificate = "mongo.adm.ClientCertificate"

	// ActiveEntityStore Configuration

	// ActiveEntityStoreBucketSize - the duration by which active entity sets are bucketed
	ActiveEntityStoreBucketSize = "active.entity.store.bucket.size"
	// ActiveEntityStoreMinBucketSize - the mininum time bucket size (duration) used to generate active entity record keys
	ActiveEntityStoreMinBucketSize = "active.entity.store.min.bucket.size"
	// ActiveEntityStoreDefaultTTL - the default duration for expiring active entity records
	ActiveEntityStoreDefaultTTL = "active.entity.store.default.ttl"
)

// InitDefaults sets defaults values for this server's configuration
func init() {
	viper.SetDefault(EventManagementEnabledConfig, true)
	viper.SetDefault(EventQueryEnabled, false)
	viper.SetDefault(MongoDBAddr, DefaultMongoDBAddr)
	viper.SetDefault(MongoDBCACertificate, DefaultMongoDBCACertificate)
	viper.SetDefault(MongoDBClientCertificate, DefaultMongoDBClientCertificate)
	viper.SetDefault(MongoDBName, DefaultMongoDBName)
	viper.SetDefault(ActiveEntityStoreDefaultTTL, DefaultActiveEntityStoreTTL)
	viper.SetDefault(ActiveEntityStoreBucketSize, DefaultActiveEntityStoreBucketSize)
	viper.SetDefault(ActiveEntityStoreMinBucketSize, DefaultActiveEntityStoreMinBucketSize)
}
