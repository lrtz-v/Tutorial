package requests

import "net/http"

func Head(url string) (*http.Response, error) {
	response, err := http.Head(url)
	return response, err
}
