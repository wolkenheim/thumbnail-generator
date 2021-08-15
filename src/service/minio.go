package service

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/spf13/viper"
)

type ImageService interface {
	Download(ctx context.Context, minioName string, localFilePath string) error
	Upload(ctx context.Context, minioName string, localFilePath string) error
	GetThumbnailPath(minioName string) string
	GetOriginalPath(minioName string) string
}

type MinioService struct{
	client *minio.Client
}

func(m *MinioService) Download(ctx context.Context, minioName string, localFilePath string) error {
	err := m.client.FGetObject(
		ctx,
		viper.GetString("minio.bucket"),
		minioName,
		localFilePath,
		minio.GetObjectOptions{},
	)
	if err != nil {
		return err
	}

	return nil
}

func(m *MinioService) Upload(ctx context.Context, minioName string, localFilePath string) error {

	_, err := m.client.FPutObject(
		ctx,
		viper.GetString("minio.bucket"),
		minioName,
		localFilePath,
		minio.PutObjectOptions{},
	)
	if err != nil {
		return err
	}

	return nil
}

func(m *MinioService) GetThumbnailPath(minioName string) string {
	prefix := "thumbnails/"
	return fmt.Sprintf("%s%s", prefix, minioName)
}

func(m *MinioService) GetOriginalPath(minioName string) string {
	prefix := ""
	return fmt.Sprintf("%s%s", prefix, minioName)
}

func NewMinioService(mc *minio.Client) *MinioService {
	return &MinioService{mc}
}