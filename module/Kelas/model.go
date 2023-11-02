package kelas

import (
	"time"

	mahasiswa "github.com/rizghz/genesys/module/Mahasiswa"
	"gorm.io/gorm"
)

type KelasModel struct {
	Kode      string                     `gorm:"type:char(10);primaryKey"`
	Nama      string                     `gorm:"type:varchar(10)"`
	Jurusan   string                     `gorm:"type:varchar(32)"`
	Grade     string                     `gorm:"type:char(2)"`
	Tahun     int                        `gorm:"type:year"`
	Mahasiswa []mahasiswa.MahasiswaModel `gorm:"foreignKey:Kelas;references:Kode;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt time.Time                  `gorm:"autoCreateTime"`
	UpdatedAt time.Time                  `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt             `gorm:"index"`
}

func (KelasModel) TableName() string {
	return "kelas"
}
