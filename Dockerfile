# Stage 1: Build the Go Binary
FROM golang:1.23.5 AS builder

WORKDIR /app

# Install SQLite dependencies for CGO
RUN apt-get update && apt-get install -y gcc libc-dev libsqlite3-dev ca-certificates

# Set environment for CGO
ENV CGO_ENABLED=1 GOOS=linux GOARCH=amd64

# Copy go modules
COPY go.mod go.sum ./
RUN go mod tidy

# Copy entire project source
COPY . .

# Build the Go binary, using main from internal directory
RUN go build -o whatsappbot ./cmd/main.go

# Stage 2: Runtime Environment
FROM debian:bookworm-slim

WORKDIR /app

# Install required libraries and CA certificates
RUN apt-get update && apt-get install -y libsqlite3-0 ca-certificates && rm -rf /var/lib/apt/lists/*

# Copy compiled binary from builder
COPY --from=builder /app/whatsappbot .

# Persist database and QR code files
VOLUME ["/app/data"]

# Expose API port (if needed)
EXPOSE 9090

# Run the bot
CMD ["./whatsappbot"]

