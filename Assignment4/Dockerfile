FROM golang:1.21-alpine3.18

WORKDIR /crud

COPY . .

RUN go mod download

EXPOSE 8080

RUN go build -o test main.go

CMD ["./test"]
