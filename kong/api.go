package main

import (
	_ "fmt"
	"net/http"

	"github.com/dghubble/sling"
	"github.com/koudaiii/kong-oauth-token-generator/config"
)

type APIs struct {
	API   []API `json:"data"`
	Total int   `json:"total"`
}

type API struct {
	CreatedAt        int    `json:"created_at,omitempty"`
	ID               string `json:"id,omitempty"`
	Name             string `json:"name"`
	PreserveHost     bool   `json:"preserve_host,omitempty"`
	RequestPath      string `json:"request_path,omitempty"`
	StripRequestPath bool   `json:"strip_request_path,omitempty"`
	UpstreamURL      string `json:"upstream_url,omitempty"`
	RequestHost      string `json:"request_host,omitempty"`
}

// Services

// APIService provides methods for creating and reading issues.
type APIService struct {
	sling *sling.Sling
}

// NewAPIService returns a new APIService.
func NewAPIService(httpClient *http.Client, config *config.KongConfiguration) *APIService {
	return &APIService{
		sling: sling.New().Client(httpClient).Base(config.KongAdminURL + "apis/"),
	}
}

func (s *APIService) Create(params *API) (*API, *http.Response, error) {
	api := new(API)
	kongError := new(KongError)
	resp, err := s.sling.New().Post("http://localhost:8001/apis").BodyJSON(params).Receive(api, kongError)
	if err == nil {
		err = kongError
	}
	return api, resp, err
}

func (s *APIService) Get(params string) (*API, *http.Response, error) {
	api := new(API)
	kongError := new(KongError)
	resp, err := s.sling.New().Path(params).Receive(api, kongError)
	if err == nil {
		err = kongError
	}
	return api, resp, err
}

func (s *APIService) List() (*APIs, *http.Response, error) {
	apis := new(APIs)
	kongError := new(KongError)
	resp, err := s.sling.New().Receive(apis, kongError)
	if err == nil {
		err = kongError
	}
	return apis, resp, err
}
