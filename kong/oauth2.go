package main

import (
	_ "fmt"
	"net/http"

	"github.com/dghubble/sling"
	"github.com/koudaiii/kong-oauth-token-generator/config"
)

type Oauth2List struct {
	Oauth2Config []Oauth2Config `json:"data,omitempty"`
	Total        int            `json:"total,omitempty"`
}

type Oauth2Config struct {
	ClientID     string `json:"client_id,omitempty"`
	ClientSecret string `json:"client_secret,omitempty"`
	ConsumerID   string `json:"consumer_id,omitempty"`
	CreatedAt    int    `json:"created_at,omitempty"`
	ID           string `json:"id,omitempty"`
	Name         string `json:"name"`
	RedirectURI  string `json:"redirect_uri"`
}

// Services

// ConfigService provides methods for creating and reading issues.
type Oauth2Service struct {
	sling *sling.Sling
}

// NewCOauth2Service returns a new Oauth2Service.
func NewOauth2Service(httpClient *http.Client, config *config.KongConfiguration) *Oauth2Service {
	return &Oauth2Service{
		sling: sling.New().Client(httpClient).Base(config.KongAdminURL + "consumers/"),
	}
}

func (s *Oauth2Service) Get(consumerName string, oauth2ID string) (Oauth2Config, *http.Response, error) {
	oauth2config := new(Oauth2Config)
	kongError := new(KongError)
	resp, err := s.sling.New().Path(consumerName+"/oauth2/"+oauth2ID).Receive(oauth2config, kongError)
	if err == nil {
		err = kongError
	}
	return *oauth2config, resp, err
}

func (s *Oauth2Service) List(consumerName string) (Oauth2List, *http.Response, error) {
	oauth2 := new(Oauth2List)
	kongError := new(KongError)
	resp, err := s.sling.New().Path(consumerName+"/oauth2").Receive(oauth2, kongError)
	if err == nil {
		err = kongError
	}
	return *oauth2, resp, err
}
