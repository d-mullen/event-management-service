package utils

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/zenoss/zenkit/v5"
	"github.com/zenoss/zingo/v4/filters"
)

// Int64Ptr takes a int64 and returns a pointer to it
func Int64Ptr(i int64) *int64 { return &i }

// BoolPtr takes a bool value and returns a pointer to that value
func BoolPtr(x bool) *bool {
	return &x
}

func ValidateIdentity(ctx context.Context) (string, error) {
	identity := zenkit.ContextTenantIdentity(ctx)
	if identity == nil {
		return "", errors.Wrap(ErrInvalidIdentity, "no identity found")
	}
	if identity.Tenant() == "" {
		return "", errors.Wrap(ErrInvalidIdentity, "no tenant found")
	}
	return identity.Tenant(), nil
}

// ErrInvalidIdentity occurs when no identity is found on a context
var ErrInvalidIdentity = errors.New("invalid identity")

type TimeRanger interface {
	StartTS() int64
	EndTS() int64
}

// EnumerateDays returns a list of YYYMMDD strings for all days in the passed in time range
func EnumerateDays(tr TimeRanger) []string {
	a, b := time.Unix(0, tr.StartTS()*1e6), time.Unix(0, tr.EndTS()*1e6)
	dur := b.Sub(a)
	numDays := int(dur.Hours() / 24)
	if numDays == 0 {
		numDays = 1
	}
	results := make([]string, numDays)
	for i := 0; i < numDays; i++ {
		c := a.Add(time.Hour * 24 * time.Duration(i+1))
		results[i] = fmt.Sprintf("%04d%02d%02d", c.Year(), c.Month(), c.Day())
	}
	return results
}

// GetRestrictedSources returns map of sources to which the context has access
func GetRestrictedSources(ctx context.Context) []interface{} {
	sourcesMap := make(map[string]bool)
	entityClauses, err := filters.GetEntityClauses(ctx, filters.Format(0))
	if err != nil {
		zenkit.ContextLogger(ctx).WithError(err).Error("error getting entity clauses")
	} else if len(entityClauses) > 0 {
		for _, clause := range entityClauses {
			sourcesMap[clause.(filters.EntityClause).Source] = true
		}
	}
	sources := make([]interface{}, 0)
	for k := range sourcesMap {
		sources = append(sources, k)
	}
	return sources
}
