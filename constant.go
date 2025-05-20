package go_praxis

const SIGN_HEAD_NAME = "Gt-Authentication" // 签名的header名字

//-----------------------------------------------------

// 提交类型
type IntentType string //签名的具体点位

const (
	IntentTypePayment       IntentType = "payment"
	IntentTypeWithdrawal    IntentType = "withdrawal"
	IntentTypeAuthorization IntentType = "authorization"
)
