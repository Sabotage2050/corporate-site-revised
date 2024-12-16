// src/hooks/forklift/useForkliftDetail.ts
import { storeToRefs } from 'pinia'
import { useForkliftStore } from '@/stores/forklift'

export function useForkliftDetail() {
  const store = useForkliftStore()
  const { currentForklift, loading, error } = storeToRefs(store)

  const actions = {
    async fetchForkliftByDetails(enginetype: string, model: string, serial: string) {
      await store.fetchForkliftByDetails(enginetype, model, serial)
    },
    clearCurrentForklift() {
      store.clearCurrentForklift()
    }
  }

  return {
    currentForklift,
    loading,
    error,
    actions
  }
}