package mongodb

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/zenoss/event-management-service/internal/instrumentation"
	"github.com/zenoss/zenkit/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.opencensus.io/trace"
)

const (
	CursorNotFound = 43
)

type (
	finder func(context.Context, any, ...*options.FindOptions) (Cursor, error)
)

func isRetryableError(err error) bool {
	commandError, ok := err.(mongo.CommandError)
	if !ok || commandError.HasErrorLabel("NetworkError") ||
		commandError.HasErrorLabel("ResumableChangeStreamError") || commandError.Code == CursorNotFound {
		return true
	}
	return false
}

func FindWithRetry[R any](
	ctx context.Context,
	filter bson.D,
	findOpts []*options.FindOptions,
	find finder,
	processResults func(r R) (bool, string, error),
) error {
	var (
		log          = zenkit.ContextLogger(ctx)
		cursors      []Cursor
		batchStartId = ""
		limit, total int64
		attempts     uint32
	)
	ctx, span := instrumentation.StartSpan(ctx, "FindWithRetry")
	defer span.End()
	defer instrumentation.AnnotateSpan(
		"FindWithRetry metrics",
		"FindWithRetry completed",
		span,
		map[string]any{"numAttempts": attempts}, nil)
	defer func() {
		for _, cursor := range cursors {
			if cursor == nil {
				continue
			}
			err := cursor.Close(ctx)
			if err != nil {
				log.WithField(logrus.ErrorKey, err).Warn("failed to close cursor")
			}
		}
	}()
	filterBytes, err := bson.MarshalExtJSON(filter, true, false)
	if err == nil {
		span.AddAttributes(trace.StringAttribute("mongoFilter", string(filterBytes)))
	}
	for _, opt := range findOpts {
		if opt != nil && opt.Limit != nil && *opt.Limit > 0 {
			limit = *opt.Limit
			break
		}
	}
OuterLoop:
	for {
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				return err
			}
			return nil
		default:
			// no op
		}
		var (
			updatedFilter bson.D
			updatedOpts   []*options.FindOptions
			count         int64
		)
		if len(batchStartId) > 0 {
			// TODO: Make this pluggable so the start key fits the schema of a given collection
			updatedFilter = append(filter, bson.E{Key: "_id", Value: bson.E{Key: "$gt", Value: batchStartId}})
		} else {
			updatedFilter = filter
		}
		if limit > 0 && total > 0 {
			updatedOpts = make([]*options.FindOptions, len(findOpts))
			for i, opt := range findOpts {
				updatedOpts[i] = findOpts[i]
				if *opt.Limit > 0 {
					updatedOpts[i].SetLimit(limit - total)
				}
			}
		} else {
			updatedOpts = findOpts
		}
		attempts++
		cursor, err := find(ctx, updatedFilter, updatedOpts...)
		if err != nil {
			return errors.Wrap(err, "failed to find documents with retry")
		}
		cursors = append(cursors, cursor)
		for cursor.Next(ctx) {
			if err = cursor.Err(); err != nil {
				if !isRetryableError(err) {
					log.WithField(logrus.ErrorKey, err).Error("failed to find documents with non-retryable error")
					return err
				}
				continue OuterLoop
			}
			var result R
			if err := cursor.Decode(&result); err != nil {
				log.Errorf("failed to code result: %q", err)
				continue
			}
			ok, resultID, err := processResults(result)
			if err != nil {
				log.WithField(logrus.ErrorKey, err).Error("failed to process findDocuments results")
				return err
			}
			if !ok {
				break OuterLoop
			}
			count++
			total++
			batchStartId = resultID
		}
		if err := cursor.Err(); err != nil {
			log.WithField(logrus.ErrorKey, err).Warn("got cursor error")
			if !isRetryableError(err) {
				log.WithField(logrus.ErrorKey, err).Error("failed to find documents with non-retryable error")
				return err
			}
		} else {
			break OuterLoop
		}
		log.WithFields(logrus.Fields{
			"totalCount":   total,
			"currentCount": count,
		}).Debug("got results during retrying-find operation")
	}
	return nil
}
