# 快速部署指南

> 💡 5 分钟内部署您的恋爱博客系统

---

## 📋 前置要求

- Docker 20.10+
- Docker Compose 2.0+
- 服务器端口 5700 开放

---

## 🚀 一键部署

### 步骤 1：上传代码到服务器

```bash
# 方式 1：使用 git clone
git clone <repository-url>
cd xiang

# 方式 2：打包上传
# 本地打包
zip -r xiang.zip .
# 上传到服务器
scp xiang.zip user@server:~/
# 服务器上解压
unzip xiang.zip
cd xiang
```

### 步骤 2：配置环境变量

```bash
# 复制环境变量文件
cp .env.example .env

# 编辑 .env 文件
nano .env

# 重要配置项：
# JWT_SECRET=生成一个 32 字符以上的随机字符串
# DB_PASSWORD=修改数据库密码
```

**生成 JWT_SECRET：**
```bash
# Linux/Mac
openssl rand -base64 32

# 输出示例：xYz123AbC456DeF789GhI012JkL345MnO678PqR=
```

### 步骤 3：启动服务

```bash
# 方式 1：使用 docker-compose
docker-compose up -d

# 方式 2：使用部署脚本
chmod +x deploy.sh
./deploy.sh
```

### 步骤 4：访问网站

```
http://服务器IP:5700
```

---

## 👥 默认账户

```
用户名：admin
密码：myhx0308
```

**⚠️ 首次登录后请立即修改密码！**

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

# 查看后端日志
docker-compose logs -f app

# 查看数据库日志
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

### 完全清理

```bash
# 停止服务并删除容器
docker-compose down

# 删除数据卷（谨慎使用！）
docker-compose down -v
```

---

## 🔍 故障排查

### 1. 端口被占用

```bash
# 检查端口占用
netstat -tunlp | grep 5700

# 修改 docker-compose.yml 中的端口映射
ports:
  - "5701:5700"  # 改为其他端口
```

### 2. 数据库连接失败

```bash
# 检查数据库容器
docker-compose ps postgres

# 查看数据库日志
docker-compose logs postgres

# 重启数据库
docker-compose restart postgres
```

### 3. 构建失败

```bash
# 清理 Docker 缓存
docker system prune -a

# 重新构建
docker-compose build --no-cache
```

### 4. 容器启动失败

```bash
# 查看详细错误
docker-compose logs app

# 检查 .env 配置
cat .env

# 确保 JWT_SECRET 至少 32 字符
```

---

## 📊 健康检查

```bash
# 检查服务是否正常
curl http://localhost:5700/healthz

# 预期输出：OK
```

---

## 🔐 安全建议

1. **修改默认密码** - 首次登录后立即修改管理员密码
2. **设置强 JWT_SECRET** - 至少 32 字符的随机字符串
3. **配置防火墙** - 仅开放必要端口（5700）
4. **定期备份** - 备份 PostgreSQL 数据卷
5. **使用 HTTPS** - 生产环境建议配置反向代理和 SSL

---

## 💾 数据备份

```bash
# 导出数据库
docker-compose exec postgres pg_dump -U postgres love_blog > backup.sql

# 导入数据库
cat backup.sql | docker-compose exec -T postgres psql -U postgres love_blog
```

---

## 📝 配置说明

### .env 文件配置项

| 变量名 | 说明 | 默认值 |
|--------|------|--------|
| DB_HOST | 数据库地址 | 127.0.0.1 |
| DB_PORT | 数据库端口 | 5432 |
| DB_USER | 数据库用户 | postgres |
| DB_PASSWORD | 数据库密码 | 需修改 |
| DB_NAME | 数据库名 | love_blog |
| JWT_SECRET | JWT 密钥 | 需修改 |
| JWT_EXPIRE_HOURS | Token 有效期 | 72 |
| ALLOW_ORIGINS | CORS 域名 | http://localhost:5700 |

---

## 💗 技术支持

如有问题，请查看：
- 项目文档：`README.md`
- 详细部署指南：`DEPLOY.md`
- API 文档：`http://IP:5700/swagger`

---

**祝您部署顺利！💕**
