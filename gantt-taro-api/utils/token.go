package utils

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
	"log"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	publicKeyRaw = os.Getenv("PUBLIC_KEY_ROW")
	publicKey    interface{}
)

func parsePublicKey() interface{} {
	publicKeyBlock, _ := pem.Decode([]byte(publicKeyRaw))

	publicKey, err := x509.ParseCertificate(publicKeyBlock.Bytes)
	if err != nil {
		log.Printf("[ERROR] %+v", err)
		return nil
	}

	return publicKey.PublicKey
}

func ValidateToken(t string) bool {
	if publicKey == nil {
		publicKey = parsePublicKey()
	}

	token, err := jwt.Parse(t, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("Method failed")
		}

		return publicKey, nil
	})

	if err != nil {
		log.Printf("[ERROR] %+v", err)
		return false
	}

	if !token.Valid {
		log.Println("[ERROR] token is invalid")
		return false
	}

	log.Printf("[INFO] %+v", token.Claims)

	return true
}
