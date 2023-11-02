package cloudinary

import (
	"os"

	log "github.com/sirupsen/logrus"
)

type CloudinaryEnv map[string]any

func NewCloudinaryEnv() (env CloudinaryEnv) {
	// create CloudinaryEnv object
	env = make(CloudinaryEnv)

	// check CLOUDINARY_CLOUD_NAME environment variable
	if value, found := os.LookupEnv("CLOUDINARY_CLOUD_NAME"); !found {
		log.Fatal("env: CLOUDINARY_CLOUD_NAME not found")
	} else {
		env["CLOUDINARY_CLOUD_NAME"] = value
	}

	// check CLOUDINARY_API_KEY environment variable
	if value, found := os.LookupEnv("CLOUDINARY_API_KEY"); !found {
		log.Fatal("env: CLOUDINARY_API_KEY not found")
	} else {
		env["CLOUDINARY_API_KEY"] = value
	}

	// check CLOUDINARY_API_SECRET environment variable
	if value, found := os.LookupEnv("CLOUDINARY_API_SECRET"); !found {
		log.Fatal("env: CLOUDINARY_API_SECRET not found")
	} else {
		env["CLOUDINARY_API_SECRET"] = value
	}

	return
}
