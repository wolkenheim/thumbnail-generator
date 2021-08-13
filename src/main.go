package main

import (
	"github.com/spf13/viper"
	"log"
	"net/http"
	"wolkenheim.cloud/thumbnail-generator/app"
	"wolkenheim.cloud/thumbnail-generator/controller"
	"wolkenheim.cloud/thumbnail-generator/service"
)

func main() {
	minioClient, err := app.MinioClientFactory()
	if err != nil {
		panic("Minio init failed")
	}

	a := &app.Application{}
	minioService := &service.MinioService{}
	minioService.SetClient(minioClient)

	process := &service.ProcessMinioFacade{}
	process.SetMinioService(minioService)
	process.SetFileService(&service.LocalFileService{})
	process.SetThumbnailGenerator(&service.VipsThumbnailGenerator{})

	c := controller.CreateController{}
	c.SetApp(&app.Application{})
	c.SetProcess(process)
	c.SetValidator(controller.ValidatorFactory())

	mux := http.NewServeMux()
	mux.HandleFunc("/readiness", a.Liveness)
	mux.HandleFunc("/liveness", a.Liveness)
	mux.Handle("/create", a.IsPostMiddleware(a.IsJSONMiddleware(http.HandlerFunc(c.Create))))

	log.Fatal(http.ListenAndServe(viper.GetString("server.port"), mux))

}

func init() {
	app.InitLogger()
	app.InitConfig()
}
