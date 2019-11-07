package basicutils

import (
	"testing"
)

func TestParseIP(t *testing.T) {
	addr := ParseIP("127.0.0.1")
	if addr == nil {
		t.Fatal("Invalid address")
	}
	if addr.String() != "127.0.0.1" {
		t.Fatalf("Expected %s, Got %s", "127.0.0.1", addr.String())
	}
}

func TestParseIP2(t *testing.T) {
	addr := ParseIP("0:0:0:0:0:0:0:1")
	if addr == nil {
		t.Fatal("Invalid address")
	}
	if addr.String() != "::1" {
		t.Fatalf("Expected %s, Got %s", "::1", addr.String())
	}
}

func TestDefaultMask(t *testing.T) {
	mask := DefaultMask("127.0.0.1")
	if mask.String() != "ff000000" {
		t.Fatalf("Expected %s, Got %s", "ff000000", mask.String())
	}
}

func TestMask(t *testing.T) {
	mask := Mask("127.0.0.1")
	if mask.String() != "127.0.0.0" {
		t.Fatalf("Expected %s, Got %s", "127.0.0.0", mask.String())
	}
}

func TestLookupPort(t *testing.T) {
	port, err := LookupPort("tcp", "https")
	if err != nil {
		t.Fatal("Error: ", err.Error())
	}
	if port != 443 {
		t.Fatalf("Expected Server port %d, Got %d", 443, port)
	}
}

func TestResolveIPAddr(t *testing.T) {
	_, err := ResolveIPAddr("ip", "www.baidu.com")
	if err != nil {
		t.Fatal("Resolution error", err.Error())
	}
}

func TestLookupHost(t *testing.T) {
	_, err := LookupHost("www.baidu.com")
	if err != nil {
		t.Fatal("Resolution error", err.Error())
	}
}

func TestGetHeadInfoGetHeadInfo(t *testing.T) {
	_, err := GetHeadInfo("www.baidu.com:80")
	if err != nil {
		t.Fatal("Fatal error: ", err.Error())
	}
}
