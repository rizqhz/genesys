package service

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
	"github.com/rizghz/genesys/module/Auth/transfer"
)

type UserService interface {
	GetSemuaUser(ctx echo.Context) []transfer.UserResponse
	GetUserSpesifik(id int) *transfer.UserResponse
	TambahUser(request *transfer.UserRequestBody) *transfer.UserResponse
	EditUser(id int, request *transfer.UserRequestBody) *transfer.UserResponse
	HapusUser(id int) bool
	UploadFoto(file *multipart.FileHeader) *string
}
