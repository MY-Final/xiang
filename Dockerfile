# ======================================================
# 恋爱博客 - 统一 Docker 构建文件
# 包含前后端服务
# ======================================================

# ======================================================
# 第一阶段：构建前端
# ======================================================
FROM node:20-alpine AS frontend-builder

WORKDIR /app/frontend

# 复制 package 文件
COPY frontend/package*.json ./

# 安装依赖
RUN npm install --frozen-lockfile

# 复制前端源码
COPY frontend/ ./

# 构建生产版本
RUN npm run build

# ======================================================
# 第二阶段：构建后端
# ======================================================
FROM golang:1.23-alpine AS backend-builder

WORKDIR /app/backend

# 安装依赖
RUN apk add --no-cache git

# 复制 go.mod 和 go.sum
COPY backend/go.mod backend/go.sum ./
RUN go mod download

# 复制后端源码
COPY backend/ ./

# 编译后端
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server ./cmd/server

# ======================================================
# 第三阶段：生产镜像
# ======================================================
FROM alpine:latest

WORKDIR /app

# 安装运行时依赖
RUN apk --no-cache add ca-certificates nginx nginx-mod-http-headers-more

# 创建必要目录
RUN mkdir -p /app/www /app/logs /app/uploads

# 复制后端服务
COPY --from=backend-builder /app/backend/server /app/server

# 复制前端构建产物
COPY --from=frontend-builder /app/frontend/dist /app/www

# 复制 Nginx 配置
COPY frontend/nginx.conf /etc/nginx/http.d/default.conf

# 复制启动脚本
COPY <<EOF /app/start.sh
#!/bin/sh

# 启动后端
/app/server &

# 启动 Nginx
nginx -g 'daemon off;'
EOF

# 添加执行权限
RUN chmod +x /app/start.sh

# 暴露端口
EXPOSE 5700

# 健康检查
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:5700/healthz || exit 1

# 启动服务
CMD ["/app/start.sh"]
