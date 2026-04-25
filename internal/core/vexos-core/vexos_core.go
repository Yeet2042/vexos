package vexosservice

import (
	"context"
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
	v.initializeRoutesV1()

	// create an errgroup with the provided context
	g, gCtx := errgroup.WithContext(ctx)

	// ----- Start Modules
	g.Go(func() error {
		return v.fiber.Start()
	})

	// receive output from g.Wait()
	done := make(chan error, 1)
	go func() {
		done <- g.Wait()
	}()

	g.Go(func() error {
		<-gCtx.Done()

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var sg errgroup.Group

		sg.Go(func() error {
			return v.fiber.ShutdownWithContext(shutdownCtx)
		})

		return sg.Wait()
	})

	return g.Wait()
}

func (v *service) initializeRoutesV1() {
	route := v.fiber.App().Group("/v1")

	// health check endpoint
	route.Get("/health", func(ctx fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"status":    "ok",
			"timestamp": time.Now().UTC(),
		})
	})

	// endpoint
	// TODO: add endpoints here, e.g.:
	// route.Get("/example", exampleHandler)
}
