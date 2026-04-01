# CLAUDE.md

本文档为 Claude Code (claude.ai/code) 在本仓库工作提供指导。

## 项目概述

"湘"遇恋爱博客系统，包含前台站点和后台管理面板。

- **前端**：Vue 3 + TypeScript + Vite + Element Plus + Pinia（规划中，尚未实现）
- **后端**：Go + Gin + GORM + JWT + PostgreSQL
- **端口**：前端 `6100`，后端 `8979`
- **API 前缀**：`/api/v1`

---

## 项目结构

```
xiang/
├── backend/                     # Go 后端
│   ├── cmd/server/              # 程序入口
│   │   ├── main.go              # 启动入口
│   │   └── docs.go              # Swagger 注释
│   ├── internal/                # 私有包（不可外部引用）
│   │   ├── config/              # 配置加载
│   │   ├── controller/          # HTTP 处理层
│   │   ├── service/             # 业务逻辑层
│   │   ├── repository/          # 数据访问层 (GORM)
│   │   ├── model/               # 数据模型
│   │   ├── middleware/          # 中间件 (JWT, CORS)
│   │   ├── router/              # 路由注册
│   │   ├── dto/                 # 请求/响应结构体
│   │   └── pkg/                 # 通用工具包
│   │       ├── response/        # 统一响应封装
│   │       ├── apperror/        # 错误码定义
│   │       └── jwt/             # JWT 工具
│   ├── docs/                    # Swagger 文档
│   ├── uploads/                 # 本地文件存储
│   ├── migrations/              # SQL 迁移脚本
│   ├── go.mod
│   └── README.md
├── frontend/                    # Vue 前端（待实现）
├── docs/                        # 项目文档
│   ├── 架构指南.md
│   └── 后端接口文档.md
├── .env.example                 # 环境变量模板
├── AGENTS.md                    # 代理工作指南
└── README.md
```

---

## 后端命令

```bash
cd backend

# 依赖管理
go mod tidy

# 运行服务
go run ./cmd/server

# 构建
go build ./cmd/server

# 格式化代码
gofmt -w .

# 静态检查
go vet ./...

# 运行所有测试
go test ./...

# 运行单个包的测试
go test ./internal/service

# 运行单个测试函数
go test ./internal/service -run TestCreatePost

# 竞态检测
go test ./... -race

# 测试覆盖率
go test ./... -cover

# Swagger 文档生成
swag init -g cmd/server/main.go -o docs
```

---

## Go 开发规范

### 代码风格

- **命名**：
  - 导出标识符：`PascalCase`（如 `PostService`）
  - 非导出标识符：`camelCase`（如 `postRepo`）
  - 接口：按行为命名（如 `PostRepository`）
  - 避免命名重复（如 `pkg/pkg`）

- **导入顺序**：
  1. 标准库
  2. 第三方库
  3. 项目内部包
  
  组之间保留空行，使用 `goimports` 自动格式化。

- **错误处理**：
  ```go
  // 必须带上下文包装错误
  if err != nil {
      return fmt.Errorf("create post: %w", err)
  }
  
  // 不向客户端暴露内部错误细节
  ```

- **注释规范**：
  - 导出的函数/类型必须写文档注释
  - 注释以标识符开头：`// NewPostService 创建文章服务实例`
  - 不解释显而易见的代码

### 分层约束

```
controller → service → repository → model
     ↓           ↓           ↓
   router    config       pkg
```

- **Controller**：仅处理 HTTP 协议层（参数解析、响应封装），不写业务逻辑
- **Service**：业务逻辑、参数校验、事务控制、调用 Repository
- **Repository**：纯数据库 CRUD，不依赖 Service
- **禁止循环依赖**：internal 包只能单向依赖

### 依赖注入

使用构造函数注入，避免包级全局状态：

```go
// ✅ 正确
postRepo := repository.NewPostRepository(db)
postService := service.NewPostService(postRepo)

// ❌ 错误：全局变量
var globalService *service.PostService
```

### 上下文传递

Service 和 Repository 层应传递 `context.Context`：

```go
func (s *postService) Create(ctx context.Context, dto *dto.CreatePostReq) (*Post, error)
func (r *postRepository) FindByID(ctx context.Context, id uint) (*Post, error)
```

### DTO 规范

- 请求和响应必须显式定义 DTO 结构体
- DTO 放在 `internal/dto/` 目录
- 使用 `json` 标签控制序列化
- 敏感字段（如密码）使用 `json:"-"` 排除

```go
type CreatePostReq struct {
    Title           string `json:"title" binding:"required"`
    Slug            string `json:"slug" binding:"required"`
    Summary         string `json:"summary"`
    ContentMarkdown string `json:"content_markdown"`
    CoverImage      string `json:"cover_image"`
    Status          string `json:"status"`
    IsTop           bool   `json:"is_top"`
    CategoryID      *uint  `json:"category_id"`
    TagIDs          []uint `json:"tag_ids"`
}
```

### 统一响应格式

使用 `internal/pkg/response` 封装：

```go
response.Success(c, data)
response.Fail(c, http.StatusBadRequest, apperror.CodeBadRequest, "参数错误")
```

响应结构：
```json
{"code": 0, "message": "success", "data": {...}}
```

错误码定义在 `internal/pkg/apperror/code.go`：
- `CodeOK = 0`
- `CodeBadRequest = 4001`
- `CodeUnauthorized = 4010`
- `CodeNotFound = 4004`
- `CodeConflict = 4009`
- `CodeInternal = 5000`

---

## 数据模型

### 模型列表 (`internal/model/models.go`)

| 模型 | 说明 | 软删除 |
|------|------|--------|
| `Admin` | 管理员账户 | 否 |
| `Post` | 文章（草稿/发布/私密） | 是 |
| `Category` | 分类 | 是 |
| `Tag` | 标签 | 是 |
| `PostTag` | 文章 - 标签多对多关联 | 否 |
| `Media` | 媒体资源 | 否 |
| `SiteSetting` | 站点配置 | 否 |

### 关键字段约定

- 所有模型使用 `ID` 作为主键
- 时间字段：`CreatedAt`, `UpdatedAt`, `DeletedAt`（GORM 自动管理）
- 外键使用 `ID` 结尾：`CategoryID`
- 关联使用指针：`Category *Category`

---

## API 接口

### 公开接口（无需认证）

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/api/v1/posts` | 文章列表 |
| GET | `/api/v1/posts/:id` | 按 ID 获取文章 |
| GET | `/api/v1/posts/slug/:slug` | 按 slug 获取文章 |
| GET | `/api/v1/tags` | 标签列表 |
| GET | `/api/v1/categories` | 分类列表 |
| GET | `/api/v1/timeline` | 时间线 |
| GET | `/api/v1/site/settings` | 站点配置 |

### 认证接口

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | `/api/v1/auth/login` | 管理员登录 |
| GET | `/api/v1/auth/me` | 获取当前管理员信息 |
| POST | `/api/v1/auth/logout` | 退出登录 |

### 管理接口（需 JWT）

所有 `/api/v1/admin/*` 路径需要 `Authorization: Bearer <token>` 请求头。

**文章管理**
- `GET/POST/PUT/DELETE /admin/posts`
- `GET /admin/posts/:id`
- `GET /admin/posts/check-slug`
- `PATCH /admin/posts/:id/publish|unpublish|top`

**标签管理**
- `GET/POST/PUT/DELETE /admin/tags`

**分类管理**
- `GET/POST/PUT/DELETE /admin/categories`

**媒体管理**
- `POST /admin/upload/image|images`
- `GET/DELETE /admin/media/:id`

**站点配置**
- `GET/PUT /admin/site/settings`

**仪表盘**
- `GET /admin/dashboard`

---

## 环境配置

复制 `.env.example` 到 `.env`：

```env
# 应用
APP_ENV=development
APP_PORT=8979

# PostgreSQL（二选一）
# 方式 1：使用各字段
DB_HOST=127.0.0.1
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=change_me
DB_NAME=love_blog
DB_SSLMODE=disable

# 方式 2：使用完整 DSN
DATABASE_URL=postgresql://user:pass@host:5432/dbname?sslmode=disable

# JWT
JWT_SECRET=change_me_to_a_long_random_string
JWT_EXPIRE_HOURS=72

# 文件上传
UPLOAD_DIR=uploads
```

### 初始账户

首次运行后，默认管理员账户：
- 用户名：`admin`
- 密码：`myhx0308`

---

## 实现模式

### Repository 模式

每个模型对应一个 Repository：

```go
// repository/post_repository.go
type PostRepository interface {
    FindByID(ctx context.Context, id uint) (*Post, error)
    FindBySlug(ctx context.Context, slug string) (*Post, error)
    Create(ctx context.Context, post *Post) error
    Update(ctx context.Context, post *Post) error
    Delete(ctx context.Context, id uint) error
    List(ctx context.Context, filter ListFilter) ([]*Post, int64, error)
}
```

### 软删除

带有 `gorm.DeletedAt` 字段的模型使用软删除：

```go
// 查询时自动过滤已删除记录
db.Where("status = ?", "published").Find(&posts)

// 强制包含已删除记录
db.Unscoped().Where("id = ?", id).Find(&post)
```

### JWT 认证

中间件验证 Token，失败返回 401：

```go
// router.go
admin := api.Group("/admin")
admin.Use(middleware.JWT(cfg.JWTSecret))  // 鉴权
{
    admin.GET("/posts", postCtl.ListAdmin)
}
```

Token 解析后存入 Context：
```go
// controller 中获取
claims, exists := c.Get("claims")
```

---

## Git 分支与工作流规范

### 分支命名

| 分支类型 | 格式 | 示例 |
|----------|------|------|
| 新功能 | `feature/功能英文描述` | `feature/post-love-diary` |
| Bug 修复 | `fix/问题描述` | `fix/login-token-expired` |
| 格式化 | `style/formatting` | 或直接在当前分支处理 |

### 工作流程

1. **从 develop 创建分支**：所有新功能必须基于 `develop` 分支
2. **原子提交**：完成一个功能点或明显改进后必须 commit
3. **合并审查**：合并到 `develop` 前需经 QA Tester 或 Team Lead 审查
4. **禁止行为**：
   - 禁止直接向 `main`/`develop` 推送大量未 commit 的修改
   - 禁止在功能分支中累积多个无关改动

### 提交规范

```
<type>(<scope>): <subject>
```

**Type 取值**：`feat`, `fix`, `refactor`, `docs`, `test`, `chore`, `build`, `ci`

**示例**：
- `feat(post): 支持文章草稿自动保存`
- `fix(auth): 修复 token 过期后重复跳转登录`
- `docs(api): 同步更新管理员文章接口说明`

---

## 相关文档

- `docs/架构指南.md` - 完整架构设计和技术选型
- `docs/后端接口文档.md` - 详细 API 接口定义
- `backend/README.md` - 后端快速启动指南
- `AGENTS.md` - 代理工作规范和代码风格要求

---

## 当前状态

**后端**：已完成文章、标签、分类、媒体、站点配置的完整 CRUD，支持 JWT 认证和 Swagger 文档。

**前端**：尚未实现。

**数据库**：PostgreSQL，启动时自动执行 AutoMigrate。
