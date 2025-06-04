FROM golang:1.24-alpine AS dev_environment

RUN apk add --no-cache git ca-certificates

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/air-verse/air@latest

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN $(go env GOPATH)/bin/swag init -g cmd/api/main.go

EXPOSE 8080

ENTRYPOINT ["air", "-c", ".air.toml"]
