# Go Image https://hub.docker.com/_/golang 
FROM golang:alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . /app

RUN go build -o main

EXPOSE 8080

CMD ["./main"]
