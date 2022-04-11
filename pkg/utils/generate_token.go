package utils

import (
	"time"

	"github.com/hako/branca"
)

const (
	tokenLifespan = time.Hour * 24 * 14
)

// TokenClaims is structure containing information about token, expires time and encrypted data
type TokenClaims[T interface{}] struct {
	Token     string    `json:"token,omitempty"`
	ExpiresAt time.Time `json:"expiresAt,omitempty"`
	Data      *T        `json:"data,omitempty"`
}

// GenerateToken generates new token and returns token claims
func GenerateToken[T interface{}](branca *branca.Branca, data string, store T) (*TokenClaims[T], error) {
	generatedToken, err := branca.EncodeToString(data)
	if err != nil {
		return nil, err
	}

	return &TokenClaims[T]{
		Token:     generatedToken,
		ExpiresAt: time.Now().Add(tokenLifespan),
	}, nil
}
