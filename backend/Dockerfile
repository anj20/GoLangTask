# Stage 1: Build the application
FROM golang:1.20 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire application source code to the container
COPY . .

# Build the Go application
RUN go build -o main .

# Stage 2: Create a lightweight image to run the application
FROM debian:bullseye-slim

# Install SQLite to support the database
RUN apt-get update && apt-get install -y sqlite3 && rm -rf /var/lib/apt/lists/*

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Copy any other necessary files (e.g., ads.db for SQLite)
COPY ads.db ./ads.db

# Expose the application port
EXPOSE 8080

# Set the entrypoint to run the Go application
ENTRYPOINT ["./main"]
