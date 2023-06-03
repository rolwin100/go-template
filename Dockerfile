FROM golang:1.17-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .


RUN go build -o /docker-run-adopler

EXPOSE 8000

CMD [ "/docker-run-adopler" ]
