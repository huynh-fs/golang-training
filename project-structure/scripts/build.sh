#!/bin/bash
# file: scripts/build.sh

echo "Building user-service..."
# Lệnh build, -o chỉ định file output vào thư mục /bin
go build -o bin/user-service ./cmd/user-api

echo "Build hoàn tất! File thực thi nằm tại: bin/user-service"