# thumbnail-generator

Cloud native service to generate thumbnails from min.io. 

First, a few packages are added and initialized. I use [viper](github.com/spf13/viper) for configuration and 
[zap](go.uber.org/zap) logger for logging.

I added the docker setup for local development. The project will use the underlying libvips library. To 
make sure all developers get the same library and avoid "works on my machine", let´s develop inside the
container. One catch tough: for unknown reasons the IDE GOLAND does not support Go runtimes inside Docker.
There are only build targets with Docker supported. So make sure you check that the
GoLand > Preferences > Go runtime is compatible with this project. Which is build on Go 1.16. 

```
cd _INFRA/dev && docker-compose up -d 
docker exec -it dev_thumbnail-generator_1 bash
go run .
```