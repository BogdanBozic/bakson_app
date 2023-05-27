# This image uses Debian as its underlying image. Normally I would use an Alpine image, but it is not stable yet.
# See https://github.com/golang/go/issues/19938
FROM golang:1.20.4

ENV GIN_MODE=release
WORKDIR /app
COPY . /app
RUN go mod download

WORKDIR /app/handler
RUN go build .

USER 1001
EXPOSE 3000

CMD ["./handler"]