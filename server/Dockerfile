# Use Go 1.22 image as the base
FROM golang:1.23 AS builder

# Set environment variables
ENV GO111MODULE=on
ENV APP_HOME=/app
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

# Install Air for live reloading
RUN go install github.com/air-verse/air@latest

# Create app directory and set it as the working directory
WORKDIR $APP_HOME

# Copy go.mod and go.sum and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the application files
COPY . .

# Expose ports for gRPC and GraphQL
EXPOSE 8100

# Command to run Air for live reload
CMD ["air"]
