package securityUtil

import (
	"testing"
)

func TestSha1String(t *testing.T) {
	s := "hello world"
	result := Sha1String(s, true)
	if result != "2AAE6C35C94FCFB415DBE95F408B9CE91EE846ED" {
		t.Errorf("Sha1String(\"hello world\") failed.Got %s, expected %s", result, "2AAE6C35C94FCFB415DBE95F408B9CE91EE846ED")
	}

	result = Sha1String(s, false)
	if result != "2aae6c35c94fcfb415dbe95f408b9ce91ee846ed" {
		t.Errorf("Sha1String(\"hello world\") failed.Got %s, expected %s", result, "2aae6c35c94fcfb415dbe95f408b9ce91ee846ed")
	}
}

func TestSha1Bytes(t *testing.T) {
	s := "hello world"
	b := []byte(s)
	result := Sha1Bytes(b, true)
	if result != "2AAE6C35C94FCFB415DBE95F408B9CE91EE846ED" {
		t.Errorf("Sha1Bytes(\"hello world\") failed.Got %s, expected %s", result, "2AAE6C35C94FCFB415DBE95F408B9CE91EE846ED")
	}

	result = Sha1Bytes(b, false)
	if result != "2aae6c35c94fcfb415dbe95f408b9ce91ee846ed" {
		t.Errorf("Sha1Bytes(\"hello world\") failed.Got %s, expected %s", result, "2aae6c35c94fcfb415dbe95f408b9ce91ee846ed")
	}
}
