# UI/UX 设计规范 - 恋爱博客

> 版本：1.0  
> 更新日期：2026-04-01

---

## 一、设计理念

**主题**：浪漫、温馨、有情感温度

**设计关键词**：
- 💕 浪漫 (Romantic)
- 🌸 温馨 (Warm)
- ✨ 精致 (Elegant)
- 💫 动感 (Dynamic)

---

## 二、配色方案

### 主色调

```
主色：#ec4899 (Pink-500) - 热情、浪漫
辅色：#a855f7 (Purple-500) - 优雅、神秘
背景：#fce7f3 (Pink-100) → #e9d5ff (Purple-100) 渐变
```

### 完整色板

| 类型 | 颜色 | 色值 | 用途 |
|------|------|------|------|
| **主色** | Pink-500 | `#ec4899` | 按钮、链接、强调 |
| **主色浅色** | Pink-300 | `#f472b6` | 悬停状态 |
| **主色深色** | Pink-700 | `#db2777` | 点击状态 |
| **辅色** | Purple-500 | `#a855f7` | 渐变、装饰 |
| **背景浅** | Pink-100 | `#fce7f3` | 页面背景起始 |
| **背景深** | Purple-100 | `#e9d5ff` | 页面背景结束 |
| **卡片背景** | White | `rgba(255,255,255,0.9)` | 卡片、导航 |
| **文字主** | Gray-800 | `#1f2937` | 主要文字 |
| **文字副** | Gray-500 | `#6b7280` | 次要文字 |
| **边框** | Pink-100 | `#f3e8e8` | 分割线、边框 |
| **错误** | Red-500 | `#ef4444` | 错误提示 |
| **成功** | Green-500 | `#22c55e` | 成功提示 |

### 渐变定义

```css
/* 主渐变 - 用于按钮、标题 */
gradient-primary: linear-gradient(135deg, #ec4899 0%, #a855f7 100%);

/* 背景渐变 - 用于页面背景 */
gradient-bg: linear-gradient(135deg, #fce7f3 0%, #f5d0fe 50%, #e9d5ff 100%);

/* 卡片渐变 - 用于悬停效果 */
gradient-card: linear-gradient(135deg, rgba(236,72,153,0.1) 0%, rgba(168,85,247,0.1) 100%);
```

---

## 三、字体排印

### 字体栈

```css
font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
```

### 字号层级

| 用途 | 字号 | 字重 | 行高 |
|------|------|------|------|
| 大标题 (H1) | 3rem (48px) | 700 | 1.2 |
| 标题 (H2) | 2rem (32px) | 600 | 1.3 |
| 小标题 (H3) | 1.25rem (20px) | 600 | 1.4 |
| 正文 | 1rem (16px) | 400 | 1.6 |
| 辅助文字 | 0.875rem (14px) | 400 | 1.5 |
| 标注 | 0.75rem (12px) | 400 | 1.4 |

---

## 四、组件样式

### 按钮

```css
/* 主按钮 */
.btn-primary {
  background: linear-gradient(135deg, #ec4899 0%, #a855f7 100%);
  color: white;
  padding: 1rem 2rem;
  border-radius: 12px;
  font-weight: 600;
  transition: all 0.3s ease;
}

.btn-primary:hover {
  transform: translateY(-2px);
  box-shadow: 0 10px 25px -5px rgba(236, 72, 153, 0.4);
}

.btn-primary:active {
  transform: translateY(0);
}
```

### 卡片

```css
.card {
  background: rgba(255, 255, 255, 0.9);
  backdrop-filter: blur(10px);
  border-radius: 16px;
  padding: 2rem;
  border: 1px solid rgba(236, 72, 153, 0.2);
  transition: all 0.3s ease;
}

.card:hover {
  transform: translateY(-5px);
  box-shadow: 0 20px 40px rgba(236, 72, 153, 0.2);
}
```

### 输入框

```css
.input {
  padding: 0.75rem 1rem;
  border: 2px solid #f3e8e8;
  border-radius: 12px;
  font-size: 1rem;
  background: #fafafa;
  transition: all 0.3s ease;
}

.input:focus {
  outline: none;
  border-color: #ec4899;
  background: white;
  box-shadow: 0 0 0 3px rgba(236, 72, 153, 0.1);
}
```

---

## 五、动画效果

### 心跳动画

```css
@keyframes heart-pulse {
  0%, 100% { transform: scale(1); }
  50% { transform: scale(1.2); }
}

.heart {
  animation: heart-pulse 1.5s infinite;
}
```

### 浮动爱心

```css
@keyframes float-up {
  0% {
    transform: translateY(0) rotate(0deg);
    opacity: 0;
  }
  10% { opacity: 1; }
  90% { opacity: 1; }
  100% {
    transform: translateY(-100vh) rotate(360deg);
    opacity: 0;
  }
}
```

### 加载动画

```css
@keyframes bounce {
  0%, 80%, 100% { transform: scale(0); }
  40% { transform: scale(1); }
}
```

### 页面过渡

```css
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
```

---

## 六、图标系统

### Emoji 图标使用

| 场景 | 图标 | 说明 |
|------|------|------|
| 爱/喜欢 | 💕 💗 ❤️ 💖 | 情感表达 |
| 日记 | 📝 📖 ✏️ | 内容创作 |
| 标签 | 🏷️ 🔖 | 分类标记 |
| 图片 | 📸 🖼️ | 媒体资源 |
| 时间 | 📅 ⏰ | 时间相关 |
| 登录 | 🔐 👤 | 认证相关 |

### 图标尺寸

| 用途 | 尺寸 |
|------|------|
| 大图标 | 3rem (48px) |
| 中图标 | 1.5rem (24px) |
| 小图标 | 1rem (16px) |

---

## 七、交互反馈

### 悬停 (Hover)
- 按钮：上移 2px + 阴影增强
- 卡片：上移 5px + 阴影增强 + 边框变色
- 链接：颜色加深

### 点击 (Active)
- 按钮：恢复原位（取消上移）
- 所有可点击元素：0.1s 内完成反馈

### 禁用 (Disabled)
- 透明度：0.7
- 光标：not-allowed
- 无悬停效果

### 加载 (Loading)
- 按钮：显示三点加载动画
- 页面：显示骨架屏
- 提交：禁用按钮防止重复提交

---

## 八、响应式断点

| 断点 | 宽度 | 适配设备 |
|------|------|----------|
| Mobile | < 768px | 手机竖屏 |
| Tablet | 768px - 1024px | 平板 |
| Desktop | > 1024px | 桌面 |

### 移动端适配要点

1. 导航栏：汉堡菜单
2. 卡片：单列布局
3. 字体：适当缩小
4. 按钮：全宽或大尺寸
5. 输入框：避免弹出键盘遮挡

---

## 九、无障碍设计

### 色彩对比度

- 文字与背景对比度至少 4.5:1
- 重要信息不单独使用颜色传达

### 键盘导航

- 所有交互元素可通过 Tab 键访问
- 焦点状态清晰可见

### 屏幕阅读器

- 所有图标和图像有 alt 文本
- 表单输入有对应 label

---

## 十、设计资源

### 灵感来源
- [Dribbble - Love Theme](https://dribbble.com/search/love)
- [Pinterest - Romantic Design](https://pinterest.com/search/romantic%20web%20design)

### 工具推荐
- Coolors.co - 配色方案生成
- Figma - 界面设计
- Lottie - 动画资源

---

**文档结束**
