package entities

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	Username string `json:"username"`
	Rol      string `json:"rol"`
	Name     string `json:"name"`
	Id_jefe  int    `json:"id_jefe"`
	jwt.RegisteredClaims
}
