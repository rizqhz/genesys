package service

import (
	"github.com/labstack/echo/v4"
	jwt "github.com/rizghz/genesys/infrastructure/middleware/JWT"
	"github.com/rizghz/genesys/internal/helpers"
	"github.com/rizghz/genesys/module/Mahasiswa/repository"
	"github.com/rizghz/genesys/module/Mahasiswa/transfer"
)

type MahasiswaServiceImpl struct {
	repo repository.MahasiswaRepository
}

func NewMahasiswaServiceImpl(r repository.MahasiswaRepository) MahasiswaService {
	return &MahasiswaServiceImpl{
		repo: r,
	}
}

func (srv *MahasiswaServiceImpl) GetSemuaMahasiswa(ctx echo.Context) []transfer.Response {
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
	var responses []transfer.Response
	for _, data := range srv.repo.Get(query) {
		response := transfer.Response(data)
		responses = append(responses, response)
	}
	return responses
}

func (srv *MahasiswaServiceImpl) GetMahasiswaSpesifik(ctx echo.Context, npm string) *transfer.Response {
	token := helpers.GetJwtToken(ctx)
	key := jwt.NewJwtKey()
	if helpers.JwtValidate(token, key.AccessKey) {
		claim := helpers.JwtDecode[jwt.JwtPayload](&token[1])
		if claim.Role != "admin" {
			ctx.Set("authorization.error", true)
			return nil
		}
	}
	result := srv.repo.Find(npm)
	if result != nil {
		response := transfer.Response(*result)
		return &response
	}
	return nil
}

func (srv *MahasiswaServiceImpl) TambahMahasiswa(ctx echo.Context, request *transfer.RequestBody) *transfer.Response {
	token := helpers.GetJwtToken(ctx)
	key := jwt.NewJwtKey()
	if helpers.JwtValidate(token, key.AccessKey) {
		claim := helpers.JwtDecode[jwt.JwtPayload](&token[1])
		if claim.Role != "admin" {
			ctx.Set("authorization.error", true)
			return nil
		}
	}
	data := &repository.MahasiswaModel{
		NPM:   request.NPM,
		Nama:  request.Nama,
		Kelas: request.Kelas,
	}
	result := srv.repo.Create(data)
	if result != nil {
		return &transfer.Response{
			NPM:   result.NPM,
			Nama:  result.Nama,
			Kelas: result.Kelas,
		}
	}
	return nil
}

func (srv *MahasiswaServiceImpl) EditMahasiswa(ctx echo.Context, npm string, request *transfer.RequestBody) *transfer.Response {
	token := helpers.GetJwtToken(ctx)
	key := jwt.NewJwtKey()
	if helpers.JwtValidate(token, key.AccessKey) {
		claim := helpers.JwtDecode[jwt.JwtPayload](&token[1])
		if claim.Role != "admin" {
			ctx.Set("authorization.error", true)
			return nil
		}
	}
	data := &repository.MahasiswaModel{
		NPM:   request.NPM,
		Nama:  request.Nama,
		Kelas: request.Kelas,
	}
	result := srv.repo.Update(npm, data)
	if result != nil {
		return &transfer.Response{
			NPM:   result.NPM,
			Nama:  result.Nama,
			Kelas: result.Kelas,
		}
	}
	return nil
}

func (srv *MahasiswaServiceImpl) HapusMahasiswa(ctx echo.Context, npm string) bool {
	token := helpers.GetJwtToken(ctx)
	key := jwt.NewJwtKey()
	if helpers.JwtValidate(token, key.AccessKey) {
		claim := helpers.JwtDecode[jwt.JwtPayload](&token[1])
		if claim.Role != "admin" {
			ctx.Set("authorization.error", true)
			return false
		}
	}
	return srv.repo.Delete(npm)
}
