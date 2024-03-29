package clients

import (
	"github.com/draco121/common/models"
	"github.com/go-resty/resty/v2"
)

type IAuthenticationServiceApiClient interface {
	Authenticate(token string) (*models.JwtCustomClaims, error)
}

type authenticationServiceApiClient struct {
	IAuthenticationServiceApiClient
	client *resty.Client
}

func NewAuthenticationServiceApiClient(baseUrl string) IAuthenticationServiceApiClient {
	c := resty.New()
	c.BaseURL = baseUrl
	client := authenticationServiceApiClient{
		client: c,
	}
	return &client
}

func (c *authenticationServiceApiClient) Authenticate(token string) (*models.JwtCustomClaims, error) {
	claims := models.JwtCustomClaims{}
	_, err := c.client.R().
		SetResult(&claims).
		SetHeader("Authentication", token).
		ForceContentType("application/json").
		Get("v1/authenticate")
	if err != nil {
		return nil, err
	} else {
		return &claims, nil
	}
}
