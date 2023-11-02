package handler

import "github.com/labstack/echo/v4"

type PenjadwalanHandler interface {
	Index() echo.HandlerFunc
	Observe() echo.HandlerFunc
	Store() echo.HandlerFunc
	Edit() echo.HandlerFunc
	Destroy() echo.HandlerFunc
}
