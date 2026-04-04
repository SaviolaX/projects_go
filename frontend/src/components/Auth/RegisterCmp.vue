<template>
  <div class="auth-container">
    <div class="auth-card">
      <div class="auth-header">
        <h2 class="title">Create an Account</h2>
        <p class="subtitle">Join the dev.blog community today</p>
      </div>

      <form @submit.prevent="handleRegister" class="auth-form">
        <div class="input-group">
          <label for="username">Username</label>
          <input
            type="text"
            id="username"
            v-model="username"
            placeholder="Your username"
            required
          />
        </div>

        <div class="input-group">
          <label for="email">Email</label>
          <input type="email" id="email" v-model="email" placeholder="name@example.com" required />
        </div>

        <div class="input-group">
          <label for="password">Password</label>
          <input
            type="password"
            id="password"
            v-model="password"
            placeholder="Minimum 8 characters"
            required
          />
          <p class="input-hint">Use a mix of letters, numbers, and symbols</p>
        </div>
        <div>
          <p v-if="error">{{ error }}</p>
        </div>

        <button type="submit" class="submit-btn">Sign Up</button>
      </form>

      <div class="auth-footer">
        <span>Already have an account?</span>
        <router-link to="/auth/login" class="login-link">Sign In</router-link>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()

const username = ref('')
const email = ref('')
const password = ref('')
const isLoading = ref(false)
const error = ref<string | null>(null)

const registerPath = 'http://localhost:3000/api/v1/auth/register'

const handleRegister = async () => {
  isLoading.value = true
  try {
    const response = await fetch(registerPath, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        username: username.value,
        email: email.value,
        password: password.value,
      }),
    })

    if (!response.ok) {
      const errorData = await response.json()
      throw new Error(errorData.error ?? `Error: ${response.status}`)
    }

    const data = await response.json()
    if (data.status == 'created') {
      router.push('/auth/login')
    }
  } catch (err) {
    error.value = err instanceof Error ? err.message : 'Something went wrong'
  } finally {
    isLoading.value = false
  }
}
</script>

<style scoped>
.auth-container {
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 40px 20px;
  min-height: 80vh;
}

.auth-card {
  width: 100%;
  max-width: 440px;
  background-color: rgba(30, 41, 59, 0.3);
  backdrop-filter: blur(12px);
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: 16px;
  padding: 40px;
  box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.5);
}

.auth-header {
  text-align: center;
  margin-bottom: 32px;
}

.title {
  font-size: 26px;
  font-weight: 700;
  color: #ffffff;
  margin-bottom: 10px;
}

.subtitle {
  color: #94a3b8;
  font-size: 14px;
}

.auth-form {
  display: flex;
  flex-direction: column;
  gap: 24px;
}

.input-group {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

label {
  font-size: 14px;
  font-weight: 500;
  color: #e2e8f0;
}

input {
  background-color: #0f172a;
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: 10px;
  padding: 14px 16px;
  color: #ffffff;
  font-size: 15px;
  transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);
}

input:focus {
  outline: none;
  border-color: #3b82f6;
  background-color: #111b2e;
  box-shadow: 0 0 0 4px rgba(59, 130, 246, 0.15);
}

.input-hint {
  font-size: 12px;
  color: #475569;
  margin-top: 4px;
}

.submit-btn {
  background-color: #3b82f6;
  color: #ffffff;
  border: none;
  border-radius: 10px;
  padding: 14px;
  font-size: 16px;
  font-weight: 600;
  cursor: pointer;
  margin-top: 8px;
  transition:
    transform 0.1s,
    background-color 0.2s;
}

.submit-btn:hover {
  background-color: #2563eb;
}

.submit-btn:active {
  transform: scale(0.98);
}

.auth-footer {
  margin-top: 32px;
  text-align: center;
  font-size: 14px;
  color: #94a3b8;
  display: flex;
  justify-content: center;
  gap: 8px;
}

.login-link {
  color: #3b82f6;
  text-decoration: none;
  font-weight: 600;
}

.login-link:hover {
  text-decoration: underline;
}

/* Ensure smooth transition for mobile */
@media (max-width: 480px) {
  .auth-card {
    padding: 30px 20px;
    border: none;
    background-color: transparent;
    backdrop-filter: none;
  }
}
</style>
