package middleware

import "github.com/dgrijalva/jwt-go"

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	Key = []byte("super-secret-key")
)

// Claims struct
type Claims struct {
	Email string
	jwt.StandardClaims
}

type ErrorResponse struct {
	// Code    int
	Message string
}
