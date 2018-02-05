package securityUtil

import (
	"testing"
)

func TestSha256(t *testing.T) {
	source := "oauthConsumerKey=1Nocz0wk0Hi8oGgSosogC4K4k&oauthToken=TOKEN_%2B8vQAR1eoD3ujiGstjzdyakEgbkyvWhfzF1fChQJ46EH07n%2FQvrazkMqy%2BhuprqU&oauthSignatureMethod=HMAC-SHA1&oauthTimestamp=1508486834&oauthNonce=5409983431934290948&oauthVersion=1.0&"
	sign := "813c8202a31c73371ae0bbe13cb49d65c94da3de2877345603271ca14e5e4bcd"

	result := Sha256String(source, false)
	if result != sign {
		t.Fatalf("Sha256编码结果不一致")
	}
}
