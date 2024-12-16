<template>
  <nav class="bg-sky-700">
    <div class="flex flex-wrap sm:flex-nowrap">
      <RouterLink 
        v-for="link in navigationLinks" 
        :key="link.path"
        :to="link.path" 
        class="mx-5 first:mx-10"
      >
        <span class="text-white text-3xl font-serif">{{ link.label }}</span>
      </RouterLink>
    </div>
  </nav>
</template>

<script setup lang="ts">
import { RouterLink } from 'vue-router'

const navigationLinks = [
  { path: '/', label: 'Home' },
  { path: '/about', label: 'About' },
  { path: '/stocklist', label: 'StockList' },
  { path: '/contact', label: 'Contact' },
  { path: '/machinery', label: 'Machinery' }
]
</script>