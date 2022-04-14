package dto

import "github.com/golang-jwt/jwt"

// TokenClaims is structure containing information about token and expires time
type TokenClaims struct {
	*jwt.StandardClaims
	UserID string `json:"userId"`
}
