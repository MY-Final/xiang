# 部署指南

本文档介绍如何部署「湘遇」恋爱博客系统。

---

## 方式一：Docker Compose 部署（推荐）

### 前置要求

- Docker 20+
- Docker Compose 2.0+

### 快速启动

```bash
# 1. 克隆项目
git clone <repository-url>
cd xiang

# 2. 配置环境变量
cp .env.example .env
# 编辑 .env 文件，修改 JWT_SECRET 和数据库密码

# 3. 启动所有服务
docker-compose up -d

# 4. 查看日志
docker-compose logs -f

# 5. 停止服务
docker-compose down
```

### 服务访问

| 服务 | 端口 | 说明 |
|------|------|------|
| 前端 | 6100 | http://localhost:6100 |
| 后端 | 8979 | http://localhost:8979 |
| PostgreSQL | 5432 | 数据库 |
| Redis | 6379 | 缓存 |

### Swagger 文档

访问 http://localhost:8979/swagger/index.html 查看 API 文档。

---

## 方式二：手动部署

### 后端部署

#### 1. 安装 Go 1.23+

```bash
# 验证安装
go version
```

#### 2. 编译后端

```bash
cd backend

# 下载依赖
go mod tidy

# 编译
go build -o server ./cmd/server

# 运行
./server
```

#### 3. 配置环境变量

```bash
# 复制环境变量文件
cp .env.production.example .env

# 编辑 .env 文件，配置数据库和 JWT
```

#### 4. 数据库初始化

启动后会自动执行 AutoMigrate 创建表结构。

#### 5. 默认管理员账户

- 用户名：`admin`
- 密码：`myhx0308`

**首次登录后请立即修改密码！**

---

### 前端部署

#### 1. 安装 Node.js 20+

```bash
node --version
npm --version
```

#### 2. 安装依赖

```bash
cd frontend
npm install
```

#### 3. 配置环境变量

```bash
# 创建 .env.production 文件
cat > .env.production << EOF
VITE_API_BASE_URL=https://api.yourdomain.com
EOF
```

#### 4. 构建生产版本

```bash
npm run build
```

#### 5. 部署到 Nginx

```bash
# 将 dist 目录复制到 Nginx 目录
sudo cp -r dist/* /usr/share/nginx/html/

# 或者使用 Docker
docker build -t love-blog-frontend .
docker run -p 6100:80 love-blog-frontend
```

---

## 方式三：云服务器部署

### 腾讯云/阿里云 ECS

```bash
# 1. 安装 Docker
curl -fsSL https://get.docker.com | bash

# 2. 启动 Docker Compose
docker-compose up -d

# 3. 配置 Nginx 反向代理
# 编辑 /etc/nginx/nginx.conf
```

### Nginx 配置示例

```nginx
server {
    listen 80;
    server_name your-domain.com;

    # 前端
    location / {
        proxy_pass http://localhost:6100;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }

    # 后端 API
    location /api/ {
        proxy_pass http://localhost:8979;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }
}
```

---

## 数据库备份

```bash
# 导出数据库
pg_dump -U postgres love_blog > backup.sql

# 导入数据库
psql -U postgres love_blog < backup.sql
```

---

## 常见问题

### 1. 端口被占用

修改 `docker-compose.yml` 中的端口映射：
```yaml
ports:
  - "8980:8979"  # 将 8979 改为其他端口
```

### 2. 数据库连接失败

检查 `.env` 文件中的数据库配置，确保 PostgreSQL 服务正常运行。

### 3. 前端无法访问后端 API

- 检查 CORS 配置
- 确保后端服务已启动
- 检查防火墙设置

---

## 性能优化建议

1. **启用 Redis 缓存** - 缓存热点数据
2. **配置 CDN** - 加速静态资源
3. **开启 Gzip** - 压缩响应数据
4. **数据库索引** - 优化查询性能
5. **日志轮转** - 防止日志文件过大

---

## 监控和日志

```bash
# 查看 Docker 日志
docker-compose logs -f backend
docker-compose logs -f frontend

# 查看容器状态
docker-compose ps

# 查看资源使用
docker stats
```

---

**文档结束**
