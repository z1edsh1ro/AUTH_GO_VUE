import { defineStore } from 'pinia'
import { ref } from 'vue'
import type { User } from '../types/user'
import { useAuthStore } from '@/stores/auth'
import { fetchUsers } from '@/services/user'

export const useUserStore = defineStore('user', () => {
  const users = ref<User[]>([])
  const useAuth = useAuthStore()

  const getUsers = async () => {
    try {
      const response = await fetchUsers(useAuth.jwt)
      users.value = response
    } catch (error) {
      console.error('Failed to fetch users:', error)
    }
  }

  return {
    users,
    getUsers,
  }
})
