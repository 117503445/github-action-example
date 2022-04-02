FROM golang:1.18.0-alpine3.15 as build
LABEL maintainer="117503445"
WORKDIR /root/project
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -ldflags="-s -w" -o server

FROM alpine:3.15 as prod
EXPOSE 8080
WORKDIR /root
COPY --from=build /root/project/server server
ENTRYPOINT server
