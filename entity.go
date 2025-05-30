package go_praxis

type PraxisInitParams struct {
	MerchantId     string `json:"merchantId" mapstructure:"merchantId" config:"merchantId"`             // merchantId
	MerchantSecret string `json:"merchantSecret" mapstructure:"merchantSecret" config:"merchantSecret"` // accessKey
	ApplicationKey string `json:"applicationKey" mapstructure:"applicationKey" config:"applicationKey"` // merchantSecret
	ApiVersion     string `json:"apiVersion" mapstructure:"apiVersion" config:"apiVersion"`
	ApiLocale      string `json:"apiLocale" mapstructure:"apiLocale" config:"apiLocale"`

	BaseUrl string `json:"baseUrl" mapstructure:"baseUrl" config:"baseUrl"`

	DepositBackUrl    string `json:"depositBackUrl" mapstructure:"depositBackUrl" config:"depositBackUrl"`
	DepositFeBackUrl  string `json:"depositFeBackUrl" mapstructure:"depositFeBackUrl" config:"depositFeBackUrl"`
	WithdrawBackUrl   string `json:"withdrawBackUrl" mapstructure:"withdrawBackUrl" config:"withdrawBackUrl"`
	WithdrawFeBackUrl string `json:"WithdrawFeBackUrl" mapstructure:"WithdrawFeBackUrl" config:"WithdrawFeBackUrl"`
}

// ----------pre generate-------------------------
// https://doc.praxiscashier.com/integration_docs/latest/cashier_api/cashier

type PraxisDepositReq struct {
	//must
	//MerchantID      string                     `json:"merchant_id"`
	//ApplicationKey  string                     `json:"application_key"`
	//Locale          string                     `json:"locale"`
	//Version         string                     `json:"version"`
	//Intent          string `json:"intent"`           //枚举: payment,withdrawal,authorization //这里sdk直接写死
	Currency string `json:"currency"` //币种
	Amount   int    `json:"amount"`   //这个是用 分 为单位的. 有的currency是100分,有的1000分. 具体要看 https://doc.praxiscashier.com/integration_docs/latest/overview/data_formats#currency_fraction
	Cid      string `json:"cid"`      //Unique customer id in your system. 业务系统里的唯一客户id
	//NotificationURL string                    `json:"notification_url"` //回调通知接口
	//ReturnURL       string                    `json:"return_url"`       //前端重定向地址
	OrderID      string                    `json:"order_id"`      //业务系统内的唯一订单id
	CustomerData PraxisDepositCustomerData `json:"customer_data"` //这个也是必填的
	//Timestamp       int64                     `json:"timestamp"`        //seconds
	//option
	CustomerToken string `json:"customer_token"` //客户身份id的HASH
	PaymentMethod string `json:"payment_method"`
	Gateway       string `json:"gateway"`
	ValidationURL string `json:"validation_url"`
}

type PraxisDepositCustomerData struct {
	Country   string `json:"country" mapstructure:"country"`
	FirstName string `json:"first_name" mapstructure:"first_name"`
	LastName  string `json:"last_name" mapstructure:"last_name"`
	DOB       string `json:"dob" mapstructure:"dob"` // Date of birth (format: YYYY-MM-DD)
	Email     string `json:"email" mapstructure:"email"`
	Phone     string `json:"phone" mapstructure:"phone"`     // Should include country code
	Zip       string `json:"zip" mapstructure:"zip"`         // Postal/ZIP code
	State     string `json:"state" mapstructure:"state"`     // State/Province
	City      string `json:"city" mapstructure:"city"`       // City
	Address   string `json:"address" mapstructure:"address"` // Street address
	Profile   int    `json:"profile" mapstructure:"profile"` // Street address
}

type PraxisDepositCardData struct {
	CardNumber string `json:"card_number"` // Encrypted card number (e.g., "ZMq4wDaiaQ/xOwMEcQ7R3ASjTnoOMu+avLuJYgAnz1Q=")
	CardExp    string `json:"card_exp"`    // Encrypted expiration date (e.g., "WI8V4bE5/l8fIhUv6aMO8w==")
	CVV        string `json:"cvv"`         // Encrypted CVV code
}

type PraxisDepositRsp struct {
	Status      int                        `json:"status"` //0是正确， 逻辑错误>0,系统错误<0
	Description string                     `json:"description"`
	Customer    *PraxisDevicePsRspCustomer `json:"customer"`
	Session     *PraxisDevicePsRspSession  `json:"session"`
	Version     string                     `json:"version"`
	Timestamp   int64                      `json:"timestamp"` //seconds
	//option
	RedirectURL string `json:"redirect_url"` //是用户要付款的psp页面的url
}

//---------------------------------------------

type PraxisWithdrawReq struct {
	//MerchantID      string                     `json:"merchant_id"`
	//ApplicationKey  string                     `json:"application_key"`
	//Locale          string                     `json:"locale"`
	//Version         string                     `json:"version"`
	//Intent          string                     `json:"intent"` //这里sdk直接写死
	Currency string `json:"currency"` //币种
	Amount   int    `json:"amount"`   //这个是用 分 为单位的. 有的currency是100分,有的1000分. 具体要看 https://doc.praxiscashier.com/integration_docs/latest/overview/data_formats#currency_fraction
	Cid      string `json:"cid"`      //Unique customer id in your system. 业务系统里的唯一客户id
	//NotificationURL string                    `json:"notification_url"` //回调通知接口
	//ReturnURL       string                    `json:"return_url"`       //前端重定向地址
	OrderID      string                    `json:"order_id"`      //业务系统内的唯一订单id
	CustomerData PraxisDepositCustomerData `json:"customer_data"` //这个也是必填的
	//Timestamp       int64                     `json:"timestamp"` //second
	Balance       int                    `json:"balance"`
	CustomerToken string                 `json:"customer_token"`
	CardData      *PraxisDepositCardData `json:"card_data"`
	PaymentMethod string                 `json:"payment_method"`
	Gateway       string                 `json:"gateway"`
	ValidationURL string                 `json:"validation_url"`
}

type PraxisWithdrawResp struct {
	Status      int                        `json:"status"`
	Description string                     `json:"description"`
	RedirectURL string                     `json:"redirect_url"`
	Customer    *PraxisDevicePsRspCustomer `json:"customer"`
	Session     *PraxisDevicePsRspSession  `json:"session"`
	Version     string                     `json:"version"`
	Timestamp   int64                      `json:"timestamp"`
}

// PraxisDevicePsRspCustomer represents customer data in response
type PraxisDevicePsRspCustomer struct {
	CustomerToken string `json:"customer_token"` //HASH value of customer's identity
	//option
	AVSAlert          int `json:"avs_alert"`
	VerificationAlert int `json:"verification_alert"`
}

// PraxisDevicePsRspSession represents session data in response
type PraxisDevicePsRspSession struct {
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

// ----------deposit callback-------------------------
// https://doc.praxiscashier.com/integration_docs/latest/webhooks/notification

// 回调的入参
type PraxisBackReq struct {
	//must
	MerchantID     string                        `json:"merchant_id"`
	ApplicationKey string                        `json:"application_key"`
	Customer       *PraxisBackReqCustomerData    `json:"customer"`
	Session        *PraxisBackReqSessionData     `json:"session"`
	Transaction    *PraxisBackReqTransactionData `json:"transaction"`
	Version        string                        `json:"version"`
	Timestamp      int64                         `json:"timestamp"`
}

type PraxisBackReqCustomerData struct {
	CustomerToken     string `json:"customer_token"` //must
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	AVSAlert          *int   `json:"avs_alert,omitempty"`          // 地址校验收通过 0成功 1失败
	VerificationAlert *int   `json:"verification_alert,omitempty"` // 客户验证是否通过  0通过
}

type PraxisBackReqSessionData struct {
	AuthToken         string  `json:"auth_token"`     //must
	Intent            string  `json:"intent"`         //must
	SessionStatus     string  `json:"session_status"` //must
	OrderID           string  `json:"order_id"`       //must
	Currency          string  `json:"currency"`       //must
	Amount            float64 `json:"amount,omitempty"`
	ConversionRate    float64 `json:"conversion_rate,omitempty"`
	ProcessedCurrency string  `json:"processed_currency"` //must
	ProcessedAmount   float64 `json:"processed_amount"`   //must
	PaymentMethod     string  `json:"payment_method,omitempty"`
	Gateway           string  `json:"gateway,omitempty"`
	Cid               string  `json:"cid"`                 //must
	Variable1         string  `json:"variable1,omitempty"` // omitempty if empty
	Variable2         string  `json:"variable2,omitempty"` // omitempty if empty
	Variable3         string  `json:"variable3,omitempty"` // omitempty if empty
}

type PraxisBackReqTransactionData struct {
	TransactionType   string                 `json:"transaction_type"`
	TransactionStatus string                 `json:"transaction_status"`
	Tid               int                    `json:"tid"`
	TransactionID     string                 `json:"transaction_id"`
	Currency          string                 `json:"currency"`
	Amount            int                    `json:"amount"`
	ConversionRate    float64                `json:"conversion_rate"` //TODO 文档中有冲突,说明是string,例子是float
	ProcessedCurrency string                 `json:"processed_currency"`
	ProcessedAmount   int                    `json:"processed_amount"`
	Fee               int                    `json:"fee"`
	FeeIncluded       int                    `json:"fee_included"`
	FeeType           string                 `json:"fee_type"`
	PaymentMethod     string                 `json:"payment_method"`
	PaymentProcessor  string                 `json:"payment_processor"`
	Gateway           string                 `json:"gateway"`
	Card              *PraxisBackReqCardData `json:"card,omitempty"`
}

// PraxisDepositBackReqCardData represents card data in transaction
type PraxisBackReqCardData struct {
	CardToken      string `json:"card_token"`
	CardType       string `json:"card_type"`
	CardNumber     string `json:"card_number"`
	CardExp        string `json:"card_exp"`
	CardIssuerName string `json:"card_issuer_name"`
	CardHolder     string `json:"card_holder"`
}

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
