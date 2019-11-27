package security

import (
	"testing"
)

func TestMd5Hash(t *testing.T) {
	str := "hello\n"
	data := Md5Hash(str)
	if data != "b1946ac92492d2347c6235b4d2611184" {
		t.Fatal("md5 hash error")
	}
}

func TestBlowfist(t *testing.T) {
	src := "hello\n\n\n"
	a := BlowfishEncrypt(src, "my key")

	b := BlowfishDecrypt(a, "my key")
	if b != src {
		t.Fatal("blowfish error")
	}
}
