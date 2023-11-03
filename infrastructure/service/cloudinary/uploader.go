package cloudinary

import (
	"context"
	"mime/multipart"
	"strings"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/admin"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/sirupsen/logrus"
)

type Uploader interface {
	Upload(file multipart.File, name string) string
	GetImageId(url string) string
	Delete(url string) bool
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
		logrus.Error("[upload.image]: ", err.Error())
	}
	return res.SecureURL
}

func (u *ImageUploader) GetImageId(url string) string {
	var buffer []string
	buffer = strings.Split(url, "/")
	buffer = strings.Split(buffer[len(buffer)-1], ".")
	return buffer[0]
}

func (u *ImageUploader) Delete(url string) bool {
	api := u.serv.Admin
	ctx := context.Background()
	id := u.GetImageId(url)
	_, err := api.DeleteAssets(ctx, admin.DeleteAssetsParams{
		PublicIDs: []string{id},
	})
	if err != nil {
		logrus.Error("[delete.image]: ", err.Error())
		return false
	}
	return true
}
