package stringUtil

import (
	"encoding/base64"
)

const (
	base64Table = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

// 	const encodeStd = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
// const encodeURL = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"
)

var coder = base64.NewEncoding(base64Table)

// 对字符串进行Base64编码
func Base64Encode(src string) string {
	if src == "" {
		return src
	}

	return base64.StdEncoding.EncodeToString([]byte(src))
}

// 对字符串进行Base64解码
func Base64Encode2(src []byte) []byte {
	if len(src) == 0 {
		return src
	}

	return []byte(base64.StdEncoding.EncodeToString(src))
}

// 对字符数组进行Base64编码
func Base64Decode(src string) (string, error) {
	if src == "" {
		return src, nil
	}

	bytes, err := coder.DecodeString(src)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

// 对字符数组进行Base64解码
func Base64Decode2(src []byte) ([]byte, error) {
	if len(src) == 0 {
		return src, nil
	}

	return coder.DecodeString(string(src))
}
