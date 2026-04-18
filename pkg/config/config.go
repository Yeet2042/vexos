package config

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

var (
	once           sync.Once
	configInstance any
	configErr      error
)

type validator interface {
	Validate() error
}

func NewConfigInstance[T any](path string) (*T, error) {
	once.Do(func() {
		if strings.TrimSpace(path) == "" {
			configErr = fmt.Errorf("[pkg/config]: config path is required")
			return
		}
		if filepath.Base(path) != "config.yml" {
			configErr = fmt.Errorf("[pkg/config]: config file must be named config.yml: %s", path)
			return
		}

		v := viper.New()
		v.SetConfigFile(path)

		if err := v.ReadInConfig(); err != nil {
			log.Println("[pkg/config]: config.yml not found, using environment variables...")
		}

		var config T
		if err := v.Unmarshal(&config); err != nil {
			configErr = fmt.Errorf("[pkg/config]: failed to unmarshal config: %w", err)
			return
		}

		if validate, ok := any(&config).(validator); ok {
			if err := validate.Validate(); err != nil {
				configErr = fmt.Errorf("[pkg/config]: config validation failed: %w", err)
				return
			}
		}

		configInstance = &config
	})

	if configErr != nil {
		return nil, configErr
	}

	return configInstance.(*T), nil
}
