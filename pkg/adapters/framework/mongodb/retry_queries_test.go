package mongodb_test

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
	"github.com/zenoss/event-management-service/pkg/adapters/framework/mongodb"
	"github.com/zenoss/event-management-service/pkg/adapters/framework/mongodb/mocks"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	doc struct {
		ID string `bson:"_id"`
	}
	mockedCursor struct {
		mocks.Cursor
		actualCursor *mongo.Cursor
	}
)

func isRetryableError(err error) bool {
	commandError, ok := err.(mongo.CommandError)
	if !ok || commandError.HasErrorLabel("NetworkError") ||
		commandError.HasErrorLabel("ResumableChangeStreamError") || commandError.Code == mongodb.CursorNotFound {
		return true
	}
	return false
}

func (_m *mockedCursor) Next(ctx context.Context) bool {
	return _m.actualCursor.Next(ctx)
}

func (_m *mockedCursor) Decode(dest any) error {
	return _m.actualCursor.Decode(dest)
}

func (_m *mockedCursor) Close(ctx context.Context) error {
	return _m.actualCursor.Close(ctx)
}

var _ = Describe("FindWithRetryTest", func() {
	var (
		ctx              context.Context
		cancel           context.CancelFunc
		ctxMockType      = mock.AnythingOfType("*context.valueCtx")
		bsonDMockType    = mock.AnythingOfType("primitive.D")
		findOptsMockType = mock.AnythingOfType("*options.FindOptions")
	)
	BeforeEach(func() {
		ctx, cancel = context.WithCancel(context.Background())
	})
	AfterEach(func() {
		cancel()
	})
	When("mongo returns no errors", func() {
		var (
			backingCursor    *mongo.Cursor
			mockedCollection *mocks.Collection
			mockedCursor     *mocks.Cursor
		)
		JustBeforeEach(func() {
			var (
				err         error
				nextRetArg1 *bool = new(bool)
			)
			docs := []any{doc{"e1:1"}, doc{"e1:2"}, doc{"e2:1"}}
			backingCursor, err = mongo.NewCursorFromDocuments(docs, nil, nil)
			Expect(err).ToNot(HaveOccurred())
			mockedCursor = mocks.NewCursor(GinkgoT())
			mockedCursor.On("Next", ctxMockType).Run(func(args mock.Arguments) {
				ctxArg, ok := args.Get(0).(context.Context)
				Expect(ok).To(BeTrue())
				*nextRetArg1 = backingCursor.Next(ctxArg)
			}).Times(3).Return(true)
			mockedCursor.On("Next", ctxMockType).Once().Return(false)
			mockedCursor.On("Decode", mock.AnythingOfType("*mongodb_test.doc")).Run(func(args mock.Arguments) {
				dest := args.Get(0)
				err := backingCursor.Decode(dest)
				Expect(err).ToNot(HaveOccurred())
			}).Return(nil)
			mockedCursor.On("Err").Return(nil)
			mockedCursor.On("Close", ctxMockType).Return(nil)

			mockedCollection = mocks.NewCollection(GinkgoT())
			mockedCollection.On("Find", ctxMockType, bsonDMockType, findOptsMockType).
				Return(mockedCursor, nil)
		})
		It("should retrieve all query results", func() {
			results := make([]*doc, 0)
			err := mongodb.FindWithRetry(
				ctx,
				bson.D{},
				[]*options.FindOptions{options.Find().SetLimit(2)},
				mockedCollection.Find,
				func(result doc) (bool, string, error) {
					results = append(results, &result)
					return true, result.ID, nil
				},
			)
			Expect(err).ToNot(HaveOccurred())
			Expect(results).To(HaveLen(3))
		})
	})
	When("network errors occur", func() {
		var (
			backingCursor, acutalCursor2 *mongo.Cursor
			mockedCollection             *mocks.Collection
			cursorMock                   *mockedCursor
		)
		JustBeforeEach(func() {
			var (
				err error
			)
			docs := []any{doc{"e1:1"}, doc{"e1:2"}, doc{"e2:1"}}
			backingCursor, err = mongo.NewCursorFromDocuments(docs[:2], nil, nil)
			Expect(err).ToNot(HaveOccurred())
			cursorMock = &mockedCursor{
				Cursor:       *mocks.NewCursor(GinkgoT()),
				actualCursor: backingCursor,
			}
			mongoTestErr := mongo.CommandError{Code: mongodb.CursorNotFound, Message: "test error"}
			Expect(isRetryableError(mongoTestErr)).To(BeTrue())
			cursorMock.On("Err").Twice().Return(nil)
			cursorMock.On("Err").Once().Return(mongoTestErr)
			cursorMock.On("Err").Return(nil)

			mockedCollection = mocks.NewCollection(GinkgoT())
			mockedCollection.On("Find", ctxMockType, bsonDMockType, findOptsMockType).
				Return(cursorMock, nil).Once()
			acutalCursor2, err = mongo.NewCursorFromDocuments(docs[2:], nil, nil)
			Expect(err).ToNot(HaveOccurred())
			cursorMock2 := &mockedCursor{
				Cursor:       *mocks.NewCursor(GinkgoT()),
				actualCursor: acutalCursor2,
			}
			cursorMock2.On("Err").Return(nil)
			mockedCollection.On("Find", ctxMockType, bsonDMockType, findOptsMockType).
				Once().Return(cursorMock2, nil)
		})
		It("should retry and retrieve all query results", func() {
			results := make([]*doc, 0)
			err := mongodb.FindWithRetry(
				ctx,
				bson.D{},
				[]*options.FindOptions{options.Find().SetLimit(3)},
				mockedCollection.Find,
				func(result doc) (bool, string, error) {
					results = append(results, &result)
					return true, result.ID, nil
				},
			)
			Expect(err).ToNot(HaveOccurred())
			Expect(results).To(HaveLen(3))
		})
	})
})
