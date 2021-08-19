package service

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type AwsService struct {
	logger *zap.SugaredLogger
	fs afero.Fs
	sess *session.Session
}

func NewAwsService(logger *zap.SugaredLogger, fs afero.Fs, sess *session.Session) *AwsService{
	return &AwsService{logger, fs, sess}
}

func (a *AwsService) Download(ctx context.Context, minioName string, localFilePath string) error {

	downloader := s3manager.NewDownloader(a.sess)

	// Create a file to write the S3 Object contents to.
	f, err := a.fs.Create(localFilePath)
	if err != nil {
		return err
	}

	// Write the contents of S3 Object to the file
	_, err = downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String(viper.GetString("minio.bucket")),
		Key:    aws.String(minioName),
	})

	if err != nil {
		return err
	}

	return nil
}

func (a *AwsService) Upload(ctx context.Context, minioName string, localFilePath string) error {
	uploader := s3manager.NewUploader(a.sess)

	f, err  := a.fs.Open(localFilePath)
	if err != nil {
		return err
	}

	// Upload the file to S3.
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(viper.GetString("minio.bucket")),
		Key:    aws.String(minioName),
		Body:   f,
	})
	if err != nil {
		return err
	}
	return nil
}

func (a *AwsService) GetThumbnailPath(minioName string) string {
	prefix := "thumbnails/"
	return fmt.Sprintf("%s%s", prefix, minioName)
}

func (a *AwsService) GetOriginalPath(minioName string) string {
	prefix := ""
	return fmt.Sprintf("%s%s", prefix, minioName)
}

func InitMinioSessionForAWSS3SDK() *session.Session{
	// Configure to use MinIO Server
	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(viper.GetString("minio.accessKeyID"), viper.GetString("minio.secretAccessKey"), ""),
		Endpoint:         aws.String(viper.GetString("minio.url")),
		Region:           aws.String("us-east-1"),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true),
	}
	return session.Must(session.NewSession(s3Config))
}