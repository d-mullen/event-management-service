package zenkit

import (
	"context"

	"go.opencensus.io/tag"
)

var (
	KeyTenant, _ = tag.NewKey(LogTenantField)
	KeyUser, _   = tag.NewKey(LogUserField)
	idTagEnabled = true
)

func DisableIdentityTags() {
	idTagEnabled = false
}
func IdentityTaggedContext(ctx context.Context) context.Context {
	if idTagEnabled {
		identity := ContextTenantIdentity(ctx)
		ctx, _ = tag.New(ctx, tag.Upsert(KeyTenant, identity.Tenant()), tag.Upsert(KeyUser, identity.ID()))
	}
	return ctx
}
