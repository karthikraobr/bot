# Build stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build cmd/server/server.go cmd/server/config.go

# Run stage
FROM alpine:3.18

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/server .

# Expose the port the app runs on
EXPOSE 8080

# Run the application
ENTRYPOINT ["./server"]
