package kong

import (
	_ "fmt"
	"net/http"

	"github.com/dghubble/sling"
	"github.com/wantedly/kong-frontend/config"
)

type APIs struct {
	API   []API `json:"data,omitempty"`
	Total int   `json:"total,omitempty"`
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
	sling  *sling.Sling
	config *config.KongConfiguration
}

// NewAPIService returns a new APIService.
func NewAPIService(httpClient *http.Client, config *config.KongConfiguration) *APIService {
	return &APIService{
		sling:  sling.New().Client(httpClient).Base(config.KongAdminURL + "apis/"),
		config: config,
	}
}

func (s *APIService) Create(params *API) (*API, *http.Response, error) {
	api := new(API)
	resp, err := s.sling.New().Post(s.config.KongAdminURL + "apis/").BodyJSON(params).ReceiveSuccess(api)
	return api, resp, err
}

func (s *APIService) Get(params string) (*API, *http.Response, error) {
	api := new(API)
	resp, err := s.sling.New().Path(params).ReceiveSuccess(api)
	return api, resp, err
}

func (s *APIService) List() (*APIs, *http.Response, error) {
	apis := new(APIs)
	resp, err := s.sling.New().ReceiveSuccess(apis)
	return apis, resp, err
}

func (s *APIService) Update(params *API) (*API, *http.Response, error) {
	api := new(API)
	resp, err := s.sling.New().Patch(s.config.KongAdminURL + "apis/" + params.ID).BodyJSON(params).ReceiveSuccess(api)
	return api, resp, err
}

func (s *APIService) Delete(apiID string) (string, *http.Response, error) {
	var message string
	resp, err := s.sling.New().Delete(s.config.KongAdminURL + "apis/" + apiID).ReceiveSuccess(message)
	return message, resp, err
}
