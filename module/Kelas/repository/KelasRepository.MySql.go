package repository

import (
	"fmt"
	"net/url"

	mysql "github.com/rizghz/genesys/infrastructure/database/MySql"
	"github.com/rizghz/genesys/internal/helpers"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type KelasMySqlRepository struct {
	driver *mysql.MySqlDriver
}

func NewKelasMySqlRepository(drv *mysql.MySqlDriver) KelasRepository {
	return &KelasMySqlRepository{
		driver: drv,
	}
}

func (repo *KelasMySqlRepository) Get(query url.Values) []Entity {
	db := repo.driver.DB.Table("kelas")
	db = helpers.QuerySorting(db, query)
	db = helpers.QueryPagination(db, query)
	db = helpers.QueryFiltering(db, query)
	data := make([]Entity, 0)
	if err := db.Find(&data).Error; err != nil {
		log.Error("[kelas.repository]: ", err.Error())
		return nil
	}
	return data
}

func (repo *KelasMySqlRepository) Find(kode string) *Entity {
	db := repo.driver.DB.Table("kelas")
	data := Entity{}
	condition := fmt.Sprintf("kode = '%s'", kode)
	if err := db.First(&data, condition).Error; err != nil {
		log.Error("[kelas.repository]: ", err.Error())
		return nil
	}
	return &data
}

func (repo *KelasMySqlRepository) Create(data *Model) *Entity {
	db := repo.driver.DB.Table("kelas")
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&data).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Error("[kelas.repository]: ", err.Error())
		return nil
	}
	return &Entity{
		Kode:    data.Kode,
		Nama:    data.Nama,
		Jurusan: data.Jurusan,
		Grade:   data.Grade,
		Tahun:   data.Tahun,
	}
}

func (repo *KelasMySqlRepository) Update(kode string, data *Model) *Entity {
	db := repo.driver.DB.Table("kelas")
	query := fmt.Sprintf("UPDATE kelas SET %s = '%s', %s = '%s', %s = '%s', %s = '%s', %s = %d WHERE %s = '%s'", []any{
		"kode", data.Kode, "nama", data.Nama, "jurusan", data.Jurusan, "grade", data.Grade, "tahun", data.Tahun,
		"kode", kode,
	}...)
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec(query).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Error("[kelas.repository]: ", err.Error())
		return nil
	}
	return &Entity{
		Kode:    data.Kode,
		Nama:    data.Nama,
		Jurusan: data.Jurusan,
		Grade:   data.Grade,
		Tahun:   data.Tahun,
	}
}

func (repo *KelasMySqlRepository) Delete(kode string) bool {
	db := repo.driver.DB.Table("kelas")
	query := fmt.Sprintf("DELETE FROM ruangan WHERE %s = '%s'", []any{
		"kode", kode,
	}...)
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec(query).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Error("[kelas.repository]: ", err.Error())
		return false
	}
	return true
}
