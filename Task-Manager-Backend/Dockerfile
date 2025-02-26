# Build stage
FROM golang:1.23 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o Task-Manager-Backend cmd/main.go
RUN chmod +x Task-Manager-Backend

# Run stage
FROM alpine:latest
RUN apk add --no-cache libc6-compat
WORKDIR /app
COPY --from=builder /app/Task-Manager-Backend .
RUN touch tasks.db && chmod 666 tasks.db
VOLUME /app
EXPOSE 8080
CMD ["./Task-Manager-Backend"]
