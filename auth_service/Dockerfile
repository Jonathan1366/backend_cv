# # Stage 1: Build the Go application
# FROM golang:1.24 AS builder

# # Setel direktori kerja di dalam container
# WORKDIR /app

# # Salin go.mod dan go.sum ke direktori kerja terlebih dahulu
# COPY go.mod go.sum ./

# # Unduh dependencies
# RUN go mod download

# # Salin semua file proyek ke dalam container
# COPY . .

# # Build aplikasi Go
# # RUN go build -o out .
# RUN go build -o /app/out ./auth_service


# # Stage 2: Jalankan aplikasi
# FROM debian:stable-slim

# # Setel direktori kerja untuk runtime
# WORKDIR /root/

# # Salin binary dari stage builder
# COPY --from=builder /app/out .

# # Expose port
# EXPOSE 8080

# # Jalankan aplikasi
# CMD ["./out"]

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
# COPY auth_service/.env .

EXPOSE 8080 # Ganti dengan port yang Anda gunakan
CMD ["./out"]


