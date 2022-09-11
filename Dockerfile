# syntax=docker/dockerfile:1]
FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./

COPY go.sum ./

RUN go mod download

COPY *.go ./

ADD app ./app

RUN go build -o /timeassignment

EXPOSE 8081

CMD ["/timeassignment"]

