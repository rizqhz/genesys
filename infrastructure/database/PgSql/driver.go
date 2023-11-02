package pgsql

import (
	"os"

	"github.com/rizghz/genesys/infrastructure/database"
	drv "gorm.io/driver/postgres"
	orm "gorm.io/gorm"
)

type PgSqlDriver struct {
	DB *orm.DB
}

func NewPgSqlDriver() *PgSqlDriver {
	env := database.NewDatabaseEnv()
	str := NewPgSqlConfig(env).ConnectStr()
	db, err := orm.Open(drv.Open(str), &orm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		TranslateError:         true,
	})
	if err != nil {
		os.Exit(1)
	}
	return &PgSqlDriver{db}
}
