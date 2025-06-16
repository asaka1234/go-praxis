package go_praxis

type PraxisInitParams struct {
	MerchantId     string `json:"merchantId" mapstructure:"merchantId" config:"merchantId"  yaml:"merchantId"`                 // merchantId
	MerchantSecret string `json:"merchantSecret" mapstructure:"merchantSecret" config:"merchantSecret"  yaml:"merchantSecret"` // accessKey
	ApplicationKey string `json:"applicationKey" mapstructure:"applicationKey" config:"applicationKey"  yaml:"applicationKey"` // merchantSecret
	ApiVersion     string `json:"apiVersion" mapstructure:"apiVersion" config:"apiVersion"  yaml:"apiVersion"`
	ApiLocale      string `json:"apiLocale" mapstructure:"apiLocale" config:"apiLocale"  yaml:"apiLocale"`

	BaseUrl string `json:"baseUrl" mapstructure:"baseUrl" config:"baseUrl"  yaml:"baseUrl"`

	DepositBackUrl    string `json:"depositBackUrl" mapstructure:"depositBackUrl" config:"depositBackUrl"  yaml:"depositBackUrl"`
	DepositFeBackUrl  string `json:"depositFeBackUrl" mapstructure:"depositFeBackUrl" config:"depositFeBackUrl"  yaml:"depositFeBackUrl"`
	WithdrawBackUrl   string `json:"withdrawBackUrl" mapstructure:"withdrawBackUrl" config:"withdrawBackUrl"  yaml:"withdrawBackUrl"`
	WithdrawFeBackUrl string `json:"WithdrawFeBackUrl" mapstructure:"WithdrawFeBackUrl" config:"WithdrawFeBackUrl"  yaml:"WithdrawFeBackUrl"`
}

// ----------pre generate-------------------------
// https://doc.praxiscashier.com/integration_docs/latest/cashier_api/cashier

type PraxisCashierReq struct {
	Currency     string                    `json:"currency" mapstructure:"currency"`
	Amount       int                       `json:"amount" mapstructure:"amount"`
	OrderID      string                    `json:"order_id" mapstructure:"order_id"` //商户订单号
	Cid          string                    `json:"cid" mapstructure:"cid"`           //商户uid
	CustomerData PraxisCashierCustomerData `json:"customer_data" mapstructure:"customer_data"`
	//如下由sdk来设置
	//MerchantId      string `json:"merchant_id"`
	//ApplicationKey  string `json:"application_key"`
	//Intent          string `json:"intent"` //枚举: payment,withdrawal,authorization //这里sdk直接写死
	//Timestamp       int    `json:"timestamp"`
	//Version         string `json:"version"`
	//Locale          string `json:"locale"`
	//NotificationUrl string `json:"notification_url"`
	//ReturnUrl       string `json:"return_url"`
}

type PraxisCashierCustomerData struct {
	Country   string `json:"country" mapstructure:"country"`
	FirstName string `json:"first_name" mapstructure:"first_name"`
	LastName  string `json:"last_name" mapstructure:"last_name"`
	Email     string `json:"email" mapstructure:"email"`
	//option
	//DOB     string `json:"dob" mapstructure:"dob"`         // Date of birth (format: YYYY-MM-DD)
	//Phone   string `json:"phone" mapstructure:"phone"`     // Should include country code
	//Zip     string `json:"zip" mapstructure:"zip"`         // Postal/ZIP code
	//State   string `json:"state" mapstructure:"state"`     // State/Province
	//City    string `json:"city" mapstructure:"city"`       // City
	//Address string `json:"address" mapstructure:"address"` // Street address
	//Profile int    `json:"profile" mapstructure:"profile"` // Street address
}

//--------------------------------------------------------------------

type PraxisCashierResp struct {
	Status      int                    `json:"status"` //0是正确， 逻辑错误>0,系统错误<0
	Description string                 `json:"description"`
	Customer    *PraxisCashierCustomer `json:"customer"`
	Session     *PraxisCashierSession  `json:"session"`
	Version     string                 `json:"version"`
	Timestamp   int64                  `json:"timestamp"` //seconds
	//option
	RedirectURL string `json:"redirect_url"` //是用户要付款的psp页面的url
}

type PraxisCashierCustomer struct {
	CustomerToken string `json:"customer_token"` //HASH value of customer's identity
	//option
	AVSAlert          int `json:"avs_alert"`
	VerificationAlert int `json:"verification_alert"`
}

type PraxisCashierSession struct {
	Amount            float64 `json:"amount"`
	AuthToken         string  `json:"auth_token"`
	Cid               string  `json:"cid"`
	Currency          string  `json:"currency"`
	Intent            string  `json:"intent"`
	OrderID           string  `json:"order_id"`
	ProcessedAmount   float64 `json:"processed_amount"`
	ProcessedCurrency string  `json:"processed_currency"`
	SessionStatus     string  `json:"session_status"`
}

// ---------- callback---------------------------------------------------------

// https://doc.praxiscashier.com/integration_docs/latest/webhooks/notification

// 回调的入参
type PraxisCashierBackReq struct {
	MerchantID     string                               `json:"merchant_id" mapstructure:"merchant_id"`
	ApplicationKey string                               `json:"application_key" mapstructure:"application_key"`
	Customer       *PraxisCashierBackReqCustomerData    `json:"customer" mapstructure:"customer"`
	Session        *PraxisCashierBackReqSessionData     `json:"session" mapstructure:"session"`
	Transaction    *PraxisCashierBackReqTransactionData `json:"transaction" mapstructure:"transaction"`
	Version        string                               `json:"version" mapstructure:"version"`
	Timestamp      int64                                `json:"timestamp" mapstructure:"timestamp"`
}

type PraxisCashierBackReqCustomerData struct {
	CustomerToken     string `json:"customer_token" mapstructure:"customer_token"` //must
	FirstName         string `json:"first_name" mapstructure:"first_name"`
	LastName          string `json:"last_name" mapstructure:"last_name"`
	AVSAlert          *int   `json:"avs_alert,omitempty" mapstructure:"avs_alert"`                   // 地址校验收通过 0成功 1失败
	VerificationAlert *int   `json:"verification_alert,omitempty" mapstructure:"verification_alert"` // 客户验证是否通过  0通过
}

type PraxisCashierBackReqSessionData struct {
	AuthToken         string  `json:"auth_token" mapstructure:"auth_token"`         //must
	Intent            string  `json:"intent" mapstructure:"intent"`                 //must 用来区分是充值还是提现. 充值->payment, 提现->withdrawal)
	SessionStatus     string  `json:"session_status" mapstructure:"session_status"` //must
	OrderID           string  `json:"order_id" mapstructure:"order_id"`             //must  是商户的订单号 (Transaction identifier in your system)
	Currency          string  `json:"currency" mapstructure:"currency"`             //must
	Amount            float64 `json:"amount,omitempty" mapstructure:"amount"`
	ConversionRate    float64 `json:"conversion_rate,omitempty" mapstructure:"conversion_rate"`
	ProcessedCurrency string  `json:"processed_currency" mapstructure:"processed_currency"` //must
	ProcessedAmount   float64 `json:"processed_amount" mapstructure:"processed_amount"`     //must
	PaymentMethod     string  `json:"payment_method,omitempty" mapstructure:"payment_method"`
	Gateway           string  `json:"gateway,omitempty" mapstructure:"gateway"`
	Cid               string  `json:"cid" mapstructure:"cid"`                       //must 商户的user id
	Variable1         string  `json:"variable1,omitempty" mapstructure:"variable1"` // omitempty if empty
	Variable2         string  `json:"variable2,omitempty" mapstructure:"variable2"` // omitempty if empty
	Variable3         string  `json:"variable3,omitempty" mapstructure:"variable3"` // omitempty if empty
}

type PraxisCashierBackReqTransactionData struct {
	TransactionType   string                 `json:"transaction_type" mapstructure:"transaction_type"`
	TransactionStatus string                 `json:"transaction_status" mapstructure:"transaction_status"`
	Tid               int                    `json:"tid" mapstructure:"tid"`
	TransactionID     string                 `json:"transaction_id" mapstructure:"transaction_id"`
	Currency          string                 `json:"currency" mapstructure:"currency"`
	Amount            int                    `json:"amount" mapstructure:"amount"`
	ConversionRate    interface{}            `json:"conversion_rate" mapstructure:"conversion_rate"` //TODO 文档中有冲突,说明是string,例子是float
	ProcessedCurrency string                 `json:"processed_currency" mapstructure:"processed_currency"`
	ProcessedAmount   int                    `json:"processed_amount" mapstructure:"processed_amount"`
	Fee               int                    `json:"fee" mapstructure:"fee"`
	FeeIncluded       int                    `json:"fee_included" mapstructure:"fee_included"`
	FeeType           string                 `json:"fee_type" mapstructure:"fee_type"`
	PaymentMethod     string                 `json:"payment_method" mapstructure:"payment_method"`
	PaymentProcessor  string                 `json:"payment_processor" mapstructure:"payment_processor"`
	Gateway           string                 `json:"gateway" mapstructure:"gateway"`
	Card              *PraxisBackReqCardData `json:"card,omitempty" mapstructure:"card"`
}

type PraxisBackReqCardData struct {
	CardToken      string `json:"card_token" mapstructure:"card_token"`
	CardType       string `json:"card_type" mapstructure:"card_type"`
	CardNumber     string `json:"card_number" mapstructure:"card_number"`
	CardExp        string `json:"card_exp" mapstructure:"card_exp"`
	CardIssuerName string `json:"card_issuer_name" mapstructure:"card_issuer_name"`
	CardHolder     string `json:"card_holder" mapstructure:"card_holder"`
}

//-------------------------------------------------------------------

/*
{
    "merchant_id": "Test-Integration-Merchant",
    "application_key": "Sandbox",
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
}

*/
// =============
// 返回给三方的
type PraxisBackResp struct {
	Status      int    `json:"status"` //0成功, 非0失败
	Description string `json:"description"`
	Version     string `json:"version"`
	Timestamp   int64  `json:"timestamp"`
}
