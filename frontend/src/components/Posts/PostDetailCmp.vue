<template>
  <div class="post-detail">
    <header class="post-header">
      <div class="header-content">
        <div class="meta-top">
          <span class="category-tag">{{ post.category }}</span>
          <span class="dot">•</span>
          <span class="publish-date">{{ post.date }}</span>
        </div>

        <h1 class="post-title">{{ post.title }}</h1>
        <p class="post-subtitle">{{ post.description }}</p>

        <div class="author-bar">
          <div class="author-info">
            <div class="avatar">{{ post.authorInitials }}</div>
            <div class="author-details">
              <span class="author-name">{{ post.author }}</span>
              <span class="author-role">Автор контенту</span>
            </div>
          </div>
          <div class="read-info">
            <span>{{ post.readTime }} хв читання</span>
          </div>
        </div>
      </div>
    </header>

    <article class="post-body">
      <div class="content-wrapper">
        <div class="prose">
          <p>
            Composition API — це набір API, який дозволяє нам створювати компоненти Vue,
            використовуючи імпортовані функції замість оголошення опцій. Це основна фіча Vue 3, яка
            вирішує проблеми масштабування коду у великих проєктах.
          </p>

          <h2>Чому це важливо?</h2>
          <p>
            На відміну від Options API, де логіка розпорошена між data, computed та methods,
            Composition API дозволяє групувати код за логічними задачами.
          </p>

          <div class="code-block">
            <code>const count = ref(0);</code>
          </div>
        </div>
      </div>
    </article>
  </div>
</template>

<script setup lang="ts">
// In a real app, you'd fetch this via props or a store using a route ID
const post = {
  category: 'Vue.js',
  title: 'Composition API vs Options API',
  description:
    'Глибоке занурення у зміну парадигми розробки на Vue 3 та аналіз того, як це впливає на архітектуру ваших додатків.',
  date: '24 Березня, 2024',
  author: 'Михайло',
  authorInitials: 'МК',
  readTime: 5,
}
</script>

<style scoped>
.post-detail {
  width: 100%;
  animation: fadeIn 0.5s ease-out;
}

/* Header Section styling */
.post-header {
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
  padding: 60px 0;
}

.meta-top {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 24px;
}

.dot {
  color: #475569;
}

.publish-date {
  color: #64748b;
  font-size: 14px;
}

.category-tag {
  color: #3b82f6;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.05em;
  font-size: 13px;
}

.post-title {
  font-size: clamp(32px, 5vw, 48px);
  font-weight: 700;
  color: #ffffff;
  line-height: 1.2;
  margin-bottom: 20px;
}

.post-subtitle {
  font-size: 20px;
  color: #94a3b8;
  line-height: 1.6;
  max-width: 800px;
  margin-bottom: 40px;
}

/* Author Bar */
.author-bar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-top: 20px;
}

.author-info {
  display: flex;
  align-items: center;
  gap: 16px;
}

.avatar {
  width: 48px;
  height: 48px;
  background: #1e293b;
  border: 1px solid rgba(59, 130, 246, 0.4);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: bold;
  color: #60a5fa;
}

.author-details {
  display: flex;
  flex-direction: column;
}

.author-name {
  color: #f1f5f9;
  font-weight: 600;
}

.author-role {
  color: #64748b;
  font-size: 13px;
}

.read-info {
  color: #475569;
  font-size: 14px;
}

/* Body Content */
.post-body {
  padding: 60px 0;
}

.content-wrapper {
  max-width: 800px;
}

.prose {
  color: #cbd5e1;
  font-size: 18px;
  line-height: 1.8;
}

.prose h2 {
  color: #ffffff;
  margin-top: 40px;
  margin-bottom: 20px;
}

.prose p {
  margin-bottom: 24px;
}

.code-block {
  background: #0f172a;
  padding: 20px;
  border-radius: 12px;
  border: 1px solid rgba(255, 255, 255, 0.05);
  margin: 30px 0;
}

code {
  color: #3b82f6;
  font-family: 'Fira Code', monospace;
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
</style>
