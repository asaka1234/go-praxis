package go_praxis

import "github.com/samber/lo"

const SIGN_HEAD_NAME = "Gt-Authentication" // 签名的header名字

//-----------------------------------------------------

// 提交类型
type IntentType string //签名的具体点位

const (
	IntentTypePayment       IntentType = "payment"
	IntentTypeWithdrawal    IntentType = "withdrawal"
	IntentTypeAuthorization IntentType = "authorization"
)

//-----------------------------------------------------

// 1元=1000分
var CurrencyFractionThousandList = []string{"BHD", "IQD", "JOD", "KWD", "LYD", "OMR", "TND"}

// 1元=1分
var CurrencyFractionOneList = []string{"BIF", "CLP", "DJF", "GNF", "ISK", "JPY", "KMF", "KRW", "PYG", "RWF", "UGX", "UYI", "VND", "VUV", "XAF", "XOF", "XPF"}

//其他都是1元=100分

func GetCurrencyFraction(symbol string) int {
	if -1 != lo.IndexOf(CurrencyFractionThousandList, symbol) {
		return 1000
	}

	if -1 != lo.IndexOf(CurrencyFractionOneList, symbol) {
		return 1
	}
	
	return 100
}
