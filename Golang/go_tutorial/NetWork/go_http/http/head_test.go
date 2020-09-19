package requests

import "testing"

func TestHead(t *testing.T) {
	response, err := Head("https://www.google.com")
	if err != nil {
		t.Fatal(err.Error())
	}
	if response.Status != "200 OK" {
		t.Fatal("Status err")
	}
}
