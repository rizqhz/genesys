package repository

import (
	"fmt"
	"net/url"

	"github.com/fatih/structs"
	mysql "github.com/rizghz/genesys/infrastructure/database/MySql"
	"github.com/rizghz/genesys/internal/helpers"
	"github.com/sirupsen/logrus"
)

type AsistenMySqlRepository struct {
	driver *mysql.MySqlDriver
}

func NewAsistenMySqlRepository(driver *mysql.MySqlDriver) AsistenRepository {
	return &AsistenMySqlRepository{
		driver: driver,
	}
}

func (repo *AsistenMySqlRepository) Get(query url.Values) []AsistenEntity {
	db := repo.driver.DB.Table("asisten")
	db = helpers.QuerySorting(db, query)
	db = helpers.QueryPagination(db, query)
	db = helpers.QueryFiltering(db, query)
	data := make([]AsistenEntity, 0)
	if err := db.Where("deleted_at IS NULL").Find(&data).Error; err != nil {
		logrus.Error("[asisten.repository]: ", err.Error())
		return nil
	}
	return data
}

func (repo *AsistenMySqlRepository) Find(nias string) *AsistenEntity {
	db := repo.driver.DB.Table("asisten")
	data := &AsistenEntity{NIAS: nias}
	cond := fmt.Sprintf("nias = '%s'", nias)
	if err := db.Where("deleted_at IS NULL").First(&data, cond).Error; err != nil {
		logrus.Error("[asisten.repository]: ", err.Error())
		return nil
	}
	return data
}

func (repo *AsistenMySqlRepository) Create(data *AsistenModel) *AsistenEntity {
	db := repo.driver.DB.Table("asisten")
	if err := db.Create(&data).Error; err != nil {
		logrus.Error("[asisten.repository]: ", err.Error())
		return nil
	}
	return &AsistenEntity{
		NIAS:    data.NIAS,
		Nama:    data.Nama,
		Jabatan: data.Jabatan,
	}
}

func (repo *AsistenMySqlRepository) Update(nias string, data *AsistenModel) *AsistenEntity {
	db := repo.driver.DB.Table("asisten")
	search := &AsistenModel{NIAS: data.NIAS}
	cond := fmt.Sprintf("nias = '%s'", nias)
	if err := db.Find(&search, cond).Error; err != nil {
		logrus.Error("[asisten.repository]: ", err.Error())
		return nil
	}
	model := func(old, new *AsistenModel) *AsistenModel {
		fields := []string{"NIAS", "Nama", "Jabatan"}
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
		old.NIAS = result["NIAS"].(string)
		old.Nama = result["Nama"].(string)
		old.Jabatan = result["Jabatan"].(string)
		return old
	}(search, data)
	if err := db.Where(cond).Save(&model).Error; err != nil {
		logrus.Error("[asisten.repository]: ", err.Error())
		return nil
	}
	return &AsistenEntity{
		NIAS:    model.NIAS,
		Nama:    model.Nama,
		Jabatan: model.Jabatan,
	}
}

func (repo *AsistenMySqlRepository) Delete(nias string) bool {
	db := repo.driver.DB.Table("asisten")
	cond := fmt.Sprintf("nias = '%s'", nias)
	if err := db.Delete(&AsistenModel{}, cond).Error; err != nil {
		logrus.Error("[asisten.repository]: ", err.Error())
		return false
	}
	return true
}
