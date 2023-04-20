FROM golang:alpine

WORKDIR /go/src

COPY . .

ENV GIN_MODE=release

RUN go build -o api .

EXPOSE 8080

CMD ["./api"]