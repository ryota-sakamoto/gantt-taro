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

type Claims struct {
	UniqueId string
}

func ValidateToken(t string) (*Claims, error) {
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
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("token is invalid")
	}

	m, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("claims error")
	}

	c := Claims{
		UniqueId: m["sub"].(string),
	}

	return &c, nil
}
