# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY /controler ./
COPY /model ./
COPY /service ./

RUN go build -o /mock-api

EXPOSE 8080

CMD [ "/mock-api" ]