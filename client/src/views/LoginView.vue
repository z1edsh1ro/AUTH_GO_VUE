<template>
  <div class="page-container">
    <Navbar />
    <div class="content">
      <div class="login-container">
        <h1>Login</h1>
        <a-form
          :model="formState"
          name="login"
          @finish="handleSubmit"
          :rules="rules"
        >
          <a-form-item name="email" label="Email">
            <a-input v-model:value="formState.email" />
          </a-form-item>

          <a-form-item name="password" label="Password">
            <a-input-password v-model:value="formState.password" />
          </a-form-item>

          <a-form-item>
            <a-button type="primary" html-type="submit" :loading="loading" block>
              Login
            </a-button>
          </a-form-item>

          <div class="form-footer">
            <span>Don't have an account?</span>
            <a-button type="link" @click="goToRegister">Register</a-button>
          </div>
        </a-form>

        <a-alert
          v-if="error"
          type="error"
          :message="error"
          show-icon
          class="error-alert"
        />
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import Navbar from '@/components/Navbar.vue'

interface FormState {
  email: string
  password: string
}

const router = useRouter()
const authStore = useAuthStore()
const loading = ref(false)
const error = ref<string | null>(null)

const formState = reactive<FormState>({
  email: '',
  password: ''
})

const rules = {
  email: [
    { required: true, message: 'Please input your email!' },
    { type: 'email', message: 'Please enter a valid email!' }
  ],
  password: [
    { required: true, message: 'Please input your password!' },
    { min: 6, message: 'Password must be at least 6 characters!' }
  ]
}

const handleSubmit = async (values: FormState) => {
  loading.value = true
  error.value = null

  try {
    const response = await fetch('http://localhost:8080/api/auth/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(values),
    })

    const data = await response.json()

    if (!response.ok) {
      throw new Error(data.message || 'Login failed')
    }

    // Set the token in the auth store
    authStore.setToken(data.token)
    
    // Redirect to user page
    router.push('/user')
  } catch (err) {
    error.value = err instanceof Error ? err.message : 'Login failed'
  } finally {
    loading.value = false
  }
}

const goToRegister = () => {
  router.push('/register')
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
  display: flex;
  justify-content: center;
  align-items: flex-start;
}

.login-container {
  width: 100%;
  max-width: 400px;
  padding: 40px;
  background: #fff;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.15);
}

.login-container h1 {
  text-align: center;
  color: #1890ff;
  margin-bottom: 32px;
}

.form-footer {
  text-align: center;
  margin-top: 16px;
}

.error-alert {
  margin-top: 16px;
}

@media (max-width: 768px) {
  .login-container {
    padding: 24px;
  }
}
</style>
