FROM golang:latest

ENV GO111MODULE=on
ENV PORT=8080
WORKDIR /app
COPY go.mod /app
COPY go.sum /app

RUN go mod download
RUN go get github.com/githubnemo/CompileDaemon
COPY . /app
ENTRYPOINT CompileDaemon --build="go build -o main" --command=./main