package repository

import (
	"net/url"

	"github.com/fatih/structs"
	mysql "github.com/rizghz/genesys/infrastructure/database/MySql"
	"github.com/rizghz/genesys/internal/helpers"
	"github.com/sirupsen/logrus"
)

type PenjadwalanMySqlRepository struct {
	driver *mysql.MySqlDriver
}

func NewPenjadwalanMySqlRepository(driver *mysql.MySqlDriver) PenjadwalanRepository {
	return &PenjadwalanMySqlRepository{
		driver: driver,
	}
}

func (r *PenjadwalanMySqlRepository) Get(query url.Values) []JadwalEntity {
	db := r.driver.DB.Table("penjadwalan")
	helpers.QuerySorting(db, query)
	helpers.QueryPagination(db, query)
	helpers.QueryFiltering(db, query)
	data := make([]JadwalEntity, 0)
	if err := db.Where("deleted_at IS NULL").Find(&data).Error; err != nil {
		logrus.Error("[jadwal.repository]: ", err.Error())
		return nil
	}
	return data
}

func (r *PenjadwalanMySqlRepository) Find(id int) *JadwalEntity {
	db := r.driver.DB.Table("penjadwalan")
	data := JadwalEntity{ID: id}
	if err := db.Where("deleted_at IS NULL").First(&data).Error; err != nil {
		logrus.Error("[jadwal.repository]: ", err.Error())
		return nil
	}
	return &data
}

func (r *PenjadwalanMySqlRepository) Create(data *JadwalModel) *JadwalEntity {
	db := r.driver.DB.Table("penjadwalan")
	if err := db.Create(&data).Error; err != nil {
		logrus.Error("[jadwal.repository]: ", err.Error())
		return nil
	}
	entity := &JadwalEntity{
		ID:      data.ID,
		Asisten: data.Asisten,
		Hari:    data.Hari,
		Jam:     data.Jam,
	}
	return entity
}

func (r *PenjadwalanMySqlRepository) Update(data *JadwalModel) *JadwalEntity {
	db := r.driver.DB.Table("penjadwalan")
	search := &JadwalModel{ID: data.ID}
	if err := db.Find(&search).Error; err != nil {
		logrus.Error("[jadwal.repository]: ", err.Error())
		return nil
	}
	model := func(old, new *JadwalModel) *JadwalModel {
		fields := []string{"Asisten", "Hari", "Jam"}
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
		old.Asisten = result["Asisten"].(string)
		old.Hari = result["Hari"].(string)
		old.Jam = result["Jam"].(string)
		return old
	}(search, data)
	if err := db.Save(&model).Error; err != nil {
		logrus.Error("[jadwal.repository]: ", err.Error())
		return nil
	}
	entity := &JadwalEntity{
		ID:      model.ID,
		Asisten: model.Asisten,
		Hari:    model.Hari,
		Jam:     model.Jam,
	}
	return entity
}

func (r *PenjadwalanMySqlRepository) Delete(id int) bool {
	db := r.driver.DB.Table("penjadwalan")
	if err := db.Delete(&JadwalModel{}, id).Error; err != nil {
		logrus.Error("[jadwal.repository]: ", err.Error())
		return false
	}
	return true
}
