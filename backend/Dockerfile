# 后端 Dockerfile
FROM golang:1.23-alpine AS builder

WORKDIR /build

# 安装依赖
RUN apk add --no-cache git

# 复制 go.mod 和 go.sum
COPY go.mod go.sum ./
RUN go mod download

# 复制源代码
COPY . .

# 编译
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server ./cmd/server

# 最终镜像
FROM alpine:latest

WORKDIR /app

# 安装 CA 证书（用于 HTTPS 请求）
RUN apk --no-cache add ca-certificates

# 从 builder 复制编译文件
COPY --from=builder /build/server .
COPY --from=builder /build/uploads ./uploads

# 暴露端口
EXPOSE 8979

# 运行
CMD ["./server"]
