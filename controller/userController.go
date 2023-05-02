package controller

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"school-manager/auth"
	"school-manager/config"
	"school-manager/model"
	"school-manager/model/payload"
)

func SignIn(c echo.Context) error {
	signIn := new(payload.SignIn)
	db := config.DB()

	var user *model.User

	// Binding data
	if err := c.Bind(signIn); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}
	message := "Username or password incorrect"

	if res := db.Where("username = ?", signIn.Username).First(&user); res.Error != nil {
		data := map[string]interface{}{
			"message": message,
		}

		return c.JSON(http.StatusUnauthorized, data)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(signIn.Password)); err != nil {
		// If the two passwords don't match, return a 401 status
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusUnauthorized, data)
	}

	// Generate encoded token and send it as response.
	t, err := auth.GenerateJWT(signIn.Username)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token":    t,
		"username": user.Username,
	})
}
