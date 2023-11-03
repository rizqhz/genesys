package service

import (
	"github.com/labstack/echo/v4"
	"github.com/rizghz/genesys/module/Asisten/transfer"
)

type AsistenService interface {
	GetSemuaAsisten(ctx echo.Context) []transfer.Response
	GetAsistenSpesifik(ctx echo.Context, nias string) *transfer.Response
	TambahAsisten(ctx echo.Context, request *transfer.RequestBody) *transfer.Response
	EditAsisten(ctx echo.Context, nias string, request *transfer.RequestBody) *transfer.Response
	HapusAsisten(ctx echo.Context, nias string) bool
}
