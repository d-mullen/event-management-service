package exporter

import (
	"context"
	"fmt"
	"os"
	"time"

	grpcRetry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/sirupsen/logrus"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/stats/view"
	"google.golang.org/api/support/bundler"
	"google.golang.org/grpc"

	"github.com/zenoss/zing-proto/v11/go/cloud/data_receiver"
)

const (
	// SourceField is the tag identifying a metric's source.
	SourceField = "source"

	// SourceTypeField is a tag identifying the type of source.
	SourceTypeField = "source-type"

	// DefaultSourceType is the default source-type identifying this OpenCensus exporter.
	DefaultSourceType = "zenoss/opencensus-go-exporter-zenoss"

	// DescriptionField is the optional tag containing a metric's description.
	DescriptionField = "description"

	// UnitsField is the optional tag containing a metric's unit of measure.
	UnitsField = "units"

	// ZenossAPIProxyAddress is the gRPC address to which metrics will be sent.
	ZenossAPIProxyAddress = "data-receiver-proxy:80"

	// K8sClusterEnv is an environment variable that may contain our Kubernetes cluster.
	K8sClusterEnv = "K8S_CLUSTER"

	// K8sClusterTagKey is the tag where Zenoss expects a Kubernetes cluster.
	K8sClusterTagKey = "k8s.cluster"

	// K8sNamespaceEnv is an environment variable that may contain our Kubernetes namespace.
	K8sNamespaceEnv = "K8S_NAMESPACE"

	// K8sNamespaceTagKey is the tag where Zenoss expects a Kubernetes namespace.
	K8sNamespaceTagKey = "k8s.namespace"

	// K8sPodEnv is an environment variable that may contain our Kubernetes pod.
	K8sPodEnv = "K8S_POD"

	// K8sPodTagKey is the tag where Zenoss expects a Kubernetes pod.
	K8sPodTagKey = "k8s.pod"
)

var (
	// Ensure we implement view.Explorer interface.
	_ view.Exporter = (*Exporter)(nil)
)

// Options defines the options for a Zenoss OpenCensus exporter.
type Options struct {
	// Source is a Zenoss tag commonly used to identify the source of data.
	Source string

	// Logger will be used for logging.
	Logger *logrus.Entry
}

type Exporter struct {
	options   Options
	logger    *logrus.Entry
	bundler   *bundler.Bundler
	extraTags map[string]string
	client    data_receiver.DataReceiverServiceClient
}

// New returns a new Zenoss OpenCensus exporter.
func New(options Options) *Exporter {
	conn, _ := grpc.Dial(ZenossAPIProxyAddress, []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithStatsHandler(&ocgrpc.ClientHandler{}),
		grpc.WithUnaryInterceptor(grpcRetry.UnaryClientInterceptor([]grpcRetry.CallOption{
			grpcRetry.WithMax(5),
			grpcRetry.WithPerRetryTimeout(1 * time.Minute),
			grpcRetry.WithBackoff(grpcRetry.BackoffLinear(200 * time.Millisecond)),
		}...)),
	}...)

	e := &Exporter{
		options:   options,
		logger:    options.Logger,
		extraTags: getKubernetesTags(),
		client:    data_receiver.NewDataReceiverServiceClient(conn),
	}

	e.bundler = bundler.NewBundler((*data_receiver.TaggedMetric)(nil), func(bundle interface{}) {
		e.putTaggedMetrics(bundle.([]*data_receiver.TaggedMetric))
	})

	e.bundler.BundleCountThreshold = 1000
	e.bundler.DelayThreshold = 1 * time.Second

	return e
}

// ExportView exports stats to Zenoss. Satisfies view.Exporter interface.
func (e *Exporter) ExportView(viewData *view.Data) {
	var err error

	timestamp := viewData.End.UnixNano() / 1e6

	addMetric := func(name string, value float64, tags map[string]string) {
		if err = e.bundler.Add(&data_receiver.TaggedMetric{
			Metric:    name,
			Timestamp: timestamp,
			Value:     value,
			Tags:      tags,
		}, 1); err != nil {
			e.logger.WithError(err).Warning("failed to bundle tagged metric")
		}
	}

	for _, viewRow := range viewData.Rows {
		tags := make(map[string]string)

		for _, rowTag := range viewRow.Tags {
			tags[rowTag.Key.Name()] = rowTag.Value
		}

		description := viewData.View.Description
		if description != "" {
			tags[DescriptionField] = description
		}

		units := viewData.View.Measure.Unit()
		if units != "" {
			tags[UnitsField] = units
		}

		// Copy extraTags. Overwrite metric tags of the same name.
		for k, v := range e.extraTags {
			tags[k] = v
		}

		// Set Source. Overwrite metric tag of the same name.
		if e.options.Source != "" {
			tags[SourceField] = e.options.Source
		}

		// SourceType. Overwrite metric tag of the same name.
		tags[SourceTypeField] = DefaultSourceType

		switch rowData := viewRow.Data.(type) {
		case *view.CountData:
			addMetric(viewData.View.Name, float64(rowData.Value), tags)
		case *view.SumData:
			addMetric(viewData.View.Name, rowData.Value, tags)
		case *view.LastValueData:
			addMetric(viewData.View.Name, rowData.Value, tags)
		case *view.DistributionData:
			params := map[string]float64{
				"count": float64(rowData.Count),
				"min":   rowData.Min,
				"max":   rowData.Max,
				"mean":  rowData.Mean,
				"ss":    rowData.SumOfSquaredDev,
			}

			for suffix, value := range params {
				addMetric(
					fmt.Sprintf("%s/%s", viewData.View.Name, suffix),
					value,
					tags)
			}
		}
	}
}

// Flush waits for exported data to be sent. Call before program termination.
func (e *Exporter) Flush() {
	e.bundler.Flush()
}

// putTaggedMetrics is called by Exporter.bundler.
func (e *Exporter) putTaggedMetrics(taggedMetrics []*data_receiver.TaggedMetric) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	if r, err := e.client.PutMetrics(ctx, &data_receiver.Metrics{
		TaggedMetrics: taggedMetrics,
	}); err != nil {
		e.logger.WithError(err).WithFields(logrus.Fields{
			"failed":    len(taggedMetrics),
			"succeeded": 0,
		}).Warning("failed to sent tagged metrics")
	} else {
		if r.GetFailed() > 0 {
			e.logger.WithFields(logrus.Fields{
				"failed":    r.GetFailed(),
				"succeeded": r.GetSucceeded(),
			}).Warning("failed to send some tagged metrics")
		} else {
			e.logger.WithFields(logrus.Fields{
				"failed":    0,
				"succeeded": r.GetSucceeded(),
			}).Debug("sent tagged metrics")
		}
	}
}

// getKubernetesTags returns a tags map containing all Kubernetes tags we could discover.
func getKubernetesTags() map[string]string {
	tags := make(map[string]string)

	if k8sCluster := os.Getenv(K8sClusterEnv); k8sCluster != "" {
		tags[K8sClusterTagKey] = k8sCluster
	}

	if k8sNamespace := os.Getenv(K8sNamespaceEnv); k8sNamespace != "" {
		tags[K8sNamespaceTagKey] = k8sNamespace
	}

	if k8sPod := os.Getenv(K8sPodEnv); k8sPod != "" {
		tags[K8sPodTagKey] = k8sPod
	}

	return tags
}
