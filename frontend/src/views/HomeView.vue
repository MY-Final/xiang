<script setup lang="ts">
import { useUserStore } from '@/stores/user'

const userStore = useUserStore()
</script>

<template>
  <main class="home">
    <div class="hero">
      <div class="hero-content">
        <h1 class="title">
          <span class="heart">💕</span>
          <span v-if="userStore.isLoggedIn">
            欢迎回来，{{ userStore.user?.nickname || userStore.user?.username }}
          </span>
          <span v-else>湘遇恋爱博客</span>
        </h1>
        <p class="subtitle">
          <span v-if="userStore.isLoggedIn">记录属于我们的浪漫时光</span>
          <span v-else>记录属于我们的浪漫时光</span>
        </p>

        <div class="actions" v-if="!userStore.isLoggedIn">
          <RouterLink to="/login" class="btn btn-primary">
            <span class="btn-icon">🔐</span>
            管理员登录
          </RouterLink>
        </div>

        <div class="actions" v-if="userStore.isLoggedIn">
          <RouterLink to="/posts" class="btn btn-primary">
            <span class="btn-icon">📝</span>
            浏览日记
          </RouterLink>
        </div>
      </div>

      <!-- 装饰性元素 -->
      <div class="decorations">
        <div
          class="floating-heart"
          v-for="i in 10"
          :key="i"
          :style="{
            left: Math.random() * 100 + '%',
            animationDelay: Math.random() * 5 + 's',
            fontSize: Math.random() * 20 + 20 + 'px',
          }"
        >
          ❤
        </div>
      </div>
    </div>

    <!-- 功能卡片 -->
    <section class="features">
      <div class="feature-card">
        <div class="feature-icon">📝</div>
        <h3>恋爱日记</h3>
        <p>记录我们的甜蜜日常</p>
      </div>
      <div class="feature-card">
        <div class="feature-icon">🏷️</div>
        <h3>情感标签</h3>
        <p>为每个回忆贴上专属标签</p>
      </div>
      <div class="feature-card">
        <div class="feature-icon">📸</div>
        <h3>美好瞬间</h3>
        <p>珍藏我们的每一张照片</p>
      </div>
      <div class="feature-card">
        <div class="feature-icon">📅</div>
        <h3>时间轴线</h3>
        <p>一起走过的每个日子</p>
      </div>
    </section>
  </main>
</template>

<style scoped>
.home {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

.hero {
  position: relative;
  padding: 4rem 2rem;
  text-align: center;
  overflow: hidden;
}

.hero-content {
  position: relative;
  z-index: 1;
}

.title {
  font-size: 3rem;
  font-weight: 700;
  background: linear-gradient(135deg, #ec4899 0%, #a855f7 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  margin-bottom: 1rem;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 1rem;
}

.heart {
  font-size: 3.5rem;
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

.subtitle {
  font-size: 1.25rem;
  color: var(--color-text-light);
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
  padding: 1rem 2rem;
  border-radius: 12px;
  font-weight: 600;
  transition: all 0.3s ease;
  text-decoration: none;
}

.btn-primary {
  background: linear-gradient(135deg, #ec4899 0%, #a855f7 100%);
  color: white;
  border: none;
}

.btn-primary:hover {
  transform: translateY(-2px);
  box-shadow: 0 10px 25px -5px rgba(236, 72, 153, 0.4);
}

.btn-icon {
  font-size: 1.25rem;
}

/* 装饰动画 */
.decorations {
  position: absolute;
  inset: 0;
  pointer-events: none;
  overflow: hidden;
}

.floating-heart {
  position: absolute;
  bottom: -30px;
  color: rgba(236, 72, 153, 0.3);
  animation: float 6s infinite;
}

@keyframes float {
  0% {
    transform: translateY(0) rotate(0deg);
    opacity: 0;
  }
  10% {
    opacity: 1;
  }
  90% {
    opacity: 1;
  }
  100% {
    transform: translateY(-100vh) rotate(360deg);
    opacity: 0;
  }
}

/* 功能卡片 */
.features {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 2rem;
  padding: 2rem;
  max-width: 1200px;
  margin: 0 auto;
}

.feature-card {
  background: rgba(255, 255, 255, 0.8);
  backdrop-filter: blur(10px);
  border-radius: 16px;
  padding: 2rem;
  text-align: center;
  transition: all 0.3s ease;
  border: 1px solid rgba(236, 72, 153, 0.2);
}

.feature-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 20px 40px rgba(236, 72, 153, 0.2);
  border-color: var(--color-primary);
}

.feature-icon {
  font-size: 3rem;
  margin-bottom: 1rem;
}

.feature-card h3 {
  color: var(--color-text);
  margin-bottom: 0.5rem;
  font-size: 1.25rem;
}

.feature-card p {
  color: var(--color-text-light);
  font-size: 0.875rem;
}

@media (max-width: 768px) {
  .title {
    font-size: 2rem;
  }

  .hero {
    padding: 2rem 1rem;
  }

  .features {
    grid-template-columns: 1fr;
    padding: 1rem;
  }
}
</style>
