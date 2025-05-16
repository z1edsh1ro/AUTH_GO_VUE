<template>
  <a-layout-header class="header">
    <div class="logo" @click="goToHome">App</div>
    <a-space>
      <template v-if="authStore.isAuthenticated">
        <a-button type="link" @click="goToUserPage">User Management</a-button>
        <a-dropdown>
          <a class="ant-dropdown-link" @click.prevent>
            <a-avatar>{{ authStore.user?.name?.[0]?.toUpperCase() }}</a-avatar>
            <span class="username">{{ authStore.user?.name }}</span>
          </a>
          <template #overlay>
            <a-menu>
              <a-menu-item key="profile">
                <div class="profile-info">
                  <div class="name">{{ authStore.user?.name }}</div>
                  <div class="email">{{ authStore.user?.email }}</div>
                </div>
              </a-menu-item>
              <a-menu-divider />
              <a-menu-item key="logout" @click="handleLogout">
                <LogoutOutlined />
                Logout
              </a-menu-item>
            </a-menu>
          </template>
        </a-dropdown>
      </template>
      <template v-else>
        <a-button type="link" @click="goToLogin">Login</a-button>
        <a-button type="link" @click="goToRegister">Register</a-button>
      </template>
    </a-space>
  </a-layout-header>
</template>

<script lang="ts" setup>
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { LogoutOutlined } from '@ant-design/icons-vue'

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

const goToHome = () => {
  router.push('/')
}
</script>

<style scoped>
.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0 24px;
  background: #fff;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.15);
}

.logo {
  font-size: 20px;
  font-weight: bold;
  color: #1890ff;
  cursor: pointer;
  transition: color 0.3s;
}

.logo:hover {
  color: #40a9ff;
}

.ant-dropdown-link {
  display: flex;
  align-items: center;
  gap: 8px;
  cursor: pointer;
}

.username {
  color: rgba(0, 0, 0, 0.85);
}

.profile-info {
  padding: 4px 0;
}

.profile-info .name {
  font-weight: 500;
  color: rgba(0, 0, 0, 0.85);
}

.profile-info .email {
  font-size: 12px;
  color: rgba(0, 0, 0, 0.45);
}
</style>
