# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY main.go .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ollama-metrics .

# Final stage
FROM scratch

# Copy the binary from builder
COPY --from=builder /app/ollama-metrics /ollama-metrics

# Command to run
ENTRYPOINT ["/ollama-metrics"]