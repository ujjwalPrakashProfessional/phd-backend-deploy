# ---- Stage 1: Build Go binary ----
FROM golang:1.19-bullseye AS builder

WORKDIR /app

# Download dependencies first (cached layer)
COPY go.mod go.sum ./
RUN go mod download

# Copy source and build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd

# ---- Stage 2: Runtime ----
FROM debian:bullseye-slim

RUN apt-get update && \
    apt-get install -y nginx ca-certificates && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/server .

# Copy nginx config
COPY container/nginx.conf /etc/nginx/sites-enabled/default
RUN rm -f /etc/nginx/sites-enabled/default.bak

# Copy config.yaml (secrets come via env vars at runtime)
COPY config.yaml .

EXPOSE 80

# Start nginx then the Go server
CMD service nginx start && ./server
