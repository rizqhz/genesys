package repository

import (
	"fmt"
	"net/url"

	"github.com/fatih/structs"
	mysql "github.com/rizghz/genesys/infrastructure/database/MySql"
	"github.com/rizghz/genesys/internal/helpers"
	"github.com/sirupsen/logrus"
)

type PraktikumMySqlRepository struct {
	driver *mysql.MySqlDriver
}

func NewPraktikumMySqlRepository(drv *mysql.MySqlDriver) PraktikumRepository {
	return &PraktikumMySqlRepository{
		driver: drv,
	}
}

func (repo *PraktikumMySqlRepository) Get(query url.Values) []PraktikumEntity {
	db := repo.driver.DB.Table("praktikum")
	db = helpers.QuerySorting(db, query)
	db = helpers.QueryPagination(db, query)
	db = helpers.QueryFiltering(db, query)
	data := make([]PraktikumEntity, 0)
	if err := db.Where("deleted_at IS NULL").Find(&data).Error; err != nil {
		logrus.Error("[praktikum.repository]: ", err.Error())
		return nil
	}
	return data
}

func (repo *PraktikumMySqlRepository) Find(id string) *PraktikumEntity {
	db := repo.driver.DB.Table("praktikum")
	data := &PraktikumEntity{}
	cond := fmt.Sprintf("id = '%s'", id)
	if err := db.Where("deleted_at IS NULL").First(&data, cond).Error; err != nil {
		logrus.Error("[praktikum.repository]: ", err.Error())
		return nil
	}
	return data
}

func (repo *PraktikumMySqlRepository) Create(data *PraktikumModel) *PraktikumEntity {
	db := repo.driver.DB.Table("praktikum")
	if err := db.Create(&data).Error; err != nil {
		logrus.Error("[praktikum.repository]: ", err.Error())
		return nil
	}
	return &PraktikumEntity{
		ID:                data.ID,
		KodeMataPraktikum: data.KodeMataPraktikum,
		KodeRuangan:       data.KodeRuangan,
		KodeKelas:         data.KodeKelas,
		JadwalID:          data.JadwalID,
	}
}

func (repo *PraktikumMySqlRepository) Update(id string, data *PraktikumModel) *PraktikumEntity {
	db := repo.driver.DB.Table("praktikum")
	search := &PraktikumModel{ID: data.ID}
	cond := fmt.Sprintf("id = '%s'", id)
	if err := db.Find(&search, cond).Error; err != nil {
		logrus.Error("[praktikum.repository]: ", err.Error())
		return nil
	}
	model := func(old, new *PraktikumModel) *PraktikumModel {
		fields := []string{"ID", "KodeMataPraktikum", "KodeRuangan", "KodeKelas", "JadwalID"}
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
		old.ID = result["ID"].(string)
		old.KodeMataPraktikum = result["KodeMataPraktikum"].(string)
		old.KodeRuangan = result["KodeRuangan"].(string)
		old.KodeKelas = result["KodeKelas"].(string)
		old.JadwalID = result["JadwalID"].(int)
		return old
	}(search, data)
	if err := db.Where(cond).Save(&model).Error; err != nil {
		logrus.Error("[praktikum.repository]: ", err.Error())
		return nil
	}
	return &PraktikumEntity{
		ID:                model.ID,
		KodeMataPraktikum: model.KodeMataPraktikum,
		KodeRuangan:       model.KodeRuangan,
		KodeKelas:         model.KodeKelas,
		JadwalID:          model.JadwalID,
	}
}

func (repo *PraktikumMySqlRepository) Delete(id string) bool {
	db := repo.driver.DB.Table("praktikum")
	cond := fmt.Sprintf("id = '%s'", id)
	if err := db.Delete(&PraktikumModel{}, cond).Error; err != nil {
		logrus.Error("[praktikum.repository]: ", err.Error())
		return false
	}
	return true
}
