<template>
  <a-layout-header class="header">
    <router-link to="/">
      <div class="logo">App</div>
    </router-link>
    <a-space>
      <template v-if="authStore.loadJwt()">
        <router-link to="/user">
          <a-button>User Management</a-button>
        </router-link>
        <a-dropdown>
          <a class="ant-dropdown-link" @click.prevent>
            <a-avatar>{{ authStore.jwtPayload?.name?.[0]?.toUpperCase() }}</a-avatar>
            <span class="username">{{ authStore.jwtPayload?.name }}</span>
          </a>
          <template #overlay>
            <a-menu>
              <a-menu-item key="profile">
                <div class="profile-info">
                  <div class="name">{{ authStore.jwtPayload?.name }}</div>
                  <div class="email">{{ authStore.jwtPayload?.email }}</div>
                  <div class="time">{{ authStore.jwtPayload?.loginTime }}</div>
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
        <router-link to="/login">
          <a-button type="link">Login</a-button>
        </router-link>
        <router-link to="/register">
          <a-button type="link">Register</a-button>
        </router-link>
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
  authStore.clearJwt()
  router.push('/login')
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

.profile-info .time {
  font-size: 10px;
  color: rgba(0, 0, 0, 0.45);
}
</style>
