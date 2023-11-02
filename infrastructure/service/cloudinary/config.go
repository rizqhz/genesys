package cloudinary

type CloudinaryConfig struct {
	cloud  string
	key    string
	secret string
}

func NewCloudinaryConfig() *CloudinaryConfig {
	env := NewCloudinaryEnv()
	return &CloudinaryConfig{
		cloud:  env["CLOUDINARY_CLOUD_NAME"].(string),
		key:    env["CLOUDINARY_API_KEY"].(string),
		secret: env["CLOUDINARY_API_SECRET"].(string),
	}
}
