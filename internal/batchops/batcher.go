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
	processResults func(ctx context.Context, result R) (bool, error),
) error {
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
	processResults func(ctx context.Context, result R) (bool, error),
) error {
	return doConcurrently(ctx, false, batchSize, numRoutines, payload, doOperation, processResults)
}

func DoConcurrentlyInOrder[E any, R any](
	ctx context.Context,
	batchSize, numRoutines int,
	payload []E,
	doOperation func(ctx context.Context, batch []E) (R, error),
	processResults func(ctx context.Context, result R) (bool, error),
) error {
	return doConcurrently(ctx, true, batchSize, numRoutines, payload, doOperation, processResults)
}

type batch[E any] struct {
	index int
	data  []E
}

func doConcurrently[E any, R any](
	ctx context.Context,
	inOrder bool,
	batchSize, numRoutines int,
	payload []E,
	doOperation func(ctx context.Context, batch []E) (R, error),
	processResults func(ctx context.Context, result R) (bool, error),
) error {
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
	queue := make(chan batch[E], numBatches)
	idx := 0
	for i := 0; i < len(payload); i += batchSize {
		j := i + batchSize
		if j > len(payload) {
			j = len(payload)
		}
		queue <- batch[E]{
			data:  payload[i:j],
			index: idx,
		}
		idx++
	}
	close(queue)
	errGroup, gCtx := errgroup.WithContext(ctx)
	gCtx, groupCancel := context.WithCancel(gCtx)
	defer groupCancel()
	log := zenkit.ContextLogger(gCtx)
	combinedResults := make([]R, numBatches)
	for g := 0; g < numRoutines; g++ {
		errGroup.Go(func() error {
			for batch := range queue {
				result, err := doOpWithSpan(batch.data)
				if err != nil {
					log.WithError(err).Error("failed to do concurrent batch operation")
					return err
				}
				if !inOrder {
					ok, err := processResults(ctx, result)
					if err != nil {
						log.WithError(err).Error("failed to process result during batch operation")
						return err
					}
					if !ok {
						// cancel context so other goroutines in group will stop
						groupCancel()
						break
					}
				} else {
					combinedResults[batch.index] = result
				}
			}
			return nil
		})
	}
	if err := errGroup.Wait(); err != nil {
		if errors.Unwrap(err) != context.Canceled {
			return err
		}
	}
	if inOrder {
		for _, result := range combinedResults {
			ok, err := processResults(ctx, result)
			if err != nil {
				log.WithError(err).Error("failed to process result during batch operation")
				return err
			}
			if !ok {
				break
			}
		}
	}
	return nil
}
