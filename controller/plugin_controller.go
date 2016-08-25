package controller

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/wantedly/kong-oauth-token-generator/kong"
	"github.com/wantedly/kong-oauth-token-generator/model/api"
	"github.com/wantedly/kong-oauth-token-generator/model/plugin"
)

type PluginController struct {
	Client *kong.Client
}

func NewPluginController(client *kong.Client) *PluginController {
	return &PluginController{client}
}

func (self *PluginController) Index(c *gin.Context) {
	apiName := c.Param("apiName")

	if !api.Exists(self.Client, apiName) {
		c.HTML(http.StatusNotFound, "api.tmpl", gin.H{
			"error":   true,
			"message": fmt.Sprintf("API %s does not exist.", apiName),
		})

		return
	}

	plugins, err := plugin.List(self.Client, apiName)
	fmt.Fprintf(os.Stderr, "%+v\n", plugins)
	fmt.Fprintf(os.Stderr, "%+v\n", err)
	c.HTML(http.StatusOK, "plugins.tmpl", gin.H{
		"alert":   false,
		"error":   false,
		"plugins": plugins.Plugin,
		"total":   plugins.Total,
		"apiName": apiName,
		"err":     "",
	})
	return
}

func (self *PluginController) Get(c *gin.Context) {
	apiName := c.Param("apiName")
	pluginID := c.Param("pluginID")

	if !api.Exists(self.Client, apiName) {
		c.HTML(http.StatusNotFound, "api.tmpl", gin.H{
			"error":   true,
			"message": fmt.Sprintf("API %s does not exist.", apiName),
		})

		return
	}

	if !plugin.Exists(self.Client, apiName, pluginID) {
		c.HTML(http.StatusNotFound, "plugin.tmpl", gin.H{
			"error":   true,
			"message": fmt.Sprintf("Plugin %s does not exist.", pluginID),
		})

		return
	}

	pluginDetail, err := plugin.Get(self.Client, apiName, pluginID)
	fmt.Fprintf(os.Stdout, "%+v\n", pluginDetail)
	if pluginDetail == nil {
		fmt.Fprintf(os.Stderr, "Err: %+v\nTarget api name: %+v, plugin name: %+v\n", err, apiName, pluginID)

		c.HTML(http.StatusInternalServerError, "plugin.tmpl", gin.H{
			"error":   true,
			"message": "Failed to list app URLs.",
		})

		return
	}

	c.HTML(http.StatusOK, "plugin.tmpl", gin.H{
		"error":        false,
		"pluginDetail": pluginDetail,
		"apiName":      apiName,
	})
	return
}

func (self *PluginController) New(c *gin.Context) {
	apiName := c.Param("apiName")
	c.HTML(http.StatusOK, "new-plugin.tmpl", gin.H{
		"alert":   false,
		"error":   false,
		"apiName": apiName,
		"plugins": kong.GetPluginList(),
		"message": "",
	})
	return
}

func (self *PluginController) Config(c *gin.Context) {
	apiName := c.Param("apiName")
	var form kong.GeneratePluginParams
	if c.Bind(&form) == nil {
		fmt.Fprintf(os.Stdout, "name %+v\n", form.Name)
		fmt.Fprintf(os.Stdout, "consumer_id %+v\n", form.ConsumerID)
	} else {
		c.HTML(http.StatusBadRequest, "new-plugin.tmpl", gin.H{
			"error":   true,
			"message": fmt.Sprintf("Please fix params"),
		})
		return
	}
	c.HTML(http.StatusOK, "new-plugin-config.tmpl", gin.H{
		"alert":        false,
		"error":        false,
		"apiName":      apiName,
		"pluginName":   form.Name,
		"consumerID":   form.ConsumerID,
		"pluginConfig": kong.GetPluginList()[form.Name],
		"message":      "",
	})
	return
}

func (self *PluginController) Delete(c *gin.Context) {
	apiName := c.Param("apiName")
	pluginID := c.Param("pluginID")
	message, err := plugin.Delete(self.Client, apiName, pluginID)
	if err != nil {
		c.HTML(http.StatusBadRequest, "plugins.tmpl", gin.H{
			"error":   true,
			"err":     fmt.Sprintf("%s", err),
			"message": fmt.Sprintf("%s", message),
		})
	}

	c.Redirect(http.StatusMovedPermanently, "/apis/"+apiName+"/plugins")

	return
}

func (self *PluginController) Create(c *gin.Context) {
	apiName := c.Param("apiName")
	pluginName, _ := c.GetQuery("name")
	consumerID, _ := c.GetQuery("consumer_id")
	form := kong.GetPluginConfig(pluginName)
	fmt.Fprintf(os.Stdout, "config %#v\n", form)
	if c.Bind(form) == nil {
		fmt.Fprintf(os.Stdout, "name %+v\n", pluginName)
		fmt.Fprintf(os.Stdout, "consumer_id %+v\n", consumerID)
		fmt.Fprintf(os.Stdout, "config %+v\n", form)
	} else {
		c.HTML(http.StatusBadRequest, "new-plugin.tmpl", gin.H{
			"error":   true,
			"message": fmt.Sprintf("Please fix params"),
		})
		return
	}

	params := kong.GeneratePluginParams{
		Name:       pluginName,
		ConsumerID: consumerID,
		Config:     form,
	}
	createdPlugin, err := plugin.Create(self.Client, apiName, &params)
	if err != nil {
		c.HTML(http.StatusServiceUnavailable, "new-plugin.tmpl", gin.H{
			"error":   true,
			"message": fmt.Sprintf("Please Check kong: %s", err),
		})
		return
	}

	c.HTML(http.StatusOK, "plugin.tmpl", gin.H{
		"error":        false,
		"pluginDetail": createdPlugin,
		"apiName":      apiName,
		"message":      fmt.Sprintf("Success"),
	})
	return
}
