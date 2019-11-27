package security

import (
	"crypto/rsa"
	"fmt"
	"testing"
)

func TestGenerateRSA(t *testing.T) {
	t.Skip()
	GenRSAKeys(512)
}

func TestLoadRSAKey(t *testing.T) {
	t.Skip()
	var key rsa.PrivateKey
	LoadRSAKey("private.key", &key)
	fmt.Println("Private key primes", key.Primes[0].String(), key.Primes[1].String())
	fmt.Println("Private key exponent", key.D.String())

	var publicKey rsa.PublicKey
	LoadRSAKey("public.key", &publicKey)
	fmt.Println("Public key modulus", publicKey.N.String())
	fmt.Println("Public key exponent", publicKey.E)

}
