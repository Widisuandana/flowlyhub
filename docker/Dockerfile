# ---- Builder Stage ----
# Menggunakan versi Go yang sesuai dengan go.mod
FROM golang:1.24.3-alpine AS builder

WORKDIR /app

# Copy go.mod dan go.sum terlebih dahulu untuk caching
COPY go.mod go.sum ./
RUN go mod download

# Copy seluruh source code
COPY . .

# Compile aplikasi.
# CGO_ENABLED=0 penting untuk static binary
# -o /app/main akan membuat file binary bernama 'main'
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main ./cmd/api

# ---- Final Stage ----
# Tahap ini untuk menjalankan aplikasi yang sudah di-compile
FROM alpine:latest

WORKDIR /app

# Copy binary dari builder stage
COPY --from=builder /app/main .

# Expose port yang digunakan aplikasi
EXPOSE 8080

# Command untuk menjalankan aplikasi saat container start
CMD ["/app/main"]
