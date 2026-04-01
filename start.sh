#!/bin/sh

# 湘遇恋爱博客 - 容器启动脚本
# 用于在 Docker 容器中启动后端和 Nginx

# 环境变量转换（Docker Compose 到后端配置）
export DB_HOST="${DB_HOST:-postgres}"
export DB_PORT="${DB_PORT:-5432}"
export DB_USER="${DB_USER:-postgres}"
export DB_PASSWORD="${DB_PASSWORD:-myhx0308}"
export DB_NAME="${DB_NAME:-love_blog}"
export DB_SSLMODE="${DB_SSLMODE:-disable}"

echo "Starting Love Blog API server..."
echo "Database: ${DB_HOST}:${DB_PORT}/${DB_NAME}"

# 启动后端（后台运行）
/app/server &

# 等待后端启动
sleep 2

# 启动 Nginx（前台运行，作为容器主进程）
echo "Starting Nginx..."
exec nginx -g 'daemon off;'
