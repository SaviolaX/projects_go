<script setup lang="ts">
import NavBarCmp from './components/NavBarCmp.vue'
</script>

<template>
  <div class="app">
    <!-- background glow -->
    <div class="bg-glow"></div>

    <!-- navbar -->
    <NavBarCmp />

    <!-- page content -->
    <main class="content">
      <router-view />
    </main>
  </div>
</template>

<style>
html,
body {
  margin: 0;
  padding: 0;
  background-color: #020617;
  width: 100%;
  height: 100%;
  overflow-x: hidden;
}

#app {
  margin: 0;
  padding: 0;
}
</style>

<style scoped>
/* ROOT */
* {
  box-sizing: border-box;
  margin: 0;
  padding: 0;
}
.app {
  min-height: 100vh;
  background: radial-gradient(circle at 20% 20%, #1e293b, #020617 60%);
  color: #e5e7eb;
  position: relative;
  overflow: hidden;
}

/* SOFT GLOW (matches your screenshot vibe) */
.bg-glow {
  position: absolute;
  pointer-events: none;
  top: -100px;
  left: -100px;
  width: 500px;
  height: 500px;
  background: radial-gradient(circle, rgba(59, 130, 246, 0.25), transparent 70%);
  filter: blur(80px);
  z-index: 0;
}

/* CONTENT AREA */
.content {
  position: relative;
  z-index: 1;

  max-width: 1100px;
  margin: 40px auto;
  padding: 0 16px;
}

/* TYPOGRAPHY */
h1 {
  font-size: 28px;
  font-weight: 600;
  margin-top: 40px;
}
</style>
