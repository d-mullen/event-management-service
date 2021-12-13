package zenkit

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/zenoss/zing-proto/v11/go/cloud/feature_flags"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"strconv"
	"strings"
)

const (
	// Prefix for HTTP header that overrides feature enabled status.
	featureHeaderPrefix = "x-feature-"
)

var (
	// Ensure TestFeatureFlagsClient implements FeatureFlagsClient interface.
	_ feature_flags.FeatureFlagsClient = (*TestFeatureFlagsClient)(nil)

	featureClient feature_flags.FeatureFlagsClient

	ErrMetadataNotFound    = errors.New("no metadata found on the context")
	ErrFeatureFlagNotFound = errors.New("feature flag not found in the context metadata")
)

// SetFeatureFlagsClient sets the global FeatureFlagsClient to use.
// NOTE: This isn't necessary if you want to use the default.
func SetFeatureFlagsClient(client feature_flags.FeatureFlagsClient) {
	featureClient = client
}

func getZingUnleashClient(ctx context.Context, opts ...grpc.DialOption) feature_flags.FeatureFlagsClient {
	log := ContextLogger(ctx)

	endpoint, err := ServiceAddress("zing-features")
	if err != nil {
		log.WithError(err).Error("zing-features endpoint not available")
		return nil
	}

	log = log.WithFields(logrus.Fields{
		"endpoint": endpoint,
	})

	log.Debug("Dialing zing-features server")
	conn, err := grpc.DialContext(ctx, endpoint, opts...)
	if err != nil {
		panic(err)
	}
	log.Info("Connected to zing-features server")

	return feature_flags.NewFeatureFlagsClient(conn)
}

// A helper function for checking whether a feature is enabled.
func FeatureIsEnabled(ctx context.Context, feature string, defaultValue bool) bool {
	if featureClient == nil {
		client := getZingUnleashClient(context.Background(), grpc.WithInsecure())
		SetFeatureFlagsClient(client)
		if featureClient == nil {
			// The error is already logged; just return the default value.
			return defaultValue
		}
	}

	// We need to ask our feature provider if the feature is on for this context.
	req := &feature_flags.FeatureEnabledRequest{
		Feature: feature,
		Default: defaultValue,
	}

	response, err := featureClient.FeatureEnabled(ctx, req)
	if err != nil {
		ContextLogger(ctx).WithError(err).Errorf("unable to query the feature flags server; returning the default (%t)", defaultValue)
		return defaultValue
	}

	return response.Enabled
}

/*
 * These helper methods are to be used for integration tests. When the zing-features service it brought up with the
 * environment variable: ZING_FEATURES_PROVIDER: "testing", the features service will use an alternate provider that
 * just checks the context for an enabled feature.  This method is *also* used when a context has specified a feature
 * to be enabled via the context (ie: smoke tests request a feature on) and the feature "AllowHeaderFeatures" is
 * enabled for that request (ie: client=smoke test).
 */
func AddFeatureFlagToContext(ctx context.Context, feature string, value bool) context.Context {
	feature = strings.ToLower(featureHeaderPrefix + feature)
	return metadata.AppendToOutgoingContext(ctx, feature, strings.ToLower(strconv.FormatBool(value)))
}

func GetFeatureFlagFromContext(ctx context.Context, feature string) (bool, error) {
	featureKey := strings.ToLower(featureHeaderPrefix + feature)
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		ContextLogger(ctx).Debugf("no metadata found on the context. returning false for feature %s", feature)
		return false, ErrMetadataNotFound
	}
	valueString := meta[featureKey]
	if valueString == nil {
		ContextLogger(ctx).Debugf("metadata was found on the context, but the return value has not been set. returning false for feature %s", feature)
		return false, ErrFeatureFlagNotFound
	}
	value := strings.Join(valueString, " ")
	return value == "true", nil
}

// SetTestFeatureFlagsClient sets the global FeatureFlagsClient to a static map of enabled features.
// NOTE: This should only be used in unit tests.
func SetTestFeatureFlagsClient(features map[string]bool) *TestFeatureFlagsClient {
	testClient := &TestFeatureFlagsClient{Features: features}
	featureClient = testClient
	return testClient
}

type TestFeatureFlagsClient struct {
	Features map[string]bool
}

func (c *TestFeatureFlagsClient) FeatureEnabled(_ context.Context, in *feature_flags.FeatureEnabledRequest, _ ...grpc.CallOption) (*feature_flags.FeatureEnabledResponse, error) {
	if enabled, ok := c.Features[in.Feature]; ok {
		return &feature_flags.FeatureEnabledResponse{Enabled: enabled}, nil
	}

	return &feature_flags.FeatureEnabledResponse{Enabled: in.Default}, nil
}

func (c *TestFeatureFlagsClient) FeaturesEnabled(_ context.Context, in *feature_flags.FeaturesEnabledRequest, _ ...grpc.CallOption) (*feature_flags.FeaturesEnabledResponse, error) {
	// Create the result structure.
	featuresEnabled := make([]*feature_flags.FeatureEnabled, len(in.Features))

	// Extract the values for each of the requested features.
	for index, feature := range in.Features {
		response, _ := c.FeatureEnabled(nil, feature)
		featuresEnabled[index] = &feature_flags.FeatureEnabled{
			Feature: feature.Feature,
			Value:   response.Enabled,
		}
	}

	// And return the results.
	return &feature_flags.FeaturesEnabledResponse{
		Enabled: featuresEnabled,
	}, nil
}
