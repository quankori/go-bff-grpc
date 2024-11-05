# Use Go 1.22 image as the base
FROM golang:1.23-alpine AS builder

# Set environment variables
ENV GO111MODULE=on
ENV APP_HOME=/app
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

# Set the working directory inside the container
WORKDIR /app

# Install any necessary packages (optional)
RUN apk add --no-cache git

# Copy go.mod and go.sum to the working directory
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the application as a static binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# Stage 2: Create a smaller image for production
FROM alpine:latest

# Set the working directory in the final image
WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/main .

# Expose the application port (adjust as needed)
EXPOSE 8100

# Set environment variables if needed
ENV APP_ENV=production

# Run the binary
CMD ["./main"]