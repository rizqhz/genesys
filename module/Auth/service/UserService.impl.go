package service

import (
	"encoding/base64"
	"mime/multipart"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rizghz/genesys/infrastructure/service/cloudinary"
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
	query := ctx.QueryParams()
	var response []transfer.UserResponse
	for _, data := range srv.repo.Get(query) {
		response = append(response, transfer.UserResponse(data))
	}
	return response
}

func (srv *UserServiceImpl) GetUserSpesifik(id int) *transfer.UserResponse {
	response := srv.repo.Find(id)
	if response != nil {
		return (*transfer.UserResponse)(response)
	}
	return nil
}

func (srv *UserServiceImpl) TambahUser(request *transfer.UserRequestBody) *transfer.UserResponse {
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

func (srv *UserServiceImpl) EditUser(id int, request *transfer.UserRequestBody) *transfer.UserResponse {
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

func (srv *UserServiceImpl) HapusUser(id int) bool {
	return srv.repo.Delete(id)
}

func (srv *UserServiceImpl) UploadFoto(file *multipart.FileHeader) *string {
	stream, err := file.Open()
	if err != nil {
		logrus.Error("[user.service]: ", err.Error())
		return nil
	}
	defer stream.Close()
	src := []byte(file.Header["Content-Type"][0] + file.Filename + time.Now().String())
	name := base64.RawURLEncoding.EncodeToString(src)
	url := srv.uploader.Upload(stream, name)
	return &url
}
