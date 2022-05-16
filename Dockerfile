FROM golang:1.18.2-alpine

WORKDIR /menu-api

COPY . .

RUN go mod tidy

RUN go build .

CMD ["/menu-api"]
