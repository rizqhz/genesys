package service

import (
	"net/url"

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

func (srv *MahasiswaServiceImpl) GetSemuaMahasiswa(query url.Values) []transfer.Response {
	response := make([]transfer.Response, 0)
	for _, res := range srv.repo.Get(query) {
		response = append(response, transfer.Response(res))
	}
	return response
}

func (srv *MahasiswaServiceImpl) GetMahasiswaSpesifik(npm string) *transfer.Response {
	data := srv.repo.Find(npm)
	return (*transfer.Response)(data)
}

func (srv *MahasiswaServiceImpl) TambahMahasiswa(data transfer.RequestBody) *transfer.Response {
	request := repository.Model{
		NPM:   data.NPM,
		Nama:  data.Nama,
		Kelas: data.Kelas,
	}
	if res := srv.repo.Create(&request); res != nil {
		return (*transfer.Response)(res)
	}
	return nil
}

func (srv *MahasiswaServiceImpl) EditMahasiswa(npm string, data transfer.RequestBody) *transfer.Response {
	request := repository.Model{
		NPM:   data.NPM,
		Nama:  data.Nama,
		Kelas: data.Kelas,
	}
	if res := srv.repo.Update(npm, &request); res != nil {
		return (*transfer.Response)(res)
	}
	return nil
}

func (srv *MahasiswaServiceImpl) HapusMahasiswa(npm string) bool {
	return srv.repo.Delete(npm)
}
