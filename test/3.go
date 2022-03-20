package main

import (
	"fmt"
	"github.com/golang-jwt/jwt"
)

// JWT 颁发token
// For HMAC signing method, the key can be any []byte. It is recommended to generate
// a key using crypto/rand or something equivalent. You need the same key for signing
// and validating.
var hmacSampleSecret1 []byte = []byte("djf")

func main() {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"djf": "dkff",
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hmacSampleSecret1)

	fmt.Println(tokenString, err)
}
