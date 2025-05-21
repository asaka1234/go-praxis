package go_praxis

import (
	"crypto/tls"
	"github.com/asaka1234/go-praxis/utils"
	"github.com/mitchellh/mapstructure"
	"time"
)

// 下单(充值/提现是同一个接口)
func (cli *Client) Deposit(req PraxisDepositReq) (*PraxisDepositRsp, error) {

	rawURL := cli.BaseURL

	//拿到签名的参数
	requestParams := cli.createDepositRequestParams(req)
	bsUtil := utils.NewBuildSignatureUtils()
	gtAuthentication := bsUtil.GetGtAuthentication(requestParams, cli.MerchantSecret, utils.SignTypeSendReq)

	//返回值会放到这里
	var result PraxisDepositRsp

	_, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetBody(requestParams).
		SetHeaders(getAuthHeaders(gtAuthentication)).
		SetResult(&result).
		Post(rawURL)

	if err != nil {
		return nil, err
	}

	return &result, err
}

func (cli *Client) createDepositRequestParams(req PraxisDepositReq) map[string]interface{} {
	params := make(map[string]interface{})

	params["merchant_id"] = cli.MerchantID // Assuming these are package-level variables
	params["application_key"] = cli.ApplicationKey
	params["intent"] = string(IntentTypePayment) //req.Intent //枚举: payment,withdrawal,authorization (这里完全可以直接写死)
	params["currency"] = req.Currency
	params["amount"] = req.Amount
	params["cid"] = req.Cid
	params["locale"] = cli.ApiLocale
	params["customer_token"] = req.CustomerToken

	// struct → map
	var userMap map[string]interface{}
	mapstructure.Decode(req.CustomerData, &userMap)
	params["customer_data"] = userMap //req.CustomerData //把这个struct转为map

	params["payment_method"] = req.PaymentMethod
	params["gateway"] = req.Gateway
	params["validation_url"] = req.ValidationURL
	params["notification_url"] = req.NotificationURL
	params["return_url"] = req.ReturnURL
	params["order_id"] = req.OrderID
	params["version"] = cli.ApiVersion
	params["timestamp"] = time.Now().Unix() // Unix timestamp in seconds

	return params
}
