package service

import (
	"encoding/base64"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/labstack/echo/v4"
	jwt "github.com/rizghz/genesys/infrastructure/middleware/JWT"
	"github.com/rizghz/genesys/infrastructure/service/cloudinary"
	"github.com/rizghz/genesys/internal/helpers"
	"github.com/rizghz/genesys/module/Auth/repository"
	"github.com/rizghz/genesys/module/Auth/transfer"
	"github.com/sirupsen/logrus"
)

type UserServiceImpl struct {
	repo     repository.UserRepository
	uploader cloudinary.Uploader
}

func NewUserServiceImpl(r repository.UserRepository, u cloudinary.Uploader) UserService {
	return &UserServiceImpl{
		repo:     r,
		uploader: u,
	}
}

func (srv *UserServiceImpl) GetSemuaUser(ctx echo.Context) []transfer.UserResponse {
	token := helpers.GetJwtToken(ctx)
	key := jwt.NewJwtKey()
	if helpers.JwtValidate(token, key.AccessKey) {
		claim := helpers.JwtDecode[jwt.JwtPayload](&token[1])
		if claim.Role != "admin" {
			ctx.Set("authorization.error", true)
			return nil
		}
	}
	query := ctx.QueryParams()
	var responses []transfer.UserResponse
	for _, data := range srv.repo.Get(query) {
		response := transfer.UserResponse(data)
		responses = append(responses, response)
	}
	return responses
}

func (srv *UserServiceImpl) GetUserSpesifik(ctx echo.Context, id int) *transfer.UserResponse {
	token := helpers.GetJwtToken(ctx)
	key := jwt.NewJwtKey()
	if helpers.JwtValidate(token, key.AccessKey) {
		claim := helpers.JwtDecode[jwt.JwtPayload](&token[1])
		if claim.Role != "admin" {
			ctx.Set("authorization.error", true)
			return nil
		}
	}
	result := srv.repo.Find(id)
	if result != nil {
		response := transfer.UserResponse(*result)
		return &response
	}
	return nil
}

func (srv *UserServiceImpl) TambahUser(ctx echo.Context, request *transfer.UserRequestBody) *transfer.UserResponse {
	token := helpers.GetJwtToken(ctx)
	key := jwt.NewJwtKey()
	if helpers.JwtValidate(token, key.AccessKey) {
		claim := helpers.JwtDecode[jwt.JwtPayload](&token[1])
		if claim.Role != "admin" {
			ctx.Set("authorization.error", true)
			return nil
		}
	}
	data := &repository.UserModel{
		Nama:    request.Nama,
		Alamat:  request.Alamat,
		Email:   request.Email,
		Telepon: request.Telepon,
		Foto:    request.Foto,
	}
	result := srv.repo.Create(data)
	if result != nil {
		return &transfer.UserResponse{
			ID:      result.ID,
			Nama:    result.Nama,
			Alamat:  result.Alamat,
			Email:   result.Email,
			Telepon: result.Telepon,
			Foto:    result.Foto,
		}
	}
	return nil
}

func (srv *UserServiceImpl) EditUser(ctx echo.Context, id int, request *transfer.UserRequestBody) *transfer.UserResponse {
	token := helpers.GetJwtToken(ctx)
	key := jwt.NewJwtKey()
	if helpers.JwtValidate(token, key.AccessKey) {
		claim := helpers.JwtDecode[jwt.JwtPayload](&token[1])
		if claim.Role != "admin" {
			ctx.Set("authorization.error", true)
			return nil
		}
	}
	data := &repository.UserModel{
		ID:      id,
		Nama:    request.Nama,
		Alamat:  request.Alamat,
		Email:   request.Email,
		Telepon: request.Telepon,
		Foto:    request.Foto,
	}
	result := srv.repo.Update(data)
	if result != nil {
		return &transfer.UserResponse{
			ID:      result.ID,
			Nama:    result.Nama,
			Alamat:  result.Alamat,
			Email:   result.Email,
			Telepon: result.Telepon,
			Foto:    result.Foto,
		}
	}
	return nil
}

func (srv *UserServiceImpl) HapusUser(ctx echo.Context, id int) bool {
	token := helpers.GetJwtToken(ctx)
	key := jwt.NewJwtKey()
	if helpers.JwtValidate(token, key.AccessKey) {
		claim := helpers.JwtDecode[jwt.JwtPayload](&token[1])
		if claim.Role != "admin" {
			ctx.Set("authorization.error", true)
			return false
		}
	}
	return srv.repo.Delete(id)
}

func (srv *UserServiceImpl) UploadFoto(file *multipart.FileHeader) *string {
	stream, err := file.Open()
	if err != nil {
		logrus.Error("[user.service]: ", err.Error())
		return nil
	}
	defer stream.Close()
	src := []byte(fmt.Sprint(time.Now().Unix(), file.Filename))
	name := base64.RawURLEncoding.EncodeToString(src)
	url := srv.uploader.Upload(stream, name)
	return &url
}

func (srv *UserServiceImpl) GantiFoto(id int, file *multipart.FileHeader) *string {
	data := srv.repo.Find(id)
	if data.Foto != "" {
		srv.uploader.Delete(data.Foto)
	}
	stream, err := file.Open()
	if err != nil {
		logrus.Error("[user.service]: ", err.Error())
		return nil
	}
	defer stream.Close()
	src := []byte(fmt.Sprint(time.Now().Unix(), file.Filename))
	name := base64.RawURLEncoding.EncodeToString(src)
	url := srv.uploader.Upload(stream, name)
	return &url
}
