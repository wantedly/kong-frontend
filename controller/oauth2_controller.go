package controller

import (
	"net/http"

	"github.com/koudaiii/kong-oauth-token-generator/model/oauth2"

	"github.com/gin-gonic/gin"
	"github.com/koudaiii/kong-oauth-token-generator/kong"
)

type OAuth2Controller struct {
	Client *kong.Client
}

func NewOAuth2Controller(client *kong.Client) *OAuth2Controller {
	return &OAuth2Controller{client}
}

func (self *OAuth2Controller) Index(c *gin.Context) {
	apis, assignees, err := oauth2.List(self.Client)
	c.HTML(http.StatusOK, "oauth2s.tmpl", gin.H{
		"alert":     false,
		"error":     false,
		"message":   "",
		"apis":      apis.API,
		"assignees": assignees.AssigneesOAuth2,
		"err":       err,
	})
	return
}

func (self *OAuth2Controller) Get(c *gin.Context) {
	return
}

func (self *OAuth2Controller) Create(c *gin.Context) {
	return
}

func (self *OAuth2Controller) Delete(c *gin.Context) {
	return
}
