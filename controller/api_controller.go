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
	c.HTML(http.StatusOK, "apis.tmpl", gin.H{
		"alert": false,
		"error": false,
		"apis":  apis.API,
		"total": apis.Total,
		"err":   "",
	})
	return
}

func (self *APIController) Get(c *gin.Context) {
	apiName := c.Param("apiName")

	if !api.Exists(self.APIService, apiName) {
		c.HTML(http.StatusNotFound, "api.tmpl", gin.H{
			"error":   true,
			"message": fmt.Sprintf("API %s does not exist.", apiName),
		})

		return
	}

	apiDetail, err := api.Get(self.APIService, apiName)
	fmt.Fprintf(os.Stdout, "%+v\n", apiDetail)
	if apiDetail == nil {
		fmt.Fprintf(os.Stderr, "Err: %+v\nTarget api name: %+v\n", err, apiName)

		c.HTML(http.StatusInternalServerError, "api.tmpl", gin.H{
			"error":   true,
			"message": "Failed to list app URLs.",
		})

		return
	}

	c.HTML(http.StatusOK, "api.tmpl", gin.H{
		"error":     false,
		"apiDetail": apiDetail,
	})
	return
}

func (self *APIController) New(c *gin.Context) {
	c.HTML(http.StatusOK, "new-api.tmpl", gin.H{
		"alert":   false,
		"error":   false,
		"message": "",
	})
	return
}

func (self *APIController) Create(c *gin.Context) {
	c.Redirect(http.StatusSeeOther, "/apis")
}
