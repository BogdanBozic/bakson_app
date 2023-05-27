# This image uses Debian as its underlying image. Normally I would use an Alpine image, but it is not stable yet.
# See https://github.com/golang/go/issues/19938
FROM golang:1.20.4

ENV GIN_MODE=release
WORKDIR /app
COPY go.* /app/
RUN go mod download

# Copy only the necessary files. In this case it is only a single module, but otherwise this approach
# would have to be revised.
WORKDIR /app/handler
COPY /handler /app/handler/
RUN go build .

# Do not run as root.
USER 1001
EXPOSE 3000

CMD ["./handler"]