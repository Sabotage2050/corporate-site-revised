// src/pages/stocklist/index.vue
<template>
  <div class="bg-gray-300 min-h-full p-4">
    <ul class="flex flex-wrap justify-center gap-4">
      <li v-for="item in STOCK_ITEMS" :key="item.path">
        <router-link :to="item.path" class="block">
          <img
            :src="item.image"
            :alt="item.alt"
            class="w-[300px] h-[250px] object-cover"
          />
          <h4 class="text-blue-800 font-serif text-l text-center mt-2">
            {{ item.label }}
          </h4>
        </router-link>
      </li>
    </ul>
  </div>
</template>

<script setup lang="ts">
import { STOCK_ITEMS } from '@/features/stocklist/constants'
</script>