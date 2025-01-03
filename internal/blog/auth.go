package blog

import "github.com/golang-jwt/jwt/v5"

var jwtSecret = []byte("Admin.Test")

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string
	jwt.RegisteredClaims
}
