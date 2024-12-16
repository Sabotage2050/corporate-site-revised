<template>
  <div class="flex flex-col min-h-screen">
    <Header />
    <Navbar />
    <main class="flex-grow">
      <RouterView />
    </main>
    <Footer />
  </div>
</template>

<script setup lang="ts">
import { Header, Navbar, Footer } from '@/components/layout'
import { RouterView } from 'vue-router'
</script>