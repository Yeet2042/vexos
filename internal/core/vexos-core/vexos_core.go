package vexosservice

import (
	"context"
	"errors"
	"time"

	config "github.com/Yeet2042/vexos/config/vexos-core"
	fiberserver "github.com/Yeet2042/vexos/pkg/fiber-server"
	"github.com/gofiber/fiber/v3"
	"golang.org/x/sync/errgroup"
)

type service struct {
	config *config.VEXOSConfig
	fiber  fiberserver.FiberServer
}

func New(
	config *config.VEXOSConfig,
	fiber fiberserver.FiberServer,
) (*service, error) {
	return &service{
		config: config,
		fiber:  fiber,
	}, nil
}

// Start implements the Service interface.
func (v *service) Start(ctx context.Context) error {
	// create an errgroup with the provided context
	g, _ := errgroup.WithContext(ctx)

	// ----- Start Modules
	// TODO: start modules here and add them to the errgroup, e.g.:
	// g.Go(func() error {
	// 	return module.Start(ctx)
	// })

	// ----- Start HTTP Server
	v.initializeRoutesV1()
	v.fiber.Start()

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

func (v *service) initializeRoutesV1() {
	route := v.fiber.App().Group("/v1")

	// endpoint
	// TODO: add endpoints here, e.g.:
	// route.Get("/example", exampleHandler)

	// health check endpoint
	route.Get("/health", func(ctx fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"status":    "ok",
			"timestamp": time.Now().UTC(),
		})
	})

}
