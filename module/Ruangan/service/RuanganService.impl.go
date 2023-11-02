package service

import (
	"net/url"

	"github.com/rizghz/genesys/module/Ruangan/repository"
	"github.com/rizghz/genesys/module/Ruangan/transfer"
)

type RuanganServiceImpl struct {
	repo repository.RuanganRepository
}

func NewRuanganServiceImpl(r repository.RuanganRepository) RuanganService {
	return &RuanganServiceImpl{
		repo: r,
	}
}

func (srv *RuanganServiceImpl) GetSemuaRuangan(query url.Values) []transfer.Response {
	response := make([]transfer.Response, 0)
	for _, res := range srv.repo.Get(query) {
		response = append(response, transfer.Response(res))
	}
	return response
}

func (srv *RuanganServiceImpl) GetRuanganSpesifik(kode string) *transfer.Response {
	data := srv.repo.Find(kode)
	return (*transfer.Response)(data)
}

func (srv *RuanganServiceImpl) TambahRuangan(data transfer.RequestBody) *transfer.Response {
	request := repository.Model{
		Kode: data.Kode,
		Nama: data.Nama,
	}
	if res := srv.repo.Create(&request); res != nil {
		return (*transfer.Response)(res)
	}
	return nil
}

func (srv *RuanganServiceImpl) EditRuangan(kode string, data transfer.RequestBody) *transfer.Response {
	request := repository.Model{
		Kode: data.Kode,
		Nama: data.Nama,
	}
	if res := srv.repo.Update(kode, &request); res != nil {
		return (*transfer.Response)(res)
	}
	return nil
}

func (srv *RuanganServiceImpl) HapusRuangan(kode string) bool {
	return srv.repo.Delete(kode)
}
