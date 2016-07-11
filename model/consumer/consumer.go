package consumer

import (
	"fmt"

	"github.com/koudaiii/kong-oauth-token-generator/config"
	"github.com/koudaiii/kong-oauth-token-generator/kong"
)

func List(config *config.KongConfiguration) (*kong.Consumers, error) {
	client := kong.NewClient(nil, config)
	consumers, _, err := client.ConsumerService.List()
	return consumers, err
}
