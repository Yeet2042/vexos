package fiberserver

import "github.com/gofiber/fiber/v3"

type FiberServer interface {
	Start() error
	Close() error
	App() *fiber.App
}
