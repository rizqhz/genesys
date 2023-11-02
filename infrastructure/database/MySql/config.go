package mysql

import (
	"fmt"

	"github.com/rizghz/genesys/infrastructure/database"
)

type MySqlConfig struct {
	host string
	port int
	user string
	pass string
	name string
}

func (c *MySqlConfig) ConnectStr() string {
	format := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	return fmt.Sprintf(format, c.user, c.pass, c.host, c.port, c.name)
}

func NewMySqlConfig(env database.DatabaseEnv) *MySqlConfig {
	return &MySqlConfig{
		host: env["DB_HOST"].(string),
		port: env["DB_PORT"].(int),
		user: env["DB_USER"].(string),
		pass: env["DB_PASS"].(string),
		name: env["DB_NAME"].(string),
	}
}
