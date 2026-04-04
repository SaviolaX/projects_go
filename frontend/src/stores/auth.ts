import { defineStore } from 'pinia'

interface AuthState {
  user: string | null
  isLoggedIn: boolean
  authToken: string | null
}

interface SetUserPayload {
  user: string
  isLoggedIn: boolean
  authToken: string
}

export const useAuthStore = defineStore('auth', {
  state: (): AuthState => ({
    user: null,
    isLoggedIn: false,
    authToken: null,
  }),
  persist: true,
  actions: {
    setUser({ user, isLoggedIn, authToken }: SetUserPayload) {
      this.user = user
      this.isLoggedIn = isLoggedIn
      this.authToken = authToken
    },
    clearUser() {
      this.user = null
      this.isLoggedIn = false
      this.authToken = null
    },
  },
})
