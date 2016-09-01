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
	APIID     string      `json:"api_id,omitempty"`
	Config    interface{} `json:"config,omitempty"`
	CreatedAt int         `json:"created_at,omitempty"`
	Enabled   bool        `json:"enabled,omitempty"`
	ID        string      `json:"id,omitempty"`
	Name      string      `json:"name,omitempty"`
}

type EnabledPlugin struct {
	EnabledPlugins []string `json:"enabled_plugins"`
}

type PluginSchema struct {
	Fields     map[string]*PluginSchemaField `json:"fields"`
	NoConsumer bool                          `json:"no_consumer,omitempty"`
}

type PluginSchemaField struct {
	Type     string       `json:"type"`
	Required bool         `json:"required,omitempty"`
	Func     string       `json:"func,omitempty"`
	Default  interface{}  `json:"default,omitempty"`
	Schema   PluginSchema `json:"schema,omitempty"`
	Name     string       `json:"name"`
}

// Services

// PluginService provides methods for creating and reading issues.
type PluginService struct {
	sling  *sling.Sling
	config *config.KongConfiguration
}

type GeneratePluginParams struct {
	Name       string      `form:"name" json:"name" binding:"required"`
	ConsumerID string      `form:"consumer_id" json:"consumer_id,omitempty" binding:"omitempty"`
	Config     interface{} `form:"config" json:"config" binding:"omitempty"`
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

func (schema *PluginSchema) setPluginSchemaName(prefix string) {
	for key, field := range schema.Fields {
		field.Name = prefix + key
		if field.Type == "table" {
			field.Schema.setPluginSchemaName(field.Name + ".")
		}
	}
}

func (s *PluginService) GetPluginSchema(name string) (*PluginSchema, *http.Response, error) {
	schema := new(PluginSchema)
	resp, err := s.sling.New().Get(s.config.KongAdminURL + "plugins/schema/" + name).ReceiveSuccess(schema)
	schema.setPluginSchemaName("")
	return schema, resp, err
}

func (s *PluginService) CreateOAuth(params *Plugin, apiName string) (*Plugin, *http.Response, error) {
	plugin := new(Plugin)
	resp, err := s.sling.New().Post(s.config.KongAdminURL + "apis/" + apiName + "/plugins").BodyJSON(params).ReceiveSuccess(plugin)
	return plugin, resp, err
}

func (s *PluginService) Create(params *GeneratePluginParams, apiName string) (*Plugin, *http.Response, error) {
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
