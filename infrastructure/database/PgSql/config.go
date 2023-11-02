package pgsql

import (
	"fmt"

	"github.com/rizghz/genesys/infrastructure/database"
)

type PgSqlConfig struct {
	host string
	port int
	user string
	pass string
	name string
}

func (c *PgSqlConfig) ConnectStr() string {
	format := "host=%s user=%s password=%s dbname=%s port=%d"
	return fmt.Sprintf(format, c.host, c.user, c.pass, c.name, c.port)
}

func NewPgSqlConfig(env database.DatabaseEnv) *PgSqlConfig {
	return &PgSqlConfig{
		host: env["DB_HOST"].(string),
		port: env["DB_PORT"].(int),
		user: env["DB_USER"].(string),
		pass: env["DB_PASS"].(string),
		name: env["DB_NAME"].(string),
	}
}
