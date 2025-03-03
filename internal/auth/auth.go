package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetApiKey extracts an API Key from the headers of an HTTP request.
// Example:
// Authorization: ApiKey {insert API key here}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("Not Authorized")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("Malformed Authorization header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("Malformed first part Authorization header")
	}
	return vals[1], nil
}
