package applemusic

import (
	"crypto/ecdsa"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"

	"github.com/lueurxax/shurpa/common/crypto"
)

const kidField = "kid"

type AppleJWTCreator interface {
	CreateToken() (token string, err error)
}

type jwt struct {
	key         *ecdsa.PrivateKey
	kid, issuer string
	ttl         time.Duration
}

func (j *jwt) CreateToken() (string, error) {
	now := time.Now()
	data := &jwtgo.StandardClaims{
		ExpiresAt: now.Add(j.ttl).Unix(),
		IssuedAt:  now.Unix(),
		Issuer:    j.issuer,
	}
	raw := jwtgo.NewWithClaims(jwtgo.SigningMethodES256, data)
	raw.Header[kidField] = j.kid

	signedToken, err := raw.SignedString(j.key)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func NewJwt(keyData []byte, kid, issuer string, ttl time.Duration) (AppleJWTCreator, error) {
	key, err := crypto.ParsePKCS8PrivateKeyFromPEM(keyData)
	if err != nil {
		return nil, err
	}
	return &jwt{key: key, kid: kid, issuer: issuer, ttl: ttl}, nil
}
