import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { decodeJwt } from '@/utils/jwt'
import type { JwtPayload } from './../types/jwtPayload '

const JWT_KEY = 'jwt'

export const useAuthStore = defineStore('auth', () => {
  const jwt = ref<string | null>(null)
  const jwtPayload = ref<JwtPayload | null>(null)
  const isAuthenticated = computed(() => loadJwt())

  const setJwt = (newjwt: string): void => {
    jwt.value = newjwt
    localStorage.setItem(JWT_KEY, newjwt)
    jwtPayload.value = decodeJwt(newjwt)
  }

  const clearJwt = (): void => {
    jwt.value = null
    jwtPayload.value = null
    localStorage.removeItem(JWT_KEY)
  }

  const loadJwt = (): boolean => {
    const loadJwt = localStorage.getItem(JWT_KEY)

    if (!loadJwt?.trim()) {
      return false
    }

    const payload = decodeJwt(loadJwt)

    if (!payload) {
      return false
    }

    jwt.value = loadJwt
    jwtPayload.value = payload

    return true
  }

  return {
    jwt,
    jwtPayload,
    isAuthenticated,
    loadJwt,
    setJwt,
    clearJwt,
  }
})
