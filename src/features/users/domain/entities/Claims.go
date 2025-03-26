package entities

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	ID       int32    `json:"id"`
	Username string `json:"username"`
	Rol      string `json:"rol"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Id_jefe  int    `json:"id_jefe"`
	jwt.RegisteredClaims
}
