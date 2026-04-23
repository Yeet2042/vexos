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

	// Validate fields here
	if c.Server == nil {
		return fmt.Errorf("[config/vexos]: server is required")
	} else {
		if c.Server.Environment == "" {
			return fmt.Errorf("[config/vexos]: server environment is required")
		}
		if c.Server.DNS == "" {
			return fmt.Errorf("[config/vexos]: server DNS is required")
		}
		if c.Server.Port <= 0 {
			return fmt.Errorf("[config/vexos]: server port must be greater than 0")
		}
		if c.Server.FrontendURL == "" {
			return fmt.Errorf("[config/vexos]: server frontend URL is required")
		}
	}

	return nil
}
