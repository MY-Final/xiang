# 湘遇恋爱博客 - 服务器部署指南

> 💡 在服务器上 5 分钟内部署您的恋爱博客系统

---

## 📋 前置要求

- 已安装 Docker 和 Docker Compose 的 Linux 服务器
- 服务器端口 5700 已开放
- 根权限或 docker 组权限

---

## 🚀 一键部署流程

### 步骤 1：克隆代码

```bash
git clone <您的仓库地址>
cd xiang
```

### 步骤 2：配置环境变量

```bash
# 复制环境变量文件
cp .env.example .env

# 编辑 .env 文件
nano .env
```

**必须修改的配置项：**

```bash
# 1. 生成 JWT_SECRET（至少 32 字符）
# 在终端执行以下命令生成随机字符串：
openssl rand -base64 32
# 输出示例：xYz123AbC456DeF789GhI012JkL345MnO678PqR=
# 将输出值复制到 .env 中的 JWT_SECRET

# 2. 修改数据库密码（可选，默认：myhx0308）
DB_PASSWORD=your_secure_password
```

### 步骤 3：启动服务

```bash
# 方式 1：使用 docker-compose（推荐）
docker-compose up -d --build

# 方式 2：使用部署脚本
chmod +x deploy.sh
./deploy.sh
```

### 步骤 4：验证部署

```bash
# 查看服务状态
docker-compose ps

# 查看应用日志
docker-compose logs -f app

# 健康检查
curl http://localhost:5700/healthz
# 应输出：OK
```

### 步骤 5：访问网站

在浏览器打开：
```
http://服务器IP:5700
```

**默认管理员账户：**
- 用户名：`admin`
- 密码：`myhx0308`

⚠️ **首次登录后请立即修改密码！**

---

## 🔧 常用命令

### 查看服务状态
```bash
docker-compose ps
```

### 查看日志
```bash
# 查看所有服务日志
docker-compose logs -f

# 仅查看后端日志
docker-compose logs -f app

# 仅查看数据库日志
docker-compose logs -f postgres
```

### 重启服务
```bash
docker-compose restart
```

### 停止服务
```bash
docker-compose down
```

### 完全清理（慎用！会删除所有数据）
```bash
docker-compose down -v
```

---

## 🔍 故障排查

### 1. 容器无法启动

```bash
# 查看详细错误
docker-compose logs app

# 检查 .env 配置
cat .env

# 确保 JWT_SECRET 至少 32 字符
```

### 2. 数据库连接失败

```bash
# 检查数据库容器状态
docker-compose ps postgres

# 查看数据库日志
docker-compose logs postgres

# 重启数据库
docker-compose restart postgres
```

### 3. 端口被占用

```bash
# 检查端口占用
netstat -tunlp | grep 5700

# 修改 docker-compose.yml 中的端口映射
ports:
  - "5701:5700"  # 改为其他端口
```

### 4. 构建失败

```bash
# 清理 Docker 缓存
docker system prune -a

# 重新构建（不使用缓存）
docker-compose build --no-cache
```

---

## 🔐 安全建议

1. **修改默认密码** - 首次登录后立即修改管理员密码
2. **设置强 JWT_SECRET** - 至少 32 字符的随机字符串
3. **配置防火墙** - 仅开放必要端口（5700）
4. **定期备份** - 备份 PostgreSQL 数据卷
5. **使用 HTTPS** - 生产环境建议配置 Nginx 反向代理和 SSL 证书

---

## 💾 数据备份

### 导出数据库
```bash
docker-compose exec postgres pg_dump -U postgres love_blog > backup_$(date +%Y%m%d).sql
```

### 导入数据库
```bash
cat backup.sql | docker-compose exec -T postgres psql -U postgres love_blog
```

### 备份上传文件
```bash
# 找到数据卷位置
docker volume inspect xiang_app_uploads

# 打包备份
tar -czf uploads_backup.tar.gz /var/lib/docker/volumes/xiang_app_uploads/_data
```

---

## 📊 配置说明

### .env 文件配置项

| 变量名 | 说明 | 默认值 |
|--------|------|--------|
| DB_HOST | 数据库地址 | postgres（Docker 内部） |
| DB_PORT | 数据库端口 | 5432 |
| DB_USER | 数据库用户 | postgres |
| DB_PASSWORD | 数据库密码 | myhx0308 |
| DB_NAME | 数据库名 | love_blog |
| JWT_SECRET | JWT 密钥 | 必须修改！ |
| JWT_EXPIRE_HOURS | Token 有效期 | 72 小时 |
| ALLOW_ORIGINS | CORS 域名 | http://localhost:5700 |

---

## 🎉 部署完成

访问地址：`http://服务器 IP:5700`

默认管理员：
- 用户名：`admin`
- 密码：`myhx0308`

**祝您使用愉快！💕**
