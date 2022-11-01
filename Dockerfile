FROM golang:1.19.2-alpine3.16 as build
LABEL maintainer="117503445"
RUN apk add --no-cache git
WORKDIR /root/project
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -ldflags="-s -w" -o server

FROM alpine:3.16 as prod
EXPOSE 8080
WORKDIR /root

# https://stackoverflow.com/questions/66963068/docker-alpine-executable-binary-not-found-even-if-in-path
RUN apk add gcompat 

COPY --from=build /root/project/server server
ENTRYPOINT ./server
