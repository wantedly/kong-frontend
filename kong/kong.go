package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/koudaiii/kong-oauth-token-generator/config"
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
func NewClient(httpClient *http.Client, config *config.KongConfiguration) *Client {
	return &Client{
		APIService:      NewAPIService(httpClient, config),
		ConsumerService: NewConsumerService(httpClient, config),
		PluginService:   NewPluginService(httpClient, config),
		Oauth2Service:   NewOauth2Service(httpClient, config),
	}
}

func (e KongError) Error() string {
	return fmt.Sprintf("kong: %v", e.Message)
}

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}

	client := NewClient(nil, config)

	//	apis, _, _ := client.APIService.List()
	//	fmt.Printf("APIs:\n%v\n", apis)
	//
	//	api, _, _ := client.APIService.Get("go-api")
	//	fmt.Printf("API:\n%v\n", api)

	//	plugins, _, _ := client.PluginService.List("go-api")
	//	fmt.Printf("Plugins:\n%v\n", plugins)

	//	plugin, _, _ := client.PluginService.Get("go-api", "ec04cf37-920d-46fd-adb5-ce33f416d88b")
	//	fmt.Printf("Plugin:\n%v\n", plugin)

	consumers, _, _ := client.ConsumerService.List()
	fmt.Printf("Consumers:\n%v\n", consumers)

	consumer, _, _ := client.ConsumerService.Get("gokun")
	fmt.Printf("Consumer:\n%v\n", consumer)

	generateConsumer := &Consumer{
		Username: "sakabe",
	}

	consumer, resp, err := client.ConsumerService.Create(generateConsumer)
	fmt.Printf("Create Consumer :\n%v\n%v\n%v\n%v\n", generateConsumer, &consumer, resp, err)

	updateConsumer := &Consumer{
		ID:       consumer.ID,
		Username: "sakabeupdate",
	}

	updatedConsumer, resp, err := client.ConsumerService.Update(updateConsumer)
	fmt.Printf("Update Consumer :\n%v\n%v\n%v\n%v\n", updateConsumer, &updatedConsumer, resp, err)

	message, resp, err := client.ConsumerService.Delete(updatedConsumer.ID)
	fmt.Printf("Delete Consumer : \n%v\n%v\n%v\n", message, resp, err)

	//	oauth2, _, _ := client.Oauth2Service.List("gokun")
	//	fmt.Printf("Oauth2:\n%v\n", oauth2)

	//	targetOauth2Config, _, _ := client.Oauth2Service.Get("gokun", "86e4f18c-00a6-403f-a526-5d8fc3dac95d")
	//	fmt.Printf("Oauth2:\n%v\n", targetOauth2Config)

	//	generateAPI := &API{
	//		Name:             "sakabe",
	//		UpstreamURL:      "http://koudaiii.com",
	//		RequestHost:      "test.com",
	//		StripRequestPath: true,
	//	}
	//	api, resp, err := client.APIService.Create(generateAPI)
	//	fmt.Printf("Create API :\n%v\n%v\n%v\n%v\n", generateAPI, &api, resp, err)

	//	updateAPI := &API{
	//		ID:               api.ID,
	//		Name:             "sakabeupdate",
	//		UpstreamURL:      "http://koudaiii.com",
	//		RequestHost:      "test.com",
	//		StripRequestPath: true,
	//	}

	//	updatedAPI, resp, err := client.APIService.Update(updateAPI)
	//	fmt.Printf("Update API :\n%v\n%v\n%v\n%v\n", updateAPI, &updatedAPI, resp, err)

	//	message, resp, err := client.APIService.Delete(updatedAPI.ID)
	//	fmt.Printf("Delete API : \n%v\n%v\n%v\n", message, resp, err)

}
