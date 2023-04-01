package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateToken(userID uint, secretKey string, expirationTimeInMinutes time.Duration) (string, error) {
	// Set claims
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Minute * expirationTimeInMinutes).Unix()

	// Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	signedToken, err := token.SignedString([]byte(secretKey))

	return signedToken, err
}
