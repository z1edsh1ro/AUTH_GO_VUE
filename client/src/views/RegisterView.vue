<template>
  <div>
    <Navbar />
    <div class="register-container">
      <div class="register-box">
        <h2>Register</h2>
        <form @submit.prevent="handleRegister" class="register-form">
          <div class="form-group">
            <label for="name">Full Name</label>
            <input
              type="text"
              id="name"
              v-model="name"
              required
              placeholder="Enter your full name"
            />
          </div>
          <div class="form-group">
            <label for="email">Email</label>
            <input
              type="email"
              id="email"
              v-model="email"
              required
              placeholder="Enter your email"
            />
          </div>
          <div class="form-group">
            <label for="password">Password</label>
            <input
              type="password"
              id="password"
              v-model="password"
              required
              placeholder="Enter your password"
            />
          </div>
          <div class="form-group">
            <label for="confirmPassword">Confirm Password</label>
            <input
              type="password"
              id="confirmPassword"
              v-model="confirmPassword"
              required
              placeholder="Confirm your password"
            />
          </div>
          <button type="submit" class="register-button">Register</button>
        </form>
        <p class="login-link">
          Already have an account? <router-link to="/login">Login here</router-link>
        </p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import Navbar from '@/components/Navbar.vue'

const router = useRouter()
const name = ref('')
const email = ref('')
const password = ref('')
const confirmPassword = ref('')

const handleRegister = async () => {
  try {
    if (password.value !== confirmPassword.value) {
      alert('Passwords do not match!')
      return
    }
    // TODO: Implement registration logic here
    console.log('Registration attempt with:', {
      name: name.value,
      email: email.value,
      password: password.value
    })
    // After successful registration, redirect to login page
    // router.push('/login')
  } catch (error) {
    console.error('Registration failed:', error)
  }
}
</script>

<style scoped>
.register-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background-color: #f5f5f5;
  padding-top: 4rem; /* Add padding to account for fixed navbar */
}

.register-box {
  background: white;
  padding: 2rem;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  width: 100%;
  max-width: 400px;
}

h2 {
  text-align: center;
  color: #333;
  margin-bottom: 1.5rem;
}

.register-form {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

label {
  font-weight: 500;
  color: #555;
}

input {
  padding: 0.75rem;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 1rem;
}

input:focus {
  outline: none;
  border-color: #4CAF50;
}

.register-button {
  background-color: #4CAF50;
  color: white;
  padding: 0.75rem;
  border: none;
  border-radius: 4px;
  font-size: 1rem;
  cursor: pointer;
  transition: background-color 0.2s;
}

.register-button:hover {
  background-color: #45a049;
}

.login-link {
  text-align: center;
  margin-top: 1rem;
  color: #666;
}

.login-link a {
  color: #4CAF50;
  text-decoration: none;
}

.login-link a:hover {
  text-decoration: underline;
}
</style> 