package securityUtil

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

var (
	// ErrInvalidBlockSize indicates hash blocksize <= 0.
	ErrInvalidBlockSize = errors.New("invalid blocksize")

	// ErrInvalidPKCS7Data indicates bad input to PKCS7 pad or unpad.
	ErrInvalidPKCS7Data = errors.New("invalid PKCS7 data (empty or not padded)")

	// ErrInvalidPKCS7Padding indicates PKCS7 unpad fails to bad input.
	ErrInvalidPKCS7Padding = errors.New("invalid padding on input")
)

// pkcs7Pad right-pads the given byte slice with 1 to n bytes, where
// n is the block size. The size of the result is x times n, where x
// is at least 1.
func pkcs7Pad(b []byte, blocksize int) ([]byte, error) {
	if b == nil || len(b) == 0 {
		return nil, ErrInvalidPKCS7Data
	}
	if blocksize <= 0 {
		return nil, ErrInvalidBlockSize
	}
	n := blocksize - (len(b) % blocksize)
	pb := make([]byte, len(b)+n)
	copy(pb, b)
	copy(pb[len(b):], bytes.Repeat([]byte{byte(n)}, n))

	return pb, nil
}

// pkcs7Unpad validates and unpads data from the given bytes slice.
// The returned value will be 1 to n bytes smaller depending on the
// amount of padding, where n is the block size.
func pkcs7Unpad(b []byte, blocksize int) ([]byte, error) {
	if b == nil || len(b) == 0 {
		return nil, ErrInvalidPKCS7Data
	}
	if blocksize <= 0 {
		return nil, ErrInvalidBlockSize
	}
	if len(b)%blocksize != 0 {
		return nil, ErrInvalidPKCS7Padding
	}

	c := b[len(b)-1]
	n := int(c)
	if n == 0 || n > len(b) {
		return nil, ErrInvalidPKCS7Padding
	}

	for i := 0; i < n; i++ {
		if b[len(b)-n+i] != c {
			return nil, ErrInvalidPKCS7Padding
		}
	}

	return b[:len(b)-n], nil
}

func AESEncrypt_CBC_Pkcs7(src []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	src, err = pkcs7Pad(src, block.BlockSize())
	if err != nil {
		return nil, err
	}

	blockmode := cipher.NewCBCEncrypter(block, key)
	blockmode.CryptBlocks(src, src)

	return src, nil
}

func AESDecrypt_CBC_Pkcs7(src []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockmode := cipher.NewCBCDecrypter(block, key)
	blockmode.CryptBlocks(src, src)
	src, err = pkcs7Unpad(src, block.BlockSize())
	if err != nil {
		return nil, err
	}

	return src, nil
}
