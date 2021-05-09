FROM golang:1.16

WORKDIR /app
COPY . /app

RUN go build -o main ./cmd/server/main.go
CMD ["./main"]