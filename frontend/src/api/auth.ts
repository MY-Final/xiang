import api from './index'

export interface LoginRequest {
  username: string
  password: string
}

export interface UserProfile {
  id: number
  username: string
  nickname: string
  avatar: string
}

export interface LoginResponse {
  token: string
  expires_in: number
  user: UserProfile
}

/**
 * 管理员登录
 */
export const login = async (data: LoginRequest): Promise<LoginResponse> => {
  return api.post('/auth/login', data)
}

/**
 * 获取当前管理员信息
 */
export const getMe = async (): Promise<UserProfile> => {
  return api.get('/auth/me')
}

/**
 * 退出登录
 */
export const logout = async (): Promise<void> => {
  return api.post('/auth/logout')
}
