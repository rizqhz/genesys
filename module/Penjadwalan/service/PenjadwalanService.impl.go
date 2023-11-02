package service

import (
	"github.com/labstack/echo/v4"
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
	query := ctx.QueryParams()
	var response []transfer.Response
	for _, data := range srv.repo.Get(query) {
		response = append(response, transfer.Response(data))
	}
	return response
}

func (srv *PenjadwalanServiceImpl) GetJadwalSpesifik(id int) *transfer.Response {
	response := srv.repo.Find(id)
	if response != nil {
		return (*transfer.Response)(response)
	}
	return nil
}

func (srv *PenjadwalanServiceImpl) TambahJadwal(request *transfer.RequestBody) *transfer.Response {
	data := &repository.Model{
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

func (srv *PenjadwalanServiceImpl) EditJadwal(id int, request *transfer.RequestBody) *transfer.Response {
	data := &repository.Model{
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

func (srv *PenjadwalanServiceImpl) HapusJadwal(id int) bool {
	return srv.repo.Delete(id)
}
