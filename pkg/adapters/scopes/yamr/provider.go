package yamr

import (
	"context"
	"encoding/json"
	"github.com/zenoss/event-management-service/internal/auth"
	"time"

	"github.com/zenoss/event-management-service/pkg/models/scopes"
	"github.com/zenoss/zenkit/v5"
	yamrPb "github.com/zenoss/zing-proto/v11/go/cloud/yamr"
	"go.opencensus.io/trace"
)

type (
	yamrProvider struct {
		client yamrPb.YamrQueryClient
	}
)

func NewAdapter(cl yamrPb.YamrQueryClient) scopes.EntityScopeProvider {
	y := &yamrProvider{
		client: cl,
	}
	return y
}

func (yp *yamrProvider) GetEntityIDs(ctx context.Context, scopeCursor string) ([]string, error) {
	ctx, span := trace.StartSpan(ctx, "zenoss.cloud/eventQuery/v2/go/yamrProvider.GetEntityIDs")
	defer span.End()
	ident, err := auth.CheckAuth(ctx)
	if err != nil {
		return nil, err
	}
	tenantId := ident.Tenant()
	log := zenkit.ContextLogger(ctx).
		WithField("cursor", scopeCursor)
	results := make([]string, 0)
	defer func() {
		attr := make(map[string]any)
		attr["entityIds"] = results
		attr["cursor"] = scopeCursor
		b, _ := json.Marshal(attr)
		if b != nil {
			span.AddAttributes(trace.StringAttribute("getEntityAttributes", string(b)))
		}
	}()
	now := time.Now()
	resp, err := yp.client.Search(ctx, &yamrPb.SearchRequest{
		Query: &yamrPb.Query{
			Type:    "entity",
			Tenants: []string{tenantId},
			When: &yamrPb.When{
				Lower: &yamrPb.When_Bound{
					Time:      now.Add(-24 * time.Hour).UnixMilli(),
					BoundType: yamrPb.BoundType_CLOSED,
				},
				Upper: &yamrPb.When_Bound{
					Time:      now.UnixMilli(),
					BoundType: yamrPb.BoundType_CLOSED,
				},
			},
			Clause: &yamrPb.Clause{
				Clause: &yamrPb.Clause_WithCursor{
					WithCursor: &yamrPb.WithCursor{
						Cursor: scopeCursor,
					},
				},
			},
		},
		PageInput: &yamrPb.SearchRequest_PageInput{
			Direction: yamrPb.SearchRequest_PageInput_FORWARD,
			Limit:     1000,
		},
	})
	if err != nil {
		log.WithError(err).Error("failed to execute cursor query")
		return nil, err
	}
	for _, result := range resp.GetResults() {
		results = append(results, result.Id.GetId())
	}
	for resp.PageInfo.HasNext {
		resp, err = yp.client.Search(ctx, &yamrPb.SearchRequest{
			PageInput: &yamrPb.SearchRequest_PageInput{
				Cursor:    resp.PageInfo.GetEndCursor(),
				Direction: yamrPb.SearchRequest_PageInput_FORWARD,
				Limit:     10000,
			},
		})
		if err != nil {
			log.WithError(err).Error("failed to execute cursor query")
			return nil, err
		}
		for _, result := range resp.GetResults() {
			results = append(results, result.Id.GetId())
		}
	}

	return results, nil
}
