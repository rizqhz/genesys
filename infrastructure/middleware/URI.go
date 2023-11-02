package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Uri() echo.MiddlewareFunc {
	return middleware.RemoveTrailingSlash()
}
