FROM golang:1.24-alpine AS builder

WORKDIR /build

COPY . .

RUN go build -o hello-world-api cmd/app/main.go

FROM alpine:latest

WORKDIR /app
COPY --from=builder /build/hello-world-api /app/hello-world-api

EXPOSE 8080

CMD ["./hello-world-api"]
