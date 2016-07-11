package api

import (
	"github.com/koudaiii/kong-oauth-token-generator/kong"
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

func List(self *kong.APIService) (*kong.APIs, error) {
	apis, _, err := self.List()
	return apis, err
}

func New(self *kong.APIService) {

}
