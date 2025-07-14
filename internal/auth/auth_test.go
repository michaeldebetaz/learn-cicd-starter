package auth

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	header := http.Header{}

	if _, err := GetAPIKey(header); err != ErrNoAuthHeaderIncluded {
		t.Errorf("Expected error %v, got %v", ErrNoAuthHeaderIncluded, err)
	}

	header.Set("Authorization", "some-random-header")
	if _, err := GetAPIKey(header); err != ErrMalformedAuthorizationHeader {
		t.Errorf("Expected error %v, got %v", ErrMalformedAuthorizationHeader, err)
	}

	APIKey := "super-secret-api-key"
	header.Set("Authorization", fmt.Sprintf("Bearer: %s", APIKey))
	if _, err := GetAPIKey(header); err != ErrMalformedAuthorizationHeader {
		t.Errorf("Expected error %v, got %v", ErrMalformedAuthorizationHeader, err)
	}

	header.Set("Authorization", fmt.Sprintf("ApiKey %s", APIKey))
	if key, err := GetAPIKey(header); err != nil {
		t.Errorf("Expected no error, got %v", err)
	} else if key != APIKey {
		t.Errorf("Expected key %s, got %s", APIKey, key)
	}

	t.Error("Forced failure for the sake of the lesson")
}
