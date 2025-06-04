FROM golang:1.24-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN $(go env GOPATH)/bin/swag init -g cmd/api/main.go

RUN CGO_ENABLED=0 GOOS=linux go build -o symphony-api ./cmd/api

FROM gcr.io/distroless/static-debian11
WORKDIR /app

COPY --from=builder /app/symphony-api .
COPY --from=builder /app/docs ./docs

EXPOSE 8080
ENTRYPOINT ["./symphony-api"]