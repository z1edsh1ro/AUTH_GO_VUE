import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useAuthStore = defineStore('auth', () => {
  const token = ref<string | null>(localStorage.getItem('token'))
  const isAuthenticated = ref<boolean>(!!token.value)

  const setToken = (newToken: string) => {
    token.value = newToken
    isAuthenticated.value = true
    localStorage.setItem('token', newToken)
  }

  const clearToken = () => {
    token.value = null
    isAuthenticated.value = false
    localStorage.removeItem('token')
  }

  const getToken = () => {
    return token.value
  }

  return {
    isAuthenticated,
    setToken,
    clearToken,
    getToken
  }
}) 