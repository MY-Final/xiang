# AGENTS 工作指南

本文档面向在 `C:\Users\34861\Desktop\xiang` 仓库内工作的智能编码代理。
在开始任何改动前，请先阅读并遵循本指南。

## 1. 仓库现状

- 当前仓库处于启动阶段，代码主体尚未完整落地。
- 现有内容以文档为主（如 `README.md`、`docs/架构指南.md`）。
- 目标架构已确定为：
  - 前端：Vue 3 + TypeScript + Vite + Element Plus + Pinia + Vue Router + Axios
  - 后端：Go + Gin + GORM + JWT
  - 数据库：PostgreSQL
  - 前端开发端口：`6100`
  - 后端开发端口：`8979`

## 2. 规则来源检查（必须关注）

已检查以下规则文件：

- `.cursorrules`：未发现
- `.cursor/rules/`：未发现
- `.github/copilot-instructions.md`：未发现

若后续新增以上任意规则文件，需将其视为高优先级约束，并及时同步更新本 `AGENTS.md`。

## 3. 代理工作原则

- 只做最小且聚焦的改动，避免无关重构。
- 严格保持前后端目录隔离（`frontend/`、`backend/`）。
- 不得引入与 `docs/架构指南.md` 冲突的架构方案。
- 优先选择清晰、可维护、可测试的实现。
- 新增业务逻辑时，若测试框架已就绪，必须补充对应测试。
- 修复缺陷时，优先补一个可复现该问题的测试用例。

## 4. 构建 / Lint / 测试命令

说明：当前仓库尚未提交完整可运行脚手架。以下命令为项目落地后的标准预期；若仓库实际脚本不同，以实际脚本为准。

### 4.1 前端（目录：`frontend/`）

- 安装依赖：`npm install`
- 启动开发环境：`npm run dev -- --port 6100`
- 生产构建：`npm run build`
- 预览构建产物：`npm run preview -- --port 6100`
- 代码检查：`npm run lint`
- 类型检查：`npm run typecheck`
- 单元测试（全部）：`npm run test:unit`
- 单元测试（单文件）：`npm run test:unit -- src/components/YourComponent.test.ts`
- 单元测试（按测试名）：`npm run test:unit -- -t "renders title"`
- E2E 测试（全部）：`npm run test:e2e`
- E2E 测试（单文件）：`npm run test:e2e -- tests/e2e/login.spec.ts`

### 4.2 后端（目录：`backend/`）

- 拉取/整理依赖：`go mod tidy`
- 本地运行服务：`go run ./cmd/server`
- 构建二进制：`go build ./cmd/server`
- 格式化代码：`gofmt -w .`
- 静态检查：`go vet ./...`
- 测试（全部）：`go test ./...`
- 测试（单包）：`go test ./internal/service`
- 测试（单用例）：`go test ./internal/service -run TestCreatePost`
- 竞态测试：`go test ./... -race`
- 覆盖率：`go test ./... -cover`

## 5. 命令发现规则（执行前先核实）

在运行任何命令之前，先检查仓库中的真实脚本与配置：

- 前端脚本：`frontend/package.json`
- 后端命令：`backend/Makefile`、`backend/go.mod`、`backend/README.md`
- CI 默认流程：`.github/workflows/*.yml`

若实际命令与本指南不一致：

- 执行时以仓库实际脚本为准。
- 完成后同步更新 `AGENTS.md`，避免文档与实现脱节。

## 6. 代码风格总则

- 文件编码使用 UTF-8。
- 在不影响语义时，优先使用简洁 ASCII 文本。
- 函数保持短小，单一职责，避免隐藏副作用。
- 变更保持最小范围，不进行大面积格式化噪声提交。
- 注释仅用于解释非显而易见的意图、约束或不变量。

## 7. 前端规范（Vue 3 + TypeScript）

- 全面使用 TypeScript（`.ts`、`<script setup lang="ts">`）。
- 首选 Composition API 与 `script setup`。
- `defineProps` / `defineEmits` 使用显式类型。
- 禁止滥用 `any`，优先 `interface` / `type` / 泛型。
- 页面组件保持轻量，复杂逻辑下沉至 composables 或 stores。
- 全局状态统一由 Pinia 管理，避免临时全局单例。
- API 调用统一收敛在 `src/api/`。
- 错误处理要统一（提示策略 + 类型化错误映射）。
- UI 组件优先采用 Element Plus，风格保持一致。

### 7.1 前端导入顺序

- 导入分组顺序：Vue/核心库 -> 第三方库 -> 别名导入（`@/`）-> 相对路径。
- 组之间保留一个空行。
- 工具函数与 composables 优先使用具名导出。

### 7.2 前端命名约定

- 组件：PascalCase（如 `PostCard.vue`）。
- 组合式函数：`useXxx`（如 `useAuth.ts`）。
- Store：`useXxxStore`（如 `usePostStore.ts`）。
- API 模块：按领域命名（如 `postApi.ts`、`authApi.ts`）。
- 路由名称：小写 kebab-case（按项目约定定义常量）。

### 7.3 前端格式化约定

- 遵循 ESLint + Prettier。
- 不允许未使用的导入与变量。
- 避免模板中出现过深嵌套逻辑，必要时提取为 `computed`。

## 8. 后端规范（Go + Gin）

- 采用分层结构：`controller`、`service`、`repository`。
- Handler 仅处理协议层，业务逻辑放 `service`。
- 数据访问细节放 `repository`，避免泄漏到上层。
- 在 service/repository 边界尽量传递 `context.Context`。
- DTO 边界必须做参数校验，不信任客户端输入。
- 使用构造函数做依赖注入，避免包级全局状态。
- 预期内错误不使用 `panic`。

### 8.1 后端导入约定

- 保持 `goimports` / `gofmt` 兼容顺序。
- 顺序为：标准库 -> 第三方库 -> 项目内部包。

### 8.2 后端命名约定

- 导出标识符：PascalCase。
- 非导出标识符：camelCase。
- 接口按行为命名（如 `PostRepository`、`TokenSigner`）。
- 避免命名重复累赘（避免 stutter）。

### 8.3 后端错误处理

- 错误需带上下文包装：`fmt.Errorf("create post: %w", err)`。
- 不向客户端直接暴露内部数据库错误。
- 若项目定义统一响应结构，所有接口保持一致。
- 日志应包含关键上下文（request id、路由、操作者信息等）。

## 9. API 与契约规则

- API 基础前缀固定为 `/api/v1`。
- 公开接口不得依赖管理员 JWT。
- 管理接口统一放在 `/api/v1/admin/*` 并要求 JWT。
- 请求/响应 DTO 必须显式定义，避免隐式字段漂移。
- 若发生响应结构变更，必须同步更新文档与调用方。

## 10. 测试要求

- 后端新增业务逻辑优先补表驱动单元测试。
- 前端新增 composable/store 逻辑补单元测试。
- 关键链路（登录、发文、展示）至少覆盖一条集成或 E2E。
- 修复 bug 时，先补失败用例，再修复实现。

## 11. 变更一致性要求

- 当端口、路由、环境变量、命令发生变化时：文档与代码必须同时更新。
- 禁止提交密钥、令牌、数据库导出、`.env` 等敏感信息。
- 提交粒度应保持单一关注点，便于审查与回滚。

## 12. Git 提交规范

- 默认在功能分支开发，避免直接在 `main` 上堆叠大量改动。
- 每次提交只做一件事（单一关注点），避免把无关改动混入同一个 commit。
- 提交前至少执行与改动相关的最小验证（如 lint、单测或构建中的一项）。
- 提交信息建议使用：`<type>(<scope>): <subject>`。
- `type` 建议值：`feat`、`fix`、`refactor`、`docs`、`test`、`chore`、`build`、`ci`。
- `subject` 使用简明祈使句，聚焦“为什么改”，避免冗长描述。
- 示例：`feat(auth): 增加管理员登录 token 刷新逻辑`。
- 若改动影响接口、端口、命令或环境变量，提交中必须包含文档同步更新。
- 禁止提交敏感信息（`.env`、密钥、令牌、数据库导出、私有证书）。
- 非明确要求下，不做 `--amend`、`push --force`、`reset --hard` 等高风险操作。
- 推送前确认 `git status` 干净且提交内容可回滚、可审查。

### 12.1 标准提交模板示例

- `feat(post): 支持文章草稿自动保存`
- `fix(auth): 修复 token 过期后重复跳转登录`
- `docs(api): 同步更新管理员文章接口说明`
- `test(service): 为文章发布流程补充表驱动测试`
- `refactor(router): 拆分路由注册并减少重复逻辑`
- `chore(ci): 调整前后端流水线缓存策略`

## 13. 不确定时怎么做

- 以 `docs/架构指南.md` 为最高架构参考。
- 选择满足需求的最简单实现，不做过度设计。
- 遇到非显而易见取舍，在提交说明中简要记录原因。
