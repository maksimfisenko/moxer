package middleware

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/maksimfisenko/moxer/internal/errorsx"
	"github.com/maksimfisenko/moxer/internal/services/jwt"
)

const (
	authHeaderVal   = "Authorization"
	bearerPrefixVal = "Bearer "
)

func JwtRequired() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get(authHeaderVal)
			if authHeader == "" || !strings.HasPrefix(authHeader, bearerPrefixVal) {
				return errorsx.ErrUnauthorizedHTTP
			}

			token := strings.TrimSpace(strings.TrimPrefix(authHeader, bearerPrefixVal))

			userId, err := jwt.ParseToken(token)
			if err != nil {
				return errorsx.ErrUnauthorizedHTTP
			}

			c.Set("userId", userId)
			return next(c)
		}
	}
}
