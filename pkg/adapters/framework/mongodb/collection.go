package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	Database interface {
		Collection(name string, opts ...*options.CollectionOptions) Collection
	}
	Collection interface {
		Find(ctx context.Context, filter any, opts ...*options.FindOptions) (Cursor, error)
		Aggregate(ctx context.Context, pipeline any, opts ...*options.AggregateOptions) (Cursor, error)
		InsertMany(ctx context.Context, documents []interface{},
			opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error)
	}
	Cursor interface {
		ID() int64
		Next(ctx context.Context) bool
		TryNext(ctx context.Context) bool
		Decode(val any) error
		Err() error
		Close(ctx context.Context) error
		All(ctx context.Context, results any) error
		RemainingBatchLength() int
	}
)

type collectionHelper struct {
	*mongo.Collection
}

func (c *collectionHelper) Find(ctx context.Context, filter any, opts ...*options.FindOptions) (Cursor, error) {
	if cur, err := c.Collection.Find(ctx, filter, opts...); err != nil {
		return nil, err
	} else {
		return cur, nil
	}
}
func (c *collectionHelper) Aggregate(ctx context.Context, pipeline any, opts ...*options.AggregateOptions) (Cursor, error) {
	if cur, err := c.Collection.Aggregate(ctx, pipeline, opts...); err != nil {
		return nil, err
	} else {
		return cur, nil
	}
}

type databaseHelper struct {
	*mongo.Database
}

func NewDatabaseHelper(db *mongo.Database) Database {
	return &databaseHelper{Database: db}
}

func (db *databaseHelper) Collection(name string, opts ...*options.CollectionOptions) Collection {
	collection := db.Database.Collection(name, opts...)
	return &collectionHelper{Collection: collection}
}
