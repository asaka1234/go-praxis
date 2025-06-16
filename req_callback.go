package go_praxis

import (
	"errors"
	"fmt"
	"github.com/asaka1234/go-praxis/utils"
)

//https://doc.praxiscashier.com/integration_docs/latest/webhooks/notification

// if your API response contains "status":-1 or responds in unrecognized format,
// the notification will be resent automatically within approximately 5 minutes
func (cli *Client) CashierCallback(req PraxisCashierBackReq, sign string, processor func(PraxisCashierBackReq) error) error {
	//验证回调合法性
	if req.MerchantID != cli.Params.MerchantId || req.ApplicationKey != cli.Params.ApplicationKey {
		return errors.New("merchantId or applicationKey is illegal!")
	}
	//自己算一下签名
	requestParams := cli.CreateCashierCallbackRequestParams(req)
	bsUtil := utils.NewBuildSignatureUtils()
	signSelf := bsUtil.GetGtAuthentication(requestParams, cli.Params.MerchantSecret, utils.SignTypeCallbackReq)
	//对比下签名正确性
	if signSelf != sign {
		fmt.Printf("sign is not equal, sign:%s, signSelf:%s\n", sign, signSelf)
		return errors.New("sign is error!")
	}

	//开始处理
	return processor(req)
}

// 获取签名的字段
func (cli *Client) CreateCashierCallbackRequestParams(req PraxisCashierBackReq) map[string]interface{} {
	params := make(map[string]interface{})

	params["merchant_id"] = cli.Params.MerchantId // Assuming these are package-level variables
	params["application_key"] = cli.Params.ApplicationKey
	params["timestamp"] = req.Timestamp //praxis传过来的.
	params["customer_token"] = req.Customer.CustomerToken
	params["order_id"] = req.Session.OrderID
	params["tid"] = req.Transaction.Tid
	params["currency"] = req.Transaction.Currency
	params["amount"] = req.Transaction.Amount
	params["conversion_rate"] = req.Transaction.ConversionRate
	params["processed_currency"] = req.Transaction.ProcessedCurrency
	params["processed_amount"] = req.Transaction.ProcessedAmount

	return params
}

//========================================================

// 计算callback的resp的签名
func (cli *Client) GenerateCallbackRespGtAuthentication(resp PraxisBackResp) string {
	requestParams := createCashierCallbackResponseParams(resp)
	bsUtil := utils.NewBuildSignatureUtils()
	gtAuthentication := bsUtil.GetGtAuthentication(requestParams, cli.Params.MerchantSecret, utils.SignTypeCallbackResp)
	return gtAuthentication
}

// 对返回的json做签名
func createCashierCallbackResponseParams(resp PraxisBackResp) map[string]interface{} {
	params := make(map[string]interface{})
	params["status"] = resp.Status
	params["timestamp"] = resp.Timestamp
	return params
}
