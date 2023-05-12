package middlewares

import (
	"github.com/labstack/echo/v4"
	"school-manager/auth"
	"strings"
)

func Auth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenString := c.Request().Header.Get("Authorization")
			if tokenString == "" {
				return c.JSON(401, map[string]string{"error": "request does not contain an access token"})
			}
			tokenString = strings.Split(tokenString, " ")[1]
			err := auth.ValidateToken(tokenString)
			if err != nil {
				return c.JSON(401, map[string]string{"error": err.Error()})
			}
			return next(c)
		}
	}
}
