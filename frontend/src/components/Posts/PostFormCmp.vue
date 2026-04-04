<template>
  <div class="form-container">
    <div class="form-card">
      <div class="form-header">
        <h1 class="title">{{ isEditing ? 'Редагувати пост' : 'Новий пост' }}</h1>
        <p class="subtitle">
          {{ isEditing ? 'Внесіть зміни до посту' : 'Поділіться своїми думками' }}
        </p>
      </div>

      <form class="post-form" @submit.prevent="handleSubmit">
        <div class="input-group">
          <label for="title">Заголовок</label>
          <input id="title" v-model="title" type="text" placeholder="Введіть заголовок..." />
        </div>

        <div class="input-group">
          <label for="entry">Текст</label>
          <textarea id="entry" v-model="entry" placeholder="Введіть текст посту..." rows="8" />
        </div>

        <div class="input-group">
          <label>Category</label>
          <input v-model="categoryName" type="text" placeholder="Or enter a new category..." />
          <select v-if="categories != null" v-model="categoryID" class="category-select">
            <option value="" disabled>Select a category</option>
            <option v-for="cat in categories" :key="cat.ID" :value="cat.ID">
              {{ cat.Name }}
            </option>
          </select>
          <p v-else class="hint">No categories yet</p>
        </div>
        <p v-if="error" class="error">{{ error }}</p>

        <button class="submit-btn" type="submit" :disabled="isLoading">
          {{ isLoading ? 'Збереження...' : isEditing ? 'Зберегти' : 'Опублікувати' }}
        </button>
      </form>
    </div>
  </div>
</template>
<script setup lang="ts">
import { useAuthStore } from '@/stores/auth'
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'

interface Category {
  ID: number
  Name: string
}

const categories = ref<Category[]>([])

onMounted(async () => {
  const response = await fetch(getAllCategoriesPath)
  const data = await response.json()
  categories.value = data.categories
})

const getAllCategoriesPath = 'http://localhost:3000/api/v1/categories'
const createPostPath = 'http://localhost:3000/api/v1/posts/create'

const auth = useAuthStore()
const router = useRouter()

const error = ref('')
const isLoading = ref(false)
const isEditing = ref(false)
const entry = ref('')
const title = ref('')
const categoryName = ref('')
const categoryID = ref(0)

const handleSubmit = async () => {
  isLoading.value = true
  try {
    const response = await fetch(createPostPath, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json', Authorization: `Bearer ${auth.authToken}` },
      body: JSON.stringify({
        title: title.value,
        entry: entry.value,
        categoryName: categoryName.value,
        categoryID: categoryID.value,
      }),
    })

    if (!response.ok) {
      const errorData = await response.json()
      throw new Error(errorData.error ?? `Error: ${response.status}`)
    }

    const data = await response.json()
    if (data.status == 'created') {
      router.push('/')
    }
  } catch (err) {
    error.value = err instanceof Error ? err.message : 'Something went wrong'
  } finally {
    isLoading.value = false
  }
}
</script>
<style scoped>
.form-container {
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 60px 20px;
  min-height: 70vh;
}

.form-card {
  width: 100%;
  max-width: 620px;
  background-color: rgba(30, 41, 59, 0.3);
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: 16px;
  padding: 40px;
}

.form-header {
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

.post-form {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.input-group {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

label {
  font-size: 14px;
  font-weight: 500;
  color: #e5e7eb;
}

input,
textarea {
  background-color: #0f172a;
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: 8px;
  padding: 12px 16px;
  color: #ffffff;
  font-size: 15px;
  transition: all 0.2s ease;
  resize: vertical;
  font-family: inherit;
}

input:focus,
textarea:focus {
  outline: none;
  border-color: #3b82f6;
  box-shadow: 0 0 0 4px rgba(59, 130, 246, 0.1);
}

.error {
  font-size: 14px;
  color: #f87171;
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

.submit-btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}
.category-select {
  background-color: #0f172a;
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: 8px;
  padding: 12px 16px;
  color: #ffffff;
  font-size: 15px;
  cursor: pointer;
  transition: all 0.2s ease;
}

.category-select:focus {
  outline: none;
  border-color: #3b82f6;
  box-shadow: 0 0 0 4px rgba(59, 130, 246, 0.1);
}

.category-select option {
  background-color: #0f172a;
  color: #ffffff;
}
</style>
