package database

import (
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
	log "github.com/sirupsen/logrus"
)

type DatabaseEnv map[string]any

func NewDatabaseEnv() (env DatabaseEnv) {
	// create DatabaseEnv object
	env = make(DatabaseEnv)

	// check DB_HOST environment variable
	if value, found := os.LookupEnv("DB_HOST"); !found {
		log.Fatal("env: DB_HOST not found")
	} else {
		env["DB_HOST"] = value
	}

	// check DB_PORT environment variable
	if value, found := os.LookupEnv("DB_PORT"); !found {
		log.Fatal("env: DB_PORT not found")
	} else {
		port, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal("env: DB_PORT not valid")
		}
		env["DB_PORT"] = port
	}

	// check DB_USER environment variable
	if value, found := os.LookupEnv("DB_USER"); !found {
		log.Fatal("env: DB_USER not found")
	} else {
		env["DB_USER"] = value
	}

	// check DB_PASS environment variable
	if value, found := os.LookupEnv("DB_PASS"); !found {
		log.Fatal("env: DB_PASS not found")
	} else {
		env["DB_PASS"] = value
	}

	// check DB_NAME environment variable
	if value, found := os.LookupEnv("DB_NAME"); !found {
		log.Fatal("env: DB_NAME not found")
	} else {
		env["DB_NAME"] = value
	}

	return
}
