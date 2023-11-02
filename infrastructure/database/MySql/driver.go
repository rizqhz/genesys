package mysql

import (
	"os"

	"github.com/rizghz/genesys/infrastructure/database"
	drv "gorm.io/driver/mysql"
	orm "gorm.io/gorm"
)

type MySqlDriver struct {
	DB *orm.DB
}

func NewMySqlDriver() *MySqlDriver {
	env := database.NewDatabaseEnv()
	str := NewMySqlConfig(env).ConnectStr()
	db, err := orm.Open(drv.Open(str), &orm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		TranslateError:         true,
	})
	if err != nil {
		os.Exit(1)
	}
	return &MySqlDriver{db}
}
