package auth

import "github.com/golang-jwt/jwt/v5"

var Salt = []byte("da455fc4-a992-4ce3-b5e9-b73d7ceb25ea")

type contextKey string

const claimsKey contextKey = "claims"

type UserClaims struct {
	Uname string `json:"uname"`
	jwt.RegisteredClaims
}
