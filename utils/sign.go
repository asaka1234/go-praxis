package utils

import (
	"crypto/sha512"
	"encoding/hex"
	"github.com/spf13/cast"
)

type SignType int //签名的具体点位

const (
	SignTypeSendReq      SignType = 1 //是发送pre-order时:对请求做签名
	SignTypeCallbackReq  SignType = 2 //是收到回调的req的签名
	SignTypeCallbackResp SignType = 3 //是收到回调后, 处理完毕再发送给psp的resp的签名
)

// BuildSignatureUtils provides signature generation functionality
type BuildSignatureUtils struct{}

// NewBuildSignatureUtils creates a new instance (though we'll mostly use static methods)
func NewBuildSignatureUtils() *BuildSignatureUtils {
	return &BuildSignatureUtils{}
}

// ---------------------------------------------
// Cashier API 1.3
// 发请求给psp三方
func (h *BuildSignatureUtils) GetRequestSignatureList() []string {
	return []string{
		"merchant_id",
		"application_key",
		"timestamp",
		"intent",
		"cid",
		"order_id",
	}
}

// psp回调通知时的签名字段，也是放在header里
func (h *BuildSignatureUtils) GetCallbackRequestSignatureList() []string {
	return []string{
		"merchant_id",
		"application_key",
		"timestamp",
		"customer_token",     //customer
		"order_id",           //session
		"tid",                //transaction
		"currency",           //transaction
		"amount",             //transaction
		"conversion_rate",    //transaction
		"processed_currency", //transaction
		"processed_amount",   //transaction
	}
}

// 对callback的response的签名
func (h *BuildSignatureUtils) GetCallbackResponseSignatureList() []string {
	return []string{
		"status",
		"timestamp",
	}
}

//------------------------------------

func (h *BuildSignatureUtils) GetConcatenatedString(data map[string]interface{}, endPoint SignType) string {
	concatenatedString := ""

	//-------------拿到签名key列表-------------------
	keyList := make([]string, 0)
	if endPoint == SignTypeSendReq {
		keyList = h.GetRequestSignatureList()
	} else if endPoint == SignTypeCallbackReq {
		keyList = h.GetCallbackRequestSignatureList()
	} else if endPoint == SignTypeCallbackResp {
		keyList = h.GetCallbackResponseSignatureList()
	}

	//------------拼凑签名的原始字符串-------------------
	for _, key := range keyList {
		if val, exists := data[key]; exists && val != nil {
			concatenatedString += cast.ToString(data[key])
		}
	}
	return concatenatedString
}

// which 决定了是哪种
func (h *BuildSignatureUtils) GetGtAuthentication(request map[string]interface{}, merchantSecret string, endPoint SignType) string {
	// Sort request array by keys ASC
	concatenatedString := h.GetConcatenatedString(request, endPoint)

	// Concatenate Merchant Secret Key with response params
	concatenatedString += merchantSecret
	//fmt.Printf("concatenatedString======>%s\n", concatenatedString)

	// Generate HASH of concatenated string
	signature := h.GenerateSignature(concatenatedString)
	return signature
}

func (h *BuildSignatureUtils) GenerateSignature(input string) string {
	hash := sha512.New384()
	hash.Write([]byte(input))
	hashText := hex.EncodeToString(hash.Sum(nil))
	return hashText
}
