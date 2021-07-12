FROM golang:1.15 AS build
WORKDIR /app
COPY . .
ENV GO111MODULE=on
RUN go mod vendor -v
CMD make start