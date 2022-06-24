package zenkit

import (
	"context"
	"strings"

	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"github.com/spf13/viper"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const AuthHeaderScheme = "bearer"

func DevIdentity(ctx context.Context) (context.Context, error) {
	//Treat Tenant ACME as legacy v1 tenant
	tenant := viper.GetString(AuthDevTenantConfig)
	user := viper.GetString(AuthDevUserConfig)
	email := viper.GetString(AuthDevEmailConfig)
	connection := viper.GetString(AuthDevConnectionConfig)
	scopes := viper.GetStringSlice(AuthDevScopesConfig)
	groups := viper.GetStringSlice(AuthDevGroupsConfig)
	roles := viper.GetStringSlice(AuthDevRolesConfig)
	clientid := viper.GetString(AuthDevClientIDConfig)
	subject := viper.GetString(AuthDevSubjectConfig)
	tenantid := viper.GetString(AuthDevTenantIDConfig)
	tenantname := viper.GetString(AuthDevTenantConfig)
	ident := &devTenantIdentity{user, email, scopes, tenant, connection, groups, roles, clientid, subject, tenantid, tenantname}
	addIdentityFieldsToTags(ctx, ident)
	addRequestIdToTags(ctx)
	return WithTenantIdentity(ctx, ident), nil
}

func UnverifiedIdentity(ctx context.Context) (context.Context, error) {
	// Set up the context
	meta := metautils.ExtractIncoming(ctx)
	ctx = meta.ToOutgoing(ctx)

	// Extract the identity from the metadata
	raw, err := grpc_auth.AuthFromMD(ctx, AuthHeaderScheme)
	if err != nil {
		return nil, wrapUnauthenticated(err)
	}
	ident, err := NewAuth0TenantIdentity(raw)
	if err != nil {
		return nil, wrapUnauthenticated(err)
	}

	addIdentityFieldsToTags(ctx, ident)
	addRequestIdToTags(ctx)

	// Add the identity to the context
	return WithTenantIdentity(ctx, ident), nil
}

func addIdentityFieldsToTags(ctx context.Context, ident TenantIdentity) {
	grpc_ctxtags.Extract(ctx).Set(LogTenantField, ident.Tenant()).Set(LogUserField, ident.ID())
}

func addRequestIdToTags(ctx context.Context) {
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return
	}
	requestId := meta[LogRequestIdField]

	if requestId != nil {
		grpc_ctxtags.Extract(ctx).Set(LogRequestIdField, strings.Join(requestId, " "))
	}
}

func wrapUnauthenticated(err error) error {
	stat := status.New(codes.Unauthenticated, err.Error())
	return stat.Err()
}
