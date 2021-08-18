FROM golang:1.17.0-bullseye as builder

ENV DEBIAN_FRONTEND=noninteractive

RUN apt update \
    && apt install -y --no-install-recommends \
    tzdata ca-certificates build-essential git curl wget \
    libvips-dev libvips-tools

ENV PATH "$PATH:/usr/local/go/bin"

WORKDIR /app

COPY ./ /app

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -a  -o /app/web-server .


FROM debian:bullseye

ENV DEBIAN_FRONTEND=noninteractive

RUN apt update \
    && apt install -y --no-install-recommends \
    tzdata ca-certificates \
    libvips-dev libvips-tools

ENV USER=appuser
ENV UID=10001

# See https://stackoverflow.com/a/55757473/12429735
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

COPY --from=builder /app/web-server /go/bin/web-server
COPY --from=builder /app/config /app/config
COPY --from=builder /app/images /app/images

RUN chown -R appuser:appuser /app

USER appuser:appuser

ENTRYPOINT ["/go/bin/web-server"]