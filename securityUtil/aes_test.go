package securityUtil

import (
	"encoding/base64"
	"fmt"
	"testing"
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

func TestEncryptAndDecrypt(t *testing.T) {
	x := []byte("6d15fbbf-4913-40f5-afd1-c02effc2373a")
	key := []byte("WB6aEKK5LoGpetJv")
	x1, _ := AESEncrypt_CBC_Pkcs7(x, key)
	x1_base64 := Base64Encode2(x1)
	x1_base64_str := string(x1_base64)
	fmt.Printf("Base64 of encrypted data:%s\n", x1_base64_str)

	// x1_base64_str = "OoV781MTCRIEKBaWDn4NDuS3Iq1stwnORQA30Ip/eewTEzaNDQl/TgVQU09Bm7pcIP6GxGfzO7vKhRITgKCghpTi9/D+oz/GdKn8KjF/gmE="
	x2_init_str, _ := Base64Decode(x1_base64_str)
	x2_init := []byte(x2_init_str)
	fmt.Printf("x2_init:%d\n", len(x2_init))

	x2, _ := AESDecrypt_CBC_Pkcs7(x2_init, key)
	fmt.Printf("Decrypted data:%s\n", string(x2))
	if string(x) != string(x2) {
		t.Errorf("Expected %s, but got %s", string(x), string(x2))
	}
}

func TestDecrypt(t *testing.T) {
	key := []byte("WB6aEKK5LoGpetJv")
	x2_init_str, _ := Base64Decode("Oit72+aWraykW7i0e/q+zZ77w5yEU/5KuNpRYoJaxpw93i6zWKiFB8c6/PIedZxz")
	x2_init := []byte(x2_init_str)
	x2, _ := AESDecrypt_CBC_Pkcs7(x2_init, key)
	fmt.Printf("Decrypted data:%s\n", string(x2))

	x2_init_str, _ = Base64Decode("r0KHpfSmQ8jmx/FR4IJPOGBLTYF9lDRWbo9P8lIwekjkkU8BOO0QvfypgHZRIJWS")
	x2_init = []byte(x2_init_str)
	x2, _ = AESDecrypt_CBC_Pkcs7(x2_init, key)
	fmt.Printf("Decrypted data:%s\n", string(x2))

	x2_init_str, _ = Base64Decode("ev2P8zTEnebpBm43Dd8YlA==")
	x2_init = []byte(x2_init_str)
	x2, _ = AESDecrypt_CBC_Pkcs7(x2_init, key)
	fmt.Printf("Decrypted data:%s\n", string(x2))
}
