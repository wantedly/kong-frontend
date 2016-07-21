package oauth2

import (
	"github.com/koudaiii/kong-oauth-token-generator/kong"
)

func List(self *kong.Client) (*kong.Consumers, error) {
	consumers, _, err := self.ConsumerService.List()
	return consumers, err
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

func Get(self *kong.Client, consumerName string) (*kong.Consumer, *kong.AssigneesOAuth2, error) {
	consumer, _, err := self.ConsumerService.Get(consumerName)
	if err != nil {
		return nil, nil, err
	}
	oauth2configs, _, err := self.OAuth2ConfigService.List(consumer.Username)
	if err != nil {
		return consumer, nil, err
	}
	assignees, _, err := self.AssigneesOAuth2Service.List()
	if err != nil {
		return consumer, nil, err
	}
	for _, assignee := range assignees.AssigneesOAuth2 {
		if assignee.CredentialID == oauth2configs.OAuth2Config[0].ID {
			return consumer, &assignee, err
		}
	}
	return consumer, nil, err
}

func Delete(self *kong.Client, consumerName string) (string, error) {
	consumer, _, err := self.ConsumerService.Get(consumerName)
	if err != nil {
		return "", err
	}
	message, _, err := self.ConsumerService.Delete(consumer.ID)
	return message, err
}

func Create(self *kong.Client, generateConsumer *kong.Consumer) (*kong.Consumer, *kong.AssigneesOAuth2, error) {
	consumer, _, err := self.ConsumerService.Create(generateConsumer)
	if err != nil {
		return nil, nil, err
	}
	generateOAuth2Config := &kong.OAuth2Config{
		Name:        consumer.Username,
		RedirectURI: "http://example.com",
	}
	oauth2, _, err := self.OAuth2ConfigService.Create(generateOAuth2Config, consumer.Username)
	if err != nil {
		return nil, nil, err
	}
	generateAssigneesOAuth2 := &kong.AssigneesOAuth2{
		TokenType:    "bearer",
		ExpiresIn:    0,
		CredentialID: oauth2.ID,
	}
	assigneesOAuth2, _, err := self.AssigneesOAuth2Service.Create(generateAssigneesOAuth2)
	if err != nil {
		return nil, nil, err
	}

	return consumer, assigneesOAuth2, err
}
