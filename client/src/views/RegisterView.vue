<template>
  <a-layout>
    <Navbar />
    <a-layout-content>
      <div class="flex justify-center mx-auto items-center px-64 h-[calc(92vh)]">
        <a-col :lg="8">
          <a-card title="Register" :bordered="false">
            <a-form
              :model="registerForm"
              name="register"
              @submit.prevent="handleSubmit"
              :rules="rules"
              layout="vertical"
            >
              <a-form-item name="name" label="Name">
                <a-input v-model:value="registerForm.name" />
              </a-form-item>

              <a-form-item name="email" label="Email">
                <a-input v-model:value="registerForm.email" />
              </a-form-item>

              <a-form-item name="password" label="Password">
                <a-input-password v-model:value="registerForm.password" />
              </a-form-item>

              <a-form-item name="confirmPassword" label="Confirm Password">
                <a-input-password v-model:value="registerForm.confirmPassword" />
              </a-form-item>

              <a-form-item>
                <a-button type="primary" html-type="submit" block> Register </a-button>
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
import Navbar from '@/components/Navbar.vue'
import type { Rule } from 'ant-design-vue/es/form'
import type { RegisterForm, RegisterRequest } from '@/types/auth.ts'
import { register } from '@/services/auth.ts'

const router = useRouter()
const error = ref<string | null>(null)

const registerForm = reactive<RegisterForm>({
  name: '',
  email: '',
  password: '',
  confirmPassword: '',
})

const validateConfirmPassword = async (_rule: Rule, value: string) => {
  if (value !== registerForm.password) {
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
    { min: 6, message: 'Email must be at least 6 characters!' },
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

const handleSubmit = async () => {
  error.value = null

  const payload: RegisterRequest = {
    name: registerForm.name,
    email: registerForm.email,
    password: registerForm.password,
  }

  console.log(payload)
  try {
    register(payload)

    alert('Register success')
    router.push('/login')
  } catch (err) {
    error.value = err instanceof Error ? err.message : 'Registration failed'
  }
}
</script>
