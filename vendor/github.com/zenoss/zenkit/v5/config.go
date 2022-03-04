package zenkit

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

const (
	LogLevelConfig       = "log.level"
	GrpcLogLevelConfig   = "grpc.log.level"
	LogStackdriverConfig = "log.stackdriver"

	TracingEnabledConfig    = "tracing.enabled"
	TracingSampleRateConfig = "tracing.samplerate"

	MetricsEnabledConfig = "metrics.enabled"

	// Control which OpenCensus metric and trace exporters are enabled.
	ExporterStackdriverEnabledConfig = "exporter.stackdriver.enabled"
	ExporterJaegerEnabledConfig      = "exporter.jaeger.enabled"
	ExporterZenossEnabledConfig      = "exporter.zenoss.enabled"
	ExporterJaegerAgentConfig        = "exporter.jaeger.agent.endpoint"
	ExporterJaegerCollectorConfig    = "exporter.jaeger.collector.endpoint"

	// Additional options for configuring the stackdriver exporter
	ExporterStackdriverTraceMaxBufferConfig = "exporter.stackdriver.trace.max.buffer"

	ServiceLabel = "service.label"

	// When our environment tells us which Kubernetes pod we're in.
	K8sClusterConfig   = "k8s.cluster"
	K8sNamespaceConfig = "k8s.namespace"
	K8sPodConfig       = "k8s.pod"

	ProfilingEnabledConfig           = "profiling.enabled"
	ProfilingServiceName             = "profiling.service.name"
	ProfilingMutexDisabledConfig     = "profiling.mutex.disabled"
	ProfilingHeapDisabledConfig      = "profiling.heap.disabled"
	ProfilingAllocDisabledConfig     = "profiling.alloc.disabled"
	ProfilingGoroutineDisabledConfig = "profiling.goroutine.disabled"
	ProfilingAllocForceGC            = "profiling.alloc.force_gc"

	JWTAudienceConfig = "jwt.audience"
	JWTIssuerConfig   = "jwt.issuer"
	JWTSkip           = "jwt.skip"
	JWKSURL           = "jwks.url"
	JWKSCacheMinutes  = "jwks.cache.minutes"

	AuthDisabledConfig      = "auth.disabled"
	AuthDevTenantConfig     = "auth.dev_tenant"
	AuthDevEmailConfig      = "auth.dev_email"
	AuthDevUserConfig       = "auth.dev_user"
	AuthDevConnectionConfig = "auth.dev_connection"
	AuthDevScopesConfig     = "auth.dev_scopes"
	AuthDevGroupsConfig     = "auth.dev_groups"
	AuthDevRolesConfig      = "auth.dev_roles"
	AuthDevClientIDConfig   = "auth.dev_clientid"
	AuthDevSubjectConfig    = "auth.dev_subject"

	GRPCMaxConcurrentRequests = "grpc.max_concurrent_requests"
	GRPCListenAddrConfig      = "grpc.listen_addr"
	GRPCHealthAddrConfig      = "grpc.health_addr"
	GRPCMaxMsgSizeSendConfig  = "grpc.max_msg_size_send"
	GRPCMaxMsgSizeRecvConfig  = "grpc.max_msg_size_recv"

	GCProjectIDConfig                     = "gcloud.project_id"
	GCDatastoreCredentialsConfig          = "gcloud.datastore.credentials"
	GCEmulatorHostConfig                  = "gcloud.emulator.host"
	GCEmulatorProjectConfig               = "gcloud.emulator.project"
	GCEmulatorBigtableConfig              = "gcloud.emulator.bigtable"
	GCEmulatorTableList                   = "gcloud.emulator.table.list"
	GCEmulatorDatastoreEnabledConfig      = "gcloud.emulator.datastore.enabled"
	GCEmulatorDatastoreHostPortConfig     = "gcloud.emulator.datastore.host_port"
	GCEmulatorPubsubConfig                = "gcloud.emulator.pubsub"
	GCBigtableInstanceIDConfig            = "gcloud.bigtable.instance_id"
	GCBigtableApplicationProfileIDConfig  = "gcloud.bigtable.application_profile_id"
	GCBigtableApplicationProfileUpsert    = "gcloud.bigtable.application_profile_upsert"
	GCBigtableTableOverrides              = "gcloud.bigtable.table_overrides"
	GCBigtableApplicationProfileOverrides = "gcloud.bigtable.application_profile_overrides"
	GCPubsubTopicConfig                   = "gcloud.pubsub.topic"
	GCBigtableSuffix                      = "gcloud.bigtable.suffix"
	GCBigtablePoolSize                    = "gcloud.bigtable.poolsize"
	GCMemstoreAddressConfig               = "gcloud.memorystore.address"
	GCMemstoreTTLConfig                   = "gcloud.memorystore.ttl"
	GCMemstoreLocalMaxLen                 = "gcloud.memorystore.local_max_len"

	HTTP2ListenAddrConfig = "http2.listen_addr"

	PolicyServiceRefreshInterval = "policy.service.refresh.interval"

	ServiceDialTimeoutConfig = "dial_timeout"

	ZINGAnomalyTableConfig           = "zing.bigtable.table.anomaly"
	ZINGDefinitionIDIndexTableConfig = "zing.bigtable.table.definition_id_index"
	ZINGFieldIndexTableConfig        = "zing.bigtable.table.field_index"
	ZINGItemDefinitionTableConfig    = "zing.bigtable.table.item_definition"
	ZINGItemInstanceTableConfig      = "zing.bigtable.table.item_instance"
	ZINGMetadataTableConfig          = "zing.bigtable.table.metadata"
	ZINGMetricsTableConfig           = "zing.bigtable.table.metrics"
	ZINGRecommendationsTableConfig   = "zing.bigtable.table.recommendations"
	ZINGTrendTableConfig             = "zing.bigtable.table.trend"
	ZINGQueryResultsTableConfig      = "zing.bigtable.table.query_results"

	ZINGProductNameConfig          = "zing.product.name"
	ZINGProductVersionConfig       = "zing.product.version"
	ZINGProductCompanyNameConfig   = "zing.product.company_name"
	ZINGProductOtherCommentsConfig = "zing.product.other_comments"

	// Auth0DomainConfig is the config to get the auth0 domain, used for logging a user out
	Auth0DomainConfig = "auth0.domain"

	// Auth0 creds for connecting to other microservices (currently, just tenant service)
	BackendClientIDConfig     = "auth0.backend.client_id"
	BackendClientSecretConfig = "auth0.backend.client_secret"
	BackendClientAudience     = "auth0.backend.client_audience"
)

var (
	globalViper         = viper.New()
	ErrNoServiceAddress = errors.New("no service address")

	initOnce = &sync.Once{}
)

func init() {
	globalViper.AutomaticEnv()
	globalViper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	globalViper.SetDefault(ServiceDialTimeoutConfig, 10*time.Second)

	globalViper.SetDefault(ZINGAnomalyTableConfig, "ANOMALY_V2")
	globalViper.SetDefault(ZINGDefinitionIDIndexTableConfig, "DEFINITION_ID_INDEX_V2")
	globalViper.SetDefault(ZINGFieldIndexTableConfig, "FIELD_INDEX_V2")
	globalViper.SetDefault(ZINGItemDefinitionTableConfig, "ITEM_DEFINITION_V2")
	globalViper.SetDefault(ZINGItemInstanceTableConfig, "ITEM_INSTANCE_V2")
	globalViper.SetDefault(ZINGMetadataTableConfig, "METADATA_V3")
	globalViper.SetDefault(ZINGMetricsTableConfig, "METRICS_V2")
	globalViper.SetDefault(ZINGRecommendationsTableConfig, "RECOMMENDATIONS_V2")
	globalViper.SetDefault(ZINGTrendTableConfig, "TREND_V2")
	globalViper.SetDefault(ZINGQueryResultsTableConfig, "QUERY_RESULTS_V2")

	globalViper.SetDefault(ZINGProductNameConfig, "Zenoss Cloud")
	globalViper.SetDefault(ZINGProductVersionConfig, "1.0")
	globalViper.SetDefault(ZINGProductCompanyNameConfig, "Zenoss")
}

var (
	configMu sync.Mutex
)

func initConfigWithName(name string) func() {
	return func() {
		configMu.Lock()
		defer configMu.Unlock()
		viper.SetDefault(LogLevelConfig, "info")
		viper.SetDefault(GrpcLogLevelConfig, "warn")
		viper.SetDefault(LogStackdriverConfig, true)
		viper.SetDefault(TracingEnabledConfig, true)
		viper.SetDefault(TracingSampleRateConfig, 1.0)
		viper.SetDefault(MetricsEnabledConfig, true)
		viper.SetDefault(ExporterStackdriverEnabledConfig, true)
		viper.SetDefault(ExporterZenossEnabledConfig, true)
		viper.SetDefault(ExporterJaegerEnabledConfig, false)
		viper.SetDefault(ExporterJaegerAgentConfig, "localhost:6831")
		viper.SetDefault(ExporterJaegerCollectorConfig, "localhost:14268")
		viper.SetDefault(AuthDevTenantConfig, "ACME")
		viper.SetDefault(AuthDevUserConfig, "zcuser@acme.example.com")
		viper.SetDefault(AuthDevEmailConfig, "zcuser@acme.example.com")
		viper.SetDefault(AuthDevClientIDConfig, "0123456789abcdef")
		viper.SetDefault(GRPCListenAddrConfig, ":8080")
		viper.SetDefault(GRPCHealthAddrConfig, ":8081")
		viper.SetDefault(GRPCMaxConcurrentRequests, 0)
		viper.SetDefault(GRPCMaxMsgSizeSendConfig, int(7e6))
		viper.SetDefault(GRPCMaxMsgSizeRecvConfig, int(7e6))
		viper.SetDefault(GCBigtableInstanceIDConfig, "zenoss-zing-bt1")
		viper.SetDefault(GCProjectIDConfig, "zenoss-zing")
		viper.SetDefault(GCBigtableSuffix, "")
		viper.SetDefault(GCBigtablePoolSize, 4)
		viper.SetDefault(GCBigtableApplicationProfileIDConfig, "")
		viper.SetDefault(GCBigtableApplicationProfileUpsert, "")
		viper.SetDefault(GCMemstoreAddressConfig, []string{})
		viper.SetDefault(GCMemstoreTTLConfig, "24h")
		viper.SetDefault(GCMemstoreLocalMaxLen, 1000000)
		viper.SetDefault(HTTP2ListenAddrConfig, ":9080")
		viper.SetDefault(JWKSCacheMinutes, 30)

		viper.SetDefault(PolicyServiceRefreshInterval, 300)

		viper.SetDefault(ProfilingEnabledConfig, false)
		viper.SetDefault(ProfilingServiceName, name)
		viper.SetDefault(ProfilingMutexDisabledConfig, false)
		viper.SetDefault(ProfilingHeapDisabledConfig, false)
		viper.SetDefault(ProfilingGoroutineDisabledConfig, false)
		viper.SetDefault(ProfilingAllocDisabledConfig, false)
		viper.SetDefault(ProfilingAllocForceGC, true)

		viper.SetDefault(GCEmulatorHostConfig, "bigtable:8080")
		viper.SetDefault(GCEmulatorProjectConfig, "zenoss-zing")
		viper.SetDefault(GCEmulatorBigtableConfig, "zing-bt-emu")

		viper.SetDefault(ServiceLabel, name)

		viper.SetEnvPrefix(name)
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
		viper.AutomaticEnv()
	}
}

func InitConfig(name string) {
	// ZING-4105: Prevent fatal error: concurrent map writes
	initOnce.Do(initConfigWithName(name))
}

// ServiceAddress returns a "<host>:<port>" address for svc service name.
// The svc string is expected to be a resolvable hostname.
//
// In a Kubernetes context it should be the unqualified name of the Service.
// When running within Kubernetes, the <host> portion will always be exactly
// the same as svc, and <port> will be looked up in the <svc>_SERVICE_PORT
// environment variable. The <svc>_SERVICE_HOST environment variable provided
// by Kubernetes is intentionally ignored in favor of the service's name to
// avoid service startup race conditions, and for dealing with services
// potentially changing IP addresses over time.
//
// When running outside of Kubernetes, the <host> portion will be looked up in
// <svc>_SERVICE_HOST, and the <port> portion will be looked up in
// <svc>_SERVICE_PORT.
func ServiceAddress(svc string) (string, error) {
	var host string
	if os.Getenv("KUBERNETES_PORT") != "" {
		host = svc
	} else {
		host = globalViper.GetString(svc + "_SERVICE_HOST")
		if host == "" {
			host = svc
		}
	}
	port := globalViper.GetString(svc + "_SERVICE_PORT")
	if host == "" || port == "" {
		return "", ErrNoServiceAddress
	}
	return fmt.Sprintf("%s:%s", host, port), nil
}

func GlobalConfig() *viper.Viper {
	return globalViper
}
