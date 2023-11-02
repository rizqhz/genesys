package service

import (
	"net/url"

	"github.com/rizghz/genesys/module/MataPraktikum/repository"
	"github.com/rizghz/genesys/module/MataPraktikum/transfer"
)

type MataPraktikumServiceImpl struct {
	repo repository.MataPraktikumRepository
}

func NewMataPraktikumServiceImpl(r repository.MataPraktikumRepository) MataPraktikumService {
	return &MataPraktikumServiceImpl{
		repo: r,
	}
}

func (srv *MataPraktikumServiceImpl) GetSemuaMataPraktikum(query url.Values) []transfer.Response {
	response := make([]transfer.Response, 0)
	for _, res := range srv.repo.Get(query) {
		response = append(response, transfer.Response(res))
	}
	return response
}

func (srv *MataPraktikumServiceImpl) GetMataPraktikumSpesifik(kode string) *transfer.Response {
	data := srv.repo.Find(kode)
	return (*transfer.Response)(data)
}

func (srv *MataPraktikumServiceImpl) TambahMataPraktikum(data transfer.RequestBody) *transfer.Response {
	request := repository.Model{
		Kode: data.Kode,
		Nama: data.Nama,
	}
	if res := srv.repo.Create(&request); res != nil {
		return (*transfer.Response)(res)
	}
	return nil
}

func (srv *MataPraktikumServiceImpl) EditMataPraktikum(kode string, data transfer.RequestBody) *transfer.Response {
	request := repository.Model{
		Kode: data.Kode,
		Nama: data.Nama,
	}
	if res := srv.repo.Update(kode, &request); res != nil {
		return (*transfer.Response)(res)
	}
	return nil
}

func (srv *MataPraktikumServiceImpl) HapusMataPraktikum(kode string) bool {
	return srv.repo.Delete(kode)
}
