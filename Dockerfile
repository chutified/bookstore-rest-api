FROM golang:latest

LABEL maintainer="tommychu2256@gmail.com"

ENV PORT=8081
EXPOSE $PORT
ENV GIN_MODE=release

WORKDIR /go/src/local/bookstore-api
COPY . .
RUN go mod vendor

RUN go build -o main .
CMD ["./main"]
