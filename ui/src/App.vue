<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
  <header>
    <div>
      <nav>
        <RouterLink to="/">Home</RouterLink>
        <RouterLink to="/about">About</RouterLink>
      </nav>
    </div>
  </header>

  <main>
  <RouterView />
  </main>
</template>

<style scoped>
nav {
  width: 100%;
  background-color: lightblue;
  margin-top: 5px;
}

nav a.router-link-exact-active {
  color: blue; 
  border-style: solid;
  margin: 10px;
}

nav a {
  color: blue; 
  border-color: blue;
  border-style: dotted;
  margin: 10px;
}
</style>
