<template>
  <div class="navbar">
    <div class="logo">Auth App</div>
    <div class="nav-items">
      <template v-if="authStore.isAuthenticated">
        <a-button type="link" @click="goToUserPage">Users</a-button>
        <a-button type="primary" danger @click="handleLogout">Logout</a-button>
      </template>
      <template v-else>
        <a-button type="link" @click="goToLogin">Login</a-button>
        <a-button type="link" @click="goToRegister">Register</a-button>
      </template>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const authStore = useAuthStore()

const handleLogout = () => {
  authStore.clearToken()
  router.push('/login')
}

const goToLogin = () => {
  router.push('/login')
}

const goToRegister = () => {
  router.push('/register')
}

const goToUserPage = () => {
  router.push('/user')
}
</script>

<style scoped>
.navbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0 24px;
  height: 64px;
  background: #fff;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.15);
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  z-index: 1000;
}

.logo {
  font-size: 20px;
  font-weight: bold;
  color: #1890ff;
}

.nav-items {
  display: flex;
  gap: 16px;
  align-items: center;
}
</style> 