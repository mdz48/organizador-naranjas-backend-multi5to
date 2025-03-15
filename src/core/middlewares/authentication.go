package middlewares

import (
	"organizador-naranjas-backend-multi5to/src/features/users/domain/entities"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = []byte(os.Getenv("SECRET_KEY"))

func GenerateToken(user *entities.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": user.Name,
		"rol": user.Rol,
		"exp": time.Now().Add(time.Hour * 24).Unix(), 
	})

	tokenString, err := token.SignedString(secretKey);

	if err != nil {
		return "", err; 
	}

	return tokenString, err; 
}

func VerifyPassword(password, hashPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}

func HashPassword(password string) (string,error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14); 

	if err != nil {
		return "", err; 
	}
	
	return string(bytes), err; 
}
