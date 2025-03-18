package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetApiKey extracts an API key from the Authorization header.
func GetApiKey(headers http.Header) (string, error) {
	val := strings.TrimSpace(headers.Get("Authorization"))

	if val == "" {
		return "", errors.New("missing Authorization header")
	}

	vals := strings.Fields(val)
	if len(vals) != 2 {
		return "", errors.New("malformed Authorization header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("invalid auth scheme, expected 'ApiKey'")
	}

	return vals[1], nil
}
