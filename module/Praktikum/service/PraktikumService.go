package service

import (
	"github.com/labstack/echo/v4"
	"github.com/rizghz/genesys/module/Praktikum/transfer"
)

type PraktikumService interface {
	GetSemuaPraktikum(ctx echo.Context) []transfer.Response
	GetPraktikumSpesifik(ctx echo.Context, id string) *transfer.Response
	TambahPraktikum(ctx echo.Context, request *transfer.RequestBody) *transfer.Response
	EditPraktikum(ctx echo.Context, id string, request *transfer.RequestBody) *transfer.Response
	HapusPraktikum(ctx echo.Context, id string) bool
}
