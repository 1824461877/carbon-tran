package paypal

import (
	"github.com/plutov/paypal"
	"time"
)

type PayClientConfig struct {
	ClientID string
	Secret   string
	APIBase  string
}

type PayClient struct {
	Client  *paypal.Client
	timeout int64
}

func (pc *PayClient) GetToken() error {
	var (
		token *paypal.TokenResponse
		err   error
	)
	if pc.timeout != 0 || time.Now().Unix() <= pc.timeout {
		return nil
	}
	if token, err = pc.Client.GetAccessToken(); err != nil {
		return err
	} else {
		pc.timeout = time.Now().Add(time.Duration(token.ExpiresIn)).Unix()
	}
	return nil
}

func NewPayClient(c *PayClientConfig) *PayClient {
	var (
		client *paypal.Client
		err    error
	)
	if client, err = paypal.NewClient(c.ClientID, c.Secret, c.APIBase); err != nil {
		panic(err)
	}
	return &PayClient{
		Client: client,
	}
}
