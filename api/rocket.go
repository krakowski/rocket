package rocket

import (
	"github.com/go-resty/resty/v2"
	"net/http"
)

type Credentials struct {
	Username string
	Password string
}

type ClientOptions struct {
	ServerUrl   string
	Credentials Credentials
}

type Client struct {
	common service

	Auth    *AuthService
	Message *MessageService
}

type service struct {
	client *Client
	resty  *resty.Client
}

func NewClient(client *http.Client, options ClientOptions) (*Client, error) {

	// Create a new resty client
	var restyClient *resty.Client
	if client == nil {
		restyClient = resty.New()
	} else {
		restyClient = resty.NewWithClient(client)
	}

	// Configure the created resty client
	restyClient.SetError(&LastError)
	restyClient.SetHostURL(options.ServerUrl)
	restyClient.SetHeaders(map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	})

	// Set service fields
	ret := &Client{}
	ret.common.client = ret
	ret.common.resty = restyClient

	// Set service references
	ret.Auth = (*AuthService)(&ret.common)
	ret.Message = (*MessageService)(&ret.common)

	// Login using the client
	err := ret.Auth.Login(options.Credentials.Username, options.Credentials.Password)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
