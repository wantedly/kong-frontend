package main

import (
	_ "fmt"
	"net/http"

	"github.com/dghubble/sling"
	"github.com/koudaiii/kong-oauth-token-generator/config"
)

type Plugins struct {
	Plugin []Plugin `json:"data,omitempty"`
	Total  int      `json:"total,omitempty"`
}

type Plugin struct {
	APIID     string       `json:"api_id,omitempty"`
	Config    PluginConfig `json:"config,omitempty"`
	CreatedAt int          `json:"created_at,omitempty"`
	Enabled   bool         `json:"enabled,omitempty"`
	ID        string       `json:"id,omitempty"`
	Name      string       `json:"name,omitempty"`
}

type PluginConfig struct {
	AcceptHTTPIfAlreadyTerminated bool   `json:"accept_http_if_already_terminated,omitempty"`
	EnableAuthorizationCode       bool   `json:"enable_authorization_code,omitempty"`
	EnableClientCredentials       bool   `json:"enable_client_credentials,omitempty"`
	EnableImplicitGrant           bool   `json:"enable_implicit_grant,omitempty"`
	EnablePasswordGrant           bool   `json:"enable_password_grant,omitempty"`
	HideCredentials               bool   `json:"hide_credentials,omitempty"`
	MandatoryScope                bool   `json:"mandatory_scope,omitempty"`
	ProvisionKey                  string `json:"provision_key,omitempty"`
	TokenExpiration               int    `json:"token_expiration,omitempty"`
}

// Services

// PluginService provides methods for creating and reading issues.
type PluginService struct {
	sling *sling.Sling
}

// NewPluginService returns a new PluginService.
func NewPluginService(httpClient *http.Client, config *config.KongConfiguration) *PluginService {
	return &PluginService{
		sling: sling.New().Client(httpClient).Base(config.KongAdminURL + "apis/"),
	}
}

func (s *PluginService) Get(apiName string, pluginID string) (*Plugin, *http.Response, error) {
	plugin := new(Plugin)
	kongError := new(KongError)
	resp, err := s.sling.New().Path(apiName+"/plugins/"+pluginID).Receive(plugin, kongError)
	if err == nil {
		err = kongError
	}
	return plugin, resp, err
}

func (s *PluginService) List(apiName string) (*Plugins, *http.Response, error) {
	plugins := new(Plugins)
	kongError := new(KongError)
	resp, err := s.sling.New().Path(apiName+"/plugins/").Receive(plugins, kongError)
	if err == nil {
		err = kongError
	}
	return plugins, resp, err
}
