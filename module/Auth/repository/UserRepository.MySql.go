package repository

import (
	"net/url"

	mysql "github.com/rizghz/genesys/infrastructure/database/MySql"
	"github.com/rizghz/genesys/internal/helpers"
	"github.com/sirupsen/logrus"
)

type UserMySqlRepository struct {
	driver *mysql.MySqlDriver
}

func NewUserMySqlRepository(driver *mysql.MySqlDriver) UserRepository {
	return &UserMySqlRepository{
		driver: driver,
	}
}

func (r *UserMySqlRepository) Get(query url.Values) []UserEntity {
	db := r.driver.DB.Table("users")
	helpers.QuerySorting(db, query)
	helpers.QueryPagination(db, query)
	helpers.QueryFiltering(db, query)
	data := make([]UserEntity, 0)
	if err := db.Find(&data).Error; err != nil {
		logrus.Error("[user.repository]: ", err.Error())
		return nil
	}
	return data
}

func (r *UserMySqlRepository) Find(id int) *UserEntity {
	db := r.driver.DB.Table("users")
	data := UserEntity{ID: id}
	// condition := fmt.Sprintf("id = '%d'", id)
	if err := db.First(&data).Error; err != nil {
		logrus.Error("[user.repository]: ", err.Error())
		return nil
	}
	return &data
}

func (r *UserMySqlRepository) Create(data *UserModel) *UserEntity {
	db := r.driver.DB.Table("users")
	if err := db.Create(&data).Error; err != nil {
		logrus.Error("[user.repository]: ", err.Error())
		return nil
	}
	entity := &UserEntity{
		ID:      data.ID,
		Nama:    data.Nama,
		Alamat:  data.Alamat,
		Email:   data.Email,
		Telepon: data.Telepon,
		Foto:    data.Foto,
	}
	return entity
}

func (r *UserMySqlRepository) Update(data *UserModel) *UserEntity {
	db := r.driver.DB.Table("users")
	if err := db.Save(&data).Error; err != nil {
		logrus.Error("[user.repository]: ", err.Error())
		return nil
	}
	entity := &UserEntity{
		ID:      data.ID,
		Nama:    data.Nama,
		Alamat:  data.Alamat,
		Email:   data.Email,
		Telepon: data.Telepon,
		Foto:    data.Foto,
	}
	return entity
}

func (r *UserMySqlRepository) Delete(id int) bool {
	db := r.driver.DB.Table("users")
	if err := db.Delete(&UserModel{ID: id}).Error; err != nil {
		logrus.Error("[user.repository]: ", err.Error())
		return false
	}
	return true
}
