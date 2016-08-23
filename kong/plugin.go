package kong

import (
	_ "fmt"
	"net/http"

	"github.com/dghubble/sling"
	"github.com/wantedly/kong-oauth-token-generator/config"
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

type EnabledPlugin struct {
	EnabledPlugins []string `json:"enabled_plugins"`
}

// Services

// PluginService provides methods for creating and reading issues.
type PluginService struct {
	sling  *sling.Sling
	config *config.KongConfiguration
}

// NewPluginService returns a new PluginService.
func NewPluginService(httpClient *http.Client, config *config.KongConfiguration) *PluginService {
	return &PluginService{
		sling:  sling.New().Client(httpClient).Base(config.KongAdminURL + "apis/"),
		config: config,
	}
}

func (s *PluginService) GetEnabledPlugins() (*EnabledPlugin, *http.Response, error) {
	plugins := new(EnabledPlugin)
	resp, err := s.sling.New().Get(s.config.KongAdminURL + "plugins/enabled").ReceiveSuccess(plugins)
	return plugins, resp, err
}

func (s *PluginService) Create(params *Plugin, apiName string) (*Plugin, *http.Response, error) {
	plugin := new(Plugin)
	resp, err := s.sling.New().Post(s.config.KongAdminURL + "apis/" + apiName + "/plugins").BodyJSON(params).ReceiveSuccess(plugin)
	return plugin, resp, err
}

func (s *PluginService) Get(pluginID string, apiName string) (*Plugin, *http.Response, error) {
	plugin := new(Plugin)
	resp, err := s.sling.New().Path(apiName + "/plugins/" + pluginID).ReceiveSuccess(plugin)
	return plugin, resp, err
}

func (s *PluginService) List(apiName string) (*Plugins, *http.Response, error) {
	plugins := new(Plugins)
	resp, err := s.sling.New().Path(apiName + "/plugins/").ReceiveSuccess(plugins)
	return plugins, resp, err
}

func (s *PluginService) Update(params *Plugin, apiName string) (*Plugin, *http.Response, error) {
	api := new(Plugin)
	resp, err := s.sling.New().Patch(s.config.KongAdminURL + "apis/" + apiName + "/plugins/" + params.ID).BodyJSON(params).ReceiveSuccess(api)
	return api, resp, err
}

func (s *PluginService) Delete(pluginID string, apiName string) (string, *http.Response, error) {
	var message string
	resp, err := s.sling.New().Delete(s.config.KongAdminURL + "apis/" + apiName + "/plugins/" + pluginID).ReceiveSuccess(message)
	return message, resp, err
}
