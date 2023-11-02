package mahasiswa

import (
	"time"

	"gorm.io/gorm"
)

type MahasiswaModel struct {
	NPM       string         `gorm:"type:char(10);primaryKey"`
	Nama      string         `gorm:"type:char(64)"`
	Kelas     string         `gorm:"type:char(10)"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (MahasiswaModel) TableName() string {
	return "mahasiswa"
}
