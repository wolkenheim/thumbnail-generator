package service

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
)

type FileService interface{
	DeleteFile(localFilePath string)
	GetLocalOriginalPath(fileName string) string
	GetLocalThumbnailPath(fileName string) string
}

type LocalFileService struct {}

func(d *LocalFileService) DeleteFile(localFilePath string)  {
	err := os.Remove(localFilePath)
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