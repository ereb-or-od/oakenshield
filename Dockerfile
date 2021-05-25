# Dockerfile Example
# https://medium.com/@petomalina/using-go-mod-download-to-speed-up-golang-docker-builds-707591336888
# Based on this image: https:/hub.docker.com/_/golang/
FROM golang:latest as builder

RUN mkdir -p /app
WORKDIR /app

# Force the go compiler to use modules
ENV GO111MODULE on
# <- COPY go.mod and go.sum files to the workspace
COPY go.mod .
COPY go.sum .
COPY /cmd .
COPY /pkg ./app/pkg
# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download

WORKDIR /app
RUN chmod +x /app

# Compile application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o id-service
#Image Diff
#(Not Scratch) 1.23GB
#(Scratch    ) 34.3MB
# <- Second step to build minimal image
FROM scratch
WORKDIR /root/
COPY --from=builder /app .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# Execite application when container is started

EXPOSE 80/tcp
CMD ["./id-service"]

EXPOSE 80