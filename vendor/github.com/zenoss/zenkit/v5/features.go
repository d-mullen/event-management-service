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
	// the name of the feature the controls whether we can set values in the header request (for automated tests).
	AllowFeatureHeaders = "AllowFeatureHeaders"
)

var (
	featureClient feature_flags.FeatureFlagsClient

	ErrMetadataNotFound    = errors.New("no metadata found on the context")
	ErrFeatureFlagNotFound = errors.New("feature flag not found in the context metadata")
)

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

// Perform a quick check to see if features have been specified in the context.  If so, we'll make the
// heavier call to see if this context is allowed to set the feature states (AllowHeaderFeatures).
// Returns the context-specified value + true if (a) the feature was set in the context and (b) the context
// is allowed to specify the feature.
func featureInContext(ctx context.Context, feature string) (bool, bool) {
	enabled, err := GetFeatureFlagFromContext(ctx, feature)
	if err != nil {
		// No feature flag specified in the context; we don't need to check if the
		// AllowHeaderFeatures flag is enabled on this context.
		return false, false
	}

	// We have the flag in the context; if we're allowed to specify flags in the request headers, we'll
	// use this value.
	req := &feature_flags.FeatureEnabledRequest{
		Feature: AllowFeatureHeaders,
		Default: false,
	}

	response, err := featureClient.FeatureEnabled(ctx, req)
	if err != nil {
		ContextLogger(ctx).WithError(err).Errorf("Feature found in the context, but unable to query the feature flags server")
		return false, false
	}

	if !response.Enabled {
		// This context isn't allowed to specify a feature flag, but it did.  We'll log this as a warning so we can
		// see it in StackDriver.
		log := ContextLogger(ctx).WithField("feature", feature)
		log.Warn("context specified the value for a feature in headers, but isn't allowed to do so")
		return false, false
	}

	// The feature has been set in a header, and the context is allowed to do so. We'll return the specified value.
	return enabled, true
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

	// If the feature was configured from the context (and is allowed to do so), return the specified value
	// without going to the feature provider.
	if enabled, found := featureInContext(ctx, feature); found {
		return enabled
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
	feature = strings.ToLower("x-feature-" + feature)
	return metadata.AppendToOutgoingContext(ctx, feature, strings.ToLower(strconv.FormatBool(value)))
}

func GetFeatureFlagFromContext(ctx context.Context, feature string) (bool, error) {
	featureKey := strings.ToLower("x-feature-" + feature)
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
