// src/hooks/forklift/useForkliftList.ts
import { computed } from 'vue'
import { storeToRefs } from 'pinia'
import { useForkliftStore } from '@/stores/forklift'
import type { ForkliftDisplay } from '@/features/stocklist/types'

export function useForkliftList() {
  const store = useForkliftStore()
  const { forklifts, loading, error } = storeToRefs(store)

const forkliftDisplayItems = computed(() => {
  return forklifts.value.map(forklift => {
    let topImage = ''
    try {
      const images = import.meta.glob('@/assets/stocklist/**/*.jpg', { eager: true })
      
      const imagePath = `/src/assets/stocklist/${forklift.enginetype}/${forklift.model}-${forklift.serialNo}/Top.jpg`
      
      const imageModule = images[imagePath]
      topImage = imageModule ? (imageModule as any).default : ''
    } catch (e) {
      console.warn(`Image not found for ${forklift.enginetype}/${forklift.model}`, e)
    }

    return {
      ...forklift,
      topImage
    } as ForkliftDisplay
  })
})

  const actions = {
    async fetchForkliftsByType(type: string) {
      await store.fetchForkliftsByType(type)
    }
  }

  return {
    forklifts: forkliftDisplayItems,
    loading,
    error,
    actions
  }
}