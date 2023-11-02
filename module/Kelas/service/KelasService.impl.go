package service

import (
	"net/url"

	"github.com/rizghz/genesys/module/Kelas/repository"
	"github.com/rizghz/genesys/module/Kelas/transfer"
)

type KelasServiceImpl struct {
	repo repository.KelasRepository
}

func NewKelasServiceImpl(r repository.KelasRepository) KelasService {
	return &KelasServiceImpl{
		repo: r,
	}
}

func (srv *KelasServiceImpl) GetSemuaKelas(query url.Values) []transfer.Response {
	response := make([]transfer.Response, 0)
	for _, res := range srv.repo.Get(query) {
		response = append(response, transfer.Response(res))
	}
	return response
}

func (srv *KelasServiceImpl) GetKelasSpesifik(kode string) *transfer.Response {
	data := srv.repo.Find(kode)
	return (*transfer.Response)(data)
}

func (srv *KelasServiceImpl) TambahKelas(data transfer.RequestBody) *transfer.Response {
	request := repository.Model{
		Kode:    data.Kode,
		Nama:    data.Nama,
		Jurusan: data.Jurusan,
		Grade:   data.Grade,
		Tahun:   data.Tahun,
	}
	if res := srv.repo.Create(&request); res != nil {
		return (*transfer.Response)(res)
	}
	return nil
}

func (srv *KelasServiceImpl) EditKelas(kode string, data transfer.RequestBody) *transfer.Response {
	request := repository.Model{
		Kode:    data.Kode,
		Nama:    data.Nama,
		Jurusan: data.Jurusan,
		Grade:   data.Grade,
		Tahun:   data.Tahun,
	}
	if res := srv.repo.Update(kode, &request); res != nil {
		return (*transfer.Response)(res)
	}
	return nil
}

func (srv *KelasServiceImpl) HapusKelas(kode string) bool {
	return srv.repo.Delete(kode)
}
