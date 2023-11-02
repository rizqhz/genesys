package auth

import (
	"time"

	"gorm.io/gorm"
)

type UserModel struct {
	ID         int             `gorm:"type:int;primaryKey;autoIncrement"`
	Nama       string          `gorm:"type:varchar(64)"`
	Alamat     string          `gorm:"type:text"`
	Email      string          `gorm:"type:varchar(32)"`
	Telepon    string          `gorm:"type:varchar(16)"`
	Foto       string          `gorm:"type:varchar(256)"`
	Credential CredentialModel `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt  time.Time       `gorm:"autoCreateTime"`
	UpdatedAt  time.Time       `gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt  `gorm:"index"`
}

func (UserModel) TableName() string {
	return "users"
}

type CredentialModel struct {
	ID        int            `gorm:"type:int;primaryKey;autoIncrement"`
	Usercode  string         `gorm:"type:char(10);unique"`
	Password  string         `gorm:"type:varchar(256)"`
	Role      string         `gorm:"type:enum('admin', 'asisten', 'mahasiswa', 'dosen')"`
	Token     string         `gorm:"type:varchar(256)"`
	UserID    int            `gorm:"type:int;unique"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (CredentialModel) TableName() string {
	return "credentials"
}
