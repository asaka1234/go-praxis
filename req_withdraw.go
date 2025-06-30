package go_praxis

import (
	"crypto/tls"
	"fmt"
	"github.com/asaka1234/go-praxis/utils"
	"github.com/mitchellh/mapstructure"
	"time"
)

// 下单(充值/提现是同一个接口)
func (cli *Client) Withdraw(req PraxisCashierReq) (*PraxisCashierResp, error) {

	rawURL := cli.Params.BaseUrl

	var params map[string]interface{}
	mapstructure.Decode(req, &params)

	//补充字段
	params["merchant_id"] = cli.Params.MerchantId
	params["application_key"] = cli.Params.ApplicationKey
	params["version"] = cli.Params.ApiVersion
	params["locale"] = cli.Params.ApiLocale
	params["notification_url"] = cli.Params.WithdrawBackUrl
	params["return_url"] = cli.Params.WithdrawFeBackUrl
	params["intent"] = string(IntentTypeWithdrawal) //决策了是 withdraw
	params["timestamp"] = time.Now().Unix()

	//计算签名
	bsUtil := utils.NewBuildSignatureUtils()
	gtAuthentication := bsUtil.GetGtAuthentication(params, cli.Params.MerchantSecret, utils.SignTypeSendReq)

	//返回值会放到这里
	var result PraxisCashierResp

	resp, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetBody(params).
		SetHeaders(getAuthHeaders(gtAuthentication)).
		SetDebug(cli.debugMode).
		SetResult(&result).
		Post(rawURL)

	//fmt.Printf("accessToken: %+v\n", resp)

	if err != nil {
		return nil, err
	}

	if resp.Error() != nil {
		//反序列化错误会在此捕捉
		return nil, fmt.Errorf("%v, body:%s", resp.Error(), resp.Body())
	}

	return &result, err
}
