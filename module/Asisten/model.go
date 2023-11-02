package asisten

import (
	"time"

	penjadwalan "github.com/rizghz/genesys/module/Penjadwalan"
	"gorm.io/gorm"
)

type AsistenModel struct {
	NIAS      string                       `gorm:"type:char(10);primaryKey"`
	Nama      string                       `gorm:"type:varchar(64)"`
	Jabatan   string                       `gorm:"type:varchar(32)"`
	Jadwal    penjadwalan.PenjadwalanModel `gorm:"foreignKey:Asisten;references:NIAS;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt time.Time                    `gorm:"autoCreateTime"`
	UpdatedAt time.Time                    `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt
}

func (AsistenModel) TableName() string {
	return "asisten"
}
