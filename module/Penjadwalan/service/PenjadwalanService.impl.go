package service

import (
	"github.com/labstack/echo/v4"
	jwt "github.com/rizghz/genesys/infrastructure/middleware/JWT"
	"github.com/rizghz/genesys/internal/helpers"
	"github.com/rizghz/genesys/module/Penjadwalan/repository"
	"github.com/rizghz/genesys/module/Penjadwalan/transfer"
)

type PenjadwalanServiceImpl struct {
	repo repository.PenjadwalanRepository
}

func NewPenjadwalanServiceImpl(r repository.PenjadwalanRepository) PenjadwalanService {
	return &PenjadwalanServiceImpl{
		repo: r,
	}
}

func (srv *PenjadwalanServiceImpl) GetSemuaJadwal(ctx echo.Context) []transfer.Response {
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

func (srv *PenjadwalanServiceImpl) GetJadwalSpesifik(ctx echo.Context, id int) *transfer.Response {
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
		response := transfer.Response(*result)
		return &response
	}
	return nil
}

func (srv *PenjadwalanServiceImpl) TambahJadwal(ctx echo.Context, request *transfer.RequestBody) *transfer.Response {
	token := helpers.GetJwtToken(ctx)
	key := jwt.NewJwtKey()
	if helpers.JwtValidate(token, key.AccessKey) {
		claim := helpers.JwtDecode[jwt.JwtPayload](&token[1])
		if claim.Role != "admin" {
			ctx.Set("authorization.error", true)
			return nil
		}
	}
	data := &repository.JadwalModel{
		Asisten: request.Asisten,
		Hari:    request.Hari,
		Jam:     request.Jam,
	}
	result := srv.repo.Create(data)
	if result != nil {
		return &transfer.Response{
			ID:      result.ID,
			Asisten: result.Asisten,
			Hari:    result.Hari,
			Jam:     result.Jam,
		}
	}
	return nil
}

func (srv *PenjadwalanServiceImpl) EditJadwal(ctx echo.Context, id int, request *transfer.RequestBody) *transfer.Response {
	token := helpers.GetJwtToken(ctx)
	key := jwt.NewJwtKey()
	if helpers.JwtValidate(token, key.AccessKey) {
		claim := helpers.JwtDecode[jwt.JwtPayload](&token[1])
		if claim.Role != "admin" {
			ctx.Set("authorization.error", true)
			return nil
		}
	}
	data := &repository.JadwalModel{
		Asisten: request.Asisten,
		Hari:    request.Hari,
		Jam:     request.Jam,
	}
	result := srv.repo.Update(data)
	if result != nil {
		return &transfer.Response{
			ID:      result.ID,
			Asisten: result.Asisten,
			Hari:    result.Hari,
			Jam:     result.Jam,
		}
	}
	return nil
}

func (srv *PenjadwalanServiceImpl) HapusJadwal(ctx echo.Context, id int) bool {
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
