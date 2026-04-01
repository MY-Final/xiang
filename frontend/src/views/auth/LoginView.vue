<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useUserStore } from '@/stores/user'

const router = useRouter()
const route = useRoute()
const userStore = useUserStore()

const username = ref('')
const password = ref('')
const loading = ref(false)
const error = ref('')
const hearts = ref<Array<{ id: number; left: string; delay: string; duration: string; size: string }>>([])

// 生成浮动爱心
onMounted(() => {
  hearts.value = Array.from({ length: 25 }, (_, i) => ({
    id: i,
    left: Math.random() * 100 + '%',
    delay: Math.random() * 5 + 's',
    duration: (Math.random() * 3 + 3) + 's',
    size: (Math.random() * 10 + 15) + 'px',
  }))
})

const handleLogin = async () => {
  error.value = ''
  loading.value = true

  try {
    await userStore.login({
      username: username.value,
      password: password.value,
    })

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

    <!-- 光晕效果 -->
    <div class="glow-effect glow-1"></div>
    <div class="glow-effect glow-2"></div>

    <div class="login-card">
      <!-- 顶部装饰 -->
      <div class="card-decoration">
        <div class="decoration-line"></div>
      </div>

      <div class="login-header">
        <div class="logo-wrapper">
          <div class="logo-heart">
            <svg viewBox="0 0 24 24" fill="currentColor">
              <path d="M12 21.35l-1.45-1.32C5.4 15.36 2 12.28 2 8.5 2 5.42 4.42 3 7.5 3c1.74 0 3.41.81 4.5 2.09C13.09 3.81 14.76 3 16.5 3 19.58 3 22 5.42 22 8.5c0 3.78-3.4 6.86-8.55 11.54L12 21.35z"/>
            </svg>
          </div>
          <div class="logo-ring"></div>
        </div>
        <h1 class="title">湘遇恋爱博客</h1>
        <p class="subtitle">记录属于我们的浪漫时光</p>
      </div>

      <form class="login-form" @submit.prevent="handleLogin">
        <div class="form-group">
          <label for="username" class="form-label">
            <svg class="label-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"/>
              <circle cx="12" cy="7" r="4"/>
            </svg>
            用户名
          </label>
          <div class="input-wrapper">
            <input
              id="username"
              v-model="username"
              type="text"
              class="form-input"
              placeholder="请输入用户名"
              required
              autocomplete="username"
            />
          </div>
        </div>

        <div class="form-group">
          <label for="password" class="form-label">
            <svg class="label-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <rect x="3" y="11" width="18" height="11" rx="2" ry="2"/>
              <path d="M7 11V7a5 5 0 0 1 10 0v4"/>
            </svg>
            密码
          </label>
          <div class="input-wrapper">
            <input
              id="password"
              v-model="password"
              type="password"
              class="form-input"
              placeholder="请输入密码"
              required
              autocomplete="current-password"
            />
          </div>
        </div>

        <div v-if="error" class="error-message">
          <svg class="error-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <circle cx="12" cy="12" r="10"/>
            <line x1="12" y1="8" x2="12" y2="12"/>
            <line x1="12" y1="16" x2="12.01" y2="16"/>
          </svg>
          {{ error }}
        </div>

        <button type="submit" class="login-btn" :disabled="loading">
          <span v-if="!loading" class="btn-content">
            <span>登录</span>
            <svg class="btn-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M5 12h14M12 5l7 7-7 7"/>
            </svg>
          </span>
          <span v-else class="loading-spinner">
            <span class="dot"></span>
            <span class="dot"></span>
            <span class="dot"></span>
          </span>
        </button>
      </form>

      <div class="login-footer">
        <div class="footer-divider"></div>
        <p class="footer-text">
          <svg class="footer-heart" viewBox="0 0 24 24" fill="currentColor">
            <path d="M12 21.35l-1.45-1.32C5.4 15.36 2 12.28 2 8.5 2 5.42 4.42 3 7.5 3c1.74 0 3.41.81 4.5 2.09C13.09 3.81 14.76 3 16.5 3 19.58 3 22 5.42 22 8.5c0 3.78-3.4 6.86-8.55 11.54L12 21.35z"/>
          </svg>
          执子之手，与子偕老
          <svg class="footer-heart" viewBox="0 0 24 24" fill="currentColor">
            <path d="M12 21.35l-1.45-1.32C5.4 15.36 2 12.28 2 8.5 2 5.42 4.42 3 7.5 3c1.74 0 3.41.81 4.5 2.09C13.09 3.81 14.76 3 16.5 3 19.58 3 22 5.42 22 8.5c0 3.78-3.4 6.86-8.55 11.54L12 21.35z"/>
          </svg>
        </p>
        <RouterLink to="/register" class="register-link">
          还没有账号？立即注册
        </RouterLink>
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

/* 背景装饰 */
.bg-hearts {
  position: absolute;
  inset: 0;
  pointer-events: none;
  overflow: hidden;
  z-index: 1;
}

.heart {
  position: absolute;
  bottom: -50px;
  color: rgba(236, 72, 153, 0.25);
  animation: float-up linear infinite;
}

@keyframes float-up {
  0% {
    transform: translateY(0) rotate(0deg) scale(0.8);
    opacity: 0;
  }
  10% {
    opacity: 0.6;
  }
  50% {
    opacity: 0.8;
  }
  90% {
    opacity: 0.6;
  }
  100% {
    transform: translateY(-100vh) rotate(360deg) scale(1);
    opacity: 0;
  }
}

/* 光晕效果 */
.glow-effect {
  position: absolute;
  border-radius: 50%;
  filter: blur(80px);
  opacity: 0.4;
  animation: glow-pulse 4s ease-in-out infinite;
}

.glow-1 {
  width: 400px;
  height: 400px;
  background: radial-gradient(circle, #f472b6 0%, transparent 70%);
  top: -100px;
  right: -100px;
  animation-delay: 0s;
}

.glow-2 {
  width: 300px;
  height: 300px;
  background: radial-gradient(circle, #c084fc 0%, transparent 70%);
  bottom: -50px;
  left: -50px;
  animation-delay: 2s;
}

@keyframes glow-pulse {
  0%, 100% { opacity: 0.3; transform: scale(1); }
  50% { opacity: 0.5; transform: scale(1.1); }
}

/* 登录卡片 */
.login-card {
  background: rgba(255, 255, 255, 0.9);
  backdrop-filter: blur(20px);
  border-radius: 32px;
  padding: 3.5rem 3rem;
  width: 100%;
  max-width: 460px;
  box-shadow:
    0 25px 50px -12px rgba(236, 72, 153, 0.25),
    0 0 0 1px rgba(255, 255, 255, 0.5) inset;
  position: relative;
  z-index: 10;
  animation: card-appear 0.6s ease-out;
}

@keyframes card-appear {
  from {
    opacity: 0;
    transform: translateY(20px) scale(0.95);
  }
  to {
    opacity: 1;
    transform: translateY(0) scale(1);
  }
}

/* 卡片顶部装饰 */
.card-decoration {
  position: absolute;
  top: 0;
  left: 50%;
  transform: translateX(-50%);
  width: 200px;
  height: 4px;
  background: linear-gradient(90deg, transparent, #ec4899, #a855f7, #ec4899, transparent);
  border-radius: 0 0 8px 8px;
}

/* Logo 区域 */
.login-header {
  text-align: center;
  margin-bottom: 2.5rem;
}

.logo-wrapper {
  position: relative;
  display: inline-block;
  margin-bottom: 1.5rem;
}

.logo-heart {
  width: 80px;
  height: 80px;
  color: #ec4899;
  animation: heart-pulse 1.5s ease-in-out infinite;
  filter: drop-shadow(0 4px 8px rgba(236, 72, 153, 0.3));
}

.logo-ring {
  position: absolute;
  inset: -10px;
  border: 2px solid rgba(236, 72, 153, 0.2);
  border-radius: 50%;
  animation: ring-rotate 8s linear infinite;
}

@keyframes ring-rotate {
  from { transform: rotate(0deg); }
  to { transform: rotate(360deg); }
}

@keyframes heart-pulse {
  0%, 100% { transform: scale(1); }
  50% { transform: scale(1.1); }
}

.title {
  font-size: 2rem;
  font-weight: 700;
  background: linear-gradient(135deg, #ec4899 0%, #a855f7 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  margin-bottom: 0.5rem;
  letter-spacing: -0.02em;
}

.subtitle {
  color: #9ca3af;
  font-size: 0.9375rem;
  font-weight: 400;
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

.form-label {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  font-size: 0.875rem;
  font-weight: 600;
  color: #4b5563;
}

.label-icon {
  width: 18px;
  height: 18px;
  color: #ec4899;
}

.input-wrapper {
  position: relative;
}

.form-input {
  width: 100%;
  padding: 0.875rem 1.25rem;
  border: 2px solid #f3e8e8;
  border-radius: 14px;
  font-size: 1rem;
  transition: all 0.25s ease;
  background: rgba(255, 255, 255, 0.8);
  color: #1f2937;
}

.form-input::placeholder {
  color: #9ca3af;
}

.form-input:focus {
  outline: none;
  border-color: #ec4899;
  background: white;
  box-shadow:
    0 0 0 4px rgba(236, 72, 153, 0.1),
    0 4px 12px rgba(236, 72, 153, 0.15);
}

/* 错误消息 */
.error-message {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  color: #ef4444;
  font-size: 0.875rem;
  padding: 0.75rem 1rem;
  background: rgba(239, 68, 68, 0.08);
  border-radius: 10px;
  border: 1px solid rgba(239, 68, 68, 0.2);
  animation: shake 0.4s ease-in-out;
}

@keyframes shake {
  0%, 100% { transform: translateX(0); }
  25% { transform: translateX(-5px); }
  75% { transform: translateX(5px); }
}

.error-icon {
  width: 18px;
  height: 18px;
  flex-shrink: 0;
}

/* 登录按钮 */
.login-btn {
  position: relative;
  width: 100%;
  background: linear-gradient(135deg, #ec4899 0%, #a855f7 100%);
  color: white;
  border: none;
  padding: 1rem 1.5rem;
  border-radius: 14px;
  font-size: 1rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
  overflow: hidden;
}

.login-btn::before {
  content: '';
  position: absolute;
  inset: 0;
  background: linear-gradient(135deg, rgba(255,255,255,0.2) 0%, transparent 50%);
  opacity: 0;
  transition: opacity 0.3s ease;
}

.login-btn:hover:not(:disabled)::before {
  opacity: 1;
}

.login-btn:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow:
    0 15px 30px -5px rgba(236, 72, 153, 0.4),
    0 0 20px rgba(236, 72, 153, 0.3);
}

.login-btn:active:not(:disabled) {
  transform: translateY(0);
}

.login-btn:disabled {
  opacity: 0.7;
  cursor: not-allowed;
}

.btn-content {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 0.5rem;
}

.btn-icon {
  width: 18px;
  height: 18px;
  transition: transform 0.3s ease;
}

.login-btn:hover .btn-icon {
  transform: translateX(4px);
}

.loading-spinner {
  display: flex;
  justify-content: center;
  gap: 6px;
}

.dot {
  width: 10px;
  height: 10px;
  background: white;
  border-radius: 50%;
  animation: bounce 1.4s ease-in-out infinite both;
}

.dot:nth-child(1) { animation-delay: -0.32s; }
.dot:nth-child(2) { animation-delay: -0.16s; }

@keyframes bounce {
  0%, 80%, 100% {
    transform: scale(0);
    opacity: 0.5;
  }
  40% {
    transform: scale(1);
    opacity: 1;
  }
}

/* 页脚 */
.login-footer {
  margin-top: 2.5rem;
  text-align: center;
}

.footer-divider {
  height: 1px;
  background: linear-gradient(90deg, transparent, #f3e8e8, transparent);
  margin-bottom: 1rem;
}

.footer-text {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 0.5rem;
  color: #9ca3af;
  font-size: 0.875rem;
  margin-bottom: 1rem;
}

.footer-heart {
  width: 16px;
  height: 16px;
  color: #ec4899;
  animation: heart-pulse 1s ease-in-out infinite;
}

.register-link {
  color: #ec4899;
  font-size: 0.875rem;
  font-weight: 600;
  text-decoration: none;
  transition: color 0.2s ease;
}

.register-link:hover {
  color: #db2777;
  text-decoration: underline;
}

/* 响应式 */
@media (max-width: 480px) {
  .login-container {
    padding: 1rem;
  }

  .login-card {
    padding: 2.5rem 1.5rem;
    border-radius: 24px;
  }

  .title {
    font-size: 1.5rem;
  }

  .logo-heart {
    width: 60px;
    height: 60px;
  }

  .form-input {
    padding: 0.75rem 1rem;
  }
}
</style>
