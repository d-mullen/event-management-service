package mongo_test

import (
	"context"
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/zenoss/event-management-service/pkg/adapters/datasources/cursors/redis"
	"github.com/zenoss/event-management-service/pkg/adapters/datasources/eventcontext/mongo"
	"github.com/zenoss/event-management-service/pkg/models/event"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	cursorRepoFake struct {
		data map[string]*event.Cursor
	}
	testCaseType3 struct {
		query    *event.Query
		expected struct {
			filter   bson.D
			findOpts *options.FindOptions
			err      error
		}
	}
)

func (fake *cursorRepoFake) New(ctx context.Context, cursor *event.Cursor) (string, error) {
	key := redis.CursorToString(cursor)
	fake.data[key] = cursor
	return key, nil
}

func (fake *cursorRepoFake) Get(ctx context.Context, cursorStr string) (*event.Cursor, error) {
	if cur, ok := fake.data[cursorStr]; ok {
		return cur, nil
	}
	return nil, fmt.Errorf("cursor not found")
}

func (fake *cursorRepoFake) Update(_ context.Context, _ string, _ *event.Cursor) error {
	panic("unimplemented")
}

var _ = Describe("Pagination Unit-Tests", func() {
	var (
		pager mongo.Pager
	)
	var _ = DescribeTable("skip-limit Unit Tests", func(testCase testCaseType3) {
		cursorRepoFake := &cursorRepoFake{
			data: map[string]*event.Cursor{
				"q1": {
					Query: event.Query{
						PageInput: &event.PageInput{
							Limit: 1,
						},
					},
					Config: map[string]any{
						"skipLimitConfig": &mongo.SkipLimitConfig{
							PageSize: 1,
							Offset:   10,
						},
					},
				},
			},
		}
		pager = mongo.NewSkipLimitPager()
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		actualFilter, actualFindOpts, err := pager.GetPaginationQuery(ctx, testCase.query, cursorRepoFake)
		if testCase.expected.err != nil {
			Expect(err).To(MatchError(testCase.expected.err.Error()))
		}
		_ = actualFilter
		_ = actualFindOpts
		// Expect(mustMarshal(actualFilter)).To(MatchJSON(mustMarshal(testCase.expected.filter)))
		Expect(actualFindOpts).To(BeEquivalentTo(testCase.expected.findOpts))
	},
		Entry("case 1: no cursor provided; query is valid", testCaseType3{
			query: &event.Query{
				PageInput: &event.PageInput{
					Limit: 1,
				},
			},
			expected: struct {
				filter   primitive.D
				findOpts *options.FindOptions
				err      error
			}{
				findOpts: options.Find().SetLimit(2), // we send limit +1 to mongo so we figure if there's a next result
			},
		}),
		Entry("case 2: query == nil; should emit invalid argument error", testCaseType3{
			query: nil,
			expected: struct {
				filter   primitive.D
				findOpts *options.FindOptions
				err      error
			}{
				err: fmt.Errorf("invalid argument: nil query"),
			},
		}),
		Entry("case 3: query != nil; cursor not-empty; should advance offset forward", testCaseType3{
			query: &event.Query{
				PageInput: &event.PageInput{
					Cursor: "q1",
					Limit:  1,
				},
			},
			expected: struct {
				filter   primitive.D
				findOpts *options.FindOptions
				err      error
			}{
				findOpts: options.Find().SetLimit(2).SetSkip(10),
			},
		}),
		Entry("case 4: query != nil; cursor non-empty; page direction is backward; should decrease offset", testCaseType3{
			query: &event.Query{
				PageInput: &event.PageInput{
					Cursor:    "q1",
					Limit:     2,
					Direction: event.PageDirectionBackward,
				},
			},
			expected: struct {
				filter   primitive.D
				findOpts *options.FindOptions
				err      error
			}{
				findOpts: options.Find().SetLimit(2).SetSkip(9),
			},
		}),
		Entry("case 5: query != nil; cursor non-empty; page direction is backward; offset - limit < 0; should not set findOpt.skip", testCaseType3{
			query: &event.Query{
				PageInput: &event.PageInput{
					Cursor:    "q1",
					Limit:     12,
					Direction: event.PageDirectionBackward,
				},
			},
			expected: struct {
				filter   primitive.D
				findOpts *options.FindOptions
				err      error
			}{
				// TODO: currently the original pageSize/Limit takes precedence over the req's pageInput.Limit
				// so skip should be updated based on the original limit
				// Change this expectation if override the original page size is allowed
				findOpts: options.Find().SetLimit(2).SetSkip(9),
			},
		}),
	)
})
