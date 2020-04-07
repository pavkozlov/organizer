FROM golang:alpine
RUN apk update && rm -rf /var/cache/apk/*
RUN mkdir -p /app
WORKDIR /app
ADD . /app
RUN go build ./main.go