<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'
import { useUserStore } from '@/stores/user'
import { RouterLink } from 'vue-router'

const userStore = useUserStore()

// 瀑布流卡片数据
const posts = ref([
  {
    id: 1,
    title: '我们的第一次相遇',
    summary: '那是一个阳光明媚的下午，咖啡厅的角落里，你穿着一件白色的衬衫...',
    image: 'https://picsum.photos/400/300?random=1',
    date: '2024-02-14',
    tags: ['相遇', '缘分'],
    height: 280,
  },
  {
    id: 2,
    title: '海边日落',
    summary: '牵手走在沙滩上，海浪轻轻拍打着脚踝，夕阳把我们的影子拉得很长...',
    image: 'https://picsum.photos/400/500?random=2',
    date: '2024-02-10',
    tags: ['旅行', '浪漫'],
    height: 420,
  },
  {
    id: 3,
    title: '一起做饭的时光',
    summary: '你在切菜，我在炒菜，厨房里弥漫着幸福的味道...',
    image: 'https://picsum.photos/400/350?random=3',
    date: '2024-02-08',
    tags: ['日常', '温馨'],
    height: 320,
  },
  {
    id: 4,
    title: '星空下的约定',
    summary: '躺在草地上看星星，你说每一颗星星都代表着我们的一个愿望...',
    image: 'https://picsum.photos/400/400?random=4',
    date: '2024-02-05',
    tags: ['浪漫', '约定'],
    height: 380,
  },
  {
    id: 5,
    title: '冬日暖阳',
    summary: '下雪的日子里，我们堆了一个大大的雪人，还给它戴上了你的围巾...',
    image: 'https://picsum.photos/400/320?random=5',
    date: '2024-01-28',
    tags: ['冬天', '回忆'],
    height: 300,
  },
  {
    id: 6,
    title: '生日惊喜',
    summary: '偷偷准备了很久的生日派对，看到你惊讶的表情，一切都值得...',
    image: 'https://picsum.photos/400/450?random=6',
    date: '2024-01-20',
    tags: ['生日', '惊喜'],
    height: 400,
  },
])

// 浮动爱心背景
const hearts = ref<Array<{ id: number; left: string; delay: string; duration: string; size: string }>>([])

onMounted(() => {
  hearts.value = Array.from({ length: 20 }, (_, i) => ({
    id: i,
    left: Math.random() * 100 + '%',
    delay: Math.random() * 5 + 's',
    duration: (Math.random() * 3 + 5) + 's',
    size: (Math.random() * 15 + 20) + 'px',
  }))
})
</script>

<template>
  <main class="home">
    <!-- 浮动爱心背景 -->
    <div class="bg-hearts">
      <div
        class="heart"
        v-for="heart in hearts"
        :key="heart.id"
        :style="{
          left: heart.left,
          animationDelay: heart.delay,
          animationDuration: heart.duration,
          fontSize: heart.size,
        }"
      >
        <svg viewBox="0 0 24 24" fill="currentColor">
          <path d="M12 21.35l-1.45-1.32C5.4 15.36 2 12.28 2 8.5 2 5.42 4.42 3 7.5 3c1.74 0 3.41.81 4.5 2.09C13.09 3.81 14.76 3 16.5 3 19.58 3 22 5.42 22 8.5c0 3.78-3.4 6.86-8.55 11.54L12 21.35z"/>
        </svg>
      </div>
    </div>

    <!-- Hero 区域 -->
    <section class="hero">
      <div class="hero-glass">
        <div class="hero-content">
          <h1 class="title">
            <span class="logo-heart">
              <svg viewBox="0 0 24 24" fill="currentColor">
                <path d="M12 21.35l-1.45-1.32C5.4 15.36 2 12.28 2 8.5 2 5.42 4.42 3 7.5 3c1.74 0 3.41.81 4.5 2.09C13.09 3.81 14.76 3 16.5 3 19.58 3 22 5.42 22 8.5c0 3.78-3.4 6.86-8.55 11.54L12 21.35z"/>
              </svg>
            </span>
            <span v-if="userStore.isLoggedIn">
              欢迎回来，{{ userStore.user?.nickname || userStore.user?.username }}
            </span>
            <span v-else>湘遇 · 恋爱博客</span>
          </h1>
          <p class="subtitle">
            记录属于我们的浪漫时光
          </p>

          <div class="actions">
            <RouterLink
              v-if="!userStore.isLoggedIn"
              to="/login"
              class="btn btn-primary"
            >
              <span>管理员登录</span>
              <svg class="btn-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M5 12h14M12 5l7 7-7 7"/>
              </svg>
            </RouterLink>
            <RouterLink
              v-else
              to="/posts/my"
              class="btn btn-primary"
            >
              <span>浏览日记</span>
              <svg class="btn-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M5 12h14M12 5l7 7-7 7"/>
              </svg>
            </RouterLink>
          </div>
        </div>

        <!-- 装饰光晕 -->
        <div class="glow-orb glow-1"></div>
        <div class="glow-orb glow-2"></div>
      </div>
    </section>

    <!-- 瀑布流卡片区域 -->
    <section class="waterfall-section">
      <div class="section-header">
        <h2 class="section-title">
          <span class="title-icon">
            <svg viewBox="0 0 24 24" fill="currentColor">
              <path d="M12 21.35l-1.45-1.32C5.4 15.36 2 12.28 2 8.5 2 5.42 4.42 3 7.5 3c1.74 0 3.41.81 4.5 2.09C13.09 3.81 14.76 3 16.5 3 19.58 3 22 5.42 22 8.5c0 3.78-3.4 6.86-8.55 11.54L12 21.35z"/>
            </svg>
          </span>
          爱的回忆
        </h2>
        <p class="section-subtitle">珍藏我们的每一个美好瞬间</p>
      </div>

      <div class="waterfall-grid">
        <article
          v-for="post in posts"
          :key="post.id"
          class="waterfall-card"
          :style="{ '--card-height': post.height + 'px' }"
        >
          <div class="card-image-wrapper">
            <img :src="post.image" :alt="post.title" class="card-image" />
            <div class="card-image-overlay"></div>
          </div>

          <div class="card-content">
            <h3 class="card-title">{{ post.title }}</h3>
            <p class="card-summary">{{ post.summary }}</p>

            <div class="card-meta">
              <div class="card-tags">
                <span
                  v-for="tag in post.tags"
                  :key="tag"
                  class="tag"
                >
                  {{ tag }}
                </span>
              </div>
              <span class="card-date">{{ post.date }}</span>
            </div>

            <RouterLink :to="`/post/${post.id}`" class="card-link">
              <span>阅读全文</span>
              <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M5 12h14M12 5l7 7-7 7"/>
              </svg>
            </RouterLink>
          </div>
        </article>
      </div>
    </section>
  </main>
</template>

<style scoped>
.home {
  min-height: 100vh;
  position: relative;
  padding-bottom: 4rem;
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
  color: rgba(236, 72, 153, 0.2);
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
   Hero 区域
   ═══════════════════════════════════════════════════════════ */
.hero {
  position: relative;
  padding: 4rem 2rem 3rem;
  display: flex;
  justify-content: center;
  z-index: 1;
}

.hero-glass {
  position: relative;
  width: 100%;
  max-width: 900px;
  background: rgba(255, 255, 255, 0.65);
  backdrop-filter: blur(20px);
  -webkit-backdrop-filter: blur(20px);
  border-radius: 32px;
  padding: 4rem 3rem;
  border: 1px solid rgba(255, 255, 255, 0.6);
  box-shadow:
    0 20px 60px rgba(236, 72, 153, 0.15),
    inset 0 1px 0 rgba(255, 255, 255, 0.8);
  overflow: hidden;
}

.hero-content {
  position: relative;
  z-index: 2;
  text-align: center;
}

.title {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 1rem;
  font-size: 2.5rem;
  font-weight: 700;
  background: linear-gradient(135deg, #ec4899 0%, #a855f7 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  margin-bottom: 1rem;
}

.logo-heart {
  width: 48px;
  height: 48px;
  color: #ec4899;
  animation: heart-pulse 1.5s ease-in-out infinite;
}

@keyframes heart-pulse {
  0%, 100% { transform: scale(1); }
  50% { transform: scale(1.15); }
}

.subtitle {
  font-size: 1.125rem;
  color: #6b7280;
  margin-bottom: 2rem;
}

.actions {
  display: flex;
  justify-content: center;
  gap: 1rem;
}

.btn {
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.875rem 1.75rem;
  border-radius: 14px;
  font-weight: 600;
  font-size: 0.9375rem;
  text-decoration: none;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  cursor: pointer;
}

.btn-primary {
  background: linear-gradient(135deg, #ec4899 0%, #a855f7 100%);
  color: white;
  border: none;
  box-shadow: 0 4px 15px rgba(236, 72, 153, 0.3);
}

.btn-primary:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 25px rgba(236, 72, 153, 0.4);
}

.btn-icon {
  width: 18px;
  height: 18px;
  transition: transform 0.3s ease;
}

.btn:hover .btn-icon {
  transform: translateX(4px);
}

/* 装饰光晕 */
.glow-orb {
  position: absolute;
  border-radius: 50%;
  filter: blur(60px);
  opacity: 0.5;
  pointer-events: none;
}

.glow-1 {
  width: 300px;
  height: 300px;
  background: radial-gradient(circle, #f472b6 0%, transparent 70%);
  top: -150px;
  right: -100px;
  animation: glow-pulse 4s ease-in-out infinite;
}

.glow-2 {
  width: 250px;
  height: 250px;
  background: radial-gradient(circle, #c084fc 0%, transparent 70%);
  bottom: -100px;
  left: -80px;
  animation: glow-pulse 4s ease-in-out infinite 2s;
}

@keyframes glow-pulse {
  0%, 100% { opacity: 0.4; transform: scale(1); }
  50% { opacity: 0.6; transform: scale(1.1); }
}

/* ═══════════════════════════════════════════════════════════
   瀑布流区域
   ═══════════════════════════════════════════════════════════ */
.waterfall-section {
  position: relative;
  z-index: 1;
  padding: 3rem 2rem;
  max-width: 1400px;
  margin: 0 auto;
}

.section-header {
  text-align: center;
  margin-bottom: 3rem;
}

.section-title {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 0.75rem;
  font-size: 2rem;
  font-weight: 700;
  color: #1f2937;
  margin-bottom: 0.5rem;
}

.title-icon {
  width: 32px;
  height: 32px;
  color: #ec4899;
}

.section-subtitle {
  font-size: 1rem;
  color: #6b7280;
}

/* 瀑布流网格 */
.waterfall-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(340px, 1fr));
  grid-auto-rows: masonry;
  gap: 1.5rem;
}

@supports not (grid-auto-rows: masonry) {
  .waterfall-grid {
    display: block;
    column-count: 3;
    column-gap: 1.5rem;
  }
}

@media (max-width: 1024px) {
  .waterfall-grid {
    column-count: 2;
  }
}

@media (max-width: 640px) {
  .waterfall-grid {
    column-count: 1;
  }
}

/* 瀑布流卡片 */
.waterfall-card {
  break-inside: avoid;
  margin-bottom: 1.5rem;
  background: rgba(255, 255, 255, 0.7);
  backdrop-filter: blur(16px);
  -webkit-backdrop-filter: blur(16px);
  border-radius: 24px;
  overflow: hidden;
  border: 1px solid rgba(255, 255, 255, 0.5);
  box-shadow: 0 8px 32px rgba(236, 72, 153, 0.12);
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
  cursor: pointer;
}

.waterfall-card:hover {
  transform: translateY(-6px);
  box-shadow: 0 16px 48px rgba(236, 72, 153, 0.2);
  border-color: rgba(236, 72, 153, 0.3);
}

.card-image-wrapper {
  position: relative;
  height: var(--card-height);
  overflow: hidden;
}

.card-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: transform 0.5s cubic-bezier(0.4, 0, 0.2, 1);
}

.waterfall-card:hover .card-image {
  transform: scale(1.08);
}

.card-image-overlay {
  position: absolute;
  inset: 0;
  background: linear-gradient(to bottom, transparent 0%, rgba(0, 0, 0, 0.3) 100%);
  pointer-events: none;
}

.card-content {
  padding: 1.5rem;
}

.card-title {
  font-size: 1.25rem;
  font-weight: 700;
  color: #1f2937;
  margin-bottom: 0.75rem;
  line-height: 1.4;
}

.card-summary {
  font-size: 0.9375rem;
  color: #6b7280;
  line-height: 1.6;
  margin-bottom: 1rem;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.card-meta {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 1rem;
  margin-bottom: 1rem;
}

.card-tags {
  display: flex;
  gap: 0.5rem;
  flex-wrap: wrap;
}

.tag {
  display: inline-flex;
  align-items: center;
  padding: 0.25rem 0.75rem;
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

.card-date {
  font-size: 0.8125rem;
  color: #9ca3af;
  white-space: nowrap;
}

.card-link {
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;
  color: #ec4899;
  font-weight: 600;
  font-size: 0.875rem;
  text-decoration: none;
  transition: all 0.2s ease;
}

.card-link svg {
  width: 16px;
  height: 16px;
  transition: transform 0.2s ease;
}

.card-link:hover {
  color: #db2777;
}

.card-link:hover svg {
  transform: translateX(4px);
}

/* ═══════════════════════════════════════════════════════════
   响应式
   ═══════════════════════════════════════════════════════════ */
@media (max-width: 1024px) {
  .waterfall-grid {
    grid-template-columns: repeat(2, 1fr);
  }
}

@media (max-width: 768px) {
  .hero {
    padding: 2rem 1rem 2rem;
  }

  .hero-glass {
    padding: 2.5rem 1.5rem;
    border-radius: 24px;
  }

  .title {
    font-size: 1.75rem;
    flex-direction: column;
    gap: 0.5rem;
  }

  .logo-heart {
    width: 40px;
    height: 40px;
  }

  .subtitle {
    font-size: 1rem;
  }

  .waterfall-section {
    padding: 2rem 1rem;
  }

  .section-title {
    font-size: 1.5rem;
  }

  .waterfall-grid {
    grid-template-columns: 1fr;
  }
}

@media (max-width: 640px) {
  .waterfall-card {
    margin-bottom: 1.5rem;
  }
}
</style>
