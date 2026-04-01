<script setup lang="ts">
import { ref } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useUserStore } from '@/stores/user'

const router = useRouter()
const route = useRoute()
const userStore = useUserStore()

const username = ref('')
const password = ref('')
const loading = ref(false)
const error = ref('')
const showHeart = ref(true)

// 心跳动画
setInterval(() => {
  showHeart.value = !showHeart.value
}, 800)

const handleLogin = async () => {
  error.value = ''
  loading.value = true

  try {
    await userStore.login({
      username: username.value,
      password: password.value,
    })

    // 登录成功，跳转到首页或重定向地址
    const redirect = route.query.redirect as string
    router.push(redirect || '/')
  } catch (err: any) {
    error.value = err.response?.data?.message || '登录失败，请检查用户名和密码'
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="login-container">
    <!-- 背景装饰 -->
    <div class="bg-hearts">
      <div
        class="heart"
        v-for="i in 20"
        :key="i"
        :style="{
          left: Math.random() * 100 + '%',
          animationDelay: Math.random() * 5 + 's',
          animationDuration: Math.random() * 3 + 2 + 's',
        }"
      >
        ❤
      </div>
    </div>

    <div class="login-card">
      <div class="login-header">
        <div class="logo-heart" :class="{ pulse: showHeart }">💕</div>
        <h1 class="title">湘遇恋爱博客</h1>
        <p class="subtitle">记录我们的浪漫时光</p>
      </div>

      <form class="login-form" @submit.prevent="handleLogin">
        <div class="form-group">
          <label for="username">用户名</label>
          <input
            id="username"
            v-model="username"
            type="text"
            placeholder="请输入用户名"
            required
            autocomplete="username"
          />
        </div>

        <div class="form-group">
          <label for="password">密码</label>
          <input
            id="password"
            v-model="password"
            type="password"
            placeholder="请输入密码"
            required
            autocomplete="current-password"
          />
        </div>

        <div v-if="error" class="error-message">{{ error }}</div>

        <button type="submit" class="login-btn" :disabled="loading">
          <span v-if="!loading">登录</span>
          <span v-else class="loading-spinner">
            <span class="dot"></span>
            <span class="dot"></span>
            <span class="dot"></span>
          </span>
        </button>
      </form>

      <div class="login-footer">
        <p>💗 执子之手，与子偕老 💗</p>
      </div>
    </div>
  </div>
</template>

<style scoped>
.login-container {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 2rem;
  background: linear-gradient(135deg, #fce7f3 0%, #f5d0fe 50%, #e9d5ff 100%);
  position: relative;
  overflow: hidden;
}

/* 背景爱心装饰 */
.bg-hearts {
  position: absolute;
  inset: 0;
  pointer-events: none;
  overflow: hidden;
}

.heart {
  position: absolute;
  bottom: -50px;
  font-size: 20px;
  color: rgba(236, 72, 153, 0.3);
  animation: float-up 5s infinite;
}

@keyframes float-up {
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

/* 登录卡片 */
.login-card {
  background: rgba(255, 255, 255, 0.95);
  border-radius: 24px;
  padding: 3rem;
  width: 100%;
  max-width: 420px;
  box-shadow: 0 25px 50px -12px rgba(236, 72, 153, 0.25);
  position: relative;
  z-index: 1;
}

.login-header {
  text-align: center;
  margin-bottom: 2rem;
}

.logo-heart {
  font-size: 4rem;
  display: inline-block;
  margin-bottom: 1rem;
}

.logo-heart.pulse {
  animation: heart-pulse 0.8s ease-in-out;
}

@keyframes heart-pulse {
  0%,
  100% {
    transform: scale(1);
  }
  50% {
    transform: scale(1.2);
  }
}

.title {
  font-size: 1.875rem;
  font-weight: 700;
  background: linear-gradient(135deg, #ec4899 0%, #a855f7 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  margin-bottom: 0.5rem;
}

.subtitle {
  color: #9ca3af;
  font-size: 0.875rem;
}

/* 表单样式 */
.login-form {
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.form-group label {
  font-size: 0.875rem;
  font-weight: 500;
  color: #4b5563;
}

.form-group input {
  padding: 0.75rem 1rem;
  border: 2px solid #f3e8e8;
  border-radius: 12px;
  font-size: 1rem;
  transition: all 0.3s ease;
  background: #fafafa;
}

.form-group input:focus {
  outline: none;
  border-color: #ec4899;
  background: white;
  box-shadow: 0 0 0 3px rgba(236, 72, 153, 0.1);
}

.error-message {
  color: #ef4444;
  font-size: 0.875rem;
  text-align: center;
  padding: 0.5rem;
  background: rgba(239, 68, 68, 0.1);
  border-radius: 8px;
}

.login-btn {
  background: linear-gradient(135deg, #ec4899 0%, #a855f7 100%);
  color: white;
  border: none;
  padding: 1rem;
  border-radius: 12px;
  font-size: 1rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
  position: relative;
  overflow: hidden;
}

.login-btn:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 10px 25px -5px rgba(236, 72, 153, 0.4);
}

.login-btn:active:not(:disabled) {
  transform: translateY(0);
}

.login-btn:disabled {
  opacity: 0.7;
  cursor: not-allowed;
}

.loading-spinner {
  display: flex;
  justify-content: center;
  gap: 4px;
}

.dot {
  width: 8px;
  height: 8px;
  background: white;
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

/* 页脚 */
.login-footer {
  margin-top: 2rem;
  text-align: center;
  color: #9ca3af;
  font-size: 0.875rem;
}

/* 响应式 */
@media (max-width: 480px) {
  .login-card {
    padding: 2rem 1.5rem;
  }

  .title {
    font-size: 1.5rem;
  }

  .logo-heart {
    font-size: 3rem;
  }
}
</style>
