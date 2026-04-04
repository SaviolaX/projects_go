<template>
  <div class="posts-grid">
    <div v-for="post in posts" :key="post.ID" class="post-card">
      <div class="post-content">
        <span class="category-tag">{{ post.Category.Name }}</span>
        <h3 class="post-title">{{ post.Title }}</h3>
        <p class="post-description">{{ post.Entry }}</p>
      </div>

      <div class="post-footer">
        <div class="author-info">
          <div class="avatar">None</div>
          <span class="author-name">{{ post.Author.Username }}</span>
        </div>
        <span class="read-time">??</span>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'

const getAllPostsPath = 'http://localhost:3000/api/v1/posts'

interface Author {
  ID: number
  Username: string
}
interface Category {
  ID: number
  Name: string
}
interface Post {
  ID: number
  Title: string
  Entry: string
  Author: Author
  Category: Category
}

const posts = ref<Post[]>([])

onMounted(async () => {
  const response = await fetch(getAllPostsPath)
  const data = await response.json()
  posts.value = data.posts
})
</script>

<style scoped>
.posts-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  /* The border-collapse effect is achieved by borders on items */
  border-top: 1px solid rgba(255, 255, 255, 0.1);
  background-color: transparent;
}

.post-card {
  padding: 32px;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  min-height: 220px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
  transition: background-color 0.2s ease;
  cursor: pointer;
}

/* Vertical line in the middle */
.post-card:nth-child(odd) {
  border-right: 1px solid rgba(255, 255, 255, 0.1);
}

.post-card:hover {
  background-color: rgba(255, 255, 255, 0.02);
}

.category-tag {
  color: #3b82f6; /* Soft blue */
  font-size: 14px;
  font-weight: 500;
  display: block;
  margin-bottom: 12px;
}

.post-title {
  font-size: 20px;
  font-weight: 600;
  color: #ffffff;
  margin-bottom: 12px;
}

.post-description {
  color: #94a3b8; /* Slate-400 */
  font-size: 15px;
  line-height: 1.5;
}

.post-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-top: 24px;
}

.author-info {
  display: flex;
  align-items: center;
  gap: 10px;
}

.avatar {
  width: 28px;
  height: 28px;
  background-color: #1e293b;
  border: 1px solid rgba(59, 130, 246, 0.5);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 11px;
  font-weight: bold;
  color: #60a5fa;
}

.author-name {
  color: #64748b;
  font-size: 14px;
}

.read-time {
  color: #475569;
  font-size: 13px;
}

/* Mobile Responsiveness */
@media (max-width: 768px) {
  .posts-grid {
    grid-template-columns: 1fr;
  }
  .post-card:nth-child(odd) {
    border-right: none;
  }
}
</style>
