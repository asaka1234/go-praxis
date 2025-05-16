package go_praxis

import (
	"github.com/asaka1234/go-praxis/utils"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	MerchantID string
	ApiVersion string

	BaseURL string

	ryClient *resty.Client
	logger   utils.Logger
}

func NewClient(logger utils.Logger, merchantID string, apiVersion string, baseURL string) *Client {
	return &Client{
		MerchantID: merchantID,
		ApiVersion: apiVersion,
		BaseURL:    baseURL,
		ryClient:   resty.New(), //client实例
		logger:     logger,
	}
}
