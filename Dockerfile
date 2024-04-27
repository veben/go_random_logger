FROM golang:latest

ARG GITHUB_USERNAME
ENV MY_ENV=$GITHUB_USERNAME

MAINTAINER "${GITHUB_USERNAME}"
LABEL org.opencontainers.image.source="https://github.com/${GITHUB_USERNAME}/go_random_logger"

WORKDIR /app
COPY . .
RUN go mod tidy && go build -o go_random_logger .
CMD ["./go_random_logger"]