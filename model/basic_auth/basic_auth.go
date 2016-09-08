package basic_auth

import (
	"errors"
	"net/http"

	"github.com/wantedly/kong-frontend/kong"
)

func List(self *kong.Client, consumerID string) (*kong.BasicAuthCredentials, error) {
	credentials, res, err := self.BasicAuthService.List(consumerID)
	if err != nil {
		return nil, err
	} else if res.StatusCode >= http.StatusBadRequest {
		return nil, errors.New(res.Status)
	}
	return credentials, err
}

func Get(self *kong.Client, consumerID , credentialID string) (*kong.BasicAuthCredential, error) {
	credential, res, err := self.BasicAuthService.Get(consumerID, credentialID)
	if err != nil {
		return nil, err
	} else if res.StatusCode >= http.StatusBadRequest {
		return nil, errors.New(res.Status)
	}
	return credential, err
}

func Delete(self *kong.Client, consumerID, credentialID string) (string, error) {
	message, res, err := self.BasicAuthService.Delete(consumerID, credentialID)
	if err != nil {
		return "", err
	} else if res.StatusCode >= http.StatusBadRequest {
		return "", errors.New(res.Status)
	}
	return message, err
}

func Update(self *kong.Client, consumerID string, params *kong.BasicAuthCredential) (*kong.BasicAuthCredential, error) {
	credential, res, err := self.BasicAuthService.Update(consumerID, params)
	if err != nil {
		return nil, err
	} else if res.StatusCode >= http.StatusBadRequest {
		return nil, errors.New(res.Status)
	}
	return credential, err
}

func Create(self *kong.Client, consumerID string, params *kong.BasicAuthCredential) (*kong.BasicAuthCredential, error) {
	credential, res, err := self.BasicAuthService.Create(consumerID, params)
	if err != nil {
		return nil, err
	} else if res.StatusCode >= http.StatusBadRequest {
		return nil, errors.New(res.Status)
	}
	return credential, err
}
