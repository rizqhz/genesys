package service

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
	"github.com/rizghz/genesys/module/Auth/transfer"
)

type UserService interface {
	GetSemuaUser(ctx echo.Context) []transfer.UserResponse
	GetUserSpesifik(ctx echo.Context, id int) *transfer.UserResponse
	TambahUser(ctx echo.Context, request *transfer.UserRequestBody) *transfer.UserResponse
	EditUser(ctx echo.Context, id int, request *transfer.UserRequestBody) *transfer.UserResponse
	HapusUser(ctx echo.Context, id int) bool
	UploadFoto(file *multipart.FileHeader) *string
	GantiFoto(id int, file *multipart.FileHeader) *string
}
