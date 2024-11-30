# stage de build
FROM golang:alpine AS build

WORKDIR /app

COPY . /app

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags '-s -w' -o api cmd/server.go

## stage imagem final
FROM scratch

WORKDIR /app

COPY --from=build /app/api /app/app.env ./
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 8080

CMD [ "./api" ]