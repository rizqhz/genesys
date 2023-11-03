package repository

import (
	"net/url"

	"github.com/fatih/structs"
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
	if err := db.Where("deleted_at IS NULL").Find(&data).Error; err != nil {
		logrus.Error("[user.repository]: ", err.Error())
		return nil
	}
	return data
}

func (r *UserMySqlRepository) Find(id int) *UserEntity {
	db := r.driver.DB.Table("users")
	data := &UserEntity{ID: id}
	if err := db.Where("deleted_at IS NULL").First(&data).Error; err != nil {
		logrus.Error("[user.repository]: ", err.Error())
		return nil
	}
	return data
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
	search := &UserModel{ID: data.ID}
	if err := db.Find(&search).Error; err != nil {
		logrus.Error("[user.repository]: ", err.Error())
		return nil
	}
	model := func(old, new *UserModel) *UserModel {
		fields := []string{"Nama", "Alamat", "Email", "Telepon", "Foto"}
		n := structs.Map(new)
		o := structs.Map(old)
		result := make(map[string]interface{})
		for _, field := range fields {
			if n[field] != "" {
				result[field] = n[field]
			} else {
				result[field] = o[field]
			}
		}
		old.Nama = result["Nama"].(string)
		old.Alamat = result["Alamat"].(string)
		old.Email = result["Email"].(string)
		old.Telepon = result["Telepon"].(string)
		old.Foto = result["Foto"].(string)
		return old
	}(search, data)
	if err := db.Save(&model).Error; err != nil {
		logrus.Error("[user.repository]: ", err.Error())
		return nil
	}
	entity := &UserEntity{
		ID:      model.ID,
		Nama:    model.Nama,
		Alamat:  model.Alamat,
		Email:   model.Email,
		Telepon: model.Telepon,
		Foto:    model.Foto,
	}
	return entity
}

func (r *UserMySqlRepository) Delete(id int) bool {
	db := r.driver.DB.Table("users")
	if err := db.Delete(&UserModel{}, id).Error; err != nil {
		logrus.Error("[user.repository]: ", err.Error())
		return false
	}
	return true
}
