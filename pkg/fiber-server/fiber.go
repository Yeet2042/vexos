package fiberserver

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
)

type FiberConfig struct {
	Port int
}

type fiberServer struct {
	app  *fiber.App
	port int
}

func New(cfg *FiberConfig) (FiberServer, error) {
	if cfg == nil || cfg.Port == 0 {
		return nil, fmt.Errorf("[pkg/fiber-server]: config is nil or port is invalid")
	}

	return &fiberServer{
		app:  fiber.New(),
		port: cfg.Port,
	}, nil
}

func (f *fiberServer) App() *fiber.App {
	return f.app
}

func (f *fiberServer) Start() error {
	address := fmt.Sprintf(":%d", f.port)
	return f.app.Listen(address)
}

func (f *fiberServer) Close() error {
	return f.app.Shutdown()
}
