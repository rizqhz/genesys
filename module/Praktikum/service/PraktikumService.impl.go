package service

import (
	"github.com/labstack/echo/v4"
	jwt "github.com/rizghz/genesys/infrastructure/middleware/JWT"
	"github.com/rizghz/genesys/internal/helpers"
	"github.com/rizghz/genesys/module/Praktikum/repository"
	"github.com/rizghz/genesys/module/Praktikum/transfer"
)

type PraktikumServiceImpl struct {
	repo repository.PraktikumRepository
}

func NewPraktikumServiceImpl(r repository.PraktikumRepository) PraktikumService {
	return &PraktikumServiceImpl{
		repo: r,
	}
}

func (srv *PraktikumServiceImpl) GetSemuaPraktikum(ctx echo.Context) []transfer.Response {
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

func (srv *PraktikumServiceImpl) GetPraktikumSpesifik(ctx echo.Context, id string) *transfer.Response {
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

func (srv *PraktikumServiceImpl) TambahPraktikum(ctx echo.Context, request *transfer.RequestBody) *transfer.Response {
	token := helpers.GetJwtToken(ctx)
	key := jwt.NewJwtKey()
	if helpers.JwtValidate(token, key.AccessKey) {
		claim := helpers.JwtDecode[jwt.JwtPayload](&token[1])
		if claim.Role != "admin" {
			ctx.Set("authorization.error", true)
			return nil
		}
	}
	data := &repository.PraktikumModel{
		ID:                request.ID,
		KodeMataPraktikum: request.MataPraktikum,
		KodeRuangan:       request.Ruangan,
		KodeKelas:         request.Kelas,
		JadwalID:          request.Jadwal,
	}
	result := srv.repo.Create(data)
	if result != nil {
		return &transfer.Response{
			ID:                result.ID,
			KodeMataPraktikum: result.KodeMataPraktikum,
			KodeRuangan:       result.KodeRuangan,
			KodeKelas:         result.KodeKelas,
			JadwalID:          result.JadwalID,
		}
	}
	return nil
}

func (srv *PraktikumServiceImpl) EditPraktikum(ctx echo.Context, id string, request *transfer.RequestBody) *transfer.Response {
	token := helpers.GetJwtToken(ctx)
	key := jwt.NewJwtKey()
	if helpers.JwtValidate(token, key.AccessKey) {
		claim := helpers.JwtDecode[jwt.JwtPayload](&token[1])
		if claim.Role != "admin" {
			ctx.Set("authorization.error", true)
			return nil
		}
	}
	data := &repository.PraktikumModel{
		ID:                request.ID,
		KodeMataPraktikum: request.MataPraktikum,
		KodeRuangan:       request.Ruangan,
		KodeKelas:         request.Kelas,
		JadwalID:          request.Jadwal,
	}
	result := srv.repo.Update(id, data)
	if result != nil {
		return &transfer.Response{
			ID:                result.ID,
			KodeMataPraktikum: result.KodeMataPraktikum,
			KodeRuangan:       result.KodeRuangan,
			KodeKelas:         result.KodeKelas,
			JadwalID:          result.JadwalID,
		}
	}
	return nil
}

func (srv *PraktikumServiceImpl) HapusPraktikum(ctx echo.Context, id string) bool {
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
