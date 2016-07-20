package api

import (
	_ "fmt"
	_ "os"

	"github.com/koudaiii/kong-oauth-token-generator/kong"
)

type API struct {
	CreatedAt        int    `json:"created_at,omitempty"`
	ID               string `json:"id,omitempty"`
	Name             string `json:"name"`
	PreserveHost     bool   `json:"preserve_host,omitempty"`
	RequestPath      string `json:"request_path,omitempty"`
	StripRequestPath bool   `json:"strip_request_path,omitempty"`
	UpstreamURL      string `json:"upstream_url,omitempty"`
	RequestHost      string `json:"request_host,omitempty"`
}

func List(self *kong.Client) (*kong.APIs, error) {
	apis, _, err := self.APIService.List()
	return apis, err
}

func Exists(self *kong.Client, apiName string) bool {
	_, resp, _ := self.APIService.Get(apiName)
	if resp == nil {
		return false
	}
	if resp.StatusCode != 404 {
		return true
	}
	return false
}

func Get(self *kong.Client, apiName string) (*kong.API, bool, error) {
	api, _, err := self.APIService.Get(apiName)
	if err != nil {
		return nil, false, err
	}
	plugins, _, err := self.PluginService.List(apiName)
	if err != nil {
		return nil, false, err
	}
	for _, plugin := range plugins.Plugin {
		if plugin.Name == "oauth2" {
			return api, true, err
		}
	}
	return api, false, err
}

func Delete(self *kong.Client, apiName string) (string, error) {
	api, _, err := self.APIService.Get(apiName)
	if err != nil {
		return "", err
	}
	message, _, err := self.APIService.Delete(api.ID)
	return message, err
}

func Create(self *kong.Client, generateAPI *kong.API, generatePlugin *kong.Plugin) (*kong.API, *kong.Plugin, error) {
	api, _, err := self.APIService.Create(generateAPI)
	if err != nil {
		return nil, nil, err
	}
	generatePlugin.ID = api.ID
	plugin, _, err := self.PluginService.Create(generatePlugin, generateAPI.Name)
	if err != nil {
		return nil, nil, err
	}
	return api, plugin, err
}
