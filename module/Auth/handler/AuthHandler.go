package handler

import "github.com/labstack/echo/v4"

type AuthHandler interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	Refresh() echo.HandlerFunc
	Logout() echo.HandlerFunc
}
