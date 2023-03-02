FROM golang:1.20.1-alpine3.16

WORKDIR /app
COPY go.mod ./
COPY problems.csv ./

RUN go mod download
COPY *.go ./
RUN go build -o ./solution
ENTRYPOINT  ["./solution"]