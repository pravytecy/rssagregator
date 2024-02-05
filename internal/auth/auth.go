package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Authorization: ApiKey {apiKey}
func GetApiKey(headers http.Header) (string,error){
	vals := headers.Get("Authorization")
	if vals == "" {
		return "",errors.New("no authentication key found")
	}
	apiKey := strings.Split(vals, " ")
	if len(apiKey) < 2 || apiKey[0] != "ApiKey" {
		return "",errors.New("malformed authentication key")
	}
	return apiKey[1],nil
}