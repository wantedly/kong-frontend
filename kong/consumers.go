package main

import (
	_ "fmt"
	"net/http"

	"github.com/dghubble/sling"
	"github.com/koudaiii/kong-oauth-token-generator/config"
)

type Consumers struct {
	Consumer []Consumer `json:"data"`
	Total    int        `json:"total"`
}

type Consumer struct {
	CreatedAt int    `json:"created_at,omitempty"`
	ID        string `json:"id,omitempty"`
	Username  string `json:"username,omitempty"`
}

// Services

// ConsumerService provides methods for creating and reading issues.
type ConsumerService struct {
	sling *sling.Sling
}

// NewConsumerService returns a new ConsumerService.
func NewConsumerService(httpClient *http.Client, config *config.KongConfiguration) *ConsumerService {
	return &ConsumerService{
		sling: sling.New().Client(httpClient).Base(config.KongAdminURL + "consumers/"),
	}
}

func (s *ConsumerService) Get(params string) (Consumer, *http.Response, error) {
	consumer := new(Consumer)
	kongError := new(KongError)
	resp, err := s.sling.New().Path(params).Receive(consumer, kongError)
	if err == nil {
		err = kongError
	}
	return *consumer, resp, err
}

func (s *ConsumerService) List() (Consumers, *http.Response, error) {
	consumers := new(Consumers)
	kongError := new(KongError)
	resp, err := s.sling.New().Receive(consumers, kongError)
	if err == nil {
		err = kongError
	}
	return *consumers, resp, err
}
