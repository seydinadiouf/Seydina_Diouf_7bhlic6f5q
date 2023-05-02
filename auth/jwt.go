package auth

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"os"
	"time"
)

type JwtCustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateJWT(username string) (tokenString string, err error) {
	// Load the environment variables from the .env file
	envErr := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", envErr)
	}

	var jwtKey = []byte(os.Getenv("JWT_SECRET"))

	expirationTime := time.Now().Add(72 * time.Hour)

	claims := &JwtCustomClaims{
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func ValidateToken(signedToken string) (err error) {
	// Load the environment variables from the .env file
	envErr := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", envErr)
	}

	var jwtKey = []byte(os.Getenv("JWT_SECRET"))

	token, err := jwt.ParseWithClaims(
		signedToken,
		&JwtCustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*JwtCustomClaims)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}

	if claims.ExpiresAt.Unix() < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}

	return

}
