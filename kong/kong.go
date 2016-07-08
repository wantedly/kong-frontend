package kong

import (
	"fmt"
	"net/http"

	"github.com/koudaiii/kong-oauth-token-generator/config"
)

type KongError struct {
	Message string `json:"message"`
}

// Client is kong client
type Client struct {
	APIService             *APIService
	ConsumerService        *ConsumerService
	PluginService          *PluginService
	OAuth2ConfigService    *OAuth2ConfigService
	AssigneesOAuth2Service *AssigneesOAuth2Service
	// other service endpoints...
}

// NewClient returns a new Client
func NewClient(httpClient *http.Client, config *config.KongConfiguration) *Client {
	return &Client{
		APIService:             NewAPIService(httpClient, config),
		ConsumerService:        NewConsumerService(httpClient, config),
		PluginService:          NewPluginService(httpClient, config),
		OAuth2ConfigService:    NewOAuth2ConfigService(httpClient, config),
		AssigneesOAuth2Service: NewAssigneesOAuth2Service(httpClient, config),
	}
}

func (e KongError) Error() string {
	return fmt.Sprintf("kong: %v", e.Message)
}

// func main() {
//	apis, _, _ := client.APIService.List()
//	fmt.Printf("APIs:\n%v\n", apis)
//
//	api, _, _ := client.APIService.Get("go-api")
//	fmt.Printf("API:\n%v\n", api)

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

//	plugins, _, _ := client.PluginService.List("go-api")
//	fmt.Printf("Plugins:\n%v\n", plugins)

//	plugin, _, _ := client.PluginService.Get("ec04cf37-920d-46fd-adb5-ce33f416d88b", "go-api")
//	fmt.Printf("Plugin:\n%v\n", plugin)

//	enablePlugins, resp, err := client.PluginService.GetEnabledPlugins()
//	fmt.Printf("Enable Plugins :\n%v\n%v\n%v\n", enablePlugins, resp, err)

//	generatePlugin := &Plugin{
//		Name: "oauth2",
//		Config: PluginConfig{
//			EnableClientCredentials: true,
//		},
//	}
//	plugin, resp, err := client.PluginService.Create(generatePlugin, "mockbin")
//	fmt.Printf("Create Plugin :\n%v\n%v\n%v\n%v\n", generatePlugin, plugin, resp, err)

//	updatePlugin := &Plugin{
//		ID:   plugin.ID,
//		Name: "oauth2",
//		Config: PluginConfig{
//			EnableClientCredentials: true,
//		},
//	}

//	updatedPlugin, resp, err := client.PluginService.Update(updatePlugin, "mockbin")
//	fmt.Printf("Update Plugin :\n%v\n%v\n%v\n%v\n", updatePlugin, updatedPlugin, resp, err)
//	fmt.Printf("%v", updatedPlugin.ID)
//	message, resp, err := client.PluginService.Delete(updatedPlugin.ID, "mockbin")
//	fmt.Printf("Delete Plugin : \n%v\n%v\n%v\n", message, resp, err)

//	consumers, _, _ := client.ConsumerService.List()
//	fmt.Printf("Consumers:\n%v\n", consumers)

//	consumer, _, _ := client.ConsumerService.Get("gokun")
//	fmt.Printf("Consumer:\n%v\n", consumer)

//	generateConsumer := &Consumer{
//		Username: "sakabe",
//	}

//	consumer, resp, err := client.ConsumerService.Create(generateConsumer)
//	fmt.Printf("Create Consumer :\n%v\n%v\n%v\n%v\n", generateConsumer, consumer, resp, err)

//	updateConsumer := &Consumer{
//		ID:       consumer.ID,
//		Username: "sakabeupdate",
//	}

//	updatedConsumer, resp, err := client.ConsumerService.Update(updateConsumer)
//	fmt.Printf("Update Consumer :\n%v\n%v\n%v\n%v\n", updateConsumer, updatedConsumer, resp, err)

//	message, resp, err := client.ConsumerService.Delete(updatedConsumer.ID)
//	fmt.Printf("Delete Consumer : \n%v\n%v\n%v\n", message, resp, err)

// oauth2, _, _ := client.OAuth2ConfigService.List("gokun")
// fmt.Printf("OAuth2:\n%v\n", oauth2)

// targetOAuth2Config, _, _ := client.OAuth2Service.Get("gokun", "86e4f18c-00a6-403f-a526-5d8fc3dac95d")
// fmt.Printf("OAuth2:\n%v\n", targetOAuth2Config)

// generateOAuth2Config := &OAuth2Config{
// 	Name:        "sakabe site",
// 	RedirectURI: "http://koudaiii.com",
// }

// oauth2config, resp, err := client.OAuth2ConfigService.Create(generateOAuth2Config, "gokun")
// fmt.Printf("Create OAuth2 :\n%v\n%v\n%v\n%v\n", generateOAuth2Config, oauth2config, resp, err)

// updateOAuth2Config := &OAuth2Config{
// 	ID:   oauth2config.ID,
// 	Name: "sakabeupdate",
// }

// updatedOAuth2Config, resp, err := client.OAuth2ConfigService.Update(updateOAuth2Config, "gokun")
// fmt.Printf("Update OAuth2Config :\n%v\n%v\n%v\n%v\n", updateOAuth2Config, updatedOAuth2Config, resp, err)

// message, resp, err := client.OAuth2ConfigService.Delete(updatedOAuth2Config.ID, "gokun")
// fmt.Printf("Delete OAuth2Config : \n%v\n%v\n%v\n", message, resp, err)

// assigneesOAuth2List, _, _ := client.AssigneesOAuth2Service.List()
// fmt.Printf("OAuth2:\n%v\n", assigneesOAuth2List)

// targetAssigneesOAuth2, _, _ := client.AssigneesOAuth2Service.Get("dd247d86-7377-48a6-8bba-7a90ae445c44")
// fmt.Printf("OAuth2:\n%v\n", targetAssigneesOAuth2)

// generateAssigneesOAuth2 := &AssigneesOAuth2{
// 	TokenType:    "bearer",
// 	ExpiresIn:    0,
// 	CredentialID: "cba41ce4-e965-4833-8d24-78171053cabd",
// }
// assigneesOAuth2, resp, err := client.AssigneesOAuth2Service.Create(generateAssigneesOAuth2)
// fmt.Printf("Create Assignees OAuth2 :\n%v\n%v\n%v\n%v\n", generateAssigneesOAuth2, assigneesOAuth2, resp, err)

// updateAssigneesOAuth2 := &AssigneesOAuth2{
// 	ID:        assigneesOAuth2.ID,
// 	ExpiresIn: 1000,
// }

// updatedAssigneesOAuth2, resp, err := client.AssigneesOAuth2Service.Update(updateAssigneesOAuth2)
// fmt.Printf("Update Assignees OAuth2 :\n%v\n%v\n%v\n%v\n", updateAssigneesOAuth2, updatedAssigneesOAuth2, resp, err)

// message, resp, err := client.AssigneesOAuth2Service.Delete(updatedAssigneesOAuth2.ID)
// fmt.Printf("Delete Assignees OAuth2 : \n%v\n%v\n%v\n", message, resp, err)
// }
