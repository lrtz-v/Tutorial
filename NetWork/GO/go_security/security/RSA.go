package security

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/gob"
	"encoding/pem"
	"os"
)

func GenRSAKeys(bitSize int) {
	reader := rand.Reader
	key, err := rsa.GenerateKey(reader, bitSize)
	checkErr(err)

	publicKey := key.PublicKey
	saveGobKey("private.key", key)
	saveGobKey("public.key", publicKey)
	savePEMKey("private.pem", key)
}

func saveGobKey(fileName string, key interface{}) {
	outFile, err := os.Create(fileName)
	checkErr(err)
	defer outFile.Close()

	encoder := gob.NewEncoder(outFile)
	err = encoder.Encode(key)
	checkErr(err)
}

func savePEMKey(fileName string, key *rsa.PrivateKey) {
	outFile, err := os.Create(fileName)
	checkErr(err)
	defer outFile.Close()

	var privateKey = &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key)}

	pem.Encode(outFile, privateKey)
}

func LoadRSAKey(fileName string, key interface{}) {
	inFile, err := os.Open(fileName)
	checkErr(err)
	defer inFile.Close()
	decoder := gob.NewDecoder(inFile)
	err = decoder.Decode(key)
	checkErr(err)
}
