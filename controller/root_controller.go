package controller

import (
	"net/http"

	"github.com/koudaiii/kong-oauth-token-generator/model/root"

	"github.com/gin-gonic/gin"
	"github.com/koudaiii/kong-oauth-token-generator/kong"
)

type RootController struct {
	Client *kong.Client
}

func NewRootController(client *kong.Client) *RootController {
	return &RootController{client}
}

func (self *RootController) Index(c *gin.Context) {
	apis, assignees, err := root.List(self.Client)
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"alert":     false,
		"error":     false,
		"message":   "",
		"apis":      apis.API,
		"assignees": assignees.AssigneesOAuth2,
		"err":       err,
	})
	return
}
