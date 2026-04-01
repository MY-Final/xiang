<script setup lang="ts">
import { ref } from 'vue'
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

const showHeart = ref(true)
setInterval(() => {
  showHeart.value = !showHeart.value
}, 800)

const handleRegister = async () => {
  if (form.value.password !== form.value.confirmPassword) {
    ElMessage.error('两次密码输入不一致')
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
    <div class="bg-hearts">
      <div
        class="heart"
        v-for="i in 15"
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

    <div class="register-card">
      <div class="register-header">
        <div class="logo-heart" :class="{ pulse: showHeart }">💕</div>
        <h1 class="title">加入我们</h1>
        <p class="subtitle">开始记录你们的浪漫时光</p>
      </div>

      <el-form :model="form" label-position="top" size="large" @submit.prevent="handleRegister">
        <el-form-item label="用户名" required>
          <el-input v-model="form.username" placeholder="3-32 位字符" />
        </el-form-item>

        <el-form-item label="昵称" required>
          <el-input v-model="form.nickname" placeholder="如何称呼你" />
        </el-form-item>

        <el-form-item label="邮箱" required>
          <el-input v-model="form.email" type="email" placeholder="your@email.com" />
        </el-form-item>

        <el-form-item label="密码" required>
          <el-input v-model="form.password" type="password" placeholder="6-32 位密码" />
        </el-form-item>

        <el-form-item label="确认密码" required>
          <el-input v-model="form.confirmPassword" type="password" placeholder="再次输入密码" />
        </el-form-item>

        <el-form-item label="纪念日（可选）">
          <el-date-picker
            v-model="form.anniversary"
            type="date"
            placeholder="你们的特别日子"
            style="width: 100%"
            format="YYYY-MM-DD"
            value-format="YYYY-MM-DD"
          />
        </el-form-item>

        <el-button type="primary" native-type="submit" :loading="loading" class="register-btn">
          注册
        </el-button>
      </el-form>

      <div class="login-link">已有账号？<RouterLink to="/login">立即登录</RouterLink></div>
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
  background: linear-gradient(135deg, #fce7f3 0%, #f5d0fe 50%, #e9d5ff 100%);
  position: relative;
  overflow: hidden;
}

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

.register-card {
  background: rgba(255, 255, 255, 0.95);
  border-radius: 24px;
  padding: 3rem;
  width: 100%;
  max-width: 480px;
  box-shadow: 0 25px 50px -12px rgba(236, 72, 153, 0.25);
  position: relative;
  z-index: 1;
}

.register-header {
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

:deep(.el-form-item__label) {
  color: #4b5563;
  font-weight: 500;
}

:deep(.el-input__wrapper) {
  border-radius: 12px;
}

:deep(.el-button.primary) {
  background: linear-gradient(135deg, #ec4899 0%, #a855f7 100%);
  border: none;
}

.register-btn {
  width: 100%;
  padding: 1rem;
  font-size: 1rem;
  font-weight: 600;
  margin-top: 1rem;
}

.login-link {
  text-align: center;
  margin-top: 1.5rem;
  color: #6b7280;
  font-size: 0.875rem;
}

.login-link a {
  color: #ec4899;
  font-weight: 600;
}

@media (max-width: 480px) {
  .register-card {
    padding: 2rem 1.5rem;
  }
}
</style>
