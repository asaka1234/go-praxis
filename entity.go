package go_praxis

// ----------pre generate-------------------------
// https://doc.praxiscashier.com/integration_docs/latest/cashier_api/cashier

type CashierRequest struct {
	MerchantId     string      `json:"merchant_id"`
	ApplicationKey string      `json:"application_key"`
	Intent         string      `json:"intent"`
	Currency       string      `json:"currency"`
	Amount         int         `json:"amount"`
	Cid            string      `json:"cid"`
	Locale         string      `json:"locale"`
	CustomerToken  interface{} `json:"customer_token"`
	CustomerData   struct {    //optional
		Country   string `json:"country"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Dob       string `json:"dob"`
		Email     string `json:"email"`
		Phone     string `json:"phone"`
		Zip       string `json:"zip"`
		State     string `json:"state"`
		City      string `json:"city"`
		Address   string `json:"address"`
	} `json:"customer_data"`
	PaymentMethod   string `json:"payment_method"` //optional
	Gateway         string `json:"gateway"`        //optional
	ValidationUrl   string `json:"validation_url"` //optional
	NotificationUrl string `json:"notification_url"`
	ReturnUrl       string `json:"return_url"`
	OrderId         string `json:"order_id"`

	Variable1 string      `json:"variable1"` //optional
	Variable2 string      `json:"variable2"` //optional
	Variable3 interface{} `json:"variable3"` //optional

	Version   string `json:"version"`
	Timestamp int    `json:"timestamp"`
}

type CashierResponse struct {
	Status      int    `json:"status"`
	Description string `json:"description"`
	RedirectUrl string `json:"redirect_url"`
	Customer    struct {
		CustomerToken     string `json:"customer_token"`
		Country           string `json:"country"`
		FirstName         string `json:"first_name"`
		LastName          string `json:"last_name"`
		AvsAlert          int    `json:"avs_alert"`
		VerificationAlert int    `json:"verification_alert"`
	} `json:"customer"`
	Session struct {
		AuthToken         string      `json:"auth_token"`
		Intent            string      `json:"intent"`
		SessionStatus     string      `json:"session_status"`
		OrderId           string      `json:"order_id"`
		Currency          string      `json:"currency"`
		Amount            int         `json:"amount"`
		ConversionRate    float64     `json:"conversion_rate"`
		ProcessedCurrency string      `json:"processed_currency"`
		ProcessedAmount   int         `json:"processed_amount"`
		PaymentMethod     string      `json:"payment_method"`
		Gateway           interface{} `json:"gateway"`
		Cid               string      `json:"cid"`
		Variable1         string      `json:"variable1"`
		Variable2         string      `json:"variable2"`
		Variable3         interface{} `json:"variable3"`
	} `json:"session"`
	Version   string `json:"version"`
	Timestamp int    `json:"timestamp"`
}

// ----------deposit callback-------------------------
// https://doc.praxiscashier.com/integration_docs/latest/webhooks/notification

type CallbackRequest struct {
	MerchantId     string  `json:"merchant_id"`
	ApplicationKey string  `json:"application_key"`
	ConversionRate float64 `json:"conversion_rate"`
	Customer       struct {
		CustomerToken     string      `json:"customer_token"`
		Country           string      `json:"country"`
		FirstName         string      `json:"first_name"`
		LastName          string      `json:"last_name"`
		AvsAlert          int         `json:"avs_alert"`
		VerificationAlert interface{} `json:"verification_alert"`
	} `json:"customer"`
	Session struct {
		AuthToken         string      `json:"auth_token"`
		Intent            string      `json:"intent"`
		SessionStatus     string      `json:"session_status"`
		OrderId           string      `json:"order_id"`
		Currency          string      `json:"currency"`
		Amount            int         `json:"amount"`
		ConversionRate    float64     `json:"conversion_rate"`
		ProcessedCurrency string      `json:"processed_currency"`
		ProcessedAmount   int         `json:"processed_amount"`
		PaymentMethod     string      `json:"payment_method"`
		Gateway           interface{} `json:"gateway"`
		Pin               string      `json:"pin"`
		Variable1         string      `json:"variable1"`
		Variable2         string      `json:"variable2"`
		Variable3         interface{} `json:"variable3"`
	} `json:"session"`
	Transaction struct {
		TransactionType   string  `json:"transaction_type"`
		TransactionStatus string  `json:"transaction_status"`
		Tid               int     `json:"tid"`
		TransactionId     string  `json:"transaction_id"`
		Currency          string  `json:"currency"`
		Amount            int     `json:"amount"`
		ConversionRate    float64 `json:"conversion_rate"`
		ProcessedCurrency string  `json:"processed_currency"`
		ProcessedAmount   int     `json:"processed_amount"`
		Fee               int     `json:"fee"`
		FeeIncluded       int     `json:"fee_included"`
		FeeType           string  `json:"fee_type"`
		PaymentMethod     string  `json:"payment_method"`
		PaymentProcessor  string  `json:"payment_processor"`
		Gateway           string  `json:"gateway"`
		Card              struct {
			CardToken         string `json:"card_token"`
			CardType          string `json:"card_type"`
			CardNumber        string `json:"card_number"`
			CardExp           string `json:"card_exp"`
			CardIssuerName    string `json:"card_issuer_name"`
			CardIssuerCountry string `json:"card_issuer_country"`
			CardHolder        string `json:"card_holder"`
		} `json:"card"`
		Wallet              interface{} `json:"wallet"`
		IsAsync             int         `json:"is_async"`
		IsCascade           int         `json:"is_cascade"`
		CascadeLevel        int         `json:"cascade_level"`
		ReferenceId         interface{} `json:"reference_id"`
		WithdrawalRequestId interface{} `json:"withdrawal_request_id"`
		CreatedBy           string      `json:"created_by"`
		EditedBy            string      `json:"edited_by"`
		StatusCode          string      `json:"status_code"`
		StatusDetails       string      `json:"status_details"`
	} `json:"transaction"`
	Version   string `json:"version"`
	Timestamp int    `json:"timestamp"`
}
