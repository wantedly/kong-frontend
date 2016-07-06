package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

type KongConfiguration struct {
	KongAdminURL string `envconfig:"KONG_URL" default:"localhost:8001"`
}

func LoadConfig() (*KongConfiguration, error) {
	var config KongConfiguration
	err := envconfig.Process("", &config)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to load config from envs.")
	}
	format := "KongAdminURL: %s\n"
	_, err = fmt.Printf(format, config.KongAdminURL)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
