# syntax=docker/dockerfile:1
# https://www.callicoder.com/docker-golang-image-container-example/

FROM golang:1.18-alpine

WORKDIR /app

# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

RUN apk update && apk add bash
RUN apk add build-base

COPY *.go ./
COPY *.json ./

RUN go mod tidy 

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
