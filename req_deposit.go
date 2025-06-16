package go_praxis

import (
	"crypto/tls"
	"github.com/asaka1234/go-praxis/utils"
	"github.com/mitchellh/mapstructure"
	"time"
)

// 下单(充值/提现是同一个接口)
func (cli *Client) Deposit(req PraxisCashierReq) (*PraxisCashierResp, error) {

	rawURL := cli.Params.BaseUrl

	var params map[string]interface{}
	mapstructure.Decode(req, &params)

	//补充字段
	params["merchant_id"] = cli.Params.MerchantId
	params["application_key"] = cli.Params.ApplicationKey
	params["version"] = cli.Params.ApiVersion
	params["locale"] = cli.Params.ApiLocale
	params["notification_url"] = cli.Params.DepositBackUrl
	params["return_url"] = cli.Params.DepositFeBackUrl
	params["intent"] = string(IntentTypePayment) //决策了是 deposit
	params["timestamp"] = time.Now().Unix()

	//计算签名
	bsUtil := utils.NewBuildSignatureUtils()
	gtAuthentication := bsUtil.GetGtAuthentication(params, cli.Params.MerchantSecret, utils.SignTypeSendReq)

	//返回值会放到这里
	var result PraxisCashierResp

	_, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetBody(params).
		SetHeaders(getAuthHeaders(gtAuthentication)).
		SetDebug(cli.debugMode).
		SetResult(&result).
		Post(rawURL)

	if err != nil {
		return nil, err
	}

	return &result, err
}
