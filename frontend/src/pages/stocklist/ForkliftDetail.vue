<template>
  <div class="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
    <div class="px-4 py-6 sm:px-0">
      <!-- Loading state -->
      <div v-if="loading" class="text-center py-8">
        <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-primary-600 mx-auto"></div>
      </div>

      <!-- Error state -->
      <div v-else-if="error" class="bg-red-50 text-red-600 p-4 rounded-lg">
        {{ error }}
      </div>

      <!-- Forklift details -->
      <div v-else-if="currentForklift" class="bg-white shadow overflow-hidden sm:rounded-lg">
        <!-- Image Gallery -->
        <div class="px-4 py-5 sm:px-6">
          <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div v-for="(image, index) in forkliftImages" 
                 :key="index"
                 class="aspect-w-16 aspect-h-9">
              <img 
                :src="image" 
                :alt="`${currentForklift.maker} ${currentForklift.model} - Image ${index + 1}`"
                class="object-cover rounded-lg shadow-lg"
              />
            </div>
          </div>
        </div>

        <!-- Details -->
        <div class="border-t border-gray-200">
          <dl>
            <div class="bg-gray-50 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
              <dt class="text-sm font-medium text-gray-500">Maker</dt>
              <dd class="mt-1 text-sm text-gray-900 sm:mt-0 sm:col-span-2">
                {{ currentForklift.maker }}
              </dd>
            </div>
            
            <div class="bg-white px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
              <dt class="text-sm font-medium text-gray-500">Model</dt>
              <dd class="mt-1 text-sm text-gray-900 sm:mt-0 sm:col-span-2">
                {{ currentForklift.model }}
              </dd>
            </div>

            <div class="bg-gray-50 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
              <dt class="text-sm font-medium text-gray-500">Serial No</dt>
              <dd class="mt-1 text-sm text-gray-900 sm:mt-0 sm:col-span-2">
                {{ currentForklift.serialNo }}
              </dd>
            </div>

            <div class="bg-white px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
              <dt class="text-sm font-medium text-gray-500">Height</dt>
              <dd class="mt-1 text-sm text-gray-900 sm:mt-0 sm:col-span-2">
                {{ currentForklift.height }}m
              </dd>
            </div>

            <div class="bg-gray-50 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
              <dt class="text-sm font-medium text-gray-500">Year</dt>
              <dd class="mt-1 text-sm text-gray-900 sm:mt-0 sm:col-span-2">
                {{ currentForklift.year }}
              </dd>
            </div>

            <div class="bg-white px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
              <dt class="text-sm font-medium text-gray-500">Hour Meter</dt>
              <dd class="mt-1 text-sm text-gray-900 sm:mt-0 sm:col-span-2">
                {{ currentForklift.hourMeter }} hours
              </dd>
            </div>

            <div class="bg-gray-50 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
              <dt class="text-sm font-medium text-gray-500">Attachment</dt>
              <dd class="mt-1 text-sm text-gray-900 sm:mt-0 sm:col-span-2">
                {{ currentForklift.attachment }}
              </dd>
            </div>

            <div class="bg-white px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
              <dt class="text-sm font-medium text-gray-500">Application</dt>
              <dd class="mt-1 text-sm text-gray-900 sm:mt-0 sm:col-span-2">
                {{ currentForklift.application }}
              </dd>
            </div>

            <div class="bg-gray-50 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
              <dt class="text-sm font-medium text-gray-500">FOB</dt>
              <dd class="mt-1 text-sm text-gray-900 sm:mt-0 sm:col-span-2">
                ¥{{ currentForklift.fob.toLocaleString() }}
              </dd>
            </div>
          </dl>
        </div>
      </div>

      <div class="mt-6">
        <router-link 
          :to="`/stocklist/forklift/${route.params.type}`"
          class="text-blue-600 hover:text-blue-900"
        >
          Back to List
        </router-link>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, onUnmounted } from 'vue'
import { useRoute } from 'vue-router'
import { useForkliftDetail } from '@/hooks/forklift/useForkliftDetail'

const route = useRoute()
const { currentForklift, loading, error, actions } = useForkliftDetail()

// 画像の取得
const forkliftImages = computed(() => {
  if (!currentForklift.value) return []
  
  const images = []
  try {
    const imageModules = import.meta.glob('/src/assets/stocklist/**/*.jpg', { eager: true })
    for (let i = 1; i <= 6; i++) {
      const imagePath = `/src/assets/stocklist/${currentForklift.value.enginetype}/${currentForklift.value.model}-${currentForklift.value.serialNo}/${currentForklift.value.model}-${currentForklift.value.serialNo}_${i}.jpg`
      if (imageModules[imagePath]) {
        images.push((imageModules[imagePath] as any).default)
      }
    }
    // Top画像も追加
    const topImagePath = `/src/assets/stocklist/${currentForklift.value.enginetype}/${currentForklift.value.model}-${currentForklift.value.serialNo}/Top.jpg`
    if (imageModules[topImagePath]) {
      images.push((imageModules[topImagePath] as any).default)
    }
  } catch (e) {
    console.warn('Error loading images:', e)
  }
  
  return images
})

onMounted(async () => {
  const { type, model, serial } = route.params
  if (typeof type === 'string' && typeof model === 'string' && typeof serial === 'string') {
    await actions.fetchForkliftByDetails(type, model, serial)
  }
})

onUnmounted(() => {
  actions.clearCurrentForklift()
})
</script>