package controller

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/wantedly/kong-oauth-token-generator/kong"
	"github.com/wantedly/kong-oauth-token-generator/model/api"
	"github.com/wantedly/kong-oauth-token-generator/model/oauth2"
	"github.com/wantedly/kong-oauth-token-generator/model/plugin"
)

type PluginController struct {
	Client *kong.Client
}

type newPluginFormBody struct {
	Name string `form:"name" json:"name" binding:"required"`
}

func parsePluginSchema(c *gin.Context, schema *kong.PluginSchema, prefix string, form map[string]interface{}) {
	var (
		value interface{}
		ok    bool
	)
	for key, field := range schema.Fields {
		name := prefix + key
		value, ok = c.GetPostForm(name)
		if !ok {
			if field.Type == "table" {
				tmp := make(map[string]interface{})
				parsePluginSchema(c, &field.Schema, name+".", tmp)
				form[key] = tmp
				continue
			}
			if !field.Required {
				continue
			}
			if field.Type == "boolean" {
				if field.Default == true {
					value = "true"
				} else {
					value = "false"
				}
			} else {
				value = field.Default
			}
		} else if value == "" {
			if !field.Required {
				continue
			}
			value = field.Default
		}
		form[key] = value
	}
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
	plugins, err := plugin.EnabledPlugins(self.Client)
	if err != nil {
		c.HTML(http.StatusBadRequest, "new-plugin.tmpl", gin.H{
			"error":   true,
			"message": "Failed to get enabled plugin list.",
		})
		return
	}
	c.HTML(http.StatusOK, "new-plugin.tmpl", gin.H{
		"alert":   false,
		"error":   false,
		"apiName": apiName,
		"plugins": plugins.EnabledPlugins,
		"message": "",
	})
	return
}

func (self *PluginController) SetConfig(c *gin.Context) {
	apiName := c.Param("apiName")
	consumers, err := oauth2.List(self.Client)
	if err != nil {
		c.HTML(http.StatusBadRequest, "new-plugin.tmpl", gin.H{
			"error":   true,
			"message": "Failed to get consumers.",
		})
		return
	}
	var form newPluginFormBody
	if c.Bind(&form) == nil {
		fmt.Fprintf(os.Stdout, "name %+v\n", form.Name)
	} else {
		c.HTML(http.StatusBadRequest, "new-plugin.tmpl", gin.H{
			"error":   true,
			"message": fmt.Sprintf("Please fix params"),
		})
		return
	}
	schema, _ := plugin.Schema(self.Client, form.Name)
	c.HTML(http.StatusOK, "new-plugin-config.tmpl", gin.H{
		"alert":      false,
		"error":      false,
		"apiName":    apiName,
		"consumers":  consumers.Consumer,
		"pluginName": form.Name,
		"noConsumer": schema.NoConsumer,
		"fields":     schema.Fields,
		"message":    "",
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
	pluginName, _ := c.GetPostForm("plugin_name")
	consumerID, _ := c.GetPostForm("plugin_consumer_id")
	schema, _ := plugin.Schema(self.Client, pluginName)
	form := make(map[string]interface{})
	parsePluginSchema(c, schema, "", form)
	fmt.Fprintf(os.Stderr, "config: %+v\n", form)
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
