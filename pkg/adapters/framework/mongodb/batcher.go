package mongodb

import (
	"context"

	"github.com/pkg/errors"
	"github.com/zenoss/zenkit/v5"
	"go.opencensus.io/trace"
	"golang.org/x/sync/errgroup"
)

func SerialBatcher[E any, R any](
	batchSize int,
	payload []E,
	doOperation func(batch []E) (R, error),
	processResults func(result R) (bool, error)) error {
	for i := 0; i < len(payload); i += batchSize {
		j := i + batchSize
		if j > len(payload) {
			j = len(payload)
		}
		batch := payload[i:j]
		result, err := doOperation(batch)
		if err != nil {
			return err
		}
		ok, err := processResults(result)
		if err != nil {
			return err
		}
		if !ok {
			break
		}

	}
	return nil
}

func ConcurrentBatcher[E any, R any](
	ctx context.Context,
	batchSize, numRoutines int,
	payload []E,
	doOperation func(batch []E) (R, error),
	processResults func(result R) (bool, error)) error {
	ctx, span := trace.StartSpan(ctx, "ConcurrentBatcher")
	defer span.End()
	if batchSize == 0 {
		return errors.New("invalid batch size")
	}
	doOpWithSpan := func(batch []E) (R, error) {
		_, childSpan := trace.StartSpan(ctx, "ConcurrentBatcher.doOperation")
		defer childSpan.End()
		return doOperation(batch)
	}
	queue := make(chan []E, len(payload)/batchSize+len(payload)%batchSize)
	for i := 0; i < len(payload); i += batchSize {
		j := i + batchSize
		if j > len(payload) {
			j = len(payload)
		}
		queue <- payload[i:j]
	}
	close(queue)
	errGroup, gCtx := errgroup.WithContext(ctx)
	for g := 0; g < numRoutines; g++ {
		errGroup.Go(func() error {
			log := zenkit.ContextLogger(gCtx)
			for batch := range queue {
				result, err := doOpWithSpan(batch)
				if err != nil {
					log.WithError(err).Error("failed to do concurrent batch operation")
					return err
				}
				ok, err := processResults(result)
				if err != nil {
					log.WithError(err).Error("failed to process result during batch operation")
					return err
				}
				if !ok {
					break
				}
			}
			return nil
		})
	}
	if err := errGroup.Wait(); err != nil {
		return err
	}
	return nil
}
