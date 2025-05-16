<template>
  <a-layout>
    <Navbar />
    <a-layout-content>
      <a-row justify="center" class="container">
        <a-col :lg="8">
          <a-card title="Register" :bordered="false">
            <a-form
              :model="formState"
              name="register"
              @finish="handleSubmit"
              :rules="rules"
              layout="vertical"
            >
              <a-form-item name="name" label="Name">
                <a-input v-model:value="formState.name" />
              </a-form-item>

              <a-form-item name="email" label="Email">
                <a-input v-model:value="formState.email" />
              </a-form-item>

              <a-form-item name="password" label="Password">
                <a-input-password v-model:value="formState.password" />
              </a-form-item>

              <a-form-item name="confirmPassword" label="Confirm Password">
                <a-input-password v-model:value="formState.confirmPassword" />
              </a-form-item>

              <a-form-item>
                <a-button type="primary" html-type="submit" :loading="loading" block>
                  Register
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
import Navbar from '@/components/Navbar.vue'

interface FormState {
  name: string
  email: string
  password: string
  confirmPassword: string
}

const router = useRouter()
const loading = ref(false)
const error = ref<string | null>(null)

const formState = reactive<FormState>({
  name: '',
  email: '',
  password: '',
  confirmPassword: '',
})

const validateConfirmPassword = async (_rule: any, value: string) => {
  if (value !== formState.password) {
    throw new Error('The two passwords do not match!')
  }
}

const rules = {
  name: [
    { required: true, message: 'Please input your name!' },
    { min: 2, message: 'Name must be at least 2 characters!' },
  ],
  email: [
    { required: true, message: 'Please input your email!' },
    { type: 'email', message: 'Please enter a valid email!' },
  ],
  password: [
    { required: true, message: 'Please input your password!' },
    { min: 6, message: 'Password must be at least 6 characters!' },
  ],
  confirmPassword: [
    { required: true, message: 'Please confirm your password!' },
    { validator: validateConfirmPassword },
  ],
}

const handleSubmit = async (values: FormState) => {
  loading.value = true
  error.value = null

  try {
    const response = await fetch('http://localhost:8080/api/auth/register', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        name: values.name,
        email: values.email,
        password: values.password,
      }),
    })

    const data = await response.json()

    if (!response.ok) {
      throw new Error(data.message || 'Registration failed')
    }

    router.push('/login')
  } catch (err) {
    error.value = err instanceof Error ? err.message : 'Registration failed'
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
