package migration

import (
	mysql "github.com/rizghz/genesys/infrastructure/database/MySql"
)

type MySqlMigrator struct {
	driver *mysql.MySqlDriver
}

func NewMySqlMigrator(drv *mysql.MySqlDriver) Migrator {
	return &MySqlMigrator{
		driver: drv,
	}
}

func (m *MySqlMigrator) CreateTable(table ...Table) {
	db := m.driver.DB
	for _, t := range table {
		db.AutoMigrate(t)
	}
}

func (m *MySqlMigrator) DropTable(table ...Table) {
	db := m.driver.DB
	for _, t := range table {
		db.Migrator().DropTable(t)
	}
}
