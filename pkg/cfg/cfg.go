package cfg

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

type validator interface {
	Validate() error
}

func New[T any](path string) (*T, error) {
	if strings.TrimSpace(path) == "" {
		return nil, fmt.Errorf("[pkg/config]: config path is required")
	}
	if filepath.Base(path) != "config.yml" {
		return nil, fmt.Errorf("[pkg/config]: config file must be named config.yml: %s", path)
	}

	v := viper.New()
	v.SetConfigFile(path)

	if err := v.ReadInConfig(); err != nil {
		os.Stdout.WriteString("[pkg/config]: config.yml not found, using environment variables...\n")
	}

	var config T
	if err := v.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("[pkg/config]: failed to unmarshal config: %w", err)
	}

	if validate, ok := any(&config).(validator); ok {
		if err := validate.Validate(); err != nil {
			return nil, fmt.Errorf("[pkg/config]: config validation failed: %w", err)
		}
	}

	return &config, nil
}
