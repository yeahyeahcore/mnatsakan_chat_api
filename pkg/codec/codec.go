package codec

import "github.com/golang-jwt/jwt"

// Codec ...
type Codec struct {
	SecretKey  string
	SignMethod *jwt.SigningMethodHMAC
}

// NewCodec ...
func NewCodec(secretKey string, signMethod *jwt.SigningMethodHMAC) *Codec {
	return &Codec{
		SecretKey:  secretKey,
		SignMethod: signMethod,
	}
}
