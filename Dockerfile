# --- Stage 1: Build generator ---
FROM golang:1.27-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o data-generator .

# --- Stage 2: Runtime ---
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/data-generator /app/data-generator

EXPOSE 8080

ENTRYPOINT ["/app/data-generator"]