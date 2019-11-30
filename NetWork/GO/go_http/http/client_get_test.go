package requests

import "testing"

func TestClientGet(t *testing.T) {
	response, err := ClientGet("https://www.google.com")
	if err != nil {
		t.Fatal(err.Error())
	}
	if response.Status != "200 OK" {
		t.Fatal(response.Status)
	}
	chSet := GetCharset(response)
	if chSet != "UTF-8" {
		t.Fatal("Charrset Cannot Handle")
	}

}

func TestProxyClientGet(t *testing.T) {
	response, err := ProxyGet("http://127.0.0.1:1087", "https://www.google.com")
	if err != nil {
		t.Fatal(err.Error())
	}
	if response.Status != "200 OK" {
		t.Fatal(response.Status)
	}
	chSet := GetCharset(response)
	if chSet != "UTF-8" {
		t.Fatal("Charrset Cannot Handle")
	}
}
