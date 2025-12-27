# ---------- Stage 1: Build ----------
#FROM golang:1.22-alpine AS builder
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy Go mod files and download deps
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o student-api ./cmd/student_api

# ---------- Stage 2: Run ----------
FROM alpine:latest

WORKDIR /app

RUN mkdir -p storage

# Copy binary from builder
COPY --from=builder /app/student-api .

# Copy config directory
COPY config ./config

# Expose application port
EXPOSE 8082

# Run the application
CMD ["./student-api", "--config=config/local.yaml"]

HEALTHCHECK --interval=30s --timeout=5s --start-period=10s --retries=3 \
  CMD wget --spider -q http://localhost:8082/api/students || exit 1

