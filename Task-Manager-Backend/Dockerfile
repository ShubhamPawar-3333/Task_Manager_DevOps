# Build Stage
FROM golang:1.21 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o Task-Manager-Backend cmd/main.go

# Run Stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/Task-Manager-Backend .
VOLUME /app
EXPOSE 8080
CMD ["./Task-Manager-Backend"]