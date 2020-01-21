package securityUtil

import (
	"crypto"
	"crypto/rand"
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

func packageData(originalData []byte, packageSize int) (r [][]byte) {
	var src = make([]byte, len(originalData))
	copy(src, originalData)

	r = make([][]byte, 0)
	if len(src) <= packageSize {
		return append(r, src)
	}
	for len(src) > 0 {
		var p = src[:packageSize]
		r = append(r, p)
		src = src[packageSize:]
		if len(src) <= packageSize {
			r = append(r, src)
			break
		}
	}
	return r
}

// RSAEncrypt 数据加密
// plaintext:待加密的数据
// key:公钥
// 返回值:
// []byte:加密后的数据
// error:错误信息
func RSAEncrypt(plaintext, key []byte) ([]byte, error) {
	var err error
	var block *pem.Block
	block, _ = pem.Decode(key)
	if block == nil {
		return nil, errors.New("public key error")
	}

	var pubInterface interface{}
	pubInterface, err = x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	var pub = pubInterface.(*rsa.PublicKey)

	var data = packageData(plaintext, pub.N.BitLen()/8-11)
	var cipherData []byte = make([]byte, 0, 0)

	for _, d := range data {
		var c, e = rsa.EncryptPKCS1v15(rand.Reader, pub, d)
		if e != nil {
			return nil, e
		}
		cipherData = append(cipherData, c...)
	}

	return cipherData, nil
}

// RSADecrypt 数据解密
// plaintext:待解密的数据
// key:私钥
// 返回值:
// []byte:解密后的数据
// error:错误信息
func RSADecrypt(ciphertext, key []byte) ([]byte, error) {
	var err error
	var block *pem.Block
	block, _ = pem.Decode(key)
	if block == nil {
		return nil, errors.New("private key error")
	}

	var pri *rsa.PrivateKey
	pri, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	var data = packageData(ciphertext, pri.PublicKey.N.BitLen()/8)
	var plainData []byte = make([]byte, 0, 0)

	for _, d := range data {
		var p, e = rsa.DecryptPKCS1v15(rand.Reader, pri, d)
		if e != nil {
			return nil, e
		}
		plainData = append(plainData, p...)
	}
	return plainData, nil
}

// 基于RSA PKCS1V15进行签名
// src:待签名的原始字符串
// key:签名用的私钥
// hash:签名配置
// 返回值:
// []byte:得到的签名字节流
// error:错误信息
func SignPKCS1v15(src, key []byte, hash crypto.Hash) ([]byte, error) {
	var h = hash.New()
	h.Write(src)
	var hashed = h.Sum(nil)

	var err error
	var block *pem.Block
	block, _ = pem.Decode(key)
	if block == nil {
		return nil, errors.New("private key error")
	}

	var pri *rsa.PrivateKey
	pri, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.SignPKCS1v15(rand.Reader, pri, hash, hashed)
}

// 验证RSA PKCS1V15进行签名
// src:待签名的原始字符串
// sig:签名字节流
// key:签名用的私钥
// hash:签名配置
// 返回值:
// error:错误信息
func VerifyPKCS1v15(src, sig, key []byte, hash crypto.Hash) error {
	var h = hash.New()
	h.Write(src)
	var hashed = h.Sum(nil)

	var err error
	var block *pem.Block
	block, _ = pem.Decode(key)
	if block == nil {
		return errors.New("public key error")
	}

	var pubInterface interface{}
	pubInterface, err = x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	var pub = pubInterface.(*rsa.PublicKey)

	return rsa.VerifyPKCS1v15(pub, hash, hashed, sig)
}
