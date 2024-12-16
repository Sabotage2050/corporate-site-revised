// src/components/forklift/Forklift.vue
<template>
  <div class="bg-white rounded-lg shadow-sm p-6">
    <h2 class="text-2xl font-bold text-gray-900 mb-6">Forklift List</h2>

    <!-- Loading state -->
    <div v-if="loading" class="text-center py-8">
      <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-primary-600 mx-auto"></div>
    </div>

    <!-- Error display -->
    <div v-else-if="error" class="bg-red-50 text-red-600 p-4 rounded-lg mb-4">
      {{ error }}
    </div>

    <!-- Forklift list -->
    <div v-else class="overflow-x-auto">
      <table class="min-w-full divide-y divide-gray-200">
        <thead class="bg-gray-50">
          <tr>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Model</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Maker</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Height</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Year</th>
            <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Details</th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-gray-200">
          <tr v-for="forklift in forklifts" :key="`${forklift.enginetype}-${forklift.serialNo}`">
            <td class="px-6 py-4 whitespace-nowrap">{{ forklift.model }}</td>
            <td class="px-6 py-4 whitespace-nowrap">{{ forklift.maker }}</td>
            <td class="px-6 py-4 whitespace-nowrap">{{ forklift.height }}m</td>
            <td class="px-6 py-4 whitespace-nowrap">{{ forklift.year }}</td>
            <td class="px-6 py-4 whitespace-nowrap">
              <RouterLink 
                :to="`/stocklist/forklift/${forklift.enginetype}/${forklift.model}/${forklift.serialNo}`"
                class="text-blue-600 hover:text-blue-900"
              >
                View Details
              </RouterLink>
            </td>
          </tr>
        </tbody>
      </table>

      <!-- Empty state -->
      <div 
        v-if="forklifts.length === 0" 
        class="text-center py-8 text-gray-500"
      >
        No forklifts available.
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue'
import { storeToRefs } from 'pinia'
import { useForkliftStore } from '@/stores/forklift'
import { RouterLink } from 'vue-router'

const store = useForkliftStore()
const { forklifts, loading, error } = storeToRefs(store)

onMounted(async () => {
  await store.fetchForkliftsByType('all')
})
</script>

