# Stage 1: Build
FROM golang:1.24 AS builder
WORKDIR /app

# Salin hanya file mod & sum untuk caching dependensi
COPY auth_service/go.mod auth_service/go.sum ./
RUN go mod download

# Salin sisa source code dari service Anda
COPY auth_service/ ./

# Build aplikasi
RUN go build -o /app/out .

# Stage 2: Final image
FROM debian:stable-slim
WORKDIR /root/

# Salin binary yang sudah dicompile dari stage builder
COPY --from=builder /app/out .

# (Opsional) Salin file lain yang mungkin dibutuhkan, contoh: .env, certs
COPY auth_service/.env .

EXPOSE 5432 # Ganti dengan port yang Anda gunakan
CMD ["./out"]


