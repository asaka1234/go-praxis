package go_praxis

import "github.com/go-resty/resty/v2"

type Client struct {
	MerchantID string
	AppKey     string

	BaseURL string

	ryClient *resty.Client
}

func NewClient(merchantID string, appKey string, baseURL string) *Client {
	return &Client{
		MerchantID: merchantID,
		AppKey:     appKey,
		BaseURL:    baseURL,
		ryClient:   resty.New(), //client实例
	}
}
