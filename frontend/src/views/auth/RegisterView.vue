<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'

const router = useRouter()
const loading = ref(false)

const form = ref({
  username: '',
  password: '',
  confirmPassword: '',
  nickname: '',
  email: '',
  anniversary: '',
})

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

const handleRegister = async () => {
  if (form.value.password !== form.value.confirmPassword) {
    ElMessage.error('两次密码输入不一致')
    return
  }

  if (form.value.password.length < 6) {
    ElMessage.error('密码至少 6 位')
    return
  }

  loading.value = true
  try {
    // TODO: 调用注册 API
    // await api.post('/user/register', form.value)

    ElMessage.success('注册成功，请登录')
    router.push('/login')
  } catch (error: any) {
    ElMessage.error(error.response?.data?.message || '注册失败')
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="register-container">
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

    <!-- 光晕装饰 -->
    <div class="glow-orb glow-1"></div>
    <div class="glow-orb glow-2"></div>

    <div class="register-card">
      <!-- 顶部装饰 -->
      <div class="card-decoration"></div>

      <div class="register-header">
        <div class="logo-wrapper">
          <div class="logo-heart">
            <svg viewBox="0 0 24 24" fill="currentColor">
              <path d="M12 21.35l-1.45-1.32C5.4 15.36 2 12.28 2 8.5 2 5.42 4.42 3 7.5 3c1.74 0 3.41.81 4.5 2.09C13.09 3.81 14.76 3 16.5 3 19.58 3 22 5.42 22 8.5c0 3.78-3.4 6.86-8.55 11.54L12 21.35z"/>
            </svg>
          </div>
          <div class="logo-ring"></div>
        </div>
        <h1 class="title">加入我们</h1>
        <p class="subtitle">开始记录你们的浪漫时光</p>
      </div>

      <el-form :model="form" label-position="top" size="large" @submit.prevent="handleRegister">
        <div class="form-row">
          <el-form-item label="用户名" required>
            <div class="input-wrapper">
              <svg class="input-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"/>
                <circle cx="12" cy="7" r="4"/>
              </svg>
              <el-input v-model="form.username" placeholder="3-32 位字符" class="glass-input" />
            </div>
          </el-form-item>
        </div>

        <el-form-item label="昵称" required>
          <div class="input-wrapper">
            <svg class="input-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"/>
              <circle cx="12" cy="7" r="4"/>
            </svg>
            <el-input v-model="form.nickname" placeholder="如何称呼你" class="glass-input" />
          </div>
        </el-form-item>

        <el-form-item label="邮箱" required>
          <div class="input-wrapper">
            <svg class="input-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M4 4h16c1.1 0 2 .9 2 2v12c0 1.1-.9 2-2 2H4c-1.1 0-2-.9-2-2V6c0-1.1.9-2 2-2z"/>
              <polyline points="22,6 12,13 2,6"/>
            </svg>
            <el-input v-model="form.email" type="email" placeholder="your@email.com" class="glass-input" />
          </div>
        </el-form-item>

        <el-form-item label="密码" required>
          <div class="input-wrapper">
            <svg class="input-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <rect x="3" y="11" width="18" height="11" rx="2" ry="2"/>
              <path d="M7 11V7a5 5 0 0 1 10 0v4"/>
            </svg>
            <el-input v-model="form.password" type="password" placeholder="6-32 位密码" class="glass-input" />
          </div>
        </el-form-item>

        <el-form-item label="确认密码" required>
          <div class="input-wrapper">
            <svg class="input-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <rect x="3" y="11" width="18" height="11" rx="2" ry="2"/>
              <path d="M7 11V7a5 5 0 0 1 10 0v4"/>
            </svg>
            <el-input v-model="form.confirmPassword" type="password" placeholder="再次输入密码" class="glass-input" />
          </div>
        </el-form-item>

        <el-form-item label="纪念日（可选）">
          <div class="input-wrapper">
            <svg class="input-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <rect x="3" y="4" width="18" height="18" rx="2" ry="2"/>
              <line x1="16" y1="2" x2="16" y2="6"/>
              <line x1="8" y1="2" x2="8" y2="6"/>
              <line x1="3" y1="10" x2="21" y2="10"/>
            </svg>
            <el-date-picker
              v-model="form.anniversary"
              type="date"
              placeholder="你们的特别日子"
              style="width: 100%"
              format="YYYY-MM-DD"
              value-format="YYYY-MM-DD"
              class="glass-input"
            />
          </div>
        </el-form-item>

        <el-button type="primary" native-type="submit" :loading="loading" class="register-btn">
          <span v-if="!loading">开始旅程</span>
          <span v-else class="loading-spinner">
            <span class="dot"></span>
            <span class="dot"></span>
            <span class="dot"></span>
          </span>
        </el-button>
      </el-form>

      <div class="login-link">
        已有账号？<RouterLink to="/login">立即登录</RouterLink>
      </div>
    </div>
  </div>
</template>

<style scoped>
.register-container {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 2rem;
  background: linear-gradient(135deg, #fdf2f8 0%, #f5d0fe 50%, #e9d5ff 100%);
  position: relative;
  overflow: hidden;
}

/* ═══════════════════════════════════════════════════════════
   背景爱心动画
   ═══════════════════════════════════════════════════════════ */
.bg-hearts {
  position: absolute;
  inset: 0;
  pointer-events: none;
  z-index: 1;
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

/* 光晕装饰 */
.glow-orb {
  position: absolute;
  border-radius: 50%;
  filter: blur(80px);
  opacity: 0.4;
  pointer-events: none;
}

.glow-1 {
  width: 400px;
  height: 400px;
  background: radial-gradient(circle, #f472b6 0%, transparent 70%);
  top: -150px;
  right: -150px;
  animation: glow-pulse 4s ease-in-out infinite;
}

.glow-2 {
  width: 300px;
  height: 300px;
  background: radial-gradient(circle, #c084fc 0%, transparent 70%);
  bottom: -100px;
  left: -100px;
  animation: glow-pulse 4s ease-in-out infinite 2s;
}

@keyframes glow-pulse {
  0%, 100% { opacity: 0.3; transform: scale(1); }
  50% { opacity: 0.5; transform: scale(1.1); }
}

/* ═══════════════════════════════════════════════════════════
   注册卡片
   ═══════════════════════════════════════════════════════════ */
.register-card {
  background: rgba(255, 255, 255, 0.75);
  backdrop-filter: blur(20px);
  -webkit-backdrop-filter: blur(20px);
  border-radius: 32px;
  padding: 3.5rem 3rem;
  width: 100%;
  max-width: 500px;
  box-shadow:
    0 25px 60px rgba(236, 72, 153, 0.2),
    inset 0 1px 0 rgba(255, 255, 255, 0.8);
  border: 1px solid rgba(255, 255, 255, 0.6);
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

/* 顶部装饰 */
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

/* ═══════════════════════════════════════════════════════════
   头部 Logo
   ═══════════════════════════════════════════════════════════ */
.register-header {
  text-align: center;
  margin-bottom: 2.5rem;
}

.logo-wrapper {
  position: relative;
  display: inline-block;
  margin-bottom: 1.5rem;
}

.logo-heart {
  width: 72px;
  height: 72px;
  color: #ec4899;
  animation: heart-pulse 1.5s ease-in-out infinite;
  filter: drop-shadow(0 4px 8px rgba(236, 72, 153, 0.3));
}

@keyframes heart-pulse {
  0%, 100% { transform: scale(1); }
  50% { transform: scale(1.12); }
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

.title {
  font-size: 1.875rem;
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
}

/* ═══════════════════════════════════════════════════════════
   表单样式
   ═══════════════════════════════════════════════════════════ */
:deep(.el-form-item) {
  margin-bottom: 1.25rem;
}

:deep(.el-form-item__label) {
  color: #4b5563;
  font-weight: 600;
  font-size: 0.875rem;
  margin-bottom: 0.5rem;
}

.input-wrapper {
  position: relative;
  display: flex;
  align-items: center;
}

.input-icon {
  position: absolute;
  left: 1rem;
  width: 20px;
  height: 20px;
  color: rgba(236, 72, 153, 0.5);
  z-index: 2;
  pointer-events: none;
  transition: color 0.2s ease;
}

.input-wrapper:focus-within .input-icon {
  color: #ec4899;
}

:deep(.el-input__wrapper),
:deep(.el-date-picker__wrapper) {
  border-radius: 14px;
  padding: 0.75rem 1rem 0.75rem 2.75rem;
  background: rgba(255, 255, 255, 0.8);
  backdrop-filter: blur(10px);
  border: 2px solid rgba(236, 72, 153, 0.15);
  box-shadow: none;
  transition: all 0.25s ease;
}

:deep(.el-input__wrapper:hover),
:deep(.el-date-picker__wrapper:hover) {
  background: rgba(255, 255, 255, 0.95);
}

:deep(.el-input__wrapper.is-focus),
:deep(.el-date-picker__wrapper:focus) {
  background: white;
  border-color: #ec4899;
  box-shadow: 0 0 0 4px rgba(236, 72, 153, 0.1);
}

/* 注册按钮 */
.register-btn {
  width: 100%;
  padding: 1rem;
  font-size: 1rem;
  font-weight: 600;
  margin-top: 0.5rem;
  background: linear-gradient(135deg, #ec4899 0%, #a855f7 100%);
  border: none;
  border-radius: 14px;
  box-shadow: 0 4px 15px rgba(236, 72, 153, 0.3);
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.register-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 25px rgba(236, 72, 153, 0.4);
}

.register-btn:active {
  transform: translateY(0);
}

/* 加载动画 */
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
.dot:nth-child(3) { animation-delay: 0s; }

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

/* ═══════════════════════════════════════════════════════════
   底部链接
   ═══════════════════════════════════════════════════════════ */
.login-link {
  text-align: center;
  margin-top: 2rem;
  padding-top: 1.5rem;
  border-top: 1px solid rgba(236, 72, 153, 0.15);
  color: #6b7280;
  font-size: 0.875rem;
}

.login-link a {
  color: #ec4899;
  font-weight: 600;
  text-decoration: none;
  transition: color 0.2s ease;
}

.login-link a:hover {
  color: #db2777;
  text-decoration: underline;
}

/* ═══════════════════════════════════════════════════════════
   响应式
   ═══════════════════════════════════════════════════════════ */
@media (max-width: 480px) {
  .register-container {
    padding: 1rem;
  }

  .register-card {
    padding: 2.5rem 1.5rem;
    border-radius: 24px;
  }

  .title {
    font-size: 1.5rem;
  }

  .logo-heart {
    width: 56px;
    height: 56px;
  }
}
</style>
