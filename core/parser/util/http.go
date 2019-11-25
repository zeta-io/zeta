package util

import (
	"net/http"
	"strings"
)

func HasRequestBody(method string) bool {
	return strings.ToUpper(method) != http.MethodPost
}