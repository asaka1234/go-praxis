package go_praxis

import (
	"encoding/json"
	"fmt"
	"testing"
)

type VLog struct {
}

func (l VLog) Debugf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Infof(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Warnf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Errorf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}

func TestCallback(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &PraxisInitParams{MERCHANT_ID, MERCHANT_SECRET, APPLICATION_KEY, API_VERSION, API_LOCALE, SANDBOX_URL, DepositBackUrl, DepositFeBackUrl, WithdrawBackUrl, WithdrawFeBackUrl})

	//1. 获取请求
	req := GenCallbackRequestDemo()
	var backReq PraxisCashierBackReq
	err := json.Unmarshal([]byte(req), &backReq)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	//2. 处理请求

	//发请求
	err = cli.CashierCallback(backReq, "sign", func(PraxisCashierBackReq) error { return nil })
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
	//fmt.Printf("resp:%+v\n", resp.Session)

}

func GenCallbackRequestDemo() string {
	return `{
    "merchant_id": "API-cptinternational2",
    "application_key": "cptinternational",
    "conversion_rate": 1.000000,
    "customer": {
        "customer_token": "87cfb23a8f1e68e162c276b754d9c061",
        "country": "GB",
        "first_name": "John",
        "last_name": "Johnson",
        "avs_alert": 0,
        "verification_alert": null
    },
    "session": {
        "auth_token": "8a7sd87a8sd778ac961062c6bedddb8",
        "intent": "payment",
        "session_status": "created",
        "order_id": "test-1560610955",
        "currency": "EUR",
        "amount": 100,
        "conversion_rate": 1.000000,
        "processed_currency": "EUR",
        "processed_amount": 100,
        "payment_method": "Credit Card",
        "gateway": null,
        "pin": "1",
        "variable1": "your variable",
        "variable2": "if that is not enough, you can pass even one more variable",
        "variable3": null
    },
    "transaction": {
        "transaction_type": "sale",
        "transaction_status": "approved",
        "tid": 756850,
        "transaction_id": "13348",
        "currency": "EUR",
        "amount": 100,
        "conversion_rate": 1.000000,
        "processed_currency": "EUR",
        "processed_amount": 100,
        "fee": 0,
        "fee_included": 0,
        "fee_type": "flat",
        "payment_method": "Credit Card",
        "payment_processor": "TestCardProcessor",
        "gateway": "s-pTSZyK23E1Ee5KZpcNbX_aFl0HuhQ0",
        "card": {
            "card_token": "J-4-a0vPhjZ9R75JP98VDUFgbh9y8sYr",
            "card_type": "VISA",
            "card_number": "411111******1111",
            "card_exp": "12\/2024",
            "card_issuer_name": "Bank of Somewhere",
            "card_issuer_country": "GB",
            "card_holder": "John Johnson"
        },
        "wallet": null,
        "is_async": 0,
        "is_cascade": 0,
        "cascade_level": 0,
        "reference_id": null,
        "withdrawal_request_id": null,
        "created_by": "INTERNET",
        "edited_by": "INTERNET",
        "status_code": "SC-002",
        "status_details": "Transaction approved"
    },
    "version": "1.3",
    "timestamp": 1590611635
}`
}
