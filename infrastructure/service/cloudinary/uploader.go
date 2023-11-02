package cloudinary

import (
	"context"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/admin"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/sirupsen/logrus"
)

type Uploader interface {
	Upload(file multipart.File, name string) string
}

type ImageUploader struct {
	serv *cloudinary.Cloudinary
}

func NewImageUploader() Uploader {
	config := NewCloudinaryConfig()
	var (
		cloud  = config.cloud
		key    = config.key
		secret = config.secret
	)
	serv, err := cloudinary.NewFromParams(cloud, key, secret)
	if err != nil {
		logrus.Fatal("[cloudinary]: ", err.Error())
	}
	return &ImageUploader{
		serv: serv,
	}
}

func (u *ImageUploader) Upload(file multipart.File, name string) string {
	api := u.serv.Upload
	ctx := context.Background()
	res, err := api.Upload(ctx, file, uploader.UploadParams{
		PublicID: name,
	})
	if err != nil {
		logrus.Error("[upload]: ", err.Error())
	}
	return res.SecureURL
}

func (u *ImageUploader) Delete() bool {
	api := u.serv.Admin
	ctx := context.Background()
	api.DeleteAssets(ctx, admin.DeleteAssetsParams{
		PublicIDs: []string{},
	})
	return false
}
