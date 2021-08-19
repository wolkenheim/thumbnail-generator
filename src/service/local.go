package service

import (
	"fmt"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type FileService interface{
	DeleteFile(localFilePath string)
	GetLocalOriginalPath(fileName string) string
	GetLocalThumbnailPath(fileName string) string
}

type LocalFileService struct {
	logger *zap.SugaredLogger
	fs afero.Fs
}

func(d *LocalFileService) DeleteFile(localFilePath string)  {
	err := d.fs.Remove(localFilePath)
	if err != nil {
		zap.S().Errorf("Could not delete local file %s", localFilePath)
	}
}

func(d *LocalFileService) GetLocalOriginalPath(fileName string) string {
	return fmt.Sprintf("%s%s%s", viper.GetString("localImageDir"), "originals/", fileName)
}

func(d *LocalFileService) GetLocalThumbnailPath(fileName string) string {
	return fmt.Sprintf("%s%s%s", viper.GetString("localImageDir"), "thumbnails/", fileName)
}

func NewLocalFileService(l *zap.SugaredLogger, fs afero.Fs) *LocalFileService{
	return &LocalFileService{l, fs}
}