FROM golang:alpine

WORKDIR /app

RUN apk add --no-cache tzdata
ENV TZ=Europe/Rome

COPY . .

RUN go mod download

RUN go build -o goose ./cmd/migrations/migrations.go

CMD ["./goose", "up"]