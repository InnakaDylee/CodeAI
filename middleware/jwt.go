package middleware

import (
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"

	"github.com/golang-jwt/jwt/v5"
)

type jwtCustomClaims struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	jwt.RegisteredClaims
}

// JWT authentication using HS256 algorithm
func CreateToken(userId int, name string) string {

	// set claims (payload)
	var payload jwtCustomClaims

	payload.ID = uint(userId)
	payload.Name = name
	payload.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Hour * 24))

	// create token with claims (header)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
		os.Exit(1)
	}
	// generate encoded token using the SECRET_KEY environment variable (signature)
	validToken, _ := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	return validToken
}