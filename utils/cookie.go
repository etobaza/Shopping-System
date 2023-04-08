package utils

import (
	"net/http"
	"time"
)

func CreateCookie(token string) *http.Cookie {
	// Set the expiration time for the cookie
	expiration := time.Now().Add(120 * time.Minute)

	// Create the cookie
	cookie := &http.Cookie{
		Name:     "jwtToken",
		Value:    token,
		Expires:  expiration,
		Path:     "/",
		HttpOnly: true, // Cookie HTTP-only
	}

	return cookie
}
