package main

import (
	"wolkenheim.cloud/thumbnail-generator/app"
	"wolkenheim.cloud/thumbnail-generator/service"
)

func main() {
	minioClient, err := app.MinioClientFactory()
	if err != nil {
		panic("Minio init failed")
	}

	process := &service.ProcessMinioFacade{}
	minioService := &service.MinioService{}
	minioService.SetClient(minioClient)
	process.SetMinioService(minioService)
	process.SetFileService(&service.LocalFileService{})
	process.SetThumbnailGenerator(&service.VipsThumbnailGenerator{})

	fileName := "livia-sAVFADKItCo-unsplash.jpg"

	process.ProcessImage(fileName)

}

func init() {
	app.InitLogger()
	app.InitConfig()
}
