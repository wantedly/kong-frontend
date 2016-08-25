package kong

import (
	_ "fmt"
	"net/http"
	"reflect"
	"strings"

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

type PluginConfigList struct {
	OAuth2       OAuth2PluginConfig       `name:"oauth2"`
	RateLimiting RateLimitingPluginConfig `name:"rate-limiting"`
}

type EnabledPlugin struct {
	EnabledPlugins []string `json:"enabled_plugins"`
}

type PluginSchema struct {
	Fields     map[string]PluginSchemaField `json:"fields"`
	NoConsumer bool                         `json:"no_consumer,omitempty"`
}

type PluginSchemaField struct {
	Type     string      `json:"type"`
	Required bool        `json:"required,omitempty"`
	Func     string      `json:"func,omitempty"`
	Default  interface{} `json:"default,omitempty"`
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

func (s *PluginService) GetPluginSchema(name string) (*PluginSchema, *http.Response, error) {
	schema := new(PluginSchema)
	resp, err := s.sling.New().Get(s.config.KongAdminURL + "plugins/schema/" + name).ReceiveSuccess(schema)
	return schema, resp, err
}

func (s *PluginService) Create(params *Plugin, apiName string) (*Plugin, *http.Response, error) {
	plugin := new(Plugin)
	resp, err := s.sling.New().Post(s.config.KongAdminURL + "apis/" + apiName + "/plugins").BodyJSON(params).ReceiveSuccess(plugin)
	return plugin, resp, err
}

func (s *PluginService) Create2(params *GeneratePluginParams, apiName string) (*Plugin, *http.Response, error) {
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

func GetPluginConfig(name string) interface{} {
	instance := &PluginConfigList{}
	types := reflect.TypeOf(*instance)
	for i := 0; i < types.NumField(); i++ {
		field := types.Field(i)
		if name == field.Tag.Get("name") {
			return reflect.New(field.Type).Interface()
		}
	}
	return nil
}

func GetPluginList() map[string]map[string]string {
	result := map[string]map[string]string{}
	instance := &PluginConfigList{}
	types := reflect.TypeOf(*instance)
	config := reflect.ValueOf(*instance)
	for i := 0; i < types.NumField(); i++ {
		field := types.Field(i)
		name := field.Tag.Get("name")
		result[name] = GetFieldTypes(config.Field(i).Interface())
	}
	return result
}

func GetFieldTypes(instance interface{}) map[string]string {
	result := map[string]string{}
	types := reflect.TypeOf(instance)
	values := reflect.ValueOf(instance)
	for i := 0; i < types.NumField(); i++ {
		field := types.Field(i)
		value := values.Field(i).Interface()
		tag := field.Tag.Get("json")
		name := strings.SplitN(tag, ",", 2)[0]
		if value == false {
			result[name] = "checkbox"
		} else if value == 0 {
			result[name] = "number"
		} else {
			result[name] = "text"
		}
	}
	return result
}
