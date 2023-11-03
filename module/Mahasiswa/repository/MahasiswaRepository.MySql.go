package repository

import (
	"fmt"
	"net/url"

	"github.com/fatih/structs"
	mysql "github.com/rizghz/genesys/infrastructure/database/MySql"
	"github.com/rizghz/genesys/internal/helpers"
	"github.com/sirupsen/logrus"
)

type MahasiswaMySqlRepository struct {
	driver *mysql.MySqlDriver
}

func NewMahasiswaMySqlRepository(drv *mysql.MySqlDriver) MahasiswaRepository {
	return &MahasiswaMySqlRepository{
		driver: drv,
	}
}

func (repo *MahasiswaMySqlRepository) Get(query url.Values) []MahasiswaEntity {
	db := repo.driver.DB.Table("mahasiswa")
	db = helpers.QuerySorting(db, query)
	db = helpers.QueryPagination(db, query)
	db = helpers.QueryFiltering(db, query)
	data := make([]MahasiswaEntity, 0)
	if err := db.Where("deleted_at IS NULL").Find(&data).Error; err != nil {
		logrus.Error("[mahasiswa.repository]: ", err.Error())
		return nil
	}
	return data
}

func (repo *MahasiswaMySqlRepository) Find(npm string) *MahasiswaEntity {
	db := repo.driver.DB.Table("mahasiswa")
	data := &MahasiswaEntity{}
	cond := fmt.Sprintf("npm = '%s'", npm)
	if err := db.Where("deleted_at IS NULL").First(&data, cond).Error; err != nil {
		logrus.Error("[mahasiswa.repository]: ", err.Error())
		return nil
	}
	return data
}

func (repo *MahasiswaMySqlRepository) Create(data *MahasiswaModel) *MahasiswaEntity {
	db := repo.driver.DB.Table("mahasiswa")
	if err := db.Create(&data).Error; err != nil {
		logrus.Error("[mahasiswa.repository]: ", err.Error())
		return nil
	}
	return &MahasiswaEntity{
		NPM:   data.NPM,
		Nama:  data.Nama,
		Kelas: data.Kelas,
	}
}

func (repo *MahasiswaMySqlRepository) Update(npm string, data *MahasiswaModel) *MahasiswaEntity {
	db := repo.driver.DB.Table("mahasiswa")
	search := &MahasiswaModel{NPM: data.NPM}
	cond := fmt.Sprintf("npm = '%s'", npm)
	if err := db.Find(&search, cond).Error; err != nil {
		logrus.Error("[mahasiswa.repository]: ", err.Error())
		return nil
	}
	model := func(old, new *MahasiswaModel) *MahasiswaModel {
		fields := []string{"NPM", "Nama", "Kelas"}
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
		old.NPM = result["NPM"].(string)
		old.Nama = result["Nama"].(string)
		old.Kelas = result["Kelas"].(string)
		return old
	}(search, data)
	if err := db.Where(cond).Save(&model).Error; err != nil {
		logrus.Error("[mahasiswa.repository]: ", err.Error())
		return nil
	}
	return &MahasiswaEntity{
		NPM:   model.NPM,
		Nama:  model.Nama,
		Kelas: model.Kelas,
	}
}

func (repo *MahasiswaMySqlRepository) Delete(npm string) bool {
	db := repo.driver.DB.Table("mahasiswa")
	cond := fmt.Sprintf("npm = '%s'", npm)
	if err := db.Delete(&MahasiswaModel{}, cond).Error; err != nil {
		logrus.Error("[mahasiswa.repository]: ", err.Error())
		return false
	}
	return true
}
