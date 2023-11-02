package jwt

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
	log "github.com/sirupsen/logrus"
)

type JwtSecret map[string]string

func NewJwtSecret() (secret JwtSecret) {
	// create JwtSecret object
	secret = make(JwtSecret)

	// check JWT_ACCESS_KEY environment variable
	if value, found := os.LookupEnv("JWT_ACCESS_KEY"); !found {
		log.Fatal("env: JWT_ACCESS_KEY not found")
	} else {
		secret["JWT_ACCESS_KEY"] = value
	}

	// check JWT_REFRESH_KEY environment variable
	if value, found := os.LookupEnv("JWT_REFRESH_KEY"); !found {
		log.Fatal("env: JWT_REFRESH_KEY not found")
	} else {
		secret["JWT_REFRESH_KEY"] = value
	}

	return
}

type JwtKey struct {
	AccessKey  string
	RefreshKey string
}

func NewJwtKey() *JwtKey {
	secret := NewJwtSecret()
	return &JwtKey{
		AccessKey:  secret["JWT_ACCESS_KEY"],
		RefreshKey: secret["JWT_REFRESH_KEY"],
	}
}
