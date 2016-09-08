package plugin

import (
	_ "fmt"
	_ "os"

	"github.com/wantedly/kong-frontend/kong"
)

type Plugin struct {
	APIID     string      `json:"api_id,omitempty"`
	Config    interface{} `json:"config,omitempty"`
	CreatedAt int         `json:"created_at,omitempty"`
	Enabled   bool        `json:"enabled,omitempty"`
	ID        string      `json:"id,omitempty"`
	Name      string      `json:"name,omitempty"`
}

func EnabledPlugins(self *kong.Client) (*kong.EnabledPlugin, error) {
	plugins, _, err := self.PluginService.GetEnabledPlugins()
	return plugins, err
}

func Schema(self *kong.Client, name string) (*kong.PluginSchema, error) {
	schema, _, err := self.PluginService.GetPluginSchema(name)
	return schema, err
}

func List(self *kong.Client, apiName string) (*kong.Plugins, error) {
	plugins, _, err := self.PluginService.List(apiName)
	return plugins, err
}

func Exists(self *kong.Client, apiName, pluginID string) bool {
	_, resp, _ := self.PluginService.Get(pluginID, apiName)
	if resp.StatusCode != 404 {
		return true
	}
	return false
}

func Get(self *kong.Client, apiName, pluginID string) (*kong.Plugin, error) {
	plugin, _, err := self.PluginService.Get(pluginID, apiName)
	if err != nil {
		return nil, err
	}
	return plugin, err
}

func Delete(self *kong.Client, apiName, pluginID string) (string, error) {
	plugin, _, err := self.PluginService.Get(pluginID, apiName)
	if err != nil {
		return "", err
	}
	message, _, err := self.PluginService.Delete(plugin.ID, apiName)
	return message, err
}

func Create(self *kong.Client, apiName string, generatePlugin *kong.GeneratePluginParams) (*kong.Plugin, error) {
	plugin, _, err := self.PluginService.Create(generatePlugin, apiName)
	if err != nil {
		return nil, err
	}
	return plugin, err
}
