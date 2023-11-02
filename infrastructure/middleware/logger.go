package middleware

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Logger() echo.MiddlewareFunc {
	format := fmt.Sprintf("%v,\t%v,\t%v,\t%v,\t%v\n", []any{
		"time=${time_rfc3339}",
		"latency=${latency_human}",
		"method=${method}",
		"uri=${uri}",
		"status=${status}",
	}...)
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: format,
	})
}
