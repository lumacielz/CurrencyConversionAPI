FROM golang:1.18.10-alpine3.17
WORKDIR /challenge-bravo
COPY . .
CMD go run main.go
EXPOSE 8080