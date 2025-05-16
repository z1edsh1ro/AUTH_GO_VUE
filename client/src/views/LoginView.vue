<template>
  <a-layout>
    <Navbar />
    <a-layout-content>
      <a-row justify="center" class="container">
        <a-col :lg="8">
          <a-card title="Login">
            <a-form
              :model="formState"
              name="login"
              @finish="handleSubmit"
              :rules="rules"
              layout="vertical"
            >
              <a-form-item name="email" label="Email">
                <a-input
                  :value="formState.email"
                  @update:value="(val: string) => (formState.email = val)"
                />
              </a-form-item>

              <a-form-item name="password" label="Password">
                <a-input-password
                  :value="formState.password"
                  @update:value="(val: string) => (formState.password = val)"
                />
              </a-form-item>

              <a-form-item>
                <a-button type="primary" html-type="submit" :loading="loading" block>
                  Login
                </a-button>
              </a-form-item>
            </a-form>

            <a-alert v-if="error" type="error" :message="error" show-icon banner />
          </a-card>
        </a-col>
      </a-row>
    </a-layout-content>
  </a-layout>
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
  password: '',
})

const rules = {
  email: [
    { required: true, message: 'Please input your email!' },
    { type: 'email', message: 'Please enter a valid email!' },
    { min: 6, message: 'Email must be at least 6 characters!' },
  ],
  password: [
    { required: true, message: 'Please input your password!' },
    { min: 6, message: 'Password must be at least 6 characters!' },
  ],
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

    authStore.setToken(data.jwt)
    if (data.data) {
      authStore.setUser(data.data)
    }

    router.push('/')
  } catch (err) {
    error.value = err instanceof Error ? err.message : 'Login failed'
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.container {
  min-height: calc(100vh - 64px);
  padding: 24px;
}
</style>
