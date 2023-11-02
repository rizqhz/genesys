package repository

import (
	"fmt"
	"net/url"

	mysql "github.com/rizghz/genesys/infrastructure/database/MySql"
	"github.com/rizghz/genesys/internal/helpers"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type MahasiswaMySqlRepository struct {
	driver *mysql.MySqlDriver
}

func NewMahasiswaMySqlRepository(drv *mysql.MySqlDriver) MahasiswaRepository {
	return &MahasiswaMySqlRepository{
		driver: drv,
	}
}

func (repo *MahasiswaMySqlRepository) Get(query url.Values) []Entity {
	db := repo.driver.DB.Table("mahasiswa")
	db = helpers.QuerySorting(db, query)
	db = helpers.QueryPagination(db, query)
	db = helpers.QueryFiltering(db, query)
	data := make([]Entity, 0)
	if err := db.Find(&data).Error; err != nil {
		log.Error("[mahasiswa.repository]: ", err.Error())
		return nil
	}
	return data
}

func (repo *MahasiswaMySqlRepository) Find(npm string) *Entity {
	db := repo.driver.DB.Table("mahasiswa")
	data := Entity{}
	condition := fmt.Sprintf("npm = '%s'", npm)
	if err := db.First(&data, condition).Error; err != nil {
		log.Error("[mahasiswa.repository]: ", err.Error())
		return nil
	}
	return &data
}

func (repo *MahasiswaMySqlRepository) Create(data *Model) *Entity {
	db := repo.driver.DB.Table("mahasiswa")
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&data).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Error("[mahasiswa.repository]: ", err.Error())
		return nil
	}
	return &Entity{
		NPM:   data.NPM,
		Nama:  data.Nama,
		Kelas: data.Kelas,
	}
}

func (repo *MahasiswaMySqlRepository) Update(npm string, data *Model) *Entity {
	db := repo.driver.DB.Table("mahasiswa")
	query := fmt.Sprintf("UPDATE mahasiswa SET %s = '%s', %s = '%s', %s = '%s' WHERE %s = '%s'", []any{
		"npm", data.NPM, "nama", data.Nama, "kelas", data.Kelas,
		"npm", npm,
	}...)
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec(query).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Error("[mahasiswa.repository]: ", err.Error())
		return nil
	}
	return &Entity{
		NPM:   data.NPM,
		Nama:  data.Nama,
		Kelas: data.Kelas,
	}
}

func (repo *MahasiswaMySqlRepository) Delete(npm string) bool {
	db := repo.driver.DB.Table("mahasiswa")
	query := fmt.Sprintf("DELETE FROM mahasiswa WHERE %s = '%s'", []any{
		"npm", npm,
	}...)
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec(query).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Error("[mahasiswa.repository]: ", err.Error())
		return false
	}
	return true
}
