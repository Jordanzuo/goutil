package securityUtil

import (
	"testing"
)

func TestVerifyRsaWithSha1(t *testing.T) {
	publicKey := "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCmreYIkPwVovKR8rLHWlFVw7YDfm9uQOJKL89Smt6ypXGVdrAKKl0wNYc3/jecAoPi2ylChfa2iRu5gunJyNmpWZzlCNRIau55fxGW0XEu553IiprOZcaw5OuYGlf60ga8QT6qToP0/dpiL/ZbmNUO9kUhosIjEu22uFgR+5cYyQIDAQAB"
	source := "notifyId=GC201710201145455410100040000&partnerOrder=601_2055_1508471145_42087&productName=元宝&productDesc=1000元宝&price=5000&count=1&attach=601_2055_1508471145_42087"
	sign := "JdzsJlVOgJ6gXTCJjWAXisyFeS0ztvB5m6WOgx9XqqdfxthLVC0gvxXdoqT1SnzzkaznebtbgvVrIeAFlyEBiVpShH76yZ9KO781wiBdJMY/BUwKkHlnMWjtFZx7pjqj6xBMoZ3HFl9j5loFYuYLMg+MDUCpvXV+Kg/wAqkkOnY="

	err := VerifyRsaWithSha1(publicKey, source, sign)
	if err != nil {
		t.Fatalf("VerifyRsaWithSha1验证错误:%v", err)
	}
}

func TestVerifyRsaWithSha256(t *testing.T) {
	publicKey := "MFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBAJm8eeTR8mPWuPdCFo5boenHe+Yj8zC82ohIuTeMu+4QJuRK/MI+wtJlYheKtE0s4lXzL0rw/KQzMB+KO9F/WM0CAwEAAQ=="
	source := "accessMode=0&amount=6.00&bankId=HuaWei&notifyTime=1500370048019&orderId=H20170718172707521578B0C13&orderTime=2017-07-18 17:27:07&payType=0&productName=元宝&requestId=606_2001_1500370020_19213&result=0&spending=0.00&tradeTime=2017-07-18 17:27:07&userName=900086000021763400"
	sign := "d9PDjSwgItg8eTTAYbP2OHD5t+F8wgrgjMOCXHXI7pe3qCe7sixraHmLrOQfFoAvxi4e2eYxZocN4QRoqD3/zw=="

	err := VerifyRsaWithSha256(publicKey, source, sign)
	if err != nil {
		t.Fatalf("VerifyRsaWithSha256验证错误:%v", err)
	}
}
