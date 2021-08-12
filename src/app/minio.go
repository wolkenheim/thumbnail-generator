package app

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)


func MinioClientFactory() (*minio.Client, error) {

	endpoint := viper.GetString("minio.url")
	accessKeyID := viper.GetString("minio.accessKeyID")
	secretAccessKey := viper.GetString("minio.secretAccessKey")
	useSSL := viper.GetBool("minio.useSSL")

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		zap.S().Fatal(err)
		return nil, err
	}

	return minioClient, nil
}
