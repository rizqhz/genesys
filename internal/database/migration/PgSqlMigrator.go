package migration

import pgsql "github.com/rizghz/genesys/infrastructure/database/PgSql"

type PgSqlMigrator struct {
	driver *pgsql.PgSqlDriver
}

func NewPgSqlMigrator(drv *pgsql.PgSqlDriver) Migrator {
	return &PgSqlMigrator{
		driver: drv,
	}
}

func (m *PgSqlMigrator) CreateTable(table ...Table) {
	db := m.driver.DB
	db.AutoMigrate(table)
}

func (m *PgSqlMigrator) DropTable(table ...Table) {
	db := m.driver.DB
	for _, t := range table {
		db.Migrator().DropTable(t)
	}
}
