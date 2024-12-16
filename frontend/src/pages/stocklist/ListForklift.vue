<template>
  <div class="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
    <div class="px-4 py-6 sm:px-0">
      <div class="flex justify-between items-center mb-6">
        <h1 class="text-3xl font-bold text-gray-900">{{ capitalizedType }} Forklift Stock</h1>
      </div>
      
      <!-- Loading state -->
      <div v-if="loading" class="text-center py-4">
        Loading...
      </div>
      <!-- Error state -->
      <div v-else-if="error" class="text-center py-4 text-red-600">
        {{ error }}
      </div>
      <!-- Content -->
      <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <div v-for="item in forkliftItems" 
             :key="`${item.enginetype}-${item.serialNo}`" 
             class="bg-white rounded-lg shadow overflow-hidden">
          <img v-if="item.topImage" 
               :src="item.topImage" 
               :alt="`${item.maker} ${item.model}`"
               class="w-full h-48 object-cover"/>
          <div v-else class="w-full h-48 bg-gray-200 flex items-center justify-center">
            <span class="text-gray-500">No image available</span>
          </div>
          <div class="p-4">
            <h3 class="text-lg font-semibold text-gray-900">{{ item.maker }} {{ item.model }}</h3>
            <p class="text-sm text-gray-600">Height: {{ item.height }}m</p>
            <p class="text-sm text-gray-600">Year: {{ item.year }}</p>
            <router-link 
              :to="`/stocklist/forklift/${item.enginetype}/${item.model}/${item.serialNo}`"
              class="mt-2 inline-block text-blue-600 hover:text-blue-800"
            >
              View Details
            </router-link>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted } from 'vue'
import { useForkliftList } from '@/hooks/forklift/useForkliftList'
import type { ForkliftType } from '@/features/stocklist/types'

const props = defineProps<{
  type: ForkliftType
}>()

const { forklifts, loading, error, actions } = useForkliftList()
const forkliftItems = computed(() => forklifts.value)
const capitalizedType = computed(() => {
  return props.type.charAt(0).toUpperCase() + props.type.slice(1)
})

onMounted(async () => {
  await actions.fetchForkliftsByType(props.type)
})
</script>