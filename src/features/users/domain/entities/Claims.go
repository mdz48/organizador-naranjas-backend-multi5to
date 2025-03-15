package entities

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	UserName string `json:"name"`
	Rol      string `json:"rol"`
	jwt.RegisteredClaims
}