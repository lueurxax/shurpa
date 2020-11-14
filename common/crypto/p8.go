package crypto

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
)

type ErrKeyMustBePEMEncoded struct{}

func (e *ErrKeyMustBePEMEncoded) Error() string {
	return "key must be pem encoded"
}

type ErrNotECPrivateKey struct{}

func (e *ErrNotECPrivateKey) Error() string {
	return "it's not elliptic curve a private key"
}

// ParsePKCS8PrivateKeyFromPEM parse PEM encoded Elliptic Curve Private Key Structure
func ParsePKCS8PrivateKeyFromPEM(key []byte) (*ecdsa.PrivateKey, error) {
	var err error

	// Parse PEM block
	var block *pem.Block
	if block, _ = pem.Decode(key); block == nil {
		return nil, &ErrKeyMustBePEMEncoded{}
	}

	// Parse the key
	var parsedKey interface{}
	if parsedKey, err = x509.ParsePKCS8PrivateKey(block.Bytes); err != nil {
		return nil, err
	}

	var pkey *ecdsa.PrivateKey
	var ok bool
	if pkey, ok = parsedKey.(*ecdsa.PrivateKey); !ok {
		return nil, &ErrNotECPrivateKey{}
	}

	return pkey, nil
}
