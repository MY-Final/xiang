# 💕 湘遇恋爱博客 - UI 改造说明

> 毛玻璃风格 (Glassmorphism) 全面升级

---

## 🎨 设计系统

### 主题文件
- `frontend/src/assets/theme.css` - 完整的设计系统

### 核心特性

#### 1. 毛玻璃效果 (Glassmorphism)
```css
--glass-bg-light: rgba(255, 255, 255, 0.75);
--glass-bg-medium: rgba(255, 255, 255, 0.6);
--glass-bg-dark: rgba(255, 255, 255, 0.4);
--blur-sm: 4px;
--blur-md: 12px;
--blur-lg: 20px;
--blur-xl: 40px;
```

#### 2. 浪漫配色
- 主色：`#ec4899` (Pink-500)
- 辅色：`#a855f7` (Purple-500)
- 背景渐变：`#fdf2f8 → #f5d0fe → #e9d5ff`

#### 3. 动画效果
- 心跳脉冲 (heart-pulse)
- 浮动爱心 (float-up)
- 光晕脉冲 (glow-pulse)
- 淡入/滑入 (fadeIn/slideUp)

---

## 📄 页面列表

### 1. 首页 (`/`)
**瀑布流布局**
- Masonry 网格展示文章卡片
- 毛玻璃 Hero 区域
- 浮动爱心背景
- 响应式：3 列 → 2 列 → 1 列

**卡片特性**
- 毛玻璃效果
- 悬停上浮动画
- 图片缩放效果
- 标签展示

### 2. 登录页 (`/login`)
**设计特点**
- 居中毛玻璃卡片
- SVG Logo 动画
- 浮动爱心背景
- 渐变文字标题
- 加载动画

### 3. 注册页 (`/register`)
**设计特点**
- 输入框图标
- 毛玻璃表单
- 光晕装饰
- 加载点动画

### 4. 文章列表页 (`/posts`)
**设计特点**
- 毛玻璃页面头部
- 网格卡片布局
- 骨架屏加载
- 悬停效果

### 5. 文章详情页 (`/post/:id`)
**设计特点**
- 封面图展示
- 毛玻璃头部
- 侧边目录
- 毛玻璃内容区
- 底部感谢区域

---

## 🎯 关键类名

### 毛玻璃类
```css
.glass              /* 基础毛玻璃 */
.glass-strong       /* 强毛玻璃 */
.glass-light        /* 轻毛玻璃 */
.glass-dark         /* 暗色毛玻璃 */
```

### 动画类
```css
.animate-heart-pulse    /* 心跳动画 */
.animate-float          /* 浮动动画 */
.animate-float-up       /* 上升爱心 */
.animate-shimmer        /* 闪烁效果 */
.animate-fade-in        /* 淡入 */
.animate-slide-up       /* 滑入上 */
.animate-glow           /* 发光 */
```

### 布局类
```css
.waterfall          /* 瀑布流布局 */
.grid               /* 网格布局 */
.flex               /* Flex 布局 */
```

---

## 📱 响应式断点

| 断点 | 宽度 | 布局 |
|------|------|------|
| 移动端 | < 640px | 单列 |
| 平板 | 641px - 1024px | 双列 |
| 桌面 | > 1025px | 三列 |

---

## 🎨 使用示例

### 创建毛玻璃卡片
```vue
<div class="glass-strong" style="border-radius: 24px; padding: 2rem;">
  <h2 class="text-gradient">标题</h2>
  <p>内容</p>
</div>
```

### 添加浮动爱心
```vue
<div class="bg-hearts">
  <div class="heart" v-for="i in 10" :key="i">
    <svg viewBox="0 0 24 24" fill="currentColor">
      <path d="M12 21.35l-1.45-1.32..."/>
    </svg>
  </div>
</div>
```

### 使用渐变文字
```vue
<h1 class="text-gradient">湘遇恋爱博客</h1>
```

---

## ✅ 检查清单

所有页面已完成：
- [x] 毛玻璃效果应用
- [x] SVG 图标替代 emoji
- [x] 浮动爱心背景
- [x] 响应式布局
- [x] 悬停动画
- [x] 加载状态
- [x] 光晕装饰
- [x] 渐变色彩

---

## 🚀 部署后查看效果

```bash
# 服务器部署
docker-compose up -d

# 访问
http://服务器 IP:5700
```

**默认账户**
- 用户名：`admin`
- 密码：`myhx0308`

---

**💗 享受全新的毛玻璃风格！**
