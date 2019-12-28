package security

import (
	"fmt"
	"testing"
)

func TestGenerateX590Cert(t *testing.T) {
	//t.Skip()
	GenX509Cert("1", "2")
}

func TestLoadX509Cert(t *testing.T) {
	cert := LoadX509Cert("github.name.cer")
	fmt.Printf("Name %s\n", cert.Subject.CommonName)
	fmt.Printf("Not before %s\n", cert.NotBefore.String())
	fmt.Printf("Not after %s\n", cert.NotAfter.String())
}
