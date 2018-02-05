package securityUtil

import (
	"encoding/base64"
	"testing"
)

func TestHmacSha1(t *testing.T) {
	source := "oauthConsumerKey=1Nocz0wk0Hi8oGgSosogC4K4k&oauthToken=TOKEN_%2B8vQAR1eoD3ujiGstjzdyakEgbkyvWhfzF1fChQJ46EH07n%2FQvrazkMqy%2BhuprqU&oauthSignatureMethod=HMAC-SHA1&oauthTimestamp=1508486834&oauthNonce=5409983431934290948&oauthVersion=1.0&"
	key := "f724054cbBF8710D2c3a2500Ec65Fa9F&"
	sign := "+2zYxghf5BOqAXp/o9yax4TI56c="

	buf, err := HmacSha1(source, key)
	if err != nil {
		t.Fatalf("Hmac-SHA1编码错误:%v", err)
	}

	base64Result := base64.StdEncoding.EncodeToString(buf)
	if base64Result != sign {
		t.Fatalf("Hmac-SHA1编码结果不一致")
	}
}

func TestHmacSha256(t *testing.T) {
	source := "oauthConsumerKey=1Nocz0wk0Hi8oGgSosogC4K4k&oauthToken=TOKEN_%2B8vQAR1eoD3ujiGstjzdyakEgbkyvWhfzF1fChQJ46EH07n%2FQvrazkMqy%2BhuprqU&oauthSignatureMethod=HMAC-SHA1&oauthTimestamp=1508486834&oauthNonce=5409983431934290948&oauthVersion=1.0&"
	key := "f724054cbBF8710D2c3a2500Ec65Fa9F&"
	sign := "6GsbOxY0ldpTqGIdIATUuJceMgCGSwcylODhXxPHZLE="

	buf, err := HmacSha256(source, key)
	if err != nil {
		t.Fatalf("Hmac-SHA256编码错误:%v", err)
	}

	base64Result := base64.StdEncoding.EncodeToString(buf)
	if base64Result != sign {
		t.Fatalf("Hmac-SHA256编码结果不一致")
	}
}
