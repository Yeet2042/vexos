package config

import (
	"fmt"
)

type (
	VEXOSConfig struct {
		Server   *Server
		Database *Database
	}

	Server struct {
		Environment string
		DNS         string
		Port        int
		FrontendURL string
	}

	Database struct {
		Path      string
		Namespace string
		Database  string
	}
)

func (c *VEXOSConfig) Validate() error {
	if c == nil {
		return fmt.Errorf("[config/vexos]: nil config")
	}

	return nil
}
