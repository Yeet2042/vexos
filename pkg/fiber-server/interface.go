package fiberserver

import (
	"context"

	"github.com/gofiber/fiber/v3"
)

type FiberServer interface {
	Start() error
	Close() error
	ShutdownWithContext(ctx context.Context) error
	App() *fiber.App
}
