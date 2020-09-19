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
	desURL, err := url.Parse(des)
	checkErr(err)

	client := &http.Client{}
	request, err := http.NewRequest("GET", desURL.String(), nil)
	request.Header.Add("Accept-Charrset", "UTF-8;q=1,ISO-8859-1;q=0")
	checkErr(err)
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
	proxyURL, err := url.Parse(proxy)
	checkErr(err)

	desURL, err := url.Parse(des)
	checkErr(err)

	transport := &http.Transport{Proxy: http.ProxyURL(proxyURL)}
	client := &http.Client{Transport: transport}

	request, err := http.NewRequest("GET", desURL.String(), nil)
	checkErr(err)

	return client.Do(request)
}

func ProxyAuthGet(proxy, des, auth string) (*http.Response, error) {
	proxyURL, err := url.Parse(proxy)
	checkErr(err)

	desURL, err := url.Parse(des)
	checkErr(err)

	basic := "Basic" + base64.StdEncoding.EncodeToString([]byte(auth))

	transport := &http.Transport{Proxy: http.ProxyURL(proxyURL)}
	client := &http.Client{Transport: transport}

	request, err := http.NewRequest("GET", desURL.String(), nil)
	checkErr(err)
	request.Header.Add("Proxy-Authorization", basic)

	return client.Do(request)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
