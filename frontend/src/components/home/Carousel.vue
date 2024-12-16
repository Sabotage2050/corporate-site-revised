<template>
  <div class="carousel">
    <swiper
      :modules="[Navigation, Pagination]"
      :slides-per-view="1"
      :pagination="{ clickable: true }"
      :navigation="true"
      :loop="true"
      class="h-[600px]"
    >
      <swiper-slide v-for="(image, index) in CAROUSEL_IMAGES" :key="index">
        <img
          :src="image.src"
          :alt="image.alt"
          class="w-full h-full object-cover"
        />
      </swiper-slide>
    </swiper>
  </div>
</template>

<script setup lang="ts">
import { Swiper, SwiperSlide } from 'swiper/vue'
import { Pagination, Navigation } from 'swiper/modules'
import { CAROUSEL_IMAGES } from '@/features/home/constants'
import 'swiper/css'
import 'swiper/css/pagination'
import 'swiper/css/navigation'
</script>