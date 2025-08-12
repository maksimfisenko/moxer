package middleware

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/maksimfisenko/moxer/internal/services/jwt"
)

func JwtMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")

			if !strings.HasPrefix(authHeader, "Bearer ") {
				c.Set("userId", "no_token")
				return next(c)
			}

			token := strings.TrimPrefix(authHeader, "Bearer ")

			userId, err := jwt.ParseToken(token)
			if err != nil {
				c.Set("userId", "invalid_token")
				return next(c)
			}

			c.Set("userId", userId)

			return next(c)
		}
	}
}
