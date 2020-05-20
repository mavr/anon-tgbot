# Собираем бинарник
FROM golang:1.14 AS build-stage

COPY . /go/src/github.com/mavr/anon-mail

WORKDIR /go/src/github.com/mavr/anon-mail

# Зашьем в бинарник ревизию, которую дернем из гита
RUN go mod download && CGO_ENABLED=0 GOOS=linux go build -ldflags "-X main.revision=$(git rev-parse --abbrev-ref HEAD)-$(git rev-parse HEAD)-$(date +%Y%m%d-%H:%M:%S) -s -w" -o bin/anon-mail cmd/anon-mail/*.go

# Запускаем бинарник
FROM alpine:3.11

# Навернем корневых сертификатов и установим таймзону
RUN apk update && apk add --no-cache ca-certificates tzdata
RUN cp /usr/share/zoneinfo/Europe/Moscow /etc/localtime

COPY --from=build-stage /go/src/github.com/mavr/anon-mail/ /app/
RUN mkdir -p /app/conf/

WORKDIR /app
ENTRYPOINT ./bin/anon-mail