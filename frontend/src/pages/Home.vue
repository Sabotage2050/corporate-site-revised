<template>
  <div class="bg-gray-300 overflow-hidden">
    <HomeCarousel />
    
    <div class="m-12">
      <h2 class="text-3xl font-serif">Corporate Info.</h2>
      <p class="mt-5 text-base font-serif">
        SANICHI CORPORATION is used forklift (diesel, petrol, battery), We are
        a machine tool sales company.
        <br />
        We have a history of 40 years as a used forklift exporter, and have
        earned the trust of all customers, both domestic and overseas,
        including buyers and end users.
      </p>
    </div>

    <div class="m-12">
      <h2 class="text-3xl font-serif mb-5">Forklift</h2>
      <div class="flex flex-wrap justify-center gap-4">
        <div v-for="thumbnail in FORKLIFT_THUMBNAILS" :key="thumbnail.alt" class="w-[300px]">
          <img
            :src="thumbnail.src"
            :alt="thumbnail.alt"
            class="w-full h-[250px] object-cover"
          />
        </div>
      </div>
    </div>

    <div class="m-12">
      <h2 class="text-3xl font-serif mb-5">Machine Tool</h2>
      <div class="flex flex-wrap justify-center gap-4">
        <div v-for="thumbnail in MACHINERY_THUMBNAILS" :key="thumbnail.alt" class="w-[300px]">
          <img
            :src="thumbnail.src"
            :alt="thumbnail.alt"
            class="w-full h-[250px] object-cover"
          />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import HomeCarousel from '@/components/home/Carousel.vue'
import { FORKLIFT_THUMBNAILS, MACHINERY_THUMBNAILS } from '@/features/home/constants'
</script>