<template>
  <div>
    <Header />
    <Navbar />
    <main>
      <slot></slot>
    </main>
    <Footer />
  </div>
</template>

<script setup lang="ts">
import Header from './Header.vue'
import Navbar from './Navbar.vue'
import Footer from './Footer.vue'
</script>