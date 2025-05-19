<template>
  <a-layout>
    <Navbar />
    <a-layout-content>
      <div class="flex justify-center mx-auto items-center px-64 h-[calc(92vh)]">
        <a-col :lg="8">
          <a-card title="Login">
            <a-form
              :model="loginForm"
              name="login"
              @submit.prevent="handleSubmit"
              :rules="rules"
              layout="vertical"
            >
              <a-form-item name="email" label="Email">
                <a-input
                  :value="loginForm.email"
                  @update:value="(val: string) => (loginForm.email = val)"
                />
              </a-form-item>

              <a-form-item name="password" label="Password">
                <a-input-password
                  :value="loginForm.password"
                  @update:value="(val: string) => (loginForm.password = val)"
                />
              </a-form-item>

              <a-form-item>
                <a-button type="primary" html-type="submit" block> Login </a-button>
              </a-form-item>
            </a-form>

            <a-alert v-if="error" type="error" :message="error" show-icon banner />
          </a-card>
        </a-col>
      </div>
    </a-layout-content>
  </a-layout>
</template>

<script lang="ts" setup>
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import Navbar from '@/components/Navbar.vue'
import { login } from '@/services/auth.ts'
import type { LoginForm } from '@/types/auth.ts'

const router = useRouter()
const authStore = useAuthStore()
const error = ref<string | null>(null)

const loginForm = reactive<LoginForm>({
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

const handleSubmit = async () => {
  error.value = null

  try {
    const jwt = await login(loginForm)

    authStore.setJwt(jwt)

    router.push('/')
  } catch (err) {
    error.value = err instanceof Error ? err.message : 'Login failed'
  }
}
</script>
