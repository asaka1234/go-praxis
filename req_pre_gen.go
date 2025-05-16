package go_praxis

import (
	"crypto/tls"
	"github.com/asaka1234/go-praxis/utils"
)

// 下单(充值/提现是同一个接口)
func (cli *Client) PreGen(req PraxisCashierRequest) (*PraxisCashierResponse, error) {

	rawURL := cli.BaseURL

	//赋值
	req.MerchantId = cli.MerchantID
	req.ApplicationKey = cli.ApplicationKey
	req.Version = cli.ApiVersion
	req.Locale = cli.ApiLocale

	signStr := utils.CalSign(req) //计算请求的签名

	//返回值会放到这里
	var result PraxisCashierResponse

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
