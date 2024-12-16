import apiClient from '@/utils/api/client'
import type {Forklift} from '../types'

export const forkliftApi = {
  async fetchForkliftsByType(type: string) {
    const { data } = await apiClient.get<Forklift[]>(`/forklifts/type/${type}`)
    return data
  },
  async getForkliftByDetails(enginetype: string, model: string, serial: string) {
    const { data } = await apiClient.get<Forklift>(
      `/forklifts/type/${enginetype}/${model}/${serial}`
    )
    return data
  }
}