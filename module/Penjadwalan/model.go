package penjadwalan

import (
	"time"

	"gorm.io/gorm"
)

type PenjadwalanModel struct {
	ID        int            `gorm:"type:int;primaryKey;autoIncrement"`
	Asisten   string         `gorm:"type:char(10)"`
	Hari      string         `gorm:"type:varchar(8)"`
	Jam       string         `gorm:"type:varchar(16)"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (PenjadwalanModel) TableName() string {
	return "penjadwalan"
}
