<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()
const post = ref<any>(null)
const loading = ref(true)

onMounted(async () => {
  loading.value = true
  try {
    // TODO: 调用 API 获取文章详情
    // post.value = await api.get(`/posts/${route.params.id}`)

    // 模拟数据
    post.value = {
      id: route.params.id,
      title: '我们的第一次相遇',
      content_markdown:
        '# 我们的第一次相遇\n\n那是一个阳光明媚的下午...\n\n## 相遇\n\n我在咖啡馆里等你...',
      published_at: '2024-02-14',
      tags: [{ name: '相遇' }, { name: '浪漫' }],
    }
  } catch (error) {
    console.error('加载文章失败:', error)
  } finally {
    loading.value = false
  }
})
</script>

<template>
  <div class="post-detail">
    <div v-if="loading" class="loading">
      <div class="loading-spinner">
        <span class="dot"></span>
        <span class="dot"></span>
        <span class="dot"></span>
      </div>
    </div>

    <article v-else-if="post" class="article">
      <header class="article-header">
        <div class="article-tags">
          <span class="tag" v-for="tag in post.tags" :key="tag.name"> #{{ tag.name }} </span>
        </div>
        <h1 class="article-title">{{ post.title }}</h1>
        <div class="article-meta">
          <span class="date">📅 {{ post.published_at }}</span>
        </div>
      </header>

      <div class="article-content">
        <div class="content-wrapper" v-html="post.content_markdown"></div>
      </div>

      <footer class="article-footer">
        <p>💗 感谢你的阅读 💗</p>
      </footer>
    </article>

    <div v-else class="not-found">
      <span class="not-found-icon">💔</span>
      <p>文章不存在</p>
    </div>
  </div>
</template>

<style scoped>
.post-detail {
  max-width: 800px;
  margin: 0 auto;
  padding: 2rem;
}

.loading {
  display: flex;
  justify-content: center;
  padding: 4rem 0;
}

.loading-spinner {
  display: flex;
  gap: 4px;
}

.dot {
  width: 12px;
  height: 12px;
  background: linear-gradient(135deg, #ec4899, #a855f7);
  border-radius: 50%;
  animation: bounce 1.4s infinite ease-in-out both;
}

.dot:nth-child(1) {
  animation-delay: -0.32s;
}
.dot:nth-child(2) {
  animation-delay: -0.16s;
}

@keyframes bounce {
  0%,
  80%,
  100% {
    transform: scale(0);
  }
  40% {
    transform: scale(1);
  }
}

.article {
  background: rgba(255, 255, 255, 0.95);
  border-radius: 24px;
  padding: 3rem;
  box-shadow: 0 25px 50px -12px rgba(236, 72, 153, 0.2);
}

.article-header {
  text-align: center;
  margin-bottom: 3rem;
  border-bottom: 2px solid rgba(236, 72, 153, 0.2);
  padding-bottom: 2rem;
}

.article-tags {
  display: flex;
  justify-content: center;
  gap: 0.5rem;
  margin-bottom: 1rem;
}

.tag {
  background: linear-gradient(135deg, rgba(236, 72, 153, 0.1), rgba(168, 85, 247, 0.1));
  color: #ec4899;
  padding: 0.5rem 1rem;
  border-radius: 9999px;
  font-size: 0.875rem;
}

.article-title {
  font-size: 2.5rem;
  font-weight: 700;
  background: linear-gradient(135deg, #ec4899 0%, #a855f7 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  margin-bottom: 1rem;
}

.article-meta {
  color: #6b7280;
  font-size: 0.875rem;
}

.article-content {
  font-size: 1.125rem;
  line-height: 1.8;
  color: #374151;
}

.article-content :deep(h2) {
  color: #ec4899;
  margin-top: 2rem;
  margin-bottom: 1rem;
}

.article-footer {
  text-align: center;
  margin-top: 3rem;
  padding-top: 2rem;
  border-top: 1px solid rgba(236, 72, 153, 0.2);
  color: #9ca3af;
}

.not-found {
  text-align: center;
  padding: 4rem 2rem;
}

.not-found-icon {
  font-size: 5rem;
  display: block;
  margin-bottom: 1rem;
}

@media (max-width: 768px) {
  .article {
    padding: 2rem 1.5rem;
  }

  .article-title {
    font-size: 1.75rem;
  }
}
</style>
