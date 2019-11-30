package requests

import (
	"testing"
)

func TestGet(t *testing.T) {
	url := "https://www.google.com"
	response, err := Get(url)
	if err != nil {
		t.Fatal(err.Error())
	}
	if response.Status != "200 OK" {
		t.Fatal("Status error")
	}

	//	b, _ := httputil.DumpResponse(response, false)
	//	fmt.Println(string(b))

	conntentTypes := response.Header["Content-Type"]
	if !AcceptableCharset(conntentTypes, "ISO-8859-1") {
		t.Fatal("Cannot Handle")
	}

	//var buf [512]byte
	//reader := response.Body
	//for {
	//	n, err1 := reader.Read(buf[0:])
	//	if err1 != nil {
	//		return
	//	}
	//fmt.Println(string(buf[0:n]))
	//}
}
