package model

import "github.com/dgrijalva/jwt-go"

type JwtCustomClaims struct {
	UserId      string `json:"userId"`
	Role        string `json:"role"`
	PhoneNumber string `json:"phoneNunber"`
	jwt.StandardClaims
}
