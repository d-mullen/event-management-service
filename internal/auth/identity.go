package auth

import (
	"context"
	"github.com/zenoss/zenkit/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func CheckAuth(ctx context.Context) (zenkit.TenantIdentity, error) {
	ident := zenkit.ContextTenantIdentity(ctx)
	if ident == nil {
		return nil, status.Error(codes.Unauthenticated, "no identity found on context")
	}
	if len(ident.TenantName()) == 0 && len(ident.Tenant()) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "not tenant claim found identity: %#v", ident)
	}
	return ident, nil
}
