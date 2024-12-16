<template>
  <div v-if="selectedTodo || editingTodo">
    <!-- 背景オーバーレイ -->
    <div 
      class="fixed inset-0 bg-black bg-opacity-50 z-40"
      @click="editingTodo ? $emit('stop-editing') : $emit('close')"
    ></div>

    <!-- モーダルコンテンツ -->
    <div class="fixed inset-0 flex items-center justify-center z-50">
      <!-- 詳細モーダル -->
      <div 
        v-if="selectedTodo"
        class="bg-white p-6 rounded-lg shadow-xl max-w-md w-full mx-4"
      >
        <h3 class="text-lg font-medium text-gray-900 mb-4">
          Todo 詳細
        </h3>
        <div class="space-y-4">
          <p><span class="font-medium">タスク:</span> {{ selectedTodo.task }}</p>
          <p><span class="font-medium">状態:</span> {{ selectedTodo.done ? '完了' : '未完了' }}</p>
        </div>
        <div class="mt-6">
          <button
            @click="$emit('close')"
            class="px-4 py-2 bg-gray-100 text-gray-700 rounded hover:bg-gray-200 transition-colors"
            data-testid="close-btn"
          >
            閉じる
          </button>
        </div>
      </div>

      <!-- 編集モーダル -->
      <div
        v-if="editingTodo"
        class="bg-white p-6 rounded-lg shadow-xl max-w-md w-full mx-4"
      >
        <h3 class="text-lg font-medium text-gray-900 mb-4">
          Todoの編集
        </h3>
        <form @submit.prevent="handleSave">
          <input
            v-model="editedTask"
            type="text"
            class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary-500 focus:border-transparent"
            required
            data-testid="edit-input"
          >
          <div class="mt-4">
            <label class="flex items-center gap-2">
              <input
                v-model="editedDone"
                type="checkbox"
                class="w-4 h-4 text-primary-600 border-gray-300 rounded focus:ring-primary-500"
                data-testid="edit-done"
              >
              <span>完了</span>
            </label>
          </div>
          <div class="mt-6 flex gap-4">
            <button
              type="submit"              
              class="px-4 py-2 bg-gray-100 text-gray-700 rounded hover:bg-gray-200 transition-colors"
              data-testid="save-btn"
            >
              保存
            </button>
            <button
              type="button"
              @click="$emit('stop-editing')"
              class="px-4 py-2 bg-gray-100 text-gray-700 rounded hover:bg-gray-200 transition-colors"
              data-testid="cancel-btn"
            >
              キャンセル
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'
import type { Todo } from '@/features/todo/types'

const props = defineProps<{
  selectedTodo: Todo | null
  editingTodo: Todo | null
}>()

const emit = defineEmits<{
  (e: 'close'): void
  (e: 'stop-editing'): void
  (e: 'save', todo: Todo): void
}>()

const editedTask = ref('')
const editedDone = ref(false)

// 編集データの初期化
watch(() => props.editingTodo, (newTodo) => {
  if (newTodo) {
    editedTask.value = newTodo.task
    editedDone.value = newTodo.done
  }
}, { immediate: true })

const handleSave = () => {
  if (!props.editingTodo || !editedTask.value.trim()) return
  
  // saveイベントを発火する前にフォームのサブミットを防ぐ
  emit('save', {
    id: props.editingTodo.id,
    task: editedTask.value.trim(),
    done: editedDone.value
  })
}
</script>