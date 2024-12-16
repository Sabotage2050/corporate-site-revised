<!-- TodoForm.vue -->
<template>
  <form @submit.prevent="handleSubmit" class="mb-6">
    <div class="flex gap-4">
      <input 
        :value="modelValue"
        @input="$emit('update:modelValue', ($event.target as HTMLInputElement).value)"
        type="text"
        placeholder="新しいタスクを入力"
        class="flex-1 px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary-500 focus:border-transparent"
        :disabled="loading"
        required
      >
      <button 
        type="submit"
        class="px-6 py-2 bg-primary-600 text-white rounded-lg hover:bg-primary-700 transition-colors"
        :disabled="loading"
      >
        追加
      </button>
    </div>
  </form>
</template>

<script setup lang="ts">
const props = defineProps<{
  modelValue: string
  loading: boolean
}>()

const emit = defineEmits<{
  (e: 'update:modelValue', value: string): void
  (e: 'submit'): void
}>()

const handleSubmit = (e: Event) => {
  if (props.loading) {
    e.preventDefault()
    return
  }
  emit('submit')
}
</script>