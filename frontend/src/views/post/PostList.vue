<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { RouterLink } from 'vue-router'

interface Post {
  id: number
  title: string
  summary: string
  cover_image: string
  published_at: string
  tags: Array<{ id: number; name: string }>
}

const posts = ref<Post[]>([])
const loading = ref(true)
const currentPage = ref(1)
const total = ref(0)

const loadPosts = async () => {
  loading.value = true
  try {
    // TODO: 调用 API 获取文章列表
    // const res = await api.get('/posts', { page: currentPage.value })
    // posts.value = res.data.list
    // total.value = res.data.total

    // 模拟数据
    posts.value = [
      {
        id: 1,
        title: '我们的第一次相遇',
        summary: '那是一个阳光明媚的下午，我在咖啡馆里等你...',
        cover_image: 'https://picsum.photos/600/400?random=1',
        published_at: '2024-02-14',
        tags: [
          { id: 1, name: '相遇' },
          { id: 2, name: '浪漫' },
        ],
      },
      {
        id: 2,
        title: '一起看过的日出',
        summary: '凌晨四点，我们手牵手站在山顶，看着太阳缓缓升起...',
        cover_image: 'https://picsum.photos/600/400?random=2',
        published_at: '2024-01-20',
        tags: [
          { id: 3, name: '旅行' },
          { id: 4, name: '日出' },
        ],
      },
      {
        id: 3,
        title: '海边日落',
        summary: '夕阳西下，海浪轻抚着沙滩，你的侧脸真美...',
        cover_image: 'https://picsum.photos/600/400?random=3',
        published_at: '2024-01-15',
        tags: [
          { id: 3, name: '旅行' },
          { id: 5, name: '日落' },
        ],
      },
      {
        id: 4,
        title: '冬日的第一杯奶茶',
        summary: '你说冬天最幸福的事，莫过于手捧一杯热奶茶...',
        cover_image: 'https://picsum.photos/600/400?random=4',
        published_at: '2024-01-10',
        tags: [
          { id: 6, name: '日常' },
          { id: 7, name: '温暖' },
        ],
      },
    ]
    total.value = posts.value.length
  } catch (error) {
    console.error('加载文章失败:', error)
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  loadPosts()
})
</script>

<template>
  <div class="posts-page">
    <!-- 浮动爱心背景 -->
    <div class="bg-hearts">
      <div class="heart" v-for="i in 15" :key="i" :style="{
        left: Math.random() * 100 + '%',
        animationDelay: Math.random() * 5 + 's',
        animationDuration: (Math.random() * 3 + 5) + 's',
      }">
        <svg viewBox="0 0 24 24" fill="currentColor">
          <path d="M12 21.35l-1.45-1.32C5.4 15.36 2 12.28 2 8.5 2 5.42 4.42 3 7.5 3c1.74 0 3.41.81 4.5 2.09C13.09 3.81 14.76 3 16.5 3 19.58 3 22 5.42 22 8.5c0 3.78-3.4 6.86-8.55 11.54L12 21.35z"/>
        </svg>
      </div>
    </div>

    <!-- 页面头部 -->
    <div class="page-header">
      <div class="header-glass">
        <h1 class="page-title">
          <span class="heart-icon">
            <svg viewBox="0 0 24 24" fill="currentColor">
              <path d="M12 21.35l-1.45-1.32C5.4 15.36 2 12.28 2 8.5 2 5.42 4.42 3 7.5 3c1.74 0 3.41.81 4.5 2.09C13.09 3.81 14.76 3 16.5 3 19.58 3 22 5.42 22 8.5c0 3.78-3.4 6.86-8.55 11.54L12 21.35z"/>
            </svg>
          </span>
          恋爱日记
        </h1>
        <p class="page-subtitle">记录我们的每一个甜蜜瞬间</p>
      </div>
    </div>

    <!-- 文章网格 -->
    <div class="posts-grid">
      <!-- 加载骨架屏 -->
      <div v-if="loading" class="loading-cards">
        <div class="loading-card" v-for="i in 4" :key="i">
          <div class="loading-image"></div>
          <div class="loading-content">
            <div class="loading-line short"></div>
            <div class="loading-line"></div>
            <div class="loading-line"></div>
          </div>
        </div>
      </div>

      <!-- 文章卡片 -->
      <RouterLink
        v-else
        :to="`/post/${post.id}`"
        class="post-card"
        v-for="post in posts"
        :key="post.id"
      >
        <div class="post-cover-wrapper">
          <div class="post-cover" v-if="post.cover_image">
            <img :src="post.cover_image" :alt="post.title" />
            <div class="cover-overlay"></div>
          </div>
          <div class="post-cover placeholder" v-else>
            <svg class="placeholder-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5">
              <path d="M12 6.042A8.967 8.967 0 0 0 6 3.75c-1.052 0-2.062.18-3 .512v14.25A8.987 8.987 0 0 1 6 18c2.305 0 4.408.867 6 2.292m0-14.25a8.966 8.966 0 0 1 6-2.292c1.052 0 2.062.18 3 .512v14.25A8.987 8.987 0 0 0 18 18a8.967 8.967 0 0 0-6 2.292m0-14.25v14.25"/>
            </svg>
          </div>
        </div>

        <div class="post-content">
          <h2 class="post-title">{{ post.title }}</h2>
          <p class="post-summary">{{ post.summary }}</p>

          <div class="post-meta">
            <div class="post-tags">
              <span class="tag" v-for="tag in post.tags" :key="tag.id">
                {{ tag.name }}
              </span>
            </div>
            <span class="post-date">{{ post.published_at }}</span>
          </div>
        </div>
      </RouterLink>
    </div>

    <!-- 分页 -->
    <div class="pagination" v-if="!loading && total > 0">
      <el-pagination
        v-model:current-page="currentPage"
        :page-size="10"
        :total="total"
        layout="prev, pager, next"
        @current-change="loadPosts"
      />
    </div>

    <!-- 空状态 -->
    <div class="empty-state" v-if="!loading && posts.length === 0">
      <div class="empty-glass">
        <svg class="empty-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5">
          <path d="M12 6.042A8.967 8.967 0 0 0 6 3.75c-1.052 0-2.062.18-3 .512v14.25A8.987 8.987 0 0 1 6 18c2.305 0 4.408.867 6 2.292m0-14.25a8.966 8.966 0 0 1 6-2.292c1.052 0 2.062.18 3 .512v14.25A8.987 8.987 0 0 0 18 18a8.967 8.967 0 0 0-6 2.292m0-14.25v14.25"/>
        </svg>
        <p>还没有日记，开始记录你们的故事吧</p>
      </div>
    </div>
  </div>
</template>

<style scoped>
.posts-page {
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
   页面头部
   ═══════════════════════════════════════════════════════════ */
.page-header {
  position: relative;
  z-index: 1;
  margin-bottom: 3rem;
  display: flex;
  justify-content: center;
}

.header-glass {
  background: rgba(255, 255, 255, 0.65);
  backdrop-filter: blur(20px);
  -webkit-backdrop-filter: blur(20px);
  border-radius: 24px;
  padding: 2.5rem 3rem;
  text-align: center;
  border: 1px solid rgba(255, 255, 255, 0.6);
  box-shadow: 0 12px 40px rgba(236, 72, 153, 0.12);
  max-width: 600px;
  width: 100%;
}

.page-title {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 0.75rem;
  font-size: 2.25rem;
  font-weight: 700;
  background: linear-gradient(135deg, #ec4899 0%, #a855f7 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  margin-bottom: 0.75rem;
}

.heart-icon {
  width: 40px;
  height: 40px;
  color: #ec4899;
  animation: heart-pulse 1.5s ease-in-out infinite;
}

@keyframes heart-pulse {
  0%, 100% { transform: scale(1); }
  50% { transform: scale(1.15); }
}

.page-subtitle {
  color: #6b7280;
  font-size: 1.0625rem;
}

/* ═══════════════════════════════════════════════════════════
   文章网格
   ═══════════════════════════════════════════════════════════ */
.posts-grid {
  position: relative;
  z-index: 1;
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(340px, 1fr));
  gap: 1.75rem;
  max-width: 1400px;
  margin: 0 auto 2rem;
}

/* 文章卡片 */
.post-card {
  background: rgba(255, 255, 255, 0.7);
  backdrop-filter: blur(16px);
  -webkit-backdrop-filter: blur(16px);
  border-radius: 20px;
  overflow: hidden;
  border: 1px solid rgba(255, 255, 255, 0.5);
  box-shadow: 0 8px 32px rgba(236, 72, 153, 0.1);
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
  text-decoration: none;
  color: inherit;
  display: flex;
  flex-direction: column;
}

.post-card:hover {
  transform: translateY(-6px);
  box-shadow: 0 16px 48px rgba(236, 72, 153, 0.2);
  border-color: rgba(236, 72, 153, 0.3);
}

/* 封面图 */
.post-cover-wrapper {
  position: relative;
  width: 100%;
  height: 220px;
  overflow: hidden;
}

.post-cover {
  width: 100%;
  height: 100%;
  position: relative;
}

.post-cover img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: transform 0.5s cubic-bezier(0.4, 0, 0.2, 1);
}

.post-card:hover .post-cover img {
  transform: scale(1.08);
}

.cover-overlay {
  position: absolute;
  inset: 0;
  background: linear-gradient(to bottom, transparent 0%, rgba(236, 72, 153, 0.1) 100%);
  pointer-events: none;
}

.post-cover.placeholder {
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, rgba(236, 72, 153, 0.05) 0%, rgba(168, 85, 247, 0.05) 100%);
}

.placeholder-icon {
  width: 64px;
  height: 64px;
  color: rgba(236, 72, 153, 0.3);
}

/* 卡片内容 */
.post-content {
  padding: 1.5rem;
  flex: 1;
  display: flex;
  flex-direction: column;
}

.post-title {
  font-size: 1.1875rem;
  font-weight: 700;
  color: #1f2937;
  margin-bottom: 0.75rem;
  line-height: 1.4;
}

.post-summary {
  color: #6b7280;
  font-size: 0.9375rem;
  line-height: 1.6;
  margin-bottom: 1rem;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
  flex: 1;
}

/* 元数据 */
.post-meta {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-top: 1rem;
  border-top: 1px solid rgba(236, 72, 153, 0.1);
}

.post-tags {
  display: flex;
  gap: 0.5rem;
  flex-wrap: wrap;
}

.tag {
  display: inline-flex;
  align-items: center;
  padding: 0.3125rem 0.875rem;
  background: linear-gradient(135deg, rgba(236, 72, 153, 0.1) 0%, rgba(168, 85, 247, 0.1) 100%);
  color: #ec4899;
  border-radius: 9999px;
  font-size: 0.75rem;
  font-weight: 500;
  transition: all 0.2s ease;
}

.tag:hover {
  background: linear-gradient(135deg, rgba(236, 72, 153, 0.2) 0%, rgba(168, 85, 247, 0.2) 100%);
  transform: scale(1.05);
}

.post-date {
  color: #9ca3af;
  font-size: 0.8125rem;
  white-space: nowrap;
}

/* ═══════════════════════════════════════════════════════════
   加载骨架屏
   ═══════════════════════════════════════════════════════════ */
.loading-cards {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(340px, 1fr));
  gap: 1.75rem;
}

.loading-card {
  border-radius: 20px;
  overflow: hidden;
  background: rgba(255, 255, 255, 0.5);
  border: 1px solid rgba(236, 72, 153, 0.1);
}

.loading-image {
  height: 220px;
  background: linear-gradient(90deg,
    rgba(252, 231, 243, 0.8) 0%,
    rgba(245, 208, 254, 0.8) 50%,
    rgba(252, 231, 243, 0.8) 100%);
  background-size: 200% 100%;
  animation: shimmer 1.5s infinite;
}

.loading-content {
  padding: 1.5rem;
}

.loading-line {
  height: 16px;
  background: linear-gradient(90deg,
    rgba(252, 231, 243, 0.8) 0%,
    rgba(245, 208, 254, 0.8) 50%,
    rgba(252, 231, 243, 0.8) 100%);
  background-size: 200% 100%;
  animation: shimmer 1.5s infinite;
  border-radius: 4px;
  margin-bottom: 0.75rem;
}

.loading-line.short {
  width: 60%;
  margin-bottom: 0.5rem;
}

@keyframes shimmer {
  0% { background-position: 200% 0; }
  100% { background-position: -200% 0; }
}

/* ═══════════════════════════════════════════════════════════
   分页
   ═══════════════════════════════════════════════════════════ */
.pagination {
  position: relative;
  z-index: 1;
  display: flex;
  justify-content: center;
  padding: 1rem 0;
}

/* ═══════════════════════════════════════════════════════════
   空状态
   ═══════════════════════════════════════════════════════════ */
.empty-state {
  position: relative;
  z-index: 1;
  display: flex;
  justify-content: center;
  padding: 3rem 1rem;
}

.empty-glass {
  background: rgba(255, 255, 255, 0.65);
  backdrop-filter: blur(16px);
  -webkit-backdrop-filter: blur(16px);
  border-radius: 24px;
  padding: 3rem;
  text-align: center;
  border: 1px solid rgba(255, 255, 255, 0.5);
  box-shadow: 0 12px 40px rgba(236, 72, 153, 0.1);
}

.empty-icon {
  width: 80px;
  height: 80px;
  color: rgba(236, 72, 153, 0.3);
  margin-bottom: 1.5rem;
}

.empty-glass p {
  color: #9ca3af;
  font-size: 1.0625rem;
}

/* ═══════════════════════════════════════════════════════════
   响应式
   ═══════════════════════════════════════════════════════════ */
@media (max-width: 1024px) {
  .posts-grid {
    grid-template-columns: repeat(2, 1fr);
  }
}

@media (max-width: 768px) {
  .posts-page {
    padding: 1rem;
  }

  .header-glass {
    padding: 2rem 1.5rem;
  }

  .page-title {
    font-size: 1.625rem;
  }

  .heart-icon {
    width: 32px;
    height: 32px;
  }

  .posts-grid {
    grid-template-columns: 1fr;
  }

  .loading-cards {
    grid-template-columns: 1fr;
  }
}
</style>
