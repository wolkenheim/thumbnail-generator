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
See [branch](https://github.com/wolkenheim/thumbnail-generator/tree/docker-working)

Next step: Add min.io. The easiest way to get started it checking the (GUI)[http://localhost:8080/login]. Create the 
bucket 
"app-images" here. Now at _INFRA/dev/minio-data a new directory should exist. Now get an image. Why not 
this one from (unsplash)[https://unsplash.com/photos/sAVFADKItCo] Either upload it via the min.io web
interface or just drop it in the "minio-data/app-images" directory. Your bucket contains one image now.
Now your config/local.yaml file should contain the absolute path to the local download directory. This is 
inside the container so the location is /app/downloaded-images.

Now we are finally ready for a test drive. Let´s download an image from min.io and store it locally. I added a 
simple fetch operation to main.go.
Run `go run .`

See [branch](https://github.com/wolkenheim/thumbnail-generator/tree/minio-added)

Now we need to talk a bit what is going to happen next. Let´s assume the application knows the name of 
the image already. These are the steps that need to be performed:
- Download image from min.io and save it locally to originals directory
- Generate image thumbnail and save it to thumbnails directory
- Upload thumbnail to "thumbnails" path in bucket in min.io
- Delete both local files (original / thumbnail)

Download and upload are happening on the same bucket hence there is no need for a separation of services. My
approach would be: building a service for all min.io operations that are needed: Download, Upload, Delete local file
and a few helper. Then use a facade as a client of that service. Now that we have everything together interfaces need
to be defined, struct written and injected from main.go. 

This took quite a while and quite a bit code changed. It makes sense to see if everything is working. Now the image
should still be downloaded and end up in the download-images/originals directory.

See [branch](https://github.com/wolkenheim/thumbnail-generator/tree/facade)

We´re getting closer. Next the wrapper for the thumbnail generator is going to get added. We chose
(libvips)[https://libvips.github.io/libvips/API/current/]. I won´t go into detail here but reason for
that was the great research and benchmarking done here (https://github.com/fawick/speedtest-resize)[https://github.
com/fawick/speedtest-resize]. It is possible to crop files in natively in Go, or use a C library with CGO
or execute a binary. The latter being the most efficient option. One caveat here: when you want to keep
the option open to use this as an AWS Lambda function this might not be a good choice. In my own use case
this was different: we control the underlying container image and hence no problem with dependencies.

I added the next steps. Thumbnails get generated, uploaded to min.io and both local files deleted.