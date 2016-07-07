package main

import (
	_ "fmt"
	"net/http"
	"os"

	"github.com/dghubble/sling"
	"github.com/koudaiii/kong-oauth-token-generator/config"
)

type APIs struct {
	API   []API `json:"data"`
	Total int   `json:"total"`
}

type API struct {
	CreatedAt        int    `json:"created_at"`
	ID               string `json:"id"`
	Name             string `json:"name"`
	PreserveHost     bool   `json:"preserve_host"`
	RequestPath      string `json:"request_path"`
	StripRequestPath bool   `json:"strip_request_path"`
	UpstreamURL      string `json:"upstream_url"`
}

// Services

// APIService provides methods for creating and reading issues.
type APIService struct {
	sling *sling.Sling
}

// NewAPIService returns a new APIService.
func NewAPIService(httpClient *http.Client) *APIService {
	config, err := config.LoadConfig()
	if err != nil {
		os.Exit(1)
	}
	return &APIService{
		sling: sling.New().Client(httpClient).Base(config.KongAdminURL + "apis/"),
	}
}

func (s *APIService) Get(params string) (API, *http.Response, error) {
	api := new(API)
	kongError := new(KongError)
	resp, err := s.sling.New().Path(params).Receive(api, kongError)
	if err == nil {
		err = kongError
	}
	return *api, resp, err
}

func (s *APIService) List() (APIs, *http.Response, error) {
	apis := new(APIs)
	kongError := new(KongError)
	resp, err := s.sling.New().Receive(apis, kongError)
	if err == nil {
		err = kongError
	}
	return *apis, resp, err
}
