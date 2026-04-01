<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { useRoute, RouterLink } from 'vue-router'

const route = useRoute()
const post = ref<any>(null)
const loading = ref(true)

// 模拟目录
const tableOfContents = ref<Array<{ level: number; text: string; href: string }>>([])

onMounted(async () => {
  loading.value = true
  try {
    // TODO: 调用 API 获取文章详情
    // post.value = await api.get(`/posts/${route.params.id}`)

    // 模拟数据
    post.value = {
      id: route.params.id,
      title: '我们的第一次相遇',
      content_markdown: `
# 我们的第一次相遇

那是一个阳光明媚的下午，咖啡厅的风铃轻轻作响...

## 等待

我在咖啡馆里等你，点了一杯拿铁，看着窗外的行人。时间过得很快，不知不觉已经过了约定的时间。

## 你出现了

然后你出现了，带着一丝歉意的微笑，手里拿着一本书。那一刻，我知道我们的故事开始了。

## 相谈甚欢

我们聊了很多，从喜欢的书到旅行过的城市，从童年趣事到未来的梦想。时间仿佛静止了，整个世界只剩下我们两个人。

## 分别时的约定

分别时，我们约好下次再见。那次的约会，成为了我们美好回忆的开始。
      `,
      published_at: '2024-02-14',
      tags: [{ name: '相遇' }, { name: '浪漫' }, { name: '初恋' }],
      cover_image: 'https://picsum.photos/800/400?random=1',
    }

    // 生成目录
    generateTableOfContents()
  } catch (error) {
    console.error('加载文章失败:', error)
  } finally {
    loading.value = false
  }
})

const generateTableOfContents = () => {
  // 简单模拟，实际应该解析 markdown
  tableOfContents.value = [
    { level: 2, text: '等待', href: '#wait' },
    { level: 2, text: '你出现了', href: '#appear' },
    { level: 2, text: '相谈甚欢', href: '#talk' },
    { level: 2, text: '分别时的约定', href: '#promise' },
  ]
}

const formattedDate = computed(() => {
  if (!post.value?.published_at) return ''
  return new Date(post.value.published_at).toLocaleDateString('zh-CN', {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
  })
})
</script>

<template>
  <div class="post-detail">
    <!-- 背景装饰 -->
    <div class="bg-hearts">
      <div class="heart" v-for="i in 10" :key="i" :style="{
        left: Math.random() * 100 + '%',
        animationDelay: Math.random() * 5 + 's',
        animationDuration: (Math.random() * 3 + 5) + 's',
      }">
        <svg viewBox="0 0 24 24" fill="currentColor">
          <path d="M12 21.35l-1.45-1.32C5.4 15.36 2 12.28 2 8.5 2 5.42 4.42 3 7.5 3c1.74 0 3.41.81 4.5 2.09C13.09 3.81 14.76 3 16.5 3 19.58 3 22 5.42 22 8.5c0 3.78-3.4 6.86-8.55 11.54L12 21.35z"/>
        </svg>
      </div>
    </div>

    <div v-if="loading" class="loading">
      <div class="loading-glass">
        <div class="loading-header">
          <div class="loading-line full"></div>
          <div class="loading-line short"></div>
          <div class="loading-meta"></div>
        </div>
        <div class="loading-content">
          <div class="loading-line full"></div>
          <div class="loading-line full"></div>
          <div class="loading-line full"></div>
          <div class="loading-line medium"></div>
        </div>
      </div>
    </div>

    <article v-else-if="post" class="article">
      <!-- 封面图 -->
      <div class="cover-image-wrapper" v-if="post.cover_image">
        <img :src="post.cover_image" :alt="post.title" class="cover-image" />
        <div class="cover-overlay"></div>
      </div>

      <!-- 文章头部 -->
      <header class="article-header">
        <div class="header-glass">
          <div class="article-tags">
            <span class="tag" v-for="tag in post.tags" :key="tag.name">
              {{ tag.name }}
            </span>
          </div>
          <h1 class="article-title">{{ post.title }}</h1>
          <div class="article-meta">
            <svg class="meta-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <rect x="3" y="4" width="18" height="18" rx="2" ry="2"/>
              <line x1="16" y1="2" x2="16" y2="6"/>
              <line x1="8" y1="2" x2="8" y2="6"/>
              <line x1="3" y1="10" x2="21" y2="10"/>
            </svg>
            <span>{{ formattedDate }}</span>
          </div>
        </div>
      </header>

      <!-- 目录（如果有） -->
      <aside class="toc-aside" v-if="tableOfContents.length > 0">
        <div class="toc-glass">
          <h3 class="toc-title">
            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M4 6h16M4 10h16M4 14h16M4 18h16"/>
            </svg>
            目录
          </h3>
          <nav class="toc-nav">
            <a
              v-for="item in tableOfContents"
              :key="item.href"
              :href="item.href"
              class="toc-link"
              :class="'level-' + item.level"
            >
              {{ item.text }}
            </a>
          </nav>
        </div>
      </aside>

      <!-- 文章内容 -->
      <div class="article-content">
        <div class="content-wrapper" v-html="post.content_markdown"></div>
      </div>

      <!-- 文章底部 -->
      <footer class="article-footer">
        <div class="footer-glass">
          <div class="footer-heart">
            <svg viewBox="0 0 24 24" fill="currentColor">
              <path d="M12 21.35l-1.45-1.32C5.4 15.36 2 12.28 2 8.5 2 5.42 4.42 3 7.5 3c1.74 0 3.41.81 4.5 2.09C13.09 3.81 14.76 3 16.5 3 19.58 3 22 5.42 22 8.5c0 3.78-3.4 6.86-8.55 11.54L12 21.35z"/>
            </svg>
          </div>
          <p>感谢你的阅读</p>
          <RouterLink to="/" class="back-home">
            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M19 12H5M12 19l-7-7 7-7"/>
            </svg>
            <span>返回首页</span>
          </RouterLink>
        </div>
      </footer>
    </article>

    <div v-else class="not-found">
      <div class="not-found-glass">
        <svg class="not-found-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5">
          <path d="M9.75 9.75l4.5 4.5m0-4.5l-4.5 4.5M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/>
        </svg>
        <h2>文章不存在</h2>
        <p>抱歉，这篇文章不存在或已被删除</p>
        <RouterLink to="/" class="back-home-btn">
          <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6"/>
          </svg>
          返回首页
        </RouterLink>
      </div>
    </div>
  </div>
</template>

<style scoped>
.post-detail {
  min-height: 100vh;
  position: relative;
  padding: 2rem;
}

/* ═══════════════════════════════════════════════════════════
   背景爱心动画
   ═══════════════════════════════════════════════════════════ */
.bg-hearts {
  position: fixed;
  inset: 0;
  pointer-events: none;
  z-index: 0;
  overflow: hidden;
}

.heart {
  position: absolute;
  bottom: -50px;
  color: rgba(236, 72, 153, 0.15);
  animation: float-up linear infinite;
}

@keyframes float-up {
  0% {
    transform: translateY(0) rotate(0deg) scale(0.5);
    opacity: 0;
  }
  10% {
    opacity: 0.5;
  }
  50% {
    opacity: 0.8;
  }
  90% {
    opacity: 0.5;
  }
  100% {
    transform: translateY(-100vh) rotate(360deg) scale(1);
    opacity: 0;
  }
}

/* ═══════════════════════════════════════════════════════════
   加载状态
   ═══════════════════════════════════════════════════════════ */
.loading {
  position: relative;
  z-index: 1;
  display: flex;
  justify-content: center;
  padding: 3rem 1rem;
}

.loading-glass {
  background: rgba(255, 255, 255, 0.65);
  backdrop-filter: blur(20px);
  -webkit-backdrop-filter: blur(20px);
  border-radius: 24px;
  padding: 3rem;
  max-width: 800px;
  width: 100%;
  border: 1px solid rgba(255, 255, 255, 0.5);
  box-shadow: 0 12px 40px rgba(236, 72, 153, 0.1);
}

.loading-header {
  margin-bottom: 2rem;
  padding-bottom: 2rem;
  border-bottom: 1px solid rgba(236, 72, 153, 0.1);
}

.loading-line {
  height: 24px;
  background: linear-gradient(90deg,
    rgba(252, 231, 243, 0.8) 0%,
    rgba(245, 208, 254, 0.8) 50%,
    rgba(252, 231, 243, 0.8) 100%);
  background-size: 200% 100%;
  animation: shimmer 1.5s infinite;
  border-radius: 8px;
  margin-bottom: 1rem;
}

.loading-line.full { width: 100%; }
.loading-line.short { width: 60%; }
.loading-line.medium { width: 80%; }

.loading-meta {
  height: 16px;
  width: 40%;
  background: linear-gradient(90deg,
    rgba(252, 231, 243, 0.8) 0%,
    rgba(245, 208, 254, 0.8) 50%,
    rgba(252, 231, 243, 0.8) 100%);
  background-size: 200% 100%;
  animation: shimmer 1.5s infinite;
  border-radius: 8px;
}

.loading-content .loading-line {
  margin-bottom: 0.75rem;
}

@keyframes shimmer {
  0% { background-position: 200% 0; }
  100% { background-position: -200% 0; }
}

/* ═══════════════════════════════════════════════════════════
   文章容器
   ═══════════════════════════════════════════════════════════ */
.article {
  position: relative;
  z-index: 1;
  max-width: 800px;
  margin: 0 auto;
}

/* 封面图 */
.cover-image-wrapper {
  position: relative;
  width: 100%;
  height: 320px;
  margin-bottom: 2rem;
  border-radius: 24px;
  overflow: hidden;
  box-shadow: 0 12px 40px rgba(236, 72, 153, 0.15);
}

.cover-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.cover-overlay {
  position: absolute;
  inset: 0;
  background: linear-gradient(to bottom, transparent 0%, rgba(236, 72, 153, 0.15) 100%);
  pointer-events: none;
}

/* 文章头部 */
.article-header {
  margin-bottom: 2rem;
  display: flex;
  justify-content: center;
}

.header-glass {
  background: rgba(255, 255, 255, 0.7);
  backdrop-filter: blur(20px);
  -webkit-backdrop-filter: blur(20px);
  border-radius: 24px;
  padding: 2.5rem 3rem;
  text-align: center;
  border: 1px solid rgba(255, 255, 255, 0.6);
  box-shadow: 0 8px 32px rgba(236, 72, 153, 0.1);
  width: 100%;
}

.article-tags {
  display: flex;
  justify-content: center;
  gap: 0.5rem;
  margin-bottom: 1.25rem;
  flex-wrap: wrap;
}

.tag {
  display: inline-flex;
  align-items: center;
  padding: 0.375rem 1rem;
  background: linear-gradient(135deg, rgba(236, 72, 153, 0.12) 0%, rgba(168, 85, 247, 0.12) 100%);
  color: #be185d;
  border-radius: 9999px;
  font-size: 0.8125rem;
  font-weight: 500;
  transition: all 0.2s ease;
}

.tag:hover {
  background: linear-gradient(135deg, rgba(236, 72, 153, 0.2) 0%, rgba(168, 85, 247, 0.2) 100%);
  transform: scale(1.05);
}

.article-title {
  font-size: 2.25rem;
  font-weight: 700;
  background: linear-gradient(135deg, #ec4899 0%, #a855f7 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  margin-bottom: 1rem;
  line-height: 1.3;
}

.article-meta {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 0.5rem;
  color: #6b7280;
  font-size: 0.9375rem;
}

.meta-icon {
  width: 18px;
  height: 18px;
}

/* ═══════════════════════════════════════════════════════════
   目录
   ═══════════════════════════════════════════════════════════ */
.toc-aside {
  position: sticky;
  top: 2rem;
  margin-bottom: 2rem;
  z-index: 1;
  display: flex;
  justify-content: center;
}

.toc-glass {
  background: rgba(255, 255, 255, 0.65);
  backdrop-filter: blur(16px);
  -webkit-backdrop-filter: blur(16px);
  border-radius: 20px;
  padding: 1.5rem;
  border: 1px solid rgba(255, 255, 255, 0.5);
  box-shadow: 0 8px 24px rgba(236, 72, 153, 0.08);
  max-width: 280px;
  width: 100%;
}

.toc-title {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  font-size: 0.9375rem;
  font-weight: 600;
  color: #1f2937;
  margin-bottom: 1rem;
  padding-bottom: 0.75rem;
  border-bottom: 1px solid rgba(236, 72, 153, 0.15);
}

.toc-title svg {
  width: 18px;
  height: 18px;
  color: #ec4899;
}

.toc-nav {
  display: flex;
  flex-direction: column;
  gap: 0.25rem;
}

.toc-link {
  color: #6b7280;
  text-decoration: none;
  font-size: 0.875rem;
  padding: 0.5rem 0.75rem;
  border-radius: 8px;
  transition: all 0.2s ease;
}

.toc-link:hover {
  color: #ec4899;
  background: rgba(236, 72, 153, 0.08);
}

.toc-link.level-2 {
  padding-left: 1rem;
}

/* ═══════════════════════════════════════════════════════════
   文章内容
   ═══════════════════════════════════════════════════════════ */
.article-content {
  position: relative;
  z-index: 1;
  background: rgba(255, 255, 255, 0.7);
  backdrop-filter: blur(20px);
  -webkit-backdrop-filter: blur(20px);
  border-radius: 24px;
  padding: 3rem;
  border: 1px solid rgba(255, 255, 255, 0.6);
  box-shadow: 0 8px 32px rgba(236, 72, 153, 0.08);
  margin-bottom: 2rem;
}

.content-wrapper {
  font-size: 1.0625rem;
  line-height: 1.8;
  color: #374151;
}

.content-wrapper :deep(h2) {
  font-size: 1.5rem;
  font-weight: 700;
  color: #be185d;
  margin-top: 2rem;
  margin-bottom: 1rem;
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.content-wrapper :deep(h2)::before {
  content: '';
  width: 4px;
  height: 24px;
  background: linear-gradient(135deg, #ec4899, #a855f7);
  border-radius: 2px;
}

.content-wrapper :deep(p) {
  margin-bottom: 1.25rem;
}

.content-wrapper :deep(strong) {
  color: #be185d;
  font-weight: 600;
}

.content-wrapper :deep(a) {
  color: #ec4899;
  text-decoration: underline;
  text-underline-offset: 2px;
}

.content-wrapper :deep(a):hover {
  color: #db2777;
}

/* ═══════════════════════════════════════════════════════════
   文章底部
   ═══════════════════════════════════════════════════════════ */
.article-footer {
  position: relative;
  z-index: 1;
  display: flex;
  justify-content: center;
  margin-bottom: 3rem;
}

.footer-glass {
  background: rgba(255, 255, 255, 0.65);
  backdrop-filter: blur(20px);
  -webkit-backdrop-filter: blur(20px);
  border-radius: 24px;
  padding: 2.5rem;
  text-align: center;
  border: 1px solid rgba(255, 255, 255, 0.6);
  box-shadow: 0 8px 32px rgba(236, 72, 153, 0.08);
  width: 100%;
  max-width: 800px;
}

.footer-heart {
  width: 48px;
  height: 48px;
  color: #ec4899;
  margin: 0 auto 1rem;
  animation: heart-pulse 1.5s ease-in-out infinite;
}

@keyframes heart-pulse {
  0%, 100% { transform: scale(1); }
  50% { transform: scale(1.15); }
}

.footer-glass p {
  color: #6b7280;
  font-size: 1rem;
  margin-bottom: 1.5rem;
}

.back-home {
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;
  color: #ec4899;
  font-weight: 600;
  font-size: 0.9375rem;
  text-decoration: none;
  padding: 0.625rem 1.25rem;
  background: rgba(236, 72, 153, 0.08);
  border-radius: 12px;
  transition: all 0.2s ease;
}

.back-home svg {
  width: 18px;
  height: 18px;
  transition: transform 0.2s ease;
}

.back-home:hover {
  background: rgba(236, 72, 153, 0.15);
}

.back-home:hover svg {
  transform: translateX(-4px);
}

/* ═══════════════════════════════════════════════════════════
   404 状态
   ═══════════════════════════════════════════════════════════ */
.not-found {
  position: relative;
  z-index: 1;
  display: flex;
  justify-content: center;
  padding: 4rem 1rem;
}

.not-found-glass {
  background: rgba(255, 255, 255, 0.65);
  backdrop-filter: blur(20px);
  -webkit-backdrop-filter: blur(20px);
  border-radius: 24px;
  padding: 3rem 2rem;
  text-align: center;
  border: 1px solid rgba(255, 255, 255, 0.6);
  box-shadow: 0 12px 40px rgba(236, 72, 153, 0.1);
  max-width: 500px;
  width: 100%;
}

.not-found-icon {
  width: 80px;
  height: 80px;
  color: rgba(236, 72, 153, 0.3);
  margin-bottom: 1rem;
}

.not-found-glass h2 {
  font-size: 1.5rem;
  font-weight: 700;
  color: #1f2937;
  margin-bottom: 0.5rem;
}

.not-found-glass p {
  color: #9ca3af;
  font-size: 1rem;
  margin-bottom: 1.5rem;
}

.back-home-btn {
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;
  color: white;
  font-weight: 600;
  font-size: 0.9375rem;
  text-decoration: none;
  padding: 0.75rem 1.5rem;
  background: linear-gradient(135deg, #ec4899 0%, #a855f7 100%);
  border-radius: 14px;
  box-shadow: 0 4px 15px rgba(236, 72, 153, 0.3);
  transition: all 0.2s ease;
}

.back-home-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 25px rgba(236, 72, 153, 0.4);
}

.back-home-btn svg {
  width: 18px;
  height: 18px;
}

/* ═══════════════════════════════════════════════════════════
   响应式
   ═══════════════════════════════════════════════════════════ */
@media (max-width: 768px) {
  .post-detail {
    padding: 1rem;
  }

  .cover-image-wrapper {
    height: 200px;
  }

  .header-glass {
    padding: 2rem 1.5rem;
  }

  .article-title {
    font-size: 1.625rem;
  }

  .article-content {
    padding: 2rem 1.5rem;
  }

  .content-wrapper :deep(h2) {
    font-size: 1.25rem;
  }

  .toc-aside {
    position: static;
  }

  .toc-glass {
    max-width: 100%;
  }
}
</style>
