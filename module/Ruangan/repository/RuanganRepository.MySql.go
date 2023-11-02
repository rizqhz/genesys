package repository

import (
	"fmt"
	"net/url"

	mysql "github.com/rizghz/genesys/infrastructure/database/MySql"
	"github.com/rizghz/genesys/internal/helpers"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type RuanganMySqlRepository struct {
	driver *mysql.MySqlDriver
}

func NewRuanganMySqlRepository(driver *mysql.MySqlDriver) RuanganRepository {
	return &RuanganMySqlRepository{
		driver: driver,
	}
}

func (repo *RuanganMySqlRepository) Get(query url.Values) []Entity {
	db := repo.driver.DB.Table("ruangan")
	db = helpers.QuerySorting(db, query)
	db = helpers.QueryPagination(db, query)
	db = helpers.QueryFiltering(db, query)
	data := make([]Entity, 0)
	if err := db.Find(&data).Error; err != nil {
		log.Error("[ruangan.repository]: ", err.Error())
		return nil
	}
	return data
}

func (repo *RuanganMySqlRepository) Find(kode string) *Entity {
	db := repo.driver.DB.Table("ruangan")
	data := Entity{}
	condition := fmt.Sprintf("kode = '%s'", kode)
	if err := db.First(&data, condition).Error; err != nil {
		log.Error("[ruangan.repository]: ", err.Error())
		return nil
	}
	return &data
}

func (repo *RuanganMySqlRepository) Create(data *Model) *Entity {
	db := repo.driver.DB.Table("ruangan")
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&data).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Error("[ruangan.repository]: ", err.Error())
		return nil
	}
	return &Entity{
		Kode: data.Kode,
		Nama: data.Nama,
	}
}

func (repo *RuanganMySqlRepository) Update(kode string, data *Model) *Entity {
	db := repo.driver.DB.Table("ruangan")
	query := fmt.Sprintf("UPDATE ruangan SET %s = '%s', %s = '%s' WHERE %s = '%s'", []any{
		"kode", data.Kode, "nama", data.Nama,
		"kode", kode,
	}...)
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec(query).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Error("[ruangan.repository]: ", err.Error())
		return nil
	}
	return &Entity{
		Kode: data.Kode,
		Nama: data.Nama,
	}
}

func (repo *RuanganMySqlRepository) Delete(kode string) bool {
	db := repo.driver.DB.Table("ruangan")
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
		log.Error("[ruangan.repository]: ", err.Error())
		return false
	}
	return true
}
