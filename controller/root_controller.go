package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/koudaiii/kong-oauth-token-generator/kong"
)

type RootController struct {
	APIService             *kong.APIService
	AssigneesOAuth2Service *kong.AssigneesOAuth2Service
}

func NewRootController(client *kong.Client) *RootController {
	return &RootController{
		client.APIService,
		client.AssigneesOAuth2Service,
	}
}

func (r *RootController) Index(c *gin.Context) {
	apis, _, err := r.APIService.List()
	assignees, _, err := r.AssigneesOAuth2Service.List()
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"alert":     false,
		"error":     false,
		"message":   "",
		"apis":      apis,
		"assignees": assignees,
		"err":       err,
	})
	return
}
