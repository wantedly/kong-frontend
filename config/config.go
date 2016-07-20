package config

import (
	_ "fmt"

	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

type KongConfiguration struct {
	KongAdminHost string `envconfig:"host" default:"localhost"`
	KongAdminPort string `envconfig:"port" default:"8001"`
	KongAdminURL  string `envconfig:"url" default:""`
}

const (
	ConfigPrefix = "kong"
)

func LoadConfig() (*KongConfiguration, error) {
	var config KongConfiguration
	err := envconfig.Process(ConfigPrefix, &config)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to load config from envs.")
	}
	if config.KongAdminURL == "" {
		config.KongAdminURL = "http://" + config.KongAdminHost + ":" + config.KongAdminPort + "/"
	}

	return &config, nil
}
