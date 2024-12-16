<!-- src/pages/contact/Result.vue -->
<template>
  <div class="flex text-2xl flex-col justify-center h-screen w-screen items-center font-serif bg-gray-300">
    <h4>Completed</h4>
    <br />
    <br />
    <h4>Thank you for inquiry</h4>
    <button
      class="text-center border bold m-10"
      @click="router.back()"
    >
      Back
    </button>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'

const router = useRouter()
</script>