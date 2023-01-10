package mongodb

import (
	"context"
	"encoding/json"
	"sync/atomic"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/zenoss/event-management-service/internal/instrumentation"
	"github.com/zenoss/event-management-service/metrics"
	"github.com/zenoss/zenkit/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.opencensus.io/stats"
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
		log                         = zenkit.ContextLogger(ctx)
		cursors                     []Cursor
		batchStartId                = ""
		limit, total, cursor_errors int64
		attempts                    = new(uint32)
	)
	ctx, span := instrumentation.StartSpan(ctx, "FindWithRetry")

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

	defer func() {
		span.End()
		stats.Record(ctx, metrics.MCursorErrorCount.M(cursor_errors))
	}()

	filterBytes, err := bson.MarshalExtJSON(filter, true, true)
	if err == nil {
		span.AddAttributes(trace.StringAttribute("mongoFilter", string(filterBytes)))
	}
	findOptsByte, err := json.Marshal(findOpts)
	if err == nil {
		span.AddAttributes(trace.StringAttribute("mongoFindOptions", string(findOptsByte)))
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
			updatedOpts   *options.FindOptions
			count         int64
		)
		if len(batchStartId) > 0 {
			var startKey any
			startKey = batchStartId
			// TODO: Make this pluggable so the start key fits the schema of a given collection
			id, err := primitive.ObjectIDFromHex(batchStartId)
			// if the key is an ObjectID hex
			if err == nil {
				startKey = id
			}
			updatedFilter = append(filter, bson.E{
				Key: "_id",
				Value: bson.D{{
					Key:   "$gt",
					Value: startKey,
				}},
			})
		} else {
			updatedFilter = filter
		}
		if total > 0 {
			opts := options.Find()
			if len(findOpts) > 0 {
				opts = options.MergeFindOptions(findOpts...)
			}
			if opts.Limit != nil && *opts.Limit > 0 && limit-total > 0 {
				opts.SetLimit(limit - total)
			}
			updatedOpts = opts
		} else {
			updatedOpts = options.MergeFindOptions(findOpts...)
		}
		atomic.AddUint32(attempts, 1)
		cursor, err := find(ctx, updatedFilter, updatedOpts)
		if err != nil {
			return errors.Wrap(err, "failed to find documents with retry")
		}
		cursors = append(cursors, cursor)
		for cursor.Next(ctx) {
			if err = cursor.Err(); err != nil {
				break
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
		log.WithFields(logrus.Fields{
			"totalCount":   total,
			"currentCount": count,
		}).Debug("got results during retrying-find operation")
		if err := cursor.Err(); err != nil {
			cursor_errors++
			log.WithField(logrus.ErrorKey, err).Warn("got cursor error")
			if !isRetryableError(err) {
				log.WithField(logrus.ErrorKey, err).Error("failed to find documents with non-retryable error")
				return err
			}
			_ = cursor.Close(ctx)
			continue OuterLoop
		}
		break OuterLoop
	}
	return nil
}
