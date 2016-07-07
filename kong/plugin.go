package main

import (
	_ "fmt"
	"net/http"
	"os"

	"github.com/dghubble/sling"
	"github.com/koudaiii/kong-oauth-token-generator/config"
)

type Plugins struct {
	Plugin []Plugin `json:"data"`
	Total  int      `json:"total"`
}

type Plugin struct {
	APIID     string       `json:"api_id"`
	Config    PluginConfig `json:"config"`
	CreatedAt int          `json:"created_at"`
	Enabled   bool         `json:"enabled"`
	ID        string       `json:"id"`
	Name      string       `json:"name"`
}

type PluginConfig struct {
	AcceptHTTPIfAlreadyTerminated bool   `json:"accept_http_if_already_terminated"`
	EnableAuthorizationCode       bool   `json:"enable_authorization_code"`
	EnableClientCredentials       bool   `json:"enable_client_credentials"`
	EnableImplicitGrant           bool   `json:"enable_implicit_grant"`
	EnablePasswordGrant           bool   `json:"enable_password_grant"`
	HideCredentials               bool   `json:"hide_credentials"`
	MandatoryScope                bool   `json:"mandatory_scope"`
	ProvisionKey                  string `json:"provision_key"`
	TokenExpiration               int    `json:"token_expiration"`
}

// Services

// PluginService provides methods for creating and reading issues.
type PluginService struct {
	sling *sling.Sling
}

// NewPluginService returns a new PluginService.
func NewPluginService(httpClient *http.Client) *PluginService {
	config, err := config.LoadConfig()
	if err != nil {
		os.Exit(1)
	}
	return &PluginService{
		sling: sling.New().Client(httpClient).Base(config.KongAdminURL + "apis/"),
	}
}

func (s *PluginService) Get(apiName string, pluginID string) (Plugin, *http.Response, error) {
	plugin := new(Plugin)
	kongError := new(KongError)
	resp, err := s.sling.New().Path(apiName+"/plugins/"+pluginID).Receive(plugin, kongError)
	if err == nil {
		err = kongError
	}
	return *plugin, resp, err
}

func (s *PluginService) List(apiName string) (Plugins, *http.Response, error) {
	plugins := new(Plugins)
	kongError := new(KongError)
	resp, err := s.sling.New().Path(apiName+"/plugins/").Receive(plugins, kongError)
	if err == nil {
		err = kongError
	}
	return *plugins, resp, err
}
