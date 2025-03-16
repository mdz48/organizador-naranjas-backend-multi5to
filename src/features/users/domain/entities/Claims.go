package entities

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	Username string `json:"name"`
	Rol      string `json:"rol"`
	jwt.RegisteredClaims
}
