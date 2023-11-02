package jwt

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func Impl() echo.MiddlewareFunc {
	key := NewJwtKey()
	return echojwt.WithConfig(echojwt.Config{
		SigningMethod: "HS256",
		SigningKey:    []byte(key.AccessKey),
	})
}
