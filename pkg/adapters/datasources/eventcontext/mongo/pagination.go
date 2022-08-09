package mongo

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/zenoss/event-management-service/internal"
	"github.com/zenoss/event-management-service/pkg/models/event"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	skipLimitPageConfigKey = "skipLimitConfig"
)

type (
	Pager interface {
		GetPaginationQuery(ctx context.Context, query *event.Query, cursorRepo event.CursorRepository) (bson.D, *options.FindOptions, error)
		NextPageCursor(ctx context.Context, direction event.PageDirection, currrentCursor *event.Cursor, lastPageResults any) (*event.Cursor, error)
	}
	skipLimitPager  struct{}
	SkipLimitConfig struct {
		Offset int64 `json:"offset"`
	}
)

var _ Pager = new(skipLimitPager)

func NewSkipLimitPager() *skipLimitPager {
	return &skipLimitPager{}
}

func (pager *skipLimitPager) GetPaginationQuery(ctx context.Context, query *event.Query, cursorRepo event.CursorRepository) (bson.D, *options.FindOptions, error) {
	if query == nil {
		return nil, nil, fmt.Errorf("invalid argument: nil query")
	}
	if cursorRepo == nil {
		return nil, nil, fmt.Errorf("invalid argument: nil cursor repository")
	}
	if pi := query.PageInput; pi != nil && len(pi.Cursor) > 0 {
		queryCursor, err := cursorRepo.Get(ctx, pi.Cursor)
		if err != nil {
			return nil, nil, errors.Wrap(err, "failed to process cursor")
		}
		filters, findOpt, err := QueryToFindArguments(&queryCursor.Query)
		if err != nil {
			return nil, nil, errors.Wrap(err, "failed to convert query to mongo find filter and options")
		}
		if queryCursor.Config != nil {
			if cfgAny, ok := queryCursor.Config[skipLimitPageConfigKey]; !ok {
				return nil, nil, fmt.Errorf("failed to extract cursor info")
			} else {
				cfgBytes, err := json.Marshal(cfgAny)
				if err != nil {
					return nil, nil, errors.Wrap(err, "failed to process skip-limit config")
				}
				cfg := SkipLimitConfig{}
				if err := json.Unmarshal(cfgBytes, &cfg); err != nil {
					return nil, nil, errors.Wrap(err, "failed to unmarshal skip-limit config")
				}
				if query.PageInput.Direction == event.PageDirectionBackward {
					limit := int64(query.PageInput.Limit)
					if findOpt.Limit != nil {
						limit = *findOpt.Limit
					}
					skip := cfg.Offset - limit + 1
					if skip > 0 {
						findOpt.SetSkip(skip)
					}
				} else {
					findOpt.SetSkip(cfg.Offset)
				}
			}
		}
		return filters, findOpt, nil
	} else {
		filters, findOpt, err := QueryToFindArguments(query)
		if err != nil {
			return nil, nil, errors.Wrap(err, "failed to convert query to mongo find filter and options")
		}
		return filters, findOpt, nil
	}
}

func (pager *skipLimitPager) NextPageCursor(ctx context.Context, direction event.PageDirection, currrentCursor *event.Cursor, lastPageResult any) (*event.Cursor, error) {
	if currrentCursor == nil {
		return nil, fmt.Errorf("failed to get next page cursor: invalid argument: nil cursor input")
	}
	query, config := currrentCursor.Query, currrentCursor.Config
	if err := query.Validate(); err != nil {
		return nil, errors.Wrap(err, "failed to get next page cursor: invalid argument: nil query")
	}
	if config == nil {
		return nil, fmt.Errorf("failed to get next page cursor: invalid argument: nil config")
	}
	var cfgStruct *SkipLimitConfig
	if cfgStructAny, ok := config[skipLimitPageConfigKey]; !ok {
		cfgStruct = &SkipLimitConfig{}
	} else {
		cfgBytes, err := json.Marshal(cfgStructAny)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal config to []byte")
		}
		if err := json.Unmarshal(cfgBytes, &cfgStruct); err != nil {
			return nil, fmt.Errorf("failed to unmarshal config bytes")
		}
	}
	resultsSlice, err := internal.AnyToSlice(lastPageResult)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get next page cursor: invalid argument")
	}
	dir := int64(1)
	if direction == event.PageDirectionBackward {
		dir = -1
	}
	cfgStruct.Offset = cfgStruct.Offset + dir*int64(len(resultsSlice))
	config[skipLimitPageConfigKey] = cfgStruct
	newID := currrentCursor.ID
	if len(newID) == 0 {
		newID = uuid.New().String()
	}
	updatedCursor := &event.Cursor{
		ID:     newID,
		Query:  currrentCursor.Query,
		Config: config,
	}
	return updatedCursor, nil
}

func UpsertCursor[R any](
	ctx context.Context,
	cursorRepo event.CursorRepository,
	pager Pager,
	query *event.Query,
	pageResults []R) (string, error) {

	var (
		cursorInput *event.Cursor
		direction   event.PageDirection
	)
	if query == nil {
		return "", fmt.Errorf("failed to create or update cursor: invalid argument: nil query")
	}
	if query.PageInput != nil && len(query.PageInput.Cursor) > 0 {
		if cur, err := cursorRepo.Get(ctx, query.PageInput.Cursor); err != nil {
			return "", errors.Wrap(err, "failed to get cursor")
		} else {
			cursorInput = cur
			direction = query.PageInput.Direction
		}
	} else {
		cursorInput = &event.Cursor{
			Query:  *query,
			Config: make(map[string]any),
		}
	}
	cursorUpsert, err := pager.NextPageCursor(ctx, direction, cursorInput, pageResults)
	if err != nil {
		return "", errors.Wrap(err, "failed to create next page cursor config")
	}
	cursorStr, err := cursorRepo.New(ctx, cursorUpsert)
	if err != nil {
		return "", errors.Wrap(err, "failed to create or udpate next page cursor")
	}
	return cursorStr, nil
}
