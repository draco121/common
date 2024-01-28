package clients

import (
	"github.com/draco121/common/models"
	"github.com/go-resty/resty/v2"
)

type IUserServiceApiClient interface {
	GetUserByEmail(email string) (*models.User, error)
	GetUserById(id string) (*models.User, error)
}

type userServiceApiClient struct {
	IUserServiceApiClient
	client *resty.Client
}

func NewUserServiceApiClient(baseUrl string) IUserServiceApiClient {
	c := resty.New()
	c.BaseURL = baseUrl
	client := userServiceApiClient{
		client: c,
	}
	return &client
}

func (c *userServiceApiClient) GetUserByEmail(email string) (*models.User, error) {
	user := models.User{}
	_, err := c.client.R().
		SetResult(&user).
		SetQueryParam("email", email).
		ForceContentType("application/json").
		Get("v1/user")
	if err != nil {
		return nil, err
	} else {
		return &user, nil
	}
}

func (c *userServiceApiClient) GetUserById(id string) (*models.User, error) {
	user := models.User{}
	_, err := c.client.R().
		SetResult(&user).
		SetQueryParam("id", id).
		ForceContentType("application/json").
		Get("v1/user")
	if err != nil {
		return nil, err
	} else {
		return &user, nil
	}
}
