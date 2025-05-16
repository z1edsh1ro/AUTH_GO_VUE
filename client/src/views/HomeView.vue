<template>
  <div class="page-container">
    <Navbar />
    <div class="content">
      <div class="hero-section">
        <h1>Welcome to Auth App</h1>
        <p class="subtitle">A secure and efficient user management system</p>
        <div class="cta-buttons">
          <template v-if="!authStore.isAuthenticated">
            <a-button type="primary" size="large" @click="goToLogin">
              Get Started
            </a-button>
            <a-button size="large" @click="goToRegister">
              Create Account
            </a-button>
          </template>
          <template v-else>
            <a-button type="primary" size="large" @click="goToUserPage">
              Go to Dashboard
            </a-button>
          </template>
        </div>
      </div>

      <div class="features-section">
        <h2>Key Features</h2>
        <div class="features-grid">
          <div class="feature-card">
            <a-icon type="safety" class="feature-icon" />
            <h3>Secure Authentication</h3>
            <p>Advanced security with JWT token-based authentication</p>
          </div>
          <div class="feature-card">
            <a-icon type="team" class="feature-icon" />
            <h3>User Management</h3>
            <p>Efficient user management and monitoring system</p>
          </div>
          <div class="feature-card">
            <a-icon type="dashboard" class="feature-icon" />
            <h3>Easy Dashboard</h3>
            <p>Clean and intuitive user interface</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import Navbar from '@/components/Navbar.vue'

const router = useRouter()
const authStore = useAuthStore()

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
.page-container {
  min-height: 100vh;
  background: #f0f2f5;
}

.content {
  padding: 24px;
  margin-top: 64px;
  min-height: calc(100vh - 64px);
}

.hero-section {
  text-align: center;
  padding: 60px 20px;
  background: #fff;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.15);
  margin-bottom: 40px;
}

.hero-section h1 {
  font-size: 48px;
  color: #1890ff;
  margin-bottom: 16px;
}

.subtitle {
  font-size: 20px;
  color: #666;
  margin-bottom: 32px;
}

.cta-buttons {
  display: flex;
  gap: 16px;
  justify-content: center;
}

.features-section {
  padding: 40px 20px;
  background: #fff;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.15);
}

.features-section h2 {
  text-align: center;
  font-size: 32px;
  color: #1890ff;
  margin-bottom: 40px;
}

.features-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 24px;
  max-width: 1200px;
  margin: 0 auto;
}

.feature-card {
  text-align: center;
  padding: 24px;
  background: #fafafa;
  border-radius: 8px;
  transition: transform 0.3s ease;
}

.feature-card:hover {
  transform: translateY(-5px);
}

.feature-icon {
  font-size: 48px;
  color: #1890ff;
  margin-bottom: 16px;
}

.feature-card h3 {
  font-size: 20px;
  color: #333;
  margin-bottom: 12px;
}

.feature-card p {
  color: #666;
  line-height: 1.5;
}

@media (max-width: 768px) {
  .hero-section h1 {
    font-size: 36px;
  }

  .subtitle {
    font-size: 18px;
  }

  .cta-buttons {
    flex-direction: column;
    align-items: center;
  }

  .features-grid {
    grid-template-columns: 1fr;
  }
}
</style>

