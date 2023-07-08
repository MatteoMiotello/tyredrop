# Stage 1: Build the Go application
FROM golang:1.20-alpine

WORKDIR /cron

# Copy and download Go module dependencies
COPY go.mod go.sum ./

RUN go mod download

# Copy the Go application source code
COPY . .

# Build the Go application
RUN go build -o /cron ./cmd/cron/cron.go

# Set the entry point for the container
CMD ["/cron"]