<template>
  <nav class="navbar">
    <div class="container">
      <RouterLink to="/" class="logo">dev.blog</RouterLink>

      <div class="center">
        <RouterLink v-if="auth.isLoggedIn" to="/posts/create">New post</RouterLink>
      </div>

      <div class="right">
        <div v-if="!auth.isLoggedIn">
          <RouterLink to="/auth/login" class="login">Login</RouterLink>
          <RouterLink to="/auth/register" class="register">Register</RouterLink>
        </div>
        <div v-else class="user-section">
          <span class="username">{{ auth.user }}</span>
          <button class="logout" @click="handleLogout">Logout</button>
        </div>
      </div>
    </div>
  </nav>
</template>

<script setup lang="ts">
import { useAuthStore } from '@/stores/auth'
import { useRouter } from 'vue-router'

const auth = useAuthStore()
const router = useRouter()

const logoutPath = 'http://localhost:3000/api/v1/auth/logout'

const handleLogout = async () => {
  try {
    await fetch(logoutPath, {
      method: 'POST',
      headers: {
        Authorization: `Bearer ${auth.authToken}`,
        'Content-Type': 'application/json',
      },
    })
  } catch (error) {
    console.log(error)
  } finally {
    auth.clearUser()
    router.push('/auth/login')
  }
}
</script>

<style scoped>
.navbar {
  padding: 12px;
  background: linear-gradient(to bottom, #0b0f19, #0a0d16);
  border-bottom: 1px solid #1f2937;
}

/* container */
.container {
  max-width: 1100px;
  margin: 0 auto;

  display: flex;
  align-items: center;
  justify-content: space-between;

  background: rgba(17, 24, 39, 0.6);
  border: 1px solid #1f2937;
  border-radius: 12px;
  padding: 10px 14px;
}

.center {
  position: absolute;
  left: 50%;
  transform: translateX(-50%);
}

.user-section {
  display: flex;
  align-items: center;
  gap: 12px;
}

.username {
  font-size: 13px;
  color: #64748b;
}

/* logo */
.logo {
  color: #e5e7eb;
  font-weight: 600;
  font-size: 14px;
  text-decoration: none;
}

/* right side */
.right {
  display: flex;
  align-items: center;
  gap: 10px;
}

/* dropdown */
.dropdown {
  position: relative;
}

.dropdown-btn {
  background: transparent;
  border: 1px solid #1f2937;
  color: #cbd5e1;
  font-size: 13px;
  padding: 6px 10px;
  border-radius: 8px;
  cursor: pointer;

  display: flex;
  align-items: center;
  gap: 6px;
}

.dropdown-btn:hover {
  background: #111827;
}

.arrow {
  font-size: 10px;
  opacity: 0.7;
}

/* dropdown menu */
.dropdown-menu {
  background-color: #1e293b;

  position: absolute;
  top: 120%;
  right: 0;
  z-index: 100;

  width: 180px;
  border: 1px solid #1f2937;
  border-radius: 10px;
  padding: 6px;

  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.4);
}

/* items */
.item {
  display: flex;
  justify-content: space-between;
  align-items: center;

  padding: 8px 10px;
  font-size: 13px;
  color: #cbd5e1;
  border-radius: 6px;
  cursor: pointer;
}

.item:hover {
  background: #111827;
}

.item span {
  opacity: 0.5;
  font-size: 12px;
}

/* active item */
.item.active {
  background: #1e3a8a;
  color: #e0e7ff;
}

/* login button */
.register,
.login {
  background: #2563eb;
  border: none;
  color: white;
  font-size: 13px;
  padding: 6px 12px;
  border-radius: 8px;
  cursor: pointer;
  text-decoration: none;
  margin-right: 5px;
}

.register:hover,
.login:hover {
  background: #1d4ed8;
}

.logout {
  background: #dc2626;
  border: none;
  color: white;
  font-size: 13px;
  padding: 6px 12px;
  border-radius: 8px;
  cursor: pointer;
  text-decoration: none;
}

.logout:hover {
  background: #b91c1c;
}
</style>
