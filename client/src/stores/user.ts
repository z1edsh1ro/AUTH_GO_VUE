import { defineStore } from 'pinia'
import { ref } from 'vue'
import { useAuthStore } from './auth'

export interface User {
  id: string
  name: string
  email: number
  password: string
  created_at: string
  updated_at: string
}

export const useUserStore = defineStore('user', () => {
  const users = ref<User[]>([])
  const loading = ref(false)
  const error = ref<string | null>(null)
  const authStore = useAuthStore()

  const fetchUsers = async () => {
    loading.value = true
    error.value = null

    try {
      const token = authStore.getToken()
      if (!token) {
        throw new Error('No authentication token found')
      }

      const endPoint = `http://localhost:8080/api/auth/user/getAll`
      const myHeaders = new Headers()
      myHeaders.append('Authorization', `Bearer ${token}`)

      const requestOptions = {
        method: 'GET',
        headers: myHeaders,
      }

      const response = await fetch(endPoint, requestOptions)

      if (!response.ok) {
        if (response.status === 401) {
          authStore.clearToken()
          throw new Error('Unauthorized access')
        }
        throw new Error(`Cannot fetch user data`)
      }

      const responseJson = await response.json()
      users.value = responseJson.data
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'An error occurred'
      console.error('Failed to fetch users:', err)
    } finally {
      loading.value = false
    }
  }

  return {
    users,
    loading,
    error,
    fetchUsers,
  }
})
