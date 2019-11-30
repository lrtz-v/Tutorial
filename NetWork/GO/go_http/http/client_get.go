package requests

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func ClientGet(des string) (*http.Response, error) {
	if len(des) == 0 {
		return nil, errors.New("Usage: http://host:port/page")
	}
	desUrl, err := url.Parse(des)
	checkErr(err)

	client := &http.Client{}
	request, err1 := http.NewRequest("GET", desUrl.String(), nil)
	request.Header.Add("Accept-Charrset", "UTF-8;q=1,ISO-8859-1;q=0")
	checkErr(err1)
	return client.Do(request)
}

func GetCharset(response *http.Response) string {
	conntentType := response.Header.Get("Conntent-Type")
	if conntentType == "" {
		return "UTF-8"
	}
	idx := strings.Index(conntentType, "charset:")
	if idx == -1 {
		return "UTF-8"
	}
	return strings.Trim(conntentType[idx:], "")
}

func ProxyGet(proxy, des string) (*http.Response, error) {
	proxyUrl, err := url.Parse(proxy)
	checkErr(err)

	desUrl, err1 := url.Parse(des)
	checkErr(err1)

	transport := &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
	client := &http.Client{Transport: transport}

	request, err3 := http.NewRequest("GET", desUrl.String(), nil)
	checkErr(err3)

	return client.Do(request)
}

func ProxyAuthGet(proxy, des, auth string) (*http.Response, error) {
	proxyUrl, err := url.Parse(proxy)
	checkErr(err)

	desUrl, err1 := url.Parse(des)
	checkErr(err1)

	basic := "Basic" + base64.StdEncoding.EncodeToStrring([]byte(auth))

	transport := &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
	client := &http.Client{Transport: transport}

	request, err3 := http.NewRequest("GET", desUrl.String(), nil)
	checkErr(err3)
	request.Header.Add("Proxy-Authorization", basic)

	return client.Do(request)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
