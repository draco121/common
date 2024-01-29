package clients

import (
	"github.com/draco121/common/models"
	"github.com/go-resty/resty/v2"
)

type IAuthorizationServiceApiClient interface {
	Authorize(input models.AuthorizationInput) (*models.AuthorizationOutput, error)
}

type authorizationServiceApiClient struct {
	IAuthorizationServiceApiClient
	client *resty.Client
}

func NewAuthorizationServiceApiClient(baseUrl string) IAuthorizationServiceApiClient {
	c := resty.New()
	c.BaseURL = baseUrl
	client := authorizationServiceApiClient{
		client: c,
	}
	return &client
}

func (c *authenticationServiceApiClient) Authorize(input models.AuthorizationInput) (*models.AuthorizationOutput, error) {
	var result models.AuthorizationOutput
	_, err := c.client.R().
		SetResult(&result).
		SetBody(input).
		ForceContentType("application/json").
		Post("v1/authorize")
	if err != nil {
		return nil, err
	} else {
		return &result, nil
	}
}
