package repository

import (
	"fmt"
	"net/url"

	"github.com/fatih/structs"
	mysql "github.com/rizghz/genesys/infrastructure/database/MySql"
	"github.com/rizghz/genesys/internal/helpers"
	"github.com/sirupsen/logrus"
)

type KelasMySqlRepository struct {
	driver *mysql.MySqlDriver
}

func NewKelasMySqlRepository(drv *mysql.MySqlDriver) KelasRepository {
	return &KelasMySqlRepository{
		driver: drv,
	}
}

func (repo *KelasMySqlRepository) Get(query url.Values) []KelasEntity {
	db := repo.driver.DB.Table("kelas")
	db = helpers.QuerySorting(db, query)
	db = helpers.QueryPagination(db, query)
	db = helpers.QueryFiltering(db, query)
	data := make([]KelasEntity, 0)
	if err := db.Where("deleted_at IS NULL").Find(&data).Error; err != nil {
		logrus.Error("[kelas.repository]: ", err.Error())
		return nil
	}
	return data
}

func (repo *KelasMySqlRepository) Find(kode string) *KelasEntity {
	db := repo.driver.DB.Table("kelas")
	data := &KelasEntity{Kode: kode}
	cond := fmt.Sprintf("kode = '%s'", kode)
	if err := db.Where("deleted_at IS NULL").First(&data, cond).Error; err != nil {
		logrus.Error("[kelas.repository]: ", err.Error())
		return nil
	}
	return data
}

func (repo *KelasMySqlRepository) Create(data *KelasModel) *KelasEntity {
	db := repo.driver.DB.Table("kelas")
	if err := db.Create(&data).Error; err != nil {
		logrus.Error("[kelas.repository]: ", err.Error())
		return nil
	}
	return &KelasEntity{
		Kode:    data.Kode,
		Nama:    data.Nama,
		Jurusan: data.Jurusan,
		Grade:   data.Grade,
		Tahun:   data.Tahun,
	}
}

func (repo *KelasMySqlRepository) Update(kode string, data *KelasModel) *KelasEntity {
	db := repo.driver.DB.Table("kelas")
	search := &KelasModel{Kode: data.Kode}
	cond := fmt.Sprintf("kode = '%s'", kode)
	if err := db.Find(&search, cond).Error; err != nil {
		logrus.Error("[kelas.repository]: ", err.Error())
		return nil
	}
	model := func(old, new *KelasModel) *KelasModel {
		fields := []string{"Kode", "Nama", "Jurusan", "Grade", "Tahun"}
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
		old.Jurusan = result["Jurusan"].(string)
		old.Grade = result["Grade"].(string)
		old.Tahun = result["Tahun"].(int)
		return old
	}(search, data)
	if err := db.Where(cond).Save(&model).Error; err != nil {
		logrus.Error("[kelas.repository]: ", err.Error())
		return nil
	}
	return &KelasEntity{
		Kode:    model.Kode,
		Nama:    model.Nama,
		Jurusan: model.Jurusan,
		Grade:   model.Grade,
		Tahun:   model.Tahun,
	}
}

func (repo *KelasMySqlRepository) Delete(kode string) bool {
	db := repo.driver.DB.Table("kelas")
	cond := fmt.Sprintf("kode = '%s'", kode)
	if err := db.Delete(&KelasModel{}, cond).Error; err != nil {
		logrus.Error("[kelas.repository]: ", err.Error())
		return false
	}
	return true
}
