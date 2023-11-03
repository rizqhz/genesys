package service

import (
	"github.com/labstack/echo/v4"
	"github.com/rizghz/genesys/module/MataPraktikum/transfer"
)

type MatkumService interface {
	GetSemuaMatkum(ctx echo.Context) []transfer.Response
	GetMatkumSpesifik(ctx echo.Context, kode string) *transfer.Response
	TambahMatkum(ctx echo.Context, request *transfer.RequestBody) *transfer.Response
	EditMatkum(ctx echo.Context, kode string, request *transfer.RequestBody) *transfer.Response
	HapusMatkum(ctx echo.Context, kode string) bool
}
