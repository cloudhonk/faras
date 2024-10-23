# Stage 1: Build the Go binary
FROM golang:1.23-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod ./

# Download all Go module dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go app with optimizations
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/bin/app .

# Stage 2: Create a small image for deployment
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the pre-built Go binary from the builder stage
COPY --from=builder /app/bin/app .

# Expose port 8080 (TCP server)
EXPOSE 8080

# Command to run the application
CMD ["./app"]
