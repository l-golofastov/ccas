FROM golang:1.19.13

WORKDIR /app

COPY go.mod .
COPY main.go .
COPY parser.go .
COPY generator.go .
COPY tests/data.txt .

RUN go build -o bin/ccas .

ENTRYPOINT ["/app/bin/ccas"]
