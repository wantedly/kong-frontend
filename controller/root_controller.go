package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/koudaiii/kong-oauth-token-generator/config"
	"github.com/koudaiii/kong-oauth-token-generator/model"
)

type RootController struct {
	*ApplicationController
}

func NewRootController(config *config.KongConfiguration) *RootController {
	return &RootController{NewApplicationController(config)}
}

func (s *RootController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"alert":     false,
		"error":     false,
		"logged_in": true,
		"message":   "",
	})
	return
}
