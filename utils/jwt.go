package utils

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"shopping-system/models"
	"time"
)

func GenerateToken(user *models.User) (string, error) {
	// Set claims
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user.ID
	claims["username"] = user.Username
	claims["usertype"] = user.UserType
	claims["exp"] = time.Now().Add(time.Minute * 120).Unix()

	// Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	signedToken, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	return signedToken, err
}

func ExtractToken(r *http.Request) string {
	// Get the cookie
	cookie, err := r.Cookie("jwtToken")
	if err != nil {
		fmt.Println("Error extracting token:", err)
		return ""
	}

	fmt.Println("Token extracted successfully:", cookie.Value)
	return cookie.Value
}

func IsTokenValid(claims jwt.MapClaims) bool {
	expirationTime := int64(claims["exp"].(float64))
	return time.Now().UTC().Unix() < expirationTime
}

func ParseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
