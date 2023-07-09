FROM golang:alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o goose ./cmd/migrations/migrations.go

CMD ["./goose", "up"]