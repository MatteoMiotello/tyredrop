# Stage 1: Build the Go application
FROM golang:1.20-alpine

WORKDIR /api

# Copy and download Go module dependencies
COPY go.mod go.sum ./

RUN go mod download
RUN go mod verify

# Copy the Go application source code
COPY . .

RUN go build -o /goose ./cmd/migrations/migrations.go

# Build the Go application
RUN go build -o /server ./cmd/server/server.go

# Expose the necessary port(s)
EXPOSE 8080

# Set the entry point for the container
CMD /bin/sh -c '/goose up && /server'