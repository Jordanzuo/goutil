package securityUtil

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
)

// Hmac-SHA1编码
// source：编码原数据
// key：编码密钥
// 返回值：编码结果
func HmacSha1(source, key string) (result []byte, err error) {
	mac := hmac.New(sha1.New, []byte(key))
	if _, err = mac.Write([]byte(source)); err != nil {
		return
	}

	return mac.Sum(nil), nil
}

// Hmac-SHA1 Base64编码
// source：编码原数据
// key：编码密钥
// 返回值：编码结果
func Base64HmacSha1(source, key string) (result string, err error) {
	mac := hmac.New(sha1.New, []byte(key))
	if _, err = mac.Write([]byte(source)); err != nil {
		return
	}

	bytes := mac.Sum(nil)
	result = base64.StdEncoding.EncodeToString(bytes)
	return
}

// Hmac-SHA256编码
// source：编码原数据
// key：编码密钥
// 返回值：编码结果
func HmacSha256(source, key string) (result []byte, err error) {
	mac := hmac.New(sha256.New, []byte(key))
	if _, err = mac.Write([]byte(source)); err != nil {
		return
	}

	return mac.Sum(nil), nil
}

// Hmac-SHA512编码
// source：编码原数据
// key：编码密钥
// 返回值：编码结果
func HmacSha512(source, key string) (result []byte, err error) {
	mac := hmac.New(sha512.New, []byte(key))
	if _, err = mac.Write([]byte(source)); err != nil {
		return
	}

	return mac.Sum(nil), nil
}
