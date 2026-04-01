<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import { useUserStore } from '@/stores/user'

const userStore = useUserStore()

const handleLogout = async () => {
  await userStore.logout()
}
</script>

<template>
  <div id="app">
    <header class="app-header" v-if="userStore.isLoggedIn">
      <div class="header-content">
        <RouterLink to="/" class="logo">
          <span class="logo-icon">💕</span>
          <span class="logo-text">湘遇</span>
        </RouterLink>

        <nav class="nav">
          <RouterLink to="/">首页</RouterLink>
          <RouterLink to="/about">关于</RouterLink>
        </nav>

        <div class="user-menu">
          <span class="username">{{ userStore.user?.nickname || userStore.user?.username }}</span>
          <button @click="handleLogout" class="logout-btn">退出</button>
        </div>
      </div>
    </header>

    <main class="main-content">
      <RouterView />
    </main>

    <footer class="app-footer" v-if="!userStore.isLoggedIn">
      <p>💗 记录我们的浪漫时光 💗</p>
    </footer>
  </div>
</template>

<style scoped>
.app-header {
  background: rgba(255, 255, 255, 0.9);
  backdrop-filter: blur(10px);
  border-bottom: 1px solid rgba(236, 72, 153, 0.2);
  position: sticky;
  top: 0;
  z-index: 100;
}

.header-content {
  max-width: 1200px;
  margin: 0 auto;
  padding: 1rem 2rem;
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.logo {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  font-size: 1.25rem;
  font-weight: 700;
  background: linear-gradient(135deg, #ec4899 0%, #a855f7 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.logo-icon {
  font-size: 1.5rem;
  -webkit-text-fill-color: initial;
}

.nav {
  display: flex;
  gap: 2rem;
}

.nav a {
  color: var(--color-text);
  font-weight: 500;
  position: relative;
  transition: color 0.2s ease;
}

.nav a:hover {
  color: var(--color-primary);
}

.nav a.router-link-exact-active {
  color: var(--color-primary);
}

.nav a.router-link-exact-active::after {
  content: '';
  position: absolute;
  bottom: -4px;
  left: 0;
  right: 0;
  height: 2px;
  background: linear-gradient(90deg, #ec4899, #a855f7);
  border-radius: 2px;
}

.user-menu {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.username {
  color: var(--color-text-light);
  font-size: 0.875rem;
}

.logout-btn {
  background: linear-gradient(135deg, #ec4899 0%, #a855f7 100%);
  color: white;
  border: none;
  padding: 0.5rem 1rem;
  border-radius: 8px;
  font-size: 0.875rem;
  cursor: pointer;
  transition: all 0.3s ease;
}

.logout-btn:hover {
  transform: translateY(-1px);
  box-shadow: 0 4px 12px rgba(236, 72, 153, 0.3);
}

.main-content {
  min-height: calc(100vh - 73px);
}

.app-footer {
  text-align: center;
  padding: 2rem;
  color: var(--color-text-light);
  font-size: 0.875rem;
}

@media (max-width: 768px) {
  .header-content {
    padding: 1rem;
  }

  .nav {
    gap: 1rem;
  }

  .username {
    display: none;
  }
}
</style>
