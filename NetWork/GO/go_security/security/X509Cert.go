package security

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"os"
	"time"
)

func GenX509Cert(commonName, organization string) {
	random := rand.Reader

	var key rsa.PrivateKey
	LoadRSAKey("private.key", &key)

	now := time.Now()
	then := now.Add(3600 * 24 * 365 * 1000 * 1000 * 1000) // one year
	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			CommonName:   commonName,
			Organization: []string{organization},
		},
		NotBefore:             now,
		NotAfter:              then,
		SubjectKeyId:          []byte{1, 2, 3, 4},
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		BasicConstraintsValid: true,
		IsCA:                  true,
		DNSNames:              []string{"github.name", "localhost"},
	}

	derBytes, err := x509.CreateCertificate(random, &template, &template, &key.PublicKey, &key)
	checkErr(err)

	certCerFile, err1 := os.Create("github.name.cer")
	checkErr(err1)
	defer certCerFile.Close()
	certCerFile.Write(derBytes)

	certPEMFile, err2 := os.Create("github.name.pem")
	checkErr(err2)
	defer certPEMFile.Close()
	pem.Encode(certPEMFile, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})

	keyPEMFile, err3 := os.Create("private.pem")
	checkErr(err3)
	defer keyPEMFile.Close()
	pem.Encode(keyPEMFile, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(&key)})
}

func LoadX509Cert(fileName string) *x509.Certificate {
	certCerFile, err := os.Open(fileName)
	checkErr(err)
	defer certCerFile.Close()

	derBytes := make([]byte, 1000)
	count, err1 := certCerFile.Read(derBytes)
	checkErr(err1)

	cert, err2 := x509.ParseCertificate(derBytes[0:count])
	checkErr(err2)
	return cert
}
