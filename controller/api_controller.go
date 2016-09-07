package controller

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/wantedly/kong-frontend/kong"
	"github.com/wantedly/kong-frontend/model/api"
)

type generateAPIParams struct {
	Name             string `form:"name" json:"name" binding:"required"`
	UpstreamURL      string `form:"upstream_url" json:"upstream_url" binding:"required"`
	RequestPath      string `form:"request_path" json:"request_path" binding:"required"`
	StripRequestPath bool   `form:"strip_request_path" json:"request_path" binding:"omitempty"`
	OAuth2           bool   `form:"oauth2" json:"oauth2" binding:"omitempty"`
}

type APIController struct {
	Client *kong.Client
}

func NewAPIController(client *kong.Client) *APIController {
	return &APIController{client}
}

func (self *APIController) Index(c *gin.Context) {
	apis, err := api.List(self.Client)
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

	if !api.Exists(self.Client, apiName) {
		c.HTML(http.StatusNotFound, "api.tmpl", gin.H{
			"error":   true,
			"message": fmt.Sprintf("API %s does not exist.", apiName),
		})

		return
	}

	apiDetail, enableOAuth2, err := api.Get(self.Client, apiName)
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
		"error":        false,
		"apiDetail":    apiDetail,
		"enableOAuth2": enableOAuth2,
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

func (self *APIController) Delete(c *gin.Context) {
	apiName := c.Param("apiName")
	message, err := api.Delete(self.Client, apiName)
	if err != nil {
		c.HTML(http.StatusBadRequest, "apis.tmpl", gin.H{
			"error":   true,
			"err":     fmt.Sprintf("%s", err),
			"message": fmt.Sprintf("%s", message),
		})
	}

	c.Redirect(http.StatusMovedPermanently, "/apis")

	return
}

func (self *APIController) Create(c *gin.Context) {
	var form generateAPIParams
	enableOAuth2 := false
	if c.Bind(&form) == nil {
		fmt.Fprintf(os.Stdout, "name %+v\n", form.Name)
		fmt.Fprintf(os.Stdout, "upstream_url %+v\n", form.UpstreamURL)
		fmt.Fprintf(os.Stdout, "request_path %+v\n", form.RequestPath)
		fmt.Fprintf(os.Stdout, "strip_request_path %+v\n", form.StripRequestPath)
	} else {
		c.HTML(http.StatusBadRequest, "new-api.tmpl", gin.H{
			"error":   true,
			"message": fmt.Sprintf("Please fix params"),
		})
		return
	}

	if api.Exists(self.Client, form.Name) {
		c.HTML(http.StatusConflict, "new-api.tmpl", gin.H{
			"error":   true,
			"message": fmt.Sprintf("API %s already exist.", form.Name),
		})
		return
	}

	generateAPI := &kong.API{
		Name:             form.Name,
		UpstreamURL:      form.UpstreamURL,
		RequestPath:      form.RequestPath,
		StripRequestPath: form.StripRequestPath,
	}

	generatePlugin := &kong.Plugin{
		Name: "oauth2",
		Config: kong.OAuth2PluginConfig{
			EnableClientCredentials: true,
		},
	}

	createdAPI, createdPlugin, err := api.Create(self.Client, generateAPI, generatePlugin)
	if err != nil {
		c.HTML(http.StatusServiceUnavailable, "new-api.tmpl", gin.H{
			"error":   true,
			"message": fmt.Sprintf("Please Check kong: %s", err),
		})
		return
	}

	if createdPlugin.Name == "oauth2" {
		enableOAuth2 = true
	}

	c.HTML(http.StatusOK, "api.tmpl", gin.H{
		"error":        false,
		"apiDetail":    createdAPI,
		"enableOAuth2": enableOAuth2,
		"message":      fmt.Sprintf("Success"),
	})
	return
}
