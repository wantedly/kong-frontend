package consumer

import (
	"errors"
	"net/http"

	"github.com/wantedly/kong-frontend/kong"
)

func List(self *kong.Client) (*kong.Consumers, error) {
	consumers, _, err := self.ConsumerService.List()
	return consumers, err
}

func Exists(self *kong.Client, consumerID string) bool {
	_, resp, _ := self.ConsumerService.Get(consumerID)
	if resp.StatusCode != 404 {
		return true
	}
	return false
}

func Get(self *kong.Client, consumerID string) (*kong.Consumer, error) {
	consumer, res, err := self.ConsumerService.Get(consumerID)
	if err != nil {
		return nil, err
	} else if res.StatusCode >= http.StatusBadRequest {
		return nil, errors.New(res.Status)
	}
	return consumer, err
}

func Delete(self *kong.Client, consumerID string) (string, error) {
	message, res, err := self.ConsumerService.Delete(consumerID)
	if err != nil {
		return "", err
	} else if res.StatusCode >= http.StatusBadRequest {
		return "", errors.New(res.Status)
	}
	return message, err
}

func Update(self *kong.Client, params *kong.Consumer) (*kong.Consumer, error) {
	consumer, res, err := self.ConsumerService.Update(params)
	if err != nil {
		return nil, err
	} else if res.StatusCode >= http.StatusBadRequest {
		return nil, errors.New(res.Status)
	}
	return consumer, err
}

func Create(self *kong.Client, params *kong.Consumer) (*kong.Consumer, error) {
	consumer, res, err := self.ConsumerService.Create(params)
	if err != nil {
		return nil, err
	} else if res.StatusCode >= http.StatusBadRequest {
		return nil, errors.New(res.Status)
	}
	return consumer, err
}
