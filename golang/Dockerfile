# Start from golang base image
FROM golang:alpine

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base

WORKDIR /app

# Copy the source from the current directory to the working Directory inside the container
COPY . .
COPY .env .

RUN apk add --no-cache tzdata
ENV TZ=Europe/Rome

RUN go install github.com/cosmtrek/air@latest

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# Build the Go app
# RUN go build -o /build

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD [ "air", "-c", ".air.toml"]