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
        cover_image: '',
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
        cover_image: '',
        published_at: '2024-01-20',
        tags: [
          { id: 3, name: '旅行' },
          { id: 4, name: '日出' },
        ],
      },
    ]
    total.value = 2
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
    <div class="page-header">
      <h1 class="page-title">
        <span class="heart-icon">💕</span>
        恋爱日记
      </h1>
      <p class="page-subtitle">记录我们的每一个甜蜜瞬间</p>
    </div>

    <div class="posts-grid">
      <div v-if="loading" class="loading-cards">
        <div class="loading-card" v-for="i in 4" :key="i"></div>
      </div>

      <RouterLink
        v-else
        :to="`/post/${post.id}`"
        class="post-card"
        v-for="post in posts"
        :key="post.id"
      >
        <div class="post-cover" v-if="post.cover_image">
          <img :src="post.cover_image" :alt="post.title" />
        </div>
        <div class="post-cover placeholder" v-else>
          <span class="placeholder-icon">📝</span>
        </div>

        <div class="post-content">
          <h2 class="post-title">{{ post.title }}</h2>
          <p class="post-summary">{{ post.summary }}</p>

          <div class="post-meta">
            <span class="post-date">{{ post.published_at }}</span>
            <div class="post-tags">
              <span class="tag" v-for="tag in post.tags" :key="tag.id"> #{{ tag.name }} </span>
            </div>
          </div>
        </div>
      </RouterLink>
    </div>

    <div class="pagination" v-if="!loading && total > 0">
      <el-pagination
        v-model:current-page="currentPage"
        :page-size="10"
        :total="total"
        layout="prev, pager, next"
        @current-change="loadPosts"
      />
    </div>

    <div class="empty-state" v-if="!loading && posts.length === 0">
      <span class="empty-icon">💭</span>
      <p>还没有日记，开始记录你们的故事吧</p>
    </div>
  </div>
</template>

<style scoped>
.posts-page {
  max-width: 1200px;
  margin: 0 auto;
  padding: 2rem;
}

.page-header {
  text-align: center;
  margin-bottom: 3rem;
}

.page-title {
  font-size: 2.5rem;
  font-weight: 700;
  background: linear-gradient(135deg, #ec4899 0%, #a855f7 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 0.75rem;
}

.heart-icon {
  font-size: 2.75rem;
  animation: pulse 1.5s infinite;
}

@keyframes pulse {
  0%,
  100% {
    transform: scale(1);
  }
  50% {
    transform: scale(1.1);
  }
}

.page-subtitle {
  color: #6b7280;
  font-size: 1.125rem;
  margin-top: 0.5rem;
}

.posts-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(320px, 1fr));
  gap: 2rem;
  margin-bottom: 2rem;
}

.post-card {
  background: rgba(255, 255, 255, 0.9);
  border-radius: 16px;
  overflow: hidden;
  transition: all 0.3s ease;
  border: 1px solid rgba(236, 72, 153, 0.2);
  text-decoration: none;
  color: inherit;
}

.post-card:hover {
  transform: translateY(-8px);
  box-shadow: 0 20px 40px rgba(236, 72, 153, 0.25);
  border-color: #ec4899;
}

.post-cover {
  width: 100%;
  height: 200px;
  overflow: hidden;
  background: #f3e8e8;
}

.post-cover img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: transform 0.3s ease;
}

.post-card:hover .post-cover img {
  transform: scale(1.05);
}

.post-cover.placeholder {
  display: flex;
  align-items: center;
  justify-content: center;
}

.placeholder-icon {
  font-size: 4rem;
  opacity: 0.5;
}

.post-content {
  padding: 1.5rem;
}

.post-title {
  font-size: 1.25rem;
  font-weight: 600;
  color: #1f2937;
  margin-bottom: 0.75rem;
}

.post-summary {
  color: #6b7280;
  font-size: 0.875rem;
  line-height: 1.6;
  margin-bottom: 1rem;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.post-meta {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.post-date {
  color: #9ca3af;
  font-size: 0.75rem;
}

.post-tags {
  display: flex;
  gap: 0.5rem;
}

.tag {
  background: linear-gradient(135deg, rgba(236, 72, 153, 0.1), rgba(168, 85, 247, 0.1));
  color: #ec4899;
  padding: 0.25rem 0.75rem;
  border-radius: 9999px;
  font-size: 0.75rem;
}

/* 加载骨架屏 */
.loading-cards {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(320px, 1fr));
  gap: 2rem;
}

.loading-card {
  height: 320px;
  background: linear-gradient(90deg, #fce7f3 25%, #f5d0fe 50%, #fce7f3 75%);
  background-size: 200% 100%;
  animation: shimmer 1.5s infinite;
  border-radius: 16px;
}

@keyframes shimmer {
  0% {
    background-position: 200% 0;
  }
  100% {
    background-position: -200% 0;
  }
}

/* 分页 */
.pagination {
  display: flex;
  justify-content: center;
  padding: 2rem 0;
}

/* 空状态 */
.empty-state {
  text-align: center;
  padding: 4rem 2rem;
}

.empty-icon {
  font-size: 4rem;
  display: block;
  margin-bottom: 1rem;
}

.empty-state p {
  color: #9ca3af;
  font-size: 1.125rem;
}

@media (max-width: 768px) {
  .posts-page {
    padding: 1rem;
  }

  .page-title {
    font-size: 1.75rem;
  }

  .posts-grid {
    grid-template-columns: 1fr;
  }
}
</style>
