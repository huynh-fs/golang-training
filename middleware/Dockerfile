# Stage 1: Build a Go binary
FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/api/main.go

# Stage 2: Create a small production image
FROM alpine:latest
WORKDIR /root/

RUN apk add --no-cache bash

COPY --from=builder /app/main .
COPY --from=builder /app/.env .

# Copy script wait-for-it vào trong image
COPY scripts/wait-for-it.sh .
# Cấp quyền thực thi (executable) cho script
RUN chmod +x ./wait-for-it.sh

# Expose port 8080
EXPOSE 8080
# Run the binary
CMD ["./main"]