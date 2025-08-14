# Stage 1: Build
FROM golang:1.24 AS builder
WORKDIR /app

# Salin hanya file mod & sum untuk caching dependensi
COPY auth_service/go.mod auth_service/go.sum ./
RUN go mod download

# Build aplikasi
RUN go build -o /app/out .

# Stage 2: Final image
FROM debian:stable-slim
WORKDIR /root/

# Salin binary yang sudah dicompile dari stage builder
COPY --from=builder /app/out .

# EXPOSE port yang akan didengarkan oleh aplikasi Go Anda (BUKAN port database)
EXPOSE 8080
CMD ["./out"]


