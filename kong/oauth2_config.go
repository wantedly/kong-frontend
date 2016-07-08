package kong

import (
	_ "fmt"
	"net/http"

	"github.com/dghubble/sling"
	"github.com/koudaiii/kong-oauth-token-generator/config"
)

type OAuth2ConfigList struct {
	OAuth2Config []OAuth2Config `json:"data,omitempty"`
	Total        int            `json:"total,omitempty"`
}

type OAuth2Config struct {
	ClientID     string `json:"client_id,omitempty"`
	ClientSecret string `json:"client_secret,omitempty"`
	ConsumerID   string `json:"consumer_id,omitempty"`
	CreatedAt    int    `json:"created_at,omitempty"`
	ID           string `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	RedirectURI  string `json:"redirect_uri,omitempty"`
}

// Services

// ConfigService provides methods for creating and reading issues.
type OAuth2ConfigService struct {
	sling *sling.Sling
}

// NewCOAuth2ConfigService returns a new OAuth2ConfigService.
func NewOAuth2ConfigService(httpClient *http.Client, config *config.KongConfiguration) *OAuth2ConfigService {
	return &OAuth2ConfigService{
		sling: sling.New().Client(httpClient).Base(config.KongAdminURL + "consumers/"),
	}
}

func (s *OAuth2ConfigService) Create(params *OAuth2Config, consumerName string) (*OAuth2Config, *http.Response, error) {
	oauth2 := new(OAuth2Config)
	kongError := new(KongError)
	resp, err := s.sling.New().Post("http://localhost:8001/consumers/"+consumerName+"/oauth2").BodyJSON(params).Receive(oauth2, kongError)
	if err == nil {
		err = kongError
	}
	return oauth2, resp, err
}

func (s *OAuth2ConfigService) Get(consumerName string, oauth2ID string) (*OAuth2Config, *http.Response, error) {
	oauth2config := new(OAuth2Config)
	kongError := new(KongError)
	resp, err := s.sling.New().Path(consumerName+"/oauth2/"+oauth2ID).Receive(oauth2config, kongError)
	if err == nil {
		err = kongError
	}
	return oauth2config, resp, err
}

func (s *OAuth2ConfigService) List(consumerName string) (*OAuth2ConfigList, *http.Response, error) {
	oauth2configlist := new(OAuth2ConfigList)
	kongError := new(KongError)
	resp, err := s.sling.New().Path(consumerName+"/oauth2").Receive(oauth2configlist, kongError)
	if err == nil {
		err = kongError
	}
	return oauth2configlist, resp, err
}

func (s *OAuth2ConfigService) Update(params *OAuth2Config, consumerName string) (*OAuth2Config, *http.Response, error) {
	oauth2 := new(OAuth2Config)
	kongError := new(KongError)
	resp, err := s.sling.New().Patch("http://localhost:8001/consumers/"+consumerName+"/oauth2/"+params.ID).BodyJSON(params).Receive(oauth2, kongError)
	if err == nil {
		err = kongError
	}
	return oauth2, resp, err
}

func (s *OAuth2ConfigService) Delete(oauth2ID string, consumerName string) (string, *http.Response, error) {
	var message string
	kongError := new(KongError)
	resp, err := s.sling.New().Delete("http://localhost:8001/consumers/"+consumerName+"/oauth2/"+oauth2ID).Receive(message, kongError)
	if err == nil {
		err = kongError
	}
	return message, resp, err
}
