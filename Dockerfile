FROM golang:1.17.0-bullseye

ENV DEBIAN_FRONTEND=noninteractive

RUN apt update \
    && apt install -y --no-install-recommends \
    tzdata ca-certificates build-essential git curl wget \
    libvips-dev libvips-tools

ENV PATH "$PATH:/usr/local/go/bin"

WORKDIR /app

COPY ./src /app

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -a  -o /app/web-server .

ENTRYPOINT ["/app/web-server"]