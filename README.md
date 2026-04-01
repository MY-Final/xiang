# 湘遇 - 恋爱博客系统

> 💕 记录属于我们的浪漫时光

![Vue 3](https://img.shields.io/badge/Vue-3.5-4FC08D?logo=vue.js)
![Go](https://img.shields.io/badge/Go-1.23-00ADD8?logo=go)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15-4169E1?logo=postgresql)
![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?logo=docker)

---

## ✨ 特性

- 💖 **浪漫主题** - 粉色/紫色渐变设计，心跳动画效果
- 📝 **日记管理** - 完整的 CRUD 功能，支持 Markdown 编辑
- 🏷️ **情感标签** - 为每个回忆贴上专属标签
- 📸 **图片上传** - 支持多图上传和封面设置
- 🔐 **用户认证** - JWT 认证，安全可靠
- 📱 **响应式** - 完美适配移动端和桌面端
- 🚀 **Docker** - 单镜像部署，一键启动

---

## 🏗️ 技术栈

### 前端
- Vue 3 + TypeScript
- Vite 5
- Pinia 状态管理
- Vue Router
- Element Plus UI

### 后端
- Go 1.23
- Gin Web 框架
- GORM ORM
- PostgreSQL 数据库
- Redis 缓存
- JWT 认证

---

## 🚀 快速开始

### 使用 Docker Compose（推荐）

```bash
# 1. 克隆项目
git clone <repository-url>
cd xiang

# 2. 配置环境变量
cp .env.example .env
# 编辑 .env 修改 JWT_SECRET 和数据库密码

# 3. 构建并启动
docker-compose up -d

# 4. 访问
# 前端：http://localhost:5700
# Swagger: http://localhost:5700/swagger
```

### 使用部署脚本

```bash
# 赋予执行权限
chmod +x deploy.sh

# 运行部署
./deploy.sh
```

### 本地开发

#### 后端
```bash
cd backend
go mod tidy
go run ./cmd/server
```

#### 前端
```bash
cd frontend
npm install
npm run dev
```

---

## 📁 项目结构

```
xiang/
├── backend/                     # Go 后端
│   ├── cmd/server/              # 程序入口
│   ├── internal/                # 私有包
│   │   ├── config/              # 配置
│   │   ├── controller/          # HTTP 处理层
│   │   ├── service/             # 业务逻辑
│   │   ├── repository/          # 数据访问
│   │   ├── model/               # 数据模型
│   │   ├── middleware/          # 中间件
│   │   ├── router/              # 路由
│   │   ├── dto/                 # 请求/响应
│   │   └── pkg/                 # 工具包
│   ├── uploads/                 # 文件上传
│   └── docs/                    # Swagger 文档
├── frontend/                    # Vue 前端
│   ├── src/
│   │   ├── api/                 # API 封装
│   │   ├── stores/              # Pinia
│   │   ├── router/              # 路由
│   │   ├── views/               # 页面
│   │   └── components/          # 组件
│   └── public/
├── docs/                        # 项目文档
├── docker-compose.yml           # Docker 编排
├── DEPLOY.md                    # 部署指南
└── README.md                    # 本文件
```

---

## 📋 API 接口

### 公开接口

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/api/v1/posts` | 文章列表 |
| GET | `/api/v1/posts/:id` | 文章详情 |
| GET | `/api/v1/tags` | 标签列表 |
| GET | `/api/v1/categories` | 分类列表 |
| GET | `/api/v1/site/settings` | 站点配置 |

### 用户接口

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | `/api/v1/user/register` | 用户注册 |
| POST | `/api/v1/user/login` | 用户登录 |
| GET | `/api/v1/user/me` | 获取用户信息 |
| GET | `/api/v1/posts/my` | 我的日记列表 |
| POST | `/api/v1/posts/my` | 创建日记 |
| PUT | `/api/v1/posts/my/:id` | 更新日记 |
| DELETE | `/api/v1/posts/my/:id` | 删除日记 |

### 管理员接口

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | `/api/v1/auth/login` | 管理员登录 |
| GET | `/api/v1/admin/posts` | 管理文章 |
| POST | `/api/v1/admin/upload/image` | 上传图片 |

详细 API 文档：http://localhost:8979/swagger

---

## 🎨 UI 设计

### 配色方案

| 颜色 | 色值 | 用途 |
|------|------|------|
| 主色 | `#ec4899` | 按钮、链接 |
| 辅色 | `#a855f7` | 渐变、装饰 |
| 背景 | `#fce7f3` → `#e9d5ff` | 页面渐变 |

### 动画效果

- 💓 心跳动画
- 💫 浮动爱心
- ✨ 加载动画
- 🌊 页面过渡

详细设计规范见 `docs/UI 设计规范.md`

---

## 📝 开发规范

### Git 分支

```
main (生产)
  └── develop (开发)
      ├── feature/xxx
      └── fix/xxx
```

### 提交规范

```
<type>(<scope>): <subject>
```

**Type**: `feat`, `fix`, `docs`, `style`, `refactor`, `test`, `chore`

**示例**:
```bash
feat(user): add user login API
fix(frontend): resolve token refresh issue
docs(api): update swagger auth endpoints
```

### 代码格式化

```bash
# Go
gofmt -w .
goimports -w .

# Vue
npm run format
```

---

## 🧪 测试

```bash
# 后端测试
cd backend
go test ./...

# 前端测试
cd frontend
npm run test
```

---

## 📦 部署

详见 [DEPLOY.md](./DEPLOY.md)

---

## 👥 默认账户

首次运行后，默认管理员账户：

| 角色 | 用户名 | 密码 |
|------|--------|------|
| 管理员 | `admin` | `myhx0308` |

**⚠️ 首次登录后请立即修改密码！**

---

## 📄 License

MIT License

---

## 💗 致谢

感谢以下开源项目：

- [Vue 3](https://vuejs.org/)
- [Element Plus](https://element-plus.org/)
- [Gin](https://gin-gonic.com/)
- [GORM](https://gorm.io/)

---

** built with 💕 by 湘遇团队 **
