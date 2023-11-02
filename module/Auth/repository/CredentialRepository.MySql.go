package repository

import (
	"fmt"

	mysql "github.com/rizghz/genesys/infrastructure/database/MySql"
	auth "github.com/rizghz/genesys/module/Auth"
	"github.com/rizghz/genesys/module/Auth/transfer"
	log "github.com/sirupsen/logrus"
)

type CredentialMySqlRepository struct {
	driver *mysql.MySqlDriver
}

func NewCredentialMySqlRepository(driver *mysql.MySqlDriver) CredentialRepository {
	return &CredentialMySqlRepository{
		driver: driver,
	}
}

func (r *CredentialMySqlRepository) Get() []auth.CredentialModel {
	db := r.driver.DB
	data := make([]auth.CredentialModel, 0)
	if err := db.Find(&data).Error; err != nil {
		log.Error("user.repository:", err.Error())
		return nil
	}
	return data
}

func (repo *CredentialMySqlRepository) Create(data *auth.UserModel) *transfer.RegisterResponse {
	db := repo.driver.DB
	if err := db.Create(&data).Error; err != nil {
		log.Error("[credential.repository]: ", err.Error())
		return nil
	}
	return &transfer.RegisterResponse{
		UserID:  data.ID,
		Nama:    data.Nama,
		Email:   data.Email,
		Telepon: data.Telepon,
		Credential: transfer.CredentialResponse{
			Usercode: data.Credential.Usercode,
			Password: data.Credential.Password,
			Role:     data.Credential.Role,
		},
	}
}

func (repo *CredentialMySqlRepository) Update(data *auth.UserModel) *auth.UserModel {
	db := repo.driver.DB
	credential := data.Credential
	if err := db.Save(&credential).Error; err != nil {
		log.Error("[credential.repository]: ", err.Error())
		return nil
	}
	return data
}

func (repo *CredentialMySqlRepository) Search(data *auth.CredentialModel) *auth.UserModel {
	db := repo.driver.DB
	credential := &auth.CredentialModel{}
	condition := fmt.Sprintf("usercode = '%s' AND password = '%s'", data.Usercode, data.Password)
	db.Where(condition).Find(&credential)
	if credential.ID != 0 {
		user := &auth.UserModel{
			Credential: *credential,
		}
		db.Preload("Credential").Find(&user, credential.UserID)
		return user
	}
	return nil
}

func (repo *CredentialMySqlRepository) Delete(data *auth.CredentialModel) bool {
	db := repo.driver.DB
	db.Find(&data)
	data.Token = ""
	if err := db.Save(&data).Error; err != nil {
		log.Error("[credential.repository]: ", err.Error())
		return false
	}
	return true
}
