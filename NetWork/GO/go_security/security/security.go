package security

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"

	"golang.org/x/crypto/blowfish"
)

func Md5Hash(str string) string {

	hash := md5.New()
	b := []byte(str)
	hash.Write(b)
	hashVal := hash.Sum(nil)
	return hex.EncodeToString(hashVal)

	//hashSize := hash.Size()
	//var buf []uint32
	//for n := 0; n < hashSize; n += 4 {
	//	var val uint32
	//	val = uint32(hashVal[n])<<24 + uint32(hashVal[n+1])<<16 + uint32(hashVal[n+2])<<8 + uint32(hashVal[n+3])
	//	buf = append(buf, val)
	//}
	//return buf
}

func BlowfishEncrypt(str, key string) string {
	encryptKey := []byte(key)
	cipher, err := blowfish.NewCipher(encryptKey)
	checkErr(err)
	src := []byte(str)

	enc := make([]byte, 512)
	cipher.Encrypt(enc[0:], src)

	return string(enc)
}

func BlowfishDecrypt(src, key string) string {
	encryptKey := []byte(key)
	cipher, err := blowfish.NewCipher(encryptKey)
	checkErr(err)
	decrypt := make([]byte, 8)
	cipher.Decrypt(decrypt[0:], []byte(src)[0:])
	result := bytes.NewBuffer(nil)
	result.Write(decrypt[0:8])
	return string(result.Bytes())
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
