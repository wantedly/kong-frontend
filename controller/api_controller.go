package controller

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/koudaiii/kong-oauth-token-generator/kong"
	"github.com/koudaiii/kong-oauth-token-generator/model/api"
)

type Params struct {
	Name             string `form:"name" json:"name" binding:"required"`
	UpstreamURL      string `form:"upstream_url" json:"upstream_url" binding:"required"`
	RequestPath      string `form:"request_path" json:"request_path" binding:"required"`
	StripRequestPath bool   `form:"strip_request_path" json:"request_path" binding:"omitempty"`
	OAuth2           bool   `form:"oauth2" json:"oauth2" binding:"omitempty"`
}

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
	var form Params
	if c.Bind(&form) == nil {
		fmt.Fprintf(os.Stdout, "name %+v\n", form.Name)
		fmt.Fprintf(os.Stdout, "upstream_url %+v\n", form.UpstreamURL)
		fmt.Fprintf(os.Stdout, "request_path %+v\n", form.RequestPath)
		fmt.Fprintf(os.Stdout, "strip_request_path %+v\n", form.StripRequestPath)
		fmt.Fprintf(os.Stdout, "oauth2 %+v\n", form.OAuth2)
	} else {
		c.HTML(http.StatusBadRequest, "new-api.tmpl", gin.H{
			"error":   true,
			"message": fmt.Sprintf("Please fix params"),
		})
		return
	}

	if api.Exists(self.APIService, form.Name) {
		c.HTML(http.StatusConflict, "new-api.tmpl", gin.H{
			"error":   true,
			"message": fmt.Sprintf("API %s already exist.", form.Name),
		})
		return
	}

	c.Redirect(http.StatusSeeOther, "/apis")

}
