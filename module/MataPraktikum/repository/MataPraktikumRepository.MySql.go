package repository

import (
	"fmt"
	"net/url"

	"github.com/fatih/structs"
	mysql "github.com/rizghz/genesys/infrastructure/database/MySql"
	"github.com/rizghz/genesys/internal/helpers"
	"github.com/sirupsen/logrus"
)

type MatkumMySqlRepository struct {
	driver *mysql.MySqlDriver
}

func NewMatkumMySqlRepository(driver *mysql.MySqlDriver) MatkumRepository {
	return &MatkumMySqlRepository{
		driver: driver,
	}
}

func (repo *MatkumMySqlRepository) Get(query url.Values) []MatkumEntity {
	db := repo.driver.DB.Table("mata_praktikum")
	db = helpers.QuerySorting(db, query)
	db = helpers.QueryPagination(db, query)
	db = helpers.QueryFiltering(db, query)
	data := make([]MatkumEntity, 0)
	if err := db.Where("deleted_at IS NULL").Find(&data).Error; err != nil {
		logrus.Error("[matkum.repository]: ", err.Error())
		return nil
	}
	return data
}

func (repo *MatkumMySqlRepository) Find(kode string) *MatkumEntity {
	db := repo.driver.DB.Table("mata_praktikum")
	data := &MatkumEntity{Kode: kode}
	cond := fmt.Sprintf("kode = '%s'", kode)
	if err := db.Where("deleted_at IS NULL").First(&data, cond).Error; err != nil {
		logrus.Error("[matkum.repository]: ", err.Error())
		return nil
	}
	return data
}

func (repo *MatkumMySqlRepository) Create(data *MatkumModel) *MatkumEntity {
	db := repo.driver.DB.Table("mata_praktikum")
	if err := db.Create(&data).Error; err != nil {
		logrus.Error("[matkum.repository]: ", err.Error())
		return nil
	}
	return &MatkumEntity{
		Kode: data.Kode,
		Nama: data.Nama,
	}
}

func (repo *MatkumMySqlRepository) Update(kode string, data *MatkumModel) *MatkumEntity {
	db := repo.driver.DB.Table("mata_praktikum")
	search := &MatkumModel{Kode: data.Kode}
	cond := fmt.Sprintf("kode = '%s'", kode)
	if err := db.Find(&search, cond).Error; err != nil {
		logrus.Error("[matkum.repository]: ", err.Error())
		return nil
	}
	model := func(old, new *MatkumModel) *MatkumModel {
		fields := []string{"Kode", "Nama"}
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
		old.Kode = result["Kode"].(string)
		old.Nama = result["Nama"].(string)
		return old
	}(search, data)
	if err := db.Where(cond).Save(&model).Error; err != nil {
		logrus.Error("[matkum.repository]: ", err.Error())
		return nil
	}
	return &MatkumEntity{
		Kode: model.Kode,
		Nama: model.Nama,
	}
}

func (repo *MatkumMySqlRepository) Delete(kode string) bool {
	db := repo.driver.DB.Table("mata_praktikum")
	cond := fmt.Sprintf("kode = '%s'", kode)
	if err := db.Delete(&MatkumModel{}, cond).Error; err != nil {
		logrus.Error("[matkum.repository]: ", err.Error())
		return false
	}
	return true
}
