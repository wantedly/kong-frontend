package controller

import (
	"github.com/koudaiii/kong-oauth-token-generator/config"
)

type ApplicationController struct {
	config *config.KongConfiguration
}

func NewApplicationController(config *config.KongConfiguration) *ApplicationController {
	return &ApplicationController{
		config: config,
	}
}
