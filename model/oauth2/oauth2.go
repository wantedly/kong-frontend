package oauth2

import (
	_ "fmt"

	"github.com/koudaiii/kong-oauth-token-generator/kong"
)

func List(self *kong.Client) (*kong.APIs, *kong.AssigneesOAuth2List, error) {
	apis, _, err := self.APIService.List()
	assignees, _, err := self.AssigneesOAuth2Service.List()
	return apis, assignees, err
}
