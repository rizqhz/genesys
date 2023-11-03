package praktikum

import (
	"time"

	kelas "github.com/rizghz/genesys/module/Kelas"
	matkum "github.com/rizghz/genesys/module/MataPraktikum"
	jadwal "github.com/rizghz/genesys/module/Penjadwalan"
	ruangan "github.com/rizghz/genesys/module/Ruangan"
	"gorm.io/gorm"
)

type PraktikumModel struct {
	ID                string                    `gorm:"type:char(32);primaryKey"`
	KodeMataPraktikum string                    `gorm:"type:char(7)"`
	MataPraktikum     matkum.MataPraktikumModel `gorm:"foreignKey:KodeMataPraktikum;references:Kode;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	KodeRuangan       string                    `gorm:"type:char(10)"`
	Ruangan           ruangan.RuanganModel      `gorm:"foreignKey:KodeRuangan;references:Kode;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	KodeKelas         string                    `gorm:"type:char(10)"`
	Kelas             kelas.KelasModel          `gorm:"foreignKey:KodeKelas;references:Kode;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	JadwalID          int                       `gorm:"type:int;unique"`
	Jadwal            jadwal.PenjadwalanModel   `gorm:"foreignKey:JadwalID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt         time.Time                 `gorm:"autoCreateTime"`
	UpdatedAt         time.Time                 `gorm:"autoUpdateTime"`
	DeletedAt         gorm.DeletedAt            `gorm:"index"`
}

func (PraktikumModel) TableName() string {
	return "praktikum"
}
