package repository

import (
	"fmt"
	"net/url"

	mysql "github.com/rizghz/genesys/infrastructure/database/MySql"
	"github.com/rizghz/genesys/internal/helpers"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AsistenMySqlRepository struct {
	driver *mysql.MySqlDriver
}

func NewAsistenMySqlRepository(driver *mysql.MySqlDriver) AsistenRepository {
	return &AsistenMySqlRepository{
		driver: driver,
	}
}

func (repo *AsistenMySqlRepository) Get(query url.Values) []Entity {
	db := repo.driver.DB.Table("asisten")
	db = helpers.QuerySorting(db, query)
	db = helpers.QueryPagination(db, query)
	db = helpers.QueryFiltering(db, query)
	data := make([]Entity, 0)
	if err := db.Find(&data).Error; err != nil {
		log.Error("[asisten.repository]: ", err.Error())
		return nil
	}
	return data
}

func (repo *AsistenMySqlRepository) Find(nias string) *Entity {
	db := repo.driver.DB.Table("asisten")
	data := Entity{}
	condition := fmt.Sprintf("nias = '%s'", nias)
	if err := db.First(&data, condition).Error; err != nil {
		log.Error("[Asisten.repository]: ", err.Error())
		return nil
	}
	return &data
}

func (repo *AsistenMySqlRepository) Create(data *Model) *Entity {
	db := repo.driver.DB.Table("asisten")
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&data).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Error("[asisten.repository]: ", err.Error())
		return nil
	}
	return &Entity{
		NIAS:    data.NIAS,
		Nama:    data.Nama,
		Jabatan: data.Jabatan,
	}
}

func (repo *AsistenMySqlRepository) Update(nias string, data *Model) *Entity {
	db := repo.driver.DB.Table("asisten")
	query := fmt.Sprintf("UPDATE asisten SET %s = '%s', %s = '%s', %s = '%s' WHERE %s = '%s'", []any{
		"nias", data.NIAS, "nama", data.Nama, "jabatan", data.Jabatan,
		"nias", nias,
	}...)
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec(query).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Error("[asisten.repository]: ", err.Error())
		return nil
	}
	return &Entity{
		NIAS:    data.NIAS,
		Nama:    data.Nama,
		Jabatan: data.Jabatan,
	}
}

func (repo *AsistenMySqlRepository) Delete(nias string) bool {
	db := repo.driver.DB.Table("asisten")
	query := fmt.Sprintf("DELETE FROM asisten WHERE %s = '%s'", []any{
		"nias", nias,
	}...)
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec(query).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Error("[asisten.repository]: ", err.Error())
		return false
	}
	return true
}
