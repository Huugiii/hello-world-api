FROM alpine:latest

ARG APP_NAME

WORKDIR /app
COPY ${APP_NAME} ./app

EXPOSE 8080

CMD ["./app"]
