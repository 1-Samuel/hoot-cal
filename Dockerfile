FROM golang:alpine

WORKDIR /go/src

COPY . .

RUN go build -o api main.go

ENV GIN_MODE=release

EXPOSE 8080

CMD ["./api"]