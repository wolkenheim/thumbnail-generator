package main

import (
	"github.com/spf13/afero"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"wolkenheim.cloud/thumbnail-generator/app"
	"wolkenheim.cloud/thumbnail-generator/handler"
	"wolkenheim.cloud/thumbnail-generator/service"
)

func main() {
	var appFs = afero.NewOsFs()


	minioClient, err := app.MinioClientFactory()
	if err != nil {
		panic("Minio init failed. Cannot start application.")
	}

	logger := app.NewZapLogger()
	a := app.NewApplication()

	minioService := service.NewMinioService(minioClient)
	process := service.NewProcessMinioFacade(minioService,&service.VipsThumbnailGenerator{},
	service.NewLocalFileService(logger.Sugar(), appFs), logger.Sugar())
	createHandler := handler.NewCreateHandler(a,process, handler.NewCreateValidator())

	mux := http.NewServeMux()
	mux.HandleFunc("/readiness", a.Liveness)
	mux.HandleFunc("/liveness", a.Liveness)
	mux.Handle("/create", a.IsPostMiddleware(a.IsJSONMiddleware(http.HandlerFunc(createHandler.Create))))

	log.Fatal(http.ListenAndServe(viper.GetString("server.port"), mux))

}

func init() {
	app.InitLogger()
	app.InitConfig()
}
