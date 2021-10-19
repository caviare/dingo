package utils

import (
	"encoding/base64"
	"log"
	"os"

	"golang.org/x/crypto/scrypt"
)

// 传入文本信息加密文本
// 默认salt：dingo
func Encrypt(text string) string {
	salt := os.Getenv("SALT")
	if salt == "" {
		salt = "dingo"
	}
	dk, err := scrypt.Key([]byte(text), []byte(salt), 32768, 8, 1, 32)
	if err != nil {
		log.Fatal(err)
	}
	return base64.StdEncoding.EncodeToString(dk)
}
