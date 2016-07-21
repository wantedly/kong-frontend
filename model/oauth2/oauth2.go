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

func Exists(self *kong.Client, consumerName string) bool {
	_, resp, _ := self.ConsumerService.Get(consumerName)
	if resp == nil {
		return false
	}
	if resp.StatusCode != 404 {
		return true
	}
	return false
}

func Get(self *kong.Client, consumerName string) (*kong.Consumer, error) {
	consumer, _, err := self.ConsumerService.Get(consumerName)
	if err != nil {
		return nil, err
	}
	return consumer, err
}

func Create(self *kong.Client, generateConsumer *kong.Consumer) (*kong.Consumer, error) {
	consumer, _, err := self.ConsumerService.Create(generateConsumer)
	if err != nil {
		return nil, err
	}
	return consumer, err
}
