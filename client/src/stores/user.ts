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

  const updateUser = (updateUser: User | null) => {
    if (!updateUser) {
      return
    }

    const index = users.value.findIndex(user => user.id === updateUser.id)

    if(index !== -1) {
      users.value[index] = { ...updateUser }
    }
  }

  return {
    users,
    getUsers,
    updateUser,
  }
})
