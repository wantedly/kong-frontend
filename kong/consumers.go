package kong

import (
	_ "fmt"
	"net/http"

	"github.com/dghubble/sling"
	"github.com/koudaiii/kong-oauth-token-generator/config"
)

type Consumers struct {
	Consumer []Consumer `json:"data,omitempty"`
	Total    int        `json:"total,omitempty"`
}

type Consumer struct {
	CreatedAt int    `json:"created_at,omitempty"`
	ID        string `json:"id,omitempty"`
	Username  string `json:"username,omitempty"`
}

// Services

// ConsumerService provides methods for creating and reading issues.
type ConsumerService struct {
	sling  *sling.Sling
	config *config.KongConfiguration
}

// NewConsumerService returns a new ConsumerService.
func NewConsumerService(httpClient *http.Client, config *config.KongConfiguration) *ConsumerService {
	return &ConsumerService{
		sling:  sling.New().Client(httpClient).Base(config.KongAdminURL + "consumers/"),
		config: config,
	}
}

func (s *ConsumerService) Create(params *Consumer) (*Consumer, *http.Response, error) {
	consumer := new(Consumer)
	resp, err := s.sling.New().Post(s.config.KongAdminURL + "consumers").BodyJSON(params).ReceiveSuccess(consumer)
	return consumer, resp, err
}

func (s *ConsumerService) Get(params string) (*Consumer, *http.Response, error) {
	consumer := new(Consumer)
	resp, err := s.sling.New().Path(params).ReceiveSuccess(consumer)
	return consumer, resp, err
}

func (s *ConsumerService) List() (*Consumers, *http.Response, error) {
	consumers := new(Consumers)
	resp, err := s.sling.New().ReceiveSuccess(consumers)
	return consumers, resp, err
}

func (s *ConsumerService) Update(params *Consumer) (*Consumer, *http.Response, error) {
	consumer := new(Consumer)
	resp, err := s.sling.New().Patch(s.config.KongAdminURL + "consumers/" + params.ID).BodyJSON(params).ReceiveSuccess(consumer)
	return consumer, resp, err
}

func (s *ConsumerService) Delete(consumerID string) (string, *http.Response, error) {
	var message string
	resp, err := s.sling.New().Delete(s.config.KongAdminURL + "consumers/" + consumerID).ReceiveSuccess(message)
	return message, resp, err
}
