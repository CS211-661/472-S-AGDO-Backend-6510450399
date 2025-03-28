# Build Stage
FROM golang:1.22 AS builder
WORKDIR /app

# Copy go.mod and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy application source
COPY . .

# Build the application
RUN go build -o main ./cmd/main.go

# Final Stage
FROM alpine:latest  # ใช้ Alpine ธรรมดาเพื่อลดขนาด
WORKDIR /root/

# Copy the built binary from builder
COPY --from=builder /app/main .

# Set the default command
CMD ["./main"]
