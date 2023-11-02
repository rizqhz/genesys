package ruangan

import (
	"time"

	"gorm.io/gorm"
)

type RuanganModel struct {
	Kode      string         `gorm:"type:char(10);primaryKey"`
	Nama      string         `gorm:"type:varchar(32)"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (RuanganModel) TableName() string {
	return "ruangan"
}
