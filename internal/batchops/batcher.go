package batchops

import (
	"context"

	"github.com/pkg/errors"
	"github.com/zenoss/zenkit/v5"
	"go.opencensus.io/trace"
	"golang.org/x/sync/errgroup"
)

// Do takes payload, a slice of E, subdivides it into batches and
// applies the func `doOperation` to that batch. If doOperation(batch)
// succeeds, then its results are passed to the `processResults` func.
// If `doOperation(result)` fails, then Do fails; if the
// callback returns false, then Do will stop processing
// batches.
func Do[E any, R any](
	ctx context.Context,
	batchSize int,
	payload []E,
	doOperation func(ctx context.Context, batch []E) (R, error),
	processResults func(ctx context.Context, result R) (bool, error)) error {
	for i := 0; i < len(payload); i += batchSize {
		j := i + batchSize
		if j > len(payload) {
			j = len(payload)
		}
		batch := payload[i:j]
		result, err := doOperation(ctx, batch)
		if err != nil {
			return err
		}
		ok, err := processResults(ctx, result)
		if err != nil {
			return err
		}
		if !ok {
			break
		}

	}
	return nil
}

// DoConcurrently takes payload, a slice of E, subdivides it into batches and
// concurrently applies the func `doOperation` to that batch. If doOperation(batch)
// succeeds, then its results are passed to the `processResults` func.
// If `doOperation(result)` fails, then DoConcurrently fails; if the
// callback returns false, then DoConcurrently will stop processing
// batches.
func DoConcurrently[E any, R any](
	ctx context.Context,
	batchSize, numRoutines int,
	payload []E,
	doOperation func(ctx context.Context, batch []E) (R, error),
	processResults func(ctx context.Context, result R) (bool, error)) error {
	ctx, span := trace.StartSpan(ctx, "DoConcurrently")
	defer span.End()
	if batchSize == 0 {
		return errors.New("invalid batch size")
	}
	doOpWithSpan := func(batch []E) (R, error) {
		_, childSpan := trace.StartSpan(ctx, "DoConcurrently.doOperation")
		defer childSpan.End()
		return doOperation(ctx, batch)
	}
	numBatches := len(payload)
	if len(payload)%batchSize > 0 {
		numBatches += 1
	}
	queue := make(chan []E, numBatches)
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
				ok, err := processResults(ctx, result)
				if err != nil {
					log.WithError(err).Error("failed to process result during batch operation")
					return err
				}
				if !ok {
					//TODO: implement mechanism to stop all goroutines in group
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
