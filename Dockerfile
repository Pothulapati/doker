FROM golang:latest

LABEL maintainer="Tarun Pothulapati <tarunpothulapati@outlook.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN go build -o main ./cmd/kubekerd

EXPOSE 8080

CMD ["./main"]