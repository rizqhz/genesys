package repository

import (
	"net/url"

	mahasiswa "github.com/rizghz/genesys/module/Mahasiswa"
)

type MahasiswaEntity mahasiswa.MahasiswaEntity
type MahasiswaModel mahasiswa.MahasiswaModel

type MahasiswaRepository interface {
	Get(query url.Values) []MahasiswaEntity
	Find(npm string) *MahasiswaEntity
	Create(data *MahasiswaModel) *MahasiswaEntity
	Update(npm string, data *MahasiswaModel) *MahasiswaEntity
	Delete(npm string) bool
}
