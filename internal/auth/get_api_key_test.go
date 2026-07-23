package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	t.Run("valid ApiKey header", func(t *testing.T) {
		headers := http.Header{}
		headers.Set("Authorization", "ApiKey my-secret-key")

		got, err := GetAPIKey(headers)
		// TODO: assert err is nil and got == "my-secret-key"
		_ = got
		_ = err
	})

	t.Run("missing Authorization header", func(t *testing.T) {
		headers := http.Header{}

		_, err := GetAPIKey(headers)
		// TODO: assert err == ErrNoAuthHeaderIncluded
		_ = err
	})

	// Challenge: add a case for a malformed header (e.g. "Bearer token" or just "ApiKey")
}
