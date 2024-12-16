// src/stores/forklift.ts
import { defineStore } from 'pinia'
import type { Forklift } from '@/features/stocklist/types'
import { forkliftApi } from '@/features/stocklist/api'

interface ForkliftState {
  forklifts: Forklift[]
  currentForklift: Forklift | null
  loading: boolean
  error: string | null
}

export const useForkliftStore = defineStore('forklift', {
  state: (): ForkliftState => ({
    forklifts: [],
    currentForklift: null,
    loading: false,
    error: null
  }),

  actions: {

    clearCurrentForklift() {
      this.currentForklift = null
    },

    async fetchForkliftsByType(type: string) {
      this.loading = true
      try {
        this.forklifts = await forkliftApi.fetchForkliftsByType(type)
      } catch (error) {
        this.error = 'Failed to fetch forklifts by type'
        throw error
      } finally {
        this.loading = false
      }
    },

    async fetchForkliftByDetails(enginetype: string, model: string, serial: string) {
      this.loading = true
      try {
        this.currentForklift = await forkliftApi.getForkliftByDetails(enginetype, model, serial)
      } catch (error) {
        this.error = 'Failed to fetch forklift details'
        throw error
      } finally {
        this.loading = false
      }
    }
  }
})