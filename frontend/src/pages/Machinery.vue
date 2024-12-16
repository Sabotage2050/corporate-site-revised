<template>
  <div class="bg-white m-8">
    <div class="m-12">
      <h2 class="text-3xl font-serif">Machinery</h2>
      <div class="mt-5 text-base font-serif">
        <p>
          We also deal in various types of used Japanese machinery and export
          them to all kinds of countries.
        </p>
        <p class="mt-5">
          If you are interested in used Japanese machine tools, please contact us.
        </p>
      </div>
    </div>

    <div v-for="category in MACHINERY_CATEGORIES" :key="category.title" class="m-12">
      <h3 class="text-2xl font-serif">{{ category.title }}</h3>
      <div class="mt-5 flex flex-wrap md:flex-auto">
        <div
          v-for="(image, index) in category.images"
          :key="index"
          class="border border-black bg-gray-200 m-2"
        >
          <div>
            <img
              :src="image.src"
              :alt="image.alt"
              class="w-[200px] h-[200px] object-cover"
            />
          </div>
          <div v-if="image.model" class="text-xl ml-2 font-serif">
            {{ image.model }}
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { MACHINERY_CATEGORIES } from '@/features/machinery/constants'
</script>