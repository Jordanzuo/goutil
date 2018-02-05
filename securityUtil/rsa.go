package securityUtil

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"strings"
)

// RsaSha256验证签名
// publicKey:公钥
// source:原字符串
// targetSign:用以验证的目标签名（会进行Base64解码）
// 返回值：返回nil为成功
func VerifyRsaWithSha256(publicKey, source, targetSign string) error {
	//未加标记的PEM密钥加上公钥标记
	if !strings.HasPrefix(publicKey, "-") {
		publicKey = fmt.Sprintf(`-----BEGIN PUBLIC KEY-----
%s
-----END PUBLIC KEY-----`, publicKey)
	}

	//base64解码目标签名
	signBuf, err := base64.StdEncoding.DecodeString(targetSign)
	if err != nil {
		return err
	}

	//解码公钥
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		return fmt.Errorf("RSA.VerifyRsaWithSha256,Block is nil,public key error!")
	}

	//转化为公钥对象
	pubObj, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}

	pub := pubObj.(*rsa.PublicKey)

	//验证签名
	err = rsa.VerifyPKCS1v15(pub, crypto.SHA256, Sha256Bytes([]byte(source)), signBuf)

	return err
}

// RsaSha1验证签名
// publicKey:公钥
// source:原字符串
// targetSign:用以验证的目标签名（会进行Base64解码）
// 返回值：返回nil为成功
func VerifyRsaWithSha1(publicKey, source, targetSign string) error {
	//未加标记的PEM密钥加上公钥标记
	if !strings.HasPrefix(publicKey, "-") {
		publicKey = fmt.Sprintf(`-----BEGIN PUBLIC KEY-----
%s
-----END PUBLIC KEY-----`, publicKey)
	}

	//base64解码目标签名
	signBuf, err := base64.StdEncoding.DecodeString(targetSign)
	if err != nil {
		return err
	}

	//解码公钥
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		return fmt.Errorf("RSA.VerifyRsaWithSha1,Block is nil,public key error!")
	}

	//转化为公钥对象
	pubObj, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}

	pub := pubObj.(*rsa.PublicKey)

	//验证签名
	err = rsa.VerifyPKCS1v15(pub, crypto.SHA1, Sha1StringToBytes(source), signBuf)

	return err
}

// 对字符数组进行SHA1加密
// b:输入字符数组
// 返回值：sha1加密后的原数据
func Sha1StringToBytes(b string) []byte {
	if len(b) == 0 {
		panic(errors.New("input []byte can't be empty"))
	}

	sha1Instance := sha1.New()
	sha1Instance.Write([]byte(b))
	result := sha1Instance.Sum(nil)

	return result
}
