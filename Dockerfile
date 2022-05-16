FROM golang:1.18.2-alpine3.15

ENV GIN_MODE=release

WORKDIR /menu-api

COPY . .

RUN go mod tidy

RUN go build .

CMD ["./menu-api"]
