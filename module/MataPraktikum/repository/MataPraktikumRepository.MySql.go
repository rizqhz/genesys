package repository

import (
	"fmt"
	"net/url"

	mysql "github.com/rizghz/genesys/infrastructure/database/MySql"
	"github.com/rizghz/genesys/internal/helpers"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type MataPraktikumMySqlRepository struct {
	driver *mysql.MySqlDriver
}

func NewMataPraktikumMySqlRepository(driver *mysql.MySqlDriver) MataPraktikumRepository {
	return &MataPraktikumMySqlRepository{
		driver: driver,
	}
}

func (repo *MataPraktikumMySqlRepository) Get(query url.Values) []Entity {
	db := repo.driver.DB.Table("mata_praktikum")
	db = helpers.QuerySorting(db, query)
	db = helpers.QueryPagination(db, query)
	db = helpers.QueryFiltering(db, query)
	data := make([]Entity, 0)
	if err := db.Find(&data).Error; err != nil {
		log.Error("[mata_praktikum.repository]: ", err.Error())
		return nil
	}
	return data
}

func (repo *MataPraktikumMySqlRepository) Find(kode string) *Entity {
	db := repo.driver.DB.Table("mata_praktikum")
	data := Entity{}
	condition := fmt.Sprintf("kode = '%s'", kode)
	if err := db.First(&data, condition).Error; err != nil {
		log.Error("[mata_praktikum.repository]: ", err.Error())
		return nil
	}
	return &data
}

func (repo *MataPraktikumMySqlRepository) Create(data *Model) *Entity {
	db := repo.driver.DB.Table("mata_praktikum")
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&data).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Error("[mata_praktikum.repository]: ", err.Error())
		return nil
	}
	return &Entity{
		Kode: data.Kode,
		Nama: data.Nama,
	}
}

func (repo *MataPraktikumMySqlRepository) Update(kode string, data *Model) *Entity {
	db := repo.driver.DB.Table("mata_praktikum")
	query := fmt.Sprintf("UPDATE mata_praktikum SET %s = '%s', %s = '%s' WHERE %s = '%s'", []any{
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
		log.Error("[mata_praktikum.repository]: ", err.Error())
		return nil
	}
	return &Entity{
		Kode: data.Kode,
		Nama: data.Nama,
	}
}

func (repo *MataPraktikumMySqlRepository) Delete(kode string) bool {
	db := repo.driver.DB.Table("mata_praktikum")
	query := fmt.Sprintf("DELETE FROM mata_praktikum WHERE %s = '%s'", []any{
		"kode", kode,
	}...)
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec(query).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Error("[mata_praktikum.repository]: ", err.Error())
		return false
	}
	return true
}
