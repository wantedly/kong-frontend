package controller

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/koudaiii/kong-oauth-token-generator/kong"
	"github.com/koudaiii/kong-oauth-token-generator/model/api"
)

type APIController struct {
	*kong.APIService
}

func NewAPIController(client *kong.Client) *APIController {
	return &APIController{client.APIService}
}

func (self *APIController) Index(c *gin.Context) {
	apis, err := api.List(self.APIService)
	fmt.Fprintf(os.Stderr, "%+v\n", apis)
	fmt.Fprintf(os.Stderr, "%+v\n", err)
	c.HTML(http.StatusOK, "api.tmpl", gin.H{
		"alert":     false,
		"error":     false,
		"logged_in": true,
		"apis":      apis.API,
		"total":     apis.Total,
		"err":       err,
	})
	return
}
