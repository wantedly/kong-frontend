package config

import (
	_ "fmt"

	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

type KongConfiguration struct {
	KongAdminHost string `envconfig:"KONG_HOST" default:"localhost"`
	KongAdminPort string `envconfig:"KONG_PORT" default:"8001"`
	KongAdminURL  string `envconfig:"KONG_URL" default:""`
}

func LoadConfig() (*KongConfiguration, error) {
	var config KongConfiguration
	err := envconfig.Process("", &config)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to load config from envs.")
	}
	if config.KongAdminURL == "" {
		config.KongAdminURL = "http://" + config.KongAdminHost + ":" + config.KongAdminPort + "/"
	}

	return &config, nil
}
