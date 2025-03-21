package entities

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	Username string `json:"username"`
	Rol      string `json:"rol"`
	Name	 string `json:"name"`
	jwt.RegisteredClaims
}
