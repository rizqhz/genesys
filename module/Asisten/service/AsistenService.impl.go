package service

import (
	"net/url"

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

func (srv *AsistenServiceImpl) GetSemuaAsisten(query url.Values) []transfer.Response {
	response := make([]transfer.Response, 0)
	for _, res := range srv.repo.Get(query) {
		response = append(response, transfer.Response(res))
	}
	return response
}

func (srv *AsistenServiceImpl) GetAsistenSpesifik(nias string) *transfer.Response {
	data := srv.repo.Find(nias)
	return (*transfer.Response)(data)
}

func (srv *AsistenServiceImpl) TambahAsisten(data transfer.RequestBody) *transfer.Response {
	request := repository.Model{
		NIAS:    data.NIAS,
		Nama:    data.Nama,
		Jabatan: data.Jabatan,
	}
	if res := srv.repo.Create(&request); res != nil {
		return (*transfer.Response)(res)
	}
	return nil
}

func (srv *AsistenServiceImpl) EditAsisten(nias string, data transfer.RequestBody) *transfer.Response {
	request := repository.Model{
		NIAS:    data.NIAS,
		Nama:    data.Nama,
		Jabatan: data.Jabatan,
	}
	if res := srv.repo.Update(nias, &request); res != nil {
		return (*transfer.Response)(res)
	}
	return nil
}

func (srv *AsistenServiceImpl) HapusAsisten(nias string) bool {
	return srv.repo.Delete(nias)
}
