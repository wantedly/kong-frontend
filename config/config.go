package config

import (
	_ "fmt"

	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

type KongConfiguration struct {
	KongAdminURL string `envconfig:"http://KONG_HOST:KONG_PORT/" default:"http://localhost:8001/"`
}

func LoadConfig() (*KongConfiguration, error) {
	var config KongConfiguration
	err := envconfig.Process("", &config)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to load config from envs.")
	}

	return &config, nil
}
