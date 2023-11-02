package repository

import (
	"net/url"

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

func (r *PenjadwalanMySqlRepository) Get(query url.Values) []Entity {
	db := r.driver.DB.Table("penjadwalans")
	helpers.QuerySorting(db, query)
	helpers.QueryPagination(db, query)
	helpers.QueryFiltering(db, query)
	data := make([]Entity, 0)
	if err := db.Find(&data).Error; err != nil {
		logrus.Error("[penjadwalan.repository]: ", err.Error())
		return nil
	}
	return data
}

func (r *PenjadwalanMySqlRepository) Find(id int) *Entity {
	db := r.driver.DB.Table("penjadwalans")
	data := Entity{ID: id}
	if err := db.First(&data).Error; err != nil {
		logrus.Error("[penjadwalan.repository]: ", err.Error())
		return nil
	}
	return &data
}

func (r *PenjadwalanMySqlRepository) Create(data *Model) *Entity {
	db := r.driver.DB.Table("penjadwalans")
	if err := db.Create(&data).Error; err != nil {
		logrus.Error("[penjadwalan.repository]: ", err.Error())
		return nil
	}
	entity := &Entity{
		ID:      data.ID,
		Asisten: data.Asisten,
		Hari:    data.Hari,
		Jam:     data.Jam,
	}
	return entity
}

func (r *PenjadwalanMySqlRepository) Update(data *Model) *Entity {
	db := r.driver.DB.Table("penjadwalans")
	if err := db.Save(&data).Error; err != nil {
		logrus.Error("[penjadwalan.repository]: ", err.Error())
		return nil
	}
	entity := &Entity{
		ID:      data.ID,
		Asisten: data.Asisten,
		Hari:    data.Hari,
		Jam:     data.Jam,
	}
	return entity
}

func (r *PenjadwalanMySqlRepository) Delete(id int) bool {
	db := r.driver.DB.Table("penjadwalans")
	if err := db.Delete(&Model{ID: id}).Error; err != nil {
		logrus.Error("[penjadwalan.repository]: ", err.Error())
		return false
	}
	return true
}
