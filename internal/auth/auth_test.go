package auth

import (
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	expected := "some-api-key"
	req := http.Request{Header: http.Header{}}
	req.Header.Add("Authorization", "ApiKey "+expected)

	key, err := GetAPIKey(req.Header)
	if err != nil {
		t.Errorf("did not expect an error, got %v", err)
	}

	if key != expected {
		t.Errorf("expected %s to be %s", key, expected)
	}
}
