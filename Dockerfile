FROM golang:latest
ENV APP_NAME=go_random_logger
MAINTAINER veben
LABEL org.opencontainers.image.source https://github.com/veben/$APP_NAME
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o $APP_NAME .
CMD ["./$APP_NAME"]