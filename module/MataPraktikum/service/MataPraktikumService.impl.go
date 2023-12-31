package service

import (
	"github.com/labstack/echo/v4"
	jwt "github.com/rizghz/genesys/infrastructure/middleware/JWT"
	"github.com/rizghz/genesys/internal/helpers"
	"github.com/rizghz/genesys/module/MataPraktikum/repository"
	"github.com/rizghz/genesys/module/MataPraktikum/transfer"
)

type MatkumServiceImpl struct {
	repo repository.MatkumRepository
}

func NewMatkumServiceImpl(r repository.MatkumRepository) MatkumService {
	return &MatkumServiceImpl{
		repo: r,
	}
}

func (srv *MatkumServiceImpl) GetSemuaMatkum(ctx echo.Context) []transfer.Response {
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

func (srv *MatkumServiceImpl) GetMatkumSpesifik(ctx echo.Context, kode string) *transfer.Response {
	token := helpers.GetJwtToken(ctx)
	key := jwt.NewJwtKey()
	if helpers.JwtValidate(token, key.AccessKey) {
		claim := helpers.JwtDecode[jwt.JwtPayload](&token[1])
		if claim.Role != "admin" {
			ctx.Set("authorization.error", true)
			return nil
		}
	}
	result := srv.repo.Find(kode)
	if result != nil {
		response := transfer.Response(*result)
		return &response
	}
	return nil
}

func (srv *MatkumServiceImpl) TambahMatkum(ctx echo.Context, request *transfer.RequestBody) *transfer.Response {
	token := helpers.GetJwtToken(ctx)
	key := jwt.NewJwtKey()
	if helpers.JwtValidate(token, key.AccessKey) {
		claim := helpers.JwtDecode[jwt.JwtPayload](&token[1])
		if claim.Role != "admin" {
			ctx.Set("authorization.error", true)
			return nil
		}
	}
	data := &repository.MatkumModel{
		Kode: request.Kode,
		Nama: request.Nama,
	}
	result := srv.repo.Create(data)
	if result != nil {
		return &transfer.Response{
			Kode: result.Kode,
			Nama: result.Nama,
		}
	}
	return nil
}

func (srv *MatkumServiceImpl) EditMatkum(ctx echo.Context, kode string, request *transfer.RequestBody) *transfer.Response {
	token := helpers.GetJwtToken(ctx)
	key := jwt.NewJwtKey()
	if helpers.JwtValidate(token, key.AccessKey) {
		claim := helpers.JwtDecode[jwt.JwtPayload](&token[1])
		if claim.Role != "admin" {
			ctx.Set("authorization.error", true)
			return nil
		}
	}
	data := &repository.MatkumModel{
		Kode: request.Kode,
		Nama: request.Nama,
	}
	result := srv.repo.Update(kode, data)
	if result != nil {
		return &transfer.Response{
			Kode: result.Kode,
			Nama: result.Nama,
		}
	}
	return nil
}

func (srv *MatkumServiceImpl) HapusMatkum(ctx echo.Context, kode string) bool {
	token := helpers.GetJwtToken(ctx)
	key := jwt.NewJwtKey()
	if helpers.JwtValidate(token, key.AccessKey) {
		claim := helpers.JwtDecode[jwt.JwtPayload](&token[1])
		if claim.Role != "admin" {
			ctx.Set("authorization.error", true)
			return false
		}
	}
	return srv.repo.Delete(kode)
}
