package oauth2

import (
	_ "fmt"

	"github.com/koudaiii/kong-oauth-token-generator/kong"
)

func List(self *kong.Client) (*kong.Consumers, *kong.AssigneesOAuth2s, error) {
	consumers, _, err := self.ConsumerService.List()
	assignees, _, err := self.AssigneesOAuth2Service.List()

	return consumers, assignees, err
}
