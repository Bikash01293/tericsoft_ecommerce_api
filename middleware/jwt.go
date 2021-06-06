package middleware

import (
	"log"

	"github.com/dgrijalva/jwt-go"
)

// GenerateJWT is used for generating a token
func GenerateJWT(email string) (string, error) {
	claims := &Claims{
		Email:          email,
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenString, err := token.SignedString(Key)

	if err != nil {
		log.Println("Error in JWT token generation")
		return "", err
	}
	return tokenString, nil
}
func VerifyToken(tokenString string) (email string, err error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return Key, nil
	})
	if token != nil {
		return claims.Email, nil
	}
	return "", err
}
