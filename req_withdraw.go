package go_praxis

import (
	"crypto/tls"
	"github.com/asaka1234/go-praxis/utils"
	"github.com/mitchellh/mapstructure"
	"time"
)

// 下单(充值/提现是同一个接口)
func (cli *Client) Withdraw(req PraxisWithdrawReq) (*PraxisWithdrawResp, error) {

	rawURL := cli.Params.BaseUrl

	//拿到签名的参数
	requestParams := cli.CreateWithdrawRequestParams(req)

	bsUtil := utils.NewBuildSignatureUtils()
	gtAuthentication := bsUtil.GetGtAuthentication(requestParams, cli.Params.MerchantSecret, utils.SignTypeSendReq)

	//返回值会放到这里
	var result PraxisWithdrawResp

	_, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetBody(requestParams).
		SetHeaders(getAuthHeaders(gtAuthentication)).
		SetResult(&result).
		Post(rawURL)

	//fmt.Printf("accessToken: %+v\n", resp)

	if err != nil {
		return nil, err
	}

	return &result, err
}

func (cli *Client) CreateWithdrawRequestParams(req PraxisWithdrawReq) map[string]interface{} {
	params := make(map[string]interface{})

	params["merchant_id"] = cli.Params.MerchantId
	params["application_key"] = cli.Params.ApplicationKey
	params["intent"] = string(IntentTypeWithdrawal) //req.Intent
	params["currency"] = req.Currency
	params["amount"] = req.Amount
	params["cid"] = req.Cid
	params["locale"] = cli.Params.ApiLocale // Assuming Locale is a package constant
	params["customer_token"] = req.CustomerToken

	// struct → map
	var userMap map[string]interface{}
	mapstructure.Decode(req.CustomerData, &userMap)
	params["customer_data"] = userMap //req.CustomerData //把这个struct转为map

	params["payment_method"] = req.PaymentMethod
	params["gateway"] = req.Gateway
	params["validation_url"] = req.ValidationURL
	params["notification_url"] = cli.Params.WithdrawBackUrl //req.NotificationURL
	params["return_url"] = cli.Params.WithdrawFeBackUrl     //req.ReturnURL
	params["order_id"] = req.OrderID
	params["version"] = cli.Params.ApiVersion // Assuming APIVersion is a package constant
	params["timestamp"] = time.Now().Unix()   // Unix timestamp in seconds

	return params
}
