package service

import (
	"github.com/labstack/echo/v4"
	jwt "github.com/rizghz/genesys/infrastructure/middleware/JWT"
	"github.com/rizghz/genesys/internal/helpers"
	"github.com/rizghz/genesys/module/Asisten/repository"
	"github.com/rizghz/genesys/module/Asisten/transfer"
)

type AsistenServiceImpl struct {
	repo repository.AsistenRepository
}

func NewAsistenServiceImpl(r repository.AsistenRepository) AsistenService {
	return &AsistenServiceImpl{
		repo: r,
	}
}

func (srv *AsistenServiceImpl) GetSemuaAsisten(ctx echo.Context) []transfer.Response {
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

func (srv *AsistenServiceImpl) GetAsistenSpesifik(ctx echo.Context, nias string) *transfer.Response {
	token := helpers.GetJwtToken(ctx)
	key := jwt.NewJwtKey()
	if helpers.JwtValidate(token, key.AccessKey) {
		claim := helpers.JwtDecode[jwt.JwtPayload](&token[1])
		if claim.Role != "admin" {
			ctx.Set("authorization.error", true)
			return nil
		}
	}
	result := srv.repo.Find(nias)
	if result != nil {
		response := transfer.Response(*result)
		return &response
	}
	return nil
}

func (srv *AsistenServiceImpl) TambahAsisten(ctx echo.Context, request *transfer.RequestBody) *transfer.Response {
	token := helpers.GetJwtToken(ctx)
	key := jwt.NewJwtKey()
	if helpers.JwtValidate(token, key.AccessKey) {
		claim := helpers.JwtDecode[jwt.JwtPayload](&token[1])
		if claim.Role != "admin" {
			ctx.Set("authorization.error", true)
			return nil
		}
	}
	data := &repository.AsistenModel{
		NIAS:    request.NIAS,
		Nama:    request.Nama,
		Jabatan: request.Jabatan,
	}
	result := srv.repo.Create(data)
	if result != nil {
		return &transfer.Response{
			NIAS:    result.NIAS,
			Nama:    result.Nama,
			Jabatan: result.Jabatan,
		}
	}
	return nil
}

func (srv *AsistenServiceImpl) EditAsisten(ctx echo.Context, nias string, request *transfer.RequestBody) *transfer.Response {
	token := helpers.GetJwtToken(ctx)
	key := jwt.NewJwtKey()
	if helpers.JwtValidate(token, key.AccessKey) {
		claim := helpers.JwtDecode[jwt.JwtPayload](&token[1])
		if claim.Role != "admin" {
			ctx.Set("authorization.error", true)
			return nil
		}
	}
	data := &repository.AsistenModel{
		NIAS:    request.NIAS,
		Nama:    request.Nama,
		Jabatan: request.Jabatan,
	}
	result := srv.repo.Update(nias, data)
	if result != nil {
		return &transfer.Response{
			NIAS:    result.NIAS,
			Nama:    result.Nama,
			Jabatan: result.Jabatan,
		}
	}
	return nil
}

func (srv *AsistenServiceImpl) HapusAsisten(ctx echo.Context, nias string) bool {
	token := helpers.GetJwtToken(ctx)
	key := jwt.NewJwtKey()
	if helpers.JwtValidate(token, key.AccessKey) {
		claim := helpers.JwtDecode[jwt.JwtPayload](&token[1])
		if claim.Role != "admin" {
			ctx.Set("authorization.error", true)
			return false
		}
	}
	return srv.repo.Delete(nias)
}
