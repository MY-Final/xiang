import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import {
  login as loginApi,
  logout as logoutApi,
  getMe,
  type LoginRequest,
  type UserProfile,
} from '@/api/auth'

export const useUserStore = defineStore('user', () => {
  const user = ref<UserProfile | null>(null)
  const token = ref<string>(localStorage.getItem('token'))

  const isLoggedIn = computed(() => !!token.value)

  /**
   * 登录
   */
  const login = async (credentials: LoginRequest) => {
    const response = await loginApi(credentials)
    token.value = response.token
    user.value = response.user
    localStorage.setItem('token', response.token)
    return response
  }

  /**
   * 获取当前用户信息
   */
  const fetchUser = async () => {
    try {
      user.value = await getMe()
    } catch {
      // Token 无效，清空状态
      logout()
    }
  }

  /**
   * 退出登录
   */
  const logout = async () => {
    try {
      await logoutApi()
    } finally {
      token.value = null
      user.value = null
      localStorage.removeItem('token')
    }
  }

  // 初始化：如果已有 token，尝试获取用户信息
  if (token.value) {
    fetchUser()
  }

  return {
    user,
    token,
    isLoggedIn,
    login,
    logout,
    fetchUser,
  }
})
