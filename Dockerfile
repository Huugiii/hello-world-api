FROM golang:1.24-alpine AS builder

ARG APP_NAME
ARG APP_VERSION

WORKDIR /build

COPY . .

RUN go build -o ${APP_NAME} cmd/app/main.go

FROM alpine:latest

ARG APP_NAME
ARG APP_VERSION

LABEL name="${APP_NAME}"
LABEL version="${APP_VERSION}"
LABEL maintainer="support@lakaz.co"

ENV APP_NAME=${APP_NAME}
ENV APP_VERSION=${APP_VERSION}

WORKDIR /app
COPY --from=builder /build/${APP_NAME} /app/${APP_NAME}

EXPOSE 8080
CMD ["/bin/sh", "-c", "/app/${APP_NAME}"]
