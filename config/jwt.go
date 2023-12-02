package config

import (
	"github.com/golang-jwt/jwt/v5"
)

var JWT_KEY = []byte("kawd9kke129i198j189")

type JWTClaim struct {
	Username string
	jwt.RegisteredClaims
}
