FROM golang:latest

LABEL maintainer="tommychu2256@gmail.com"

ENV PORT=8081
EXPOSE $PORT

ENV GOPROXY=https://proxy.golang.org
ENV GO111MODULE=on

ENV GIN_MODE=release

RUN mkdir /app
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download
ADD . .

RUN go build -o main .

CMD ["/app/main"]
