package cloudinary

import (
	"context"
	"errors"
	"mime/multipart"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type MyCloudinary struct {
	Cloud *cloudinary.Cloudinary
}

func NewMyCloudinary(cloudName string, cloudKey string, cloudSecret string, uploadFolder string) (*MyCloudinary, error) {
	// cldSecret := viper.GetString("CLOUDINARY_SECRET_KEY")
	// cldName := viper.GetString("CLOUDINARY_CLOUD_NAME")
	// cldKey := viper.GetString("CLOUDINARY_API_KEY")

	cld, err := cloudinary.NewFromParams(cloudName, cloudKey, cloudSecret)
	if err != nil {
		return nil, err
	}

	return &MyCloudinary{
		Cloud: cld,
	}, nil
}

func (m *MyCloudinary) UploadToCloudinary(file multipart.File, filePath string) (string, error) {
	if m.Cloud == nil {
		return "", errors.New("cloudinary connection is not initialized")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uploadParams := uploader.UploadParams{
		PublicID:     filePath,
		ResourceType: "image",
		Folder:       "HaiwanGram",
	}

	result, err := m.Cloud.Upload.Upload(ctx, file, uploadParams)
	if err != nil {
		return "", err
	}

	imageUrl := result.SecureURL
	return imageUrl, nil
}
