import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import type { User } from './user'

export const useAuthStore = defineStore('auth', () => {
  const token = ref<string | null>(localStorage.getItem('token'))
  const storedUser = localStorage.getItem('user')
  const user = ref<User | null>(storedUser ? JSON.parse(storedUser) : null)

  const isAuthenticated = computed(() => !!token.value && !!user.value)

  const setToken = (newToken: string) => {
    token.value = newToken
    localStorage.setItem('token', newToken)
  }

  const setUser = (userData: User) => {
    user.value = userData
    localStorage.setItem('user', JSON.stringify(userData))
  }

  const clearToken = () => {
    token.value = null
    user.value = null
    localStorage.removeItem('token')
    localStorage.removeItem('user')
  }

  const getToken = () => {
    return token.value
  }

  const getUser = () => {
    return user.value
  }

  return {
    token,
    user,
    isAuthenticated,
    setToken,
    setUser,
    clearToken,
    getToken,
    getUser,
  }
})
