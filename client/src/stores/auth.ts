import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import type { User } from './user'

export const useAuthStore = defineStore('auth', () => {
  const token = ref<string | null>(localStorage.getItem('token'))
  const user = ref<User | null>(null)

  const isAuthenticated = computed(() => !!token.value)

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
    const storedToken = localStorage.getItem('token')
    if (storedToken !== token.value) {
      token.value = storedToken
    }
    return token.value
  }

  const getUser = () => {
    const storedUser = localStorage.getItem('user')
    if (storedUser) {
      user.value = JSON.parse(storedUser)
    }
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
