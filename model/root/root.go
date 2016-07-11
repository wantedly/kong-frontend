package root

import (
	"fmt"

	"github.com/koudaiii/kong-oauth-token-generator/config"
	"github.com/koudaiii/kong-oauth-token-generator/kong"
)

func List(config *config.KongConfiguration) (*kong.APIs, error) {
	client := kong.NewClient(nil, config)
	apis, _, err := client.APIService.List()
	fmt.Printf("APIs:\n%v\n", apis)
	return apis, err
}
