package requests

import (
	"net/http"
	"strings"
)

func AcceptableCharset(contentTypes []string, charset string) bool {
	for _, cType := range contentTypes {
		if strings.Index(cType, charset) != -1 {
			return true
		}
	}
	return false
}

func Get(url string) (*http.Response, error) {
	return http.Get(url)
}
