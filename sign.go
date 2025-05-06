package go_praxis

import (
	"crypto/sha512"
	"encoding/hex"
	"github.com/mitchellh/mapstructure"
	"github.com/samber/lo"
	"github.com/spf13/cast"
)

// 参与签名的所有字段
func getRequestSignatureList() []string {
	return []string{
		"merchant_id",
		"application_key",
		"timestamp",
		"intent",
		"cid",
		"order_id",
	}
}

func getConcatenatedString(req CashierRequest, requestSignatureList []string) string {

	// 将结构体转换为map
	var result map[string]interface{}
	err := mapstructure.Decode(req, &result)
	if err != nil {
		panic(err)
	}

	//找出对应key的值来
	rawString := ""
	for _, param := range requestSignatureList {
		if lo.HasKey(result, param) && result[param] != nil {
			rawString = rawString + cast.ToString(result[param]) //timestamp是int，这里可能转换失败
		}
	}

	return rawString
}

// 用sha 384计算签名
func generateSignature(input string) string {
	hash := sha512.New384()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}

// export ========>
func CalSign(req CashierRequest) string {
	raw := getConcatenatedString(req, getRequestSignatureList())
	return generateSignature(raw)
}
