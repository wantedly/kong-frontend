package main

import (
	"fmt"
	"net/http"
)

type KongError struct {
	Message string `json:"message"`
}

// Client is kong client
type Client struct {
	APIService      *APIService
	ConsumerService *ConsumerService
	PluginService   *PluginService
	Oauth2Service   *Oauth2Service
	// other service endpoints...
}

// NewClient returns a new Client
func NewClient(httpClient *http.Client) *Client {
	return &Client{
		APIService:      NewAPIService(httpClient),
		ConsumerService: NewConsumerService(httpClient),
		PluginService:   NewPluginService(httpClient),
		Oauth2Service:   NewOauth2Service(httpClient),
	}
}

func (e KongError) Error() string {
	return fmt.Sprintf("kong: %v", e.Message)
}

func main() {
	client := NewClient(nil)

	apis, _, _ := client.APIService.List()
	fmt.Printf("APIs:\n%v\n", apis)

	api, _, _ := client.APIService.Get("go-api")
	fmt.Printf("API:\n%v\n", api)

	plugins, _, _ := client.PluginService.List("go-api")
	fmt.Printf("Plugins:\n%v\n", plugins)

	plugin, _, _ := client.PluginService.Get("go-api", "ec04cf37-920d-46fd-adb5-ce33f416d88b")
	fmt.Printf("Plugin:\n%v\n", plugin)

	consumers, _, _ := client.ConsumerService.List()
	fmt.Printf("Consumers:\n%v\n", consumers)

	consumer, _, _ := client.ConsumerService.Get("gokun")
	fmt.Printf("Consumer:\n%v\n", consumer)

	oauth2, _, _ := client.Oauth2Service.List("gokun")
	fmt.Printf("Oauth2:\n%v\n", oauth2)

	targetOauth2Config, _, _ := client.Oauth2Service.Get("gokun", "86e4f18c-00a6-403f-a526-5d8fc3dac95d")
	fmt.Printf("Oauth2:\n%v\n", targetOauth2Config)

	generateAPI := &API{
		Name:             "sakabe",
		UpstreamURL:      "http://koudaiii.com",
		RequestHost:      "test.com",
		StripRequestPath: true,
	}
	api, resp, err := client.APIService.Create(generateAPI)
	fmt.Printf(":\n%v\n%v\n%v\n%v\n", generateAPI, api, resp, err)
}
