FROM golang:1.24 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o symphony-api ./cmd/api

FROM gcr.io/distroless/static-debian11
WORKDIR /app

COPY --from=builder /app/symphony-api .

EXPOSE 8080
ENTRYPOINT ["./symphony-api"]
