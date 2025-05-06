package go_praxis

import (
	"crypto/tls"
)

// 下单
func (cli *Client) PreGen(req CashierRequest) (*CashierResponse, error) {

	rawURL := cli.BaseURL

	signStr := CalSign(req) //计算请求的签名

	//返回值会放到这里
	var result CashierResponse

	_, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetBody(req).
		SetHeaders(getAuthHeaders(signStr)).
		SetResult(&result).
		Post(rawURL)

	//fmt.Printf("accessToken: %+v\n", resp)

	if err != nil {
		return nil, err
	}

	return &result, err
}
