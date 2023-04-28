package controller

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"school-manager/config"
	"school-manager/model"
	"school-manager/model/payload"
	"time"
)

type jwtCustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func SignIn(c echo.Context) error {
	signIn := new(payload.SignInRequest)
	db := config.DB()

	var user *model.User

	// Binding data
	if err := c.Bind(signIn); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	if res := db.Where("username = ?", signIn.Username).First(&user); res.Error != nil {
		data := map[string]interface{}{
			"message": res.Error.Error(),
		}

		return c.JSON(http.StatusOK, data)
	}

	// Set custom claims
	claims := &jwtCustomClaims{
		user.Username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("my_jwt_secret_key"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token":    t,
		"username": user.Username,
	})
}