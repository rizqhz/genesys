package repository

import (
	"fmt"
	"net/url"

	"github.com/fatih/structs"
	mysql "github.com/rizghz/genesys/infrastructure/database/MySql"
	"github.com/rizghz/genesys/internal/helpers"
	"github.com/sirupsen/logrus"
)

type RuanganMySqlRepository struct {
	driver *mysql.MySqlDriver
}

func NewRuanganMySqlRepository(driver *mysql.MySqlDriver) RuanganRepository {
	return &RuanganMySqlRepository{
		driver: driver,
	}
}

func (repo *RuanganMySqlRepository) Get(query url.Values) []RuanganEntity {
	db := repo.driver.DB.Table("ruangan")
	db = helpers.QuerySorting(db, query)
	db = helpers.QueryPagination(db, query)
	db = helpers.QueryFiltering(db, query)
	data := make([]RuanganEntity, 0)
	if err := db.Where("deleted_at IS NULL").Find(&data).Error; err != nil {
		logrus.Error("[ruangan.repository]: ", err.Error())
		return nil
	}
	return data
}

func (repo *RuanganMySqlRepository) Find(kode string) *RuanganEntity {
	db := repo.driver.DB.Table("ruangan")
	data := &RuanganEntity{Kode: kode}
	cond := fmt.Sprintf("kode = '%s'", kode)
	if err := db.Where("deleted_at IS NULL").First(&data, cond).Error; err != nil {
		logrus.Error("[ruangan.repository]: ", err.Error())
		return nil
	}
	return data
}

func (repo *RuanganMySqlRepository) Create(data *RuanganModel) *RuanganEntity {
	db := repo.driver.DB.Table("ruangan")
	if err := db.Create(&data).Error; err != nil {
		logrus.Error("[ruangan.repository]: ", err.Error())
		return nil
	}
	return &RuanganEntity{
		Kode: data.Kode,
		Nama: data.Nama,
	}
}

func (repo *RuanganMySqlRepository) Update(kode string, data *RuanganModel) *RuanganEntity {
	db := repo.driver.DB.Table("ruangan")
	search := &RuanganModel{Kode: data.Kode}
	cond := fmt.Sprintf("kode = '%s'", kode)
	if err := db.Find(&search, cond).Error; err != nil {
		logrus.Error("[ruangan.repository]: ", err.Error())
		return nil
	}
	model := func(old, new *RuanganModel) *RuanganModel {
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
		logrus.Error("[ruangan.repository]: ", err.Error())
		return nil
	}
	return &RuanganEntity{
		Kode: model.Kode,
		Nama: model.Nama,
	}
}

func (repo *RuanganMySqlRepository) Delete(kode string) bool {
	db := repo.driver.DB.Table("ruangan")
	cond := fmt.Sprintf("kode = '%s'", kode)
	if err := db.Delete(&RuanganModel{}, cond).Error; err != nil {
		logrus.Error("[ruangan.repository]: ", err.Error())
		return false
	}
	return true
}
