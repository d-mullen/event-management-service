package mongo_test

import (
	"encoding/json"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/zenoss/event-management-service/pkg/adapters/datasources/eventcontext/mongo"
	"github.com/zenoss/event-management-service/pkg/models/event"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func mustMarshal(val any) []byte {
	if val == nil {
		return nil
	}
	b, err := json.Marshal(val)
	if err != nil {
		panic(err)
	}
	return b
}

var _ = Describe("Mongo Filter Helpers Tests", func() {
	// It("...", func() {
	// 	Expect(true).To(Equal(true))

	// 	q1 := &mongo.MongoQuery{
	// 		Filter:   primitive.D{},
	// 		FindOpts: options.Find().SetSort(&bson.D{{"f1", event.SortOrderAscending}}),
	// 	}

	// 	paginatedQuery := mongo.GeneratePaginationQuery(nil, q1)
	// 	Expect(paginatedQuery).ToNot(BeNil())
	// 	fmt.Fprintf(GinkgoWriter, "got paginatedQuery(%#v) #1: %#v\n", q1, paginatedQuery)
	// 	paginatedQuery = mongo.GeneratePaginationQuery(&bson.D{{"_id", "e1"}, {"f1", 1}}, q1)
	// 	Expect(paginatedQuery).ToNot(BeNil())
	// 	fmt.Fprintf(GinkgoWriter, "got paginatedQuery(%#v) #2: %#v\n", q1, paginatedQuery)
	// 	nextKey := mongo.NextKey("f1", []*bson.M{{"_id": "e1", "f1": 1}, {"_id": "e2", "f1": 2}})
	// 	Expect(nextKey).ToNot(BeNil())
	// 	fmt.Fprintf(GinkgoWriter, "got nextKey: %#v\n", nextKey)
	// 	paginatedQuery2 := mongo.GeneratePaginationQuery(nextKey, q1)
	// 	Expect(paginatedQuery2).ToNot(BeNil())
	// 	fmt.Fprintf(GinkgoWriter, "got paginatedQuery(%#v) #3: %#v\n", q1, paginatedQuery2)

	// 	nextKey = mongo.NextKey("", []*bson.M{{"_id": "e1", "f1": 1}, {"_id": "e2", "f1": 2}, {"_id": "e3", "f1": 3}})
	// 	Expect(nextKey).ToNot(BeNil())
	// 	fmt.Fprintf(GinkgoWriter, "got nextKey: %#v\n", nextKey)
	// })
	type (
		testCase struct {
			sortField string
			query     *bson.D
			sort      *bson.D
			nextKey   *bson.D
			items     []*bson.M
			expected  struct {
				query   *bson.D
				nextKey *bson.D
			}
		}
	)
	DescribeTable(
		"GeneratePaginationQuery Table-driven Tests",
		func(tc testCase) {
			query, sort, nextKey := tc.query, tc.sort, tc.nextKey
			sortField, items, expected := tc.sortField, tc.items, tc.expected

			actualQuery := mongo.GeneratePaginationQuery(query, sort, nextKey)
			Expect(mustMarshal(actualQuery)).To(MatchJSON(mustMarshal(expected.query)))
			actualNextKey := mongo.NextKey(sortField, items)
			Expect(mustMarshal(actualNextKey)).To(MatchJSON(mustMarshal(expected.nextKey)))
		},
		Entry("case 1", testCase{
			sortField: "",
			query:     nil,
			sort:      nil,
			nextKey:   nil,
			items:     []*primitive.M{},
			expected: struct {
				query   *primitive.D
				nextKey *primitive.D
			}{
				query:   nil,
				nextKey: nil,
			},
		}),
		Entry("case 2: non-nil query should return the same query", testCase{
			sortField: "",
			query:     &bson.D{{"f1", 1}},
			sort:      nil,
			nextKey:   nil,
			items:     []*primitive.M{},
			expected: struct {
				query   *primitive.D
				nextKey *primitive.D
			}{
				query:   &bson.D{{"f1", 1}},
				nextKey: nil,
			},
		}),
		Entry("case 3: (non-nil query, non-empty items) query == paginatedQuery; nextKey == items[-1]", testCase{
			sortField: "",
			query:     &bson.D{{"f1", 1}},
			sort:      nil,
			nextKey:   nil,
			items: []*primitive.M{
				{"_id": 1},
			},
			expected: struct {
				query   *primitive.D
				nextKey *primitive.D
			}{
				query:   &bson.D{{"f1", 1}},
				nextKey: &bson.D{{"_id", 1}},
			},
		}),
		Entry("case 4: (non-nil query, non-empty items) query == paginatedQuery; nextKey == items[-1]", testCase{
			sortField: "",
			query:     &bson.D{{"f1", 1}},
			sort:      nil,
			nextKey:   nil,
			items: []*primitive.M{
				{"_id": 1},
				{"_id": 2},
			},
			expected: struct {
				query   *primitive.D
				nextKey *primitive.D
			}{
				query:   &bson.D{{"f1", 1}},
				nextKey: &bson.D{{"_id", 2}},
			},
		}),
		Entry("case 5: (non-nil query, non-empty items) query == paginatedQuery; nextKey == items[-1]", testCase{
			sortField: "f1",
			query:     &bson.D{{"f1", 1}},
			sort:      nil,
			nextKey:   nil,
			items: []*primitive.M{
				{"_id": 1, "f1": 1},
			},
			expected: struct {
				query   *primitive.D
				nextKey *primitive.D
			}{
				query:   &bson.D{{"f1", 1}},
				nextKey: &bson.D{{"_id", 1}, {"f1", 1}},
			},
		}),
		Entry("case 6: (query != nil, sort == nil, nextKey != nil) query == paginatedQuery; nextKey == items[-1]", testCase{
			sortField: "",
			query:     &bson.D{{"f1", 1}},
			sort:      nil,
			nextKey:   &bson.D{{"_id", 1}},
			items: []*primitive.M{
				{"_id": 1, "f1": 1},
				{"_id": 2, "f1": 1},
			},
			expected: struct {
				query   *primitive.D
				nextKey *primitive.D
			}{
				query:   &bson.D{{"f1", 1}, {"_id", bson.D{{mongo.OpGreaterThan, 1}}}},
				nextKey: &bson.D{{"_id", 2}},
			},
		}),
		Entry("case 7: (query != nil, sort != nil, nextKey != nil) query == paginatedQuery; nextKey == items[-1]", testCase{
			sortField: "f1",
			query:     &bson.D{{"f1", 1}},
			sort:      &bson.D{{"f1", event.SortOrderAscending}},
			nextKey:   &bson.D{{"_id", 1}, {"f1", 1}},
			items: []*primitive.M{
				{"_id": 1, "f1": 1},
				{"_id": 2, "f1": 1},
			},
			expected: struct {
				query   *primitive.D
				nextKey *primitive.D
			}{
				query: &bson.D{
					{"f1", 1},
					{
						mongo.OpOr,
						bson.D{
							{"f1", bson.D{{mongo.OpGreaterThan, 1}}},
							{mongo.OpAnd, bson.D{
								{"f1", 1},
								{"_id", bson.D{{mongo.OpGreaterThan, 1}}},
							}},
						},
					},
				},
				nextKey: &bson.D{{"_id", 2}, {"f1", 1}},
			},
		}),
	)

})
