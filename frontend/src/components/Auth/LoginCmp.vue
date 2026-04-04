<template>
  <div class="login-container">
    <div class="login-card">
      <div class="login-header">
        <h2 class="title">Welcome back!</h2>
        <p class="subtitle">Log into your account dev.blog</p>
      </div>

      <form @submit.prevent="handleLogin" class="login-form">
        <div class="input-group">
          <label for="username">Username</label>
          <input
            type="text"
            id="username"
            v-model="username"
            placeholder="your username"
            required
          />
        </div>

        <div class="input-group">
          <div class="label-row">
            <label for="password">Password</label>
          </div>
          <input type="password" id="password" v-model="password" placeholder="••••••••" required />
        </div>
        <a href="#" class="forgot-link">Forgot ?</a>

        <button type="submit" class="submit-btn">Login</button>
      </form>

      <div>
        <p v-if="error">{{ error }}</p>
      </div>

      <div class="login-footer">
        <span>Don't have an account ?</span>
        <router-link to="/auth/register" class="register-link">Create now</router-link>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useAuthStore } from '@/stores/auth'
import { useRouter } from 'vue-router'

const router = useRouter()
const auth = useAuthStore()

const username = ref('')
const password = ref('')
const isLoading = ref(false)
const error = ref<string | null>(null)

const loginPath = 'http://localhost:3000/api/v1/auth/login'

const handleLogin = async () => {
  isLoading.value = true
  try {
    const response = await fetch(loginPath, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        username: username.value,
        password: password.value,
      }),
    })

    if (!response.ok) {
      const errorData = await response.json()
      throw new Error(errorData.error ?? `Error: ${response.status}`)
    }

    const data = await response.json()
    auth.setUser({
      user: data.username,
      authToken: data.token,
      isLoggedIn: true,
    })

    router.push('/')
  } catch (err) {
    error.value = err instanceof Error ? err.message : 'Something went wrong'
  } finally {
    isLoading.value = false
  }
}
</script>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 60px 20px;
  min-height: 70vh;
}

.login-card {
  width: 100%;
  max-width: 420px;
  background-color: rgba(30, 41, 59, 0.3); /* Transparent slate */
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: 16px;
  padding: 40px;
}

.login-header {
  text-align: center;
  margin-bottom: 32px;
}

.title {
  font-size: 24px;
  font-weight: 600;
  color: #ffffff;
  margin-bottom: 8px;
}

.subtitle {
  color: #94a3b8;
  font-size: 14px;
}

.login-form {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.input-group {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.label-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

label {
  font-size: 14px;
  font-weight: 500;
  color: #e5e7eb;
}

input {
  background-color: #0f172a; /* Darker input background */
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: 8px;
  padding: 12px 16px;
  color: #ffffff;
  font-size: 15px;
  transition: all 0.2s ease;
}

input:focus {
  outline: none;
  border-color: #3b82f6;
  box-shadow: 0 0 0 4px rgba(59, 130, 246, 0.1);
}

.forgot-link {
  font-size: 12px;
  color: #3b82f6;
  text-decoration: none;
}

.submit-btn {
  background-color: #3b82f6;
  color: white;
  border: none;
  border-radius: 8px;
  padding: 12px;
  font-size: 16px;
  font-weight: 600;
  cursor: pointer;
  margin-top: 10px;
  transition: background-color 0.2s;
}

.submit-btn:hover {
  background-color: #2563eb;
}

.login-footer {
  margin-top: 32px;
  text-align: center;
  font-size: 14px;
  color: #64748b;
  display: flex;
  justify-content: center;
  gap: 6px;
}

.register-link {
  color: #ffffff;
  text-decoration: none;
  font-weight: 500;
}

.register-link:hover {
  text-decoration: underline;
}
</style>
