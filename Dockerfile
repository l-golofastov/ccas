FROM golang:1.19.13

WORKDIR /app

COPY go.mod .
COPY src/main.go .
COPY src/parser.go .
COPY src/generator.go .
COPY tests/data.txt .

RUN go build -o bin/ccas .

ENTRYPOINT ["/app/bin/ccas"]
