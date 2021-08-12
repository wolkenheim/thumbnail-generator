package main

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/spf13/viper"
	"wolkenheim.cloud/thumbnail-generator/app"
	"fmt"
)

func main(){
	minioClient, err := app.MinioClientFactory()
	if err != nil {
		panic("Minio init failed")
	}

	err = minioClient.FGetObject(
		context.TODO(),
		viper.GetString("minio.bucket"),
		"livia-sAVFADKItCo-unsplash.jpg",
		fmt.Sprintf("%s%s", viper.GetString("localImageDir"), "livia-sAVFADKItCo-unsplash.jpg"),
		minio.GetObjectOptions{},
	)
	if err != nil {
		fmt.Printf("Not able to get image! %s", err.Error())
	}
}

func init()  {
	app.InitLogger()
	app.InitConfig()
}