package vexosservice

import (
	"context"
	"errors"
	"time"

	config "github.com/Yeet2042/vexos/config/vexos-core"
	"golang.org/x/sync/errgroup"
)

type v1 struct {
	config *config.VEXOSConfig
}

func NewV1(
	config *config.VEXOSConfig,
) (V1, error) {
	return &v1{
		config: config,
	}, nil
}

// Start implements [V1].
func (v *v1) Start(ctx context.Context) error {
	// create an errgroup with the provided context
	g, _ := errgroup.WithContext(ctx)

	// ----- Start Modules

	// receive output from g.Wait()
	done := make(chan error, 1)
	go func() {
		done <- g.Wait()
	}()

	// wait until either g.Wait() returns or the context is canceled
	<-ctx.Done()

	// start a timer to enforce a maximum shutdown time
	shutdownTimer := time.NewTimer(10 * time.Second)
	defer shutdownTimer.Stop()

	select {
	case err := <-done:
		// all module have completed or one has returned an error
		return err
	case <-shutdownTimer.C:
		// timeout waiting for modules to shut down
		return errors.New("[Timeout] Graceful shutdown failed: timed out after 10 seconds")
	}
}
