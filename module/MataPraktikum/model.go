package mata_praktikum

import (
	"time"

	"gorm.io/gorm"
)

type MataPraktikumModel struct {
	Kode      string    `gorm:"type:char(7);primaryKey"`
	Nama      string    `gorm:"type:varchar(32)"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt
}

func (MataPraktikumModel) TableName() string {
	return "mata_praktikum"
}
