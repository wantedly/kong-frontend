package kong

import (
	"net/http"

	"github.com/dghubble/sling"
	"github.com/wantedly/kong-frontend/config"
)

type BasicAuthCredentials struct {
	Data  []BasicAuthCredential `json:"data,omitempty"`
	Total int                   `json:"total,omitempty"`
}

type BasicAuthCredential struct {
	Password   string `json:"password,omitempty" form:"password"`
	ConsumerID string `json:"consumer_id,omitempty" form:"consumer_id"`
	ID         string `json:"id,omitempty"`
	Username   string `json:"username,omitempty" form:"username"`
	CreatedAt  int    `json:"created_at,omitempty"`
}

type BasicAuthService struct {
	sling  *sling.Sling
	config *config.KongConfiguration
}

func NewBasicAuthService(httpClient *http.Client, config *config.KongConfiguration) *BasicAuthService {
	return &BasicAuthService{
		sling:  sling.New().Client(httpClient).Base(config.KongAdminURL + "consumers/"),
		config: config,
	}
}

func (s *BasicAuthService) Create(consumerID string, params *BasicAuthCredential) (*BasicAuthCredential, *http.Response, error) {
	credential := new(BasicAuthCredential)
	resp, err := s.sling.New().Post(s.config.KongAdminURL + "consumers/" + consumerID + "/basic-auth").BodyJSON(params).ReceiveSuccess(credential)
	return credential, resp, err
}

func (s *BasicAuthService) Get(consumerID, credentialID string) (*BasicAuthCredential, *http.Response, error) {
	credential := new(BasicAuthCredential)
	resp, err := s.sling.New().Path(consumerID + "/basic-auth/" + credentialID).ReceiveSuccess(credential)
	return credential, resp, err
}

func (s *BasicAuthService) List(consumerID string) (*BasicAuthCredentials, *http.Response, error) {
	credentials := new(BasicAuthCredentials)
	resp, err := s.sling.New().Path(consumerID + "/basic-auth").ReceiveSuccess(credentials)
	return credentials, resp, err
}

func (s *BasicAuthService) Update(consumerID string, params *BasicAuthCredential) (*BasicAuthCredential, *http.Response, error) {
	credential := new(BasicAuthCredential)
	resp, err := s.sling.New().Patch(s.config.KongAdminURL + "consumers/" + consumerID + "/basic-auth/" + params.ID).BodyJSON(params).ReceiveSuccess(credential)
	return credential, resp, err
}

func (s *BasicAuthService) Delete(consumerID, credentialID string) (string, *http.Response, error) {
	var message string
	resp, err := s.sling.New().Delete(s.config.KongAdminURL + "consumers/" + consumerID + "/basic-auth/" + credentialID).ReceiveSuccess(message)
	return message, resp, err
}
