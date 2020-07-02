FROM golang

LABEL maintainer="tommychu2256@gmail.com"

ENV PORT=8081
EXPOSE $PORT

ENV GOPROXY=https://proxy.golang.org
ENV GO111MODULE=on

ENV GIN_MODE=release

RUN mkdir /app
WORKDIR /app
COPY . /app

RUN install github.com/chutified/bookstore-api
RUN go build -o bookstore-api .

CMD ["/app/bookstore-api"]
