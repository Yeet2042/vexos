package config

import (
	"fmt"
)

type (
	VEXOSConfig struct {
		Server *Server
	}

	Server struct {
		Environment string
		DNS         string
		Port        int
		FrontendURL string
	}
)

func (c *VEXOSConfig) Validate() error {
	if c == nil {
		return fmt.Errorf("[config/vexos]: nil config")
	}

	return nil
}
