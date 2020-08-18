FROM golang:alpine

LABEL maintainer="tommychu2256@gmail.com"

ENV GO111MODULE="on"

ENV PORT=8081
EXPOSE $PORT
ENV GIN_MODE=release

WORKDIR /go/src/local/bookstore-api
COPY . .
RUN go mod download

RUN go build -o main .
CMD ["./main"]
