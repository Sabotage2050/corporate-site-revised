<template>
  <li class="py-4 flex items-center gap-4 group hover:bg-gray-50">
    <button
      @click="handleUpdate"
      class="w-6 h-6 rounded-full border-2 border-primary-600 flex items-center justify-center transition-colors"
      :class="{ 'bg-primary-600': todo.done }"
    >
      <span v-if="todo.done" class="text-white text-sm">✓</span>
    </button>
    
    <span
      class="flex-1"
      :class="{ 'line-through text-gray-400': todo.done }"
    >
      {{ todo.task }}
    </span>

    <!-- opacity-0を削除し、常に表示されるように変更 -->
    <div class="flex items-center gap-2 transition-opacity">
      <button
        @click="$emit('view', todo.id)"
        class="px-3 py-1 text-sm bg-blue-500 text-white rounded hover:bg-blue-600 transition-colors"
        data-testid="view-btn"
      >
        詳細
      </button>
      <button
        @click="$emit('edit', todo.id)"
        class="px-3 py-1 text-sm bg-green-500 text-white rounded hover:bg-green-600 transition-colors"
        data-testid="edit-btn"
      >
        編集
      </button>
      <button
        @click="$emit('delete', todo.id)"
        class="px-3 py-1 text-sm bg-red-500 text-white rounded hover:bg-red-600 transition-colors"
        data-testid="delete-btn"
      >
        削除
      </button>
    </div>
  </li>
</template>

<script setup lang="ts">
import { Todo } from '@/features/todo/types';

const props = defineProps<{
  todo: Todo
}>()

const emit = defineEmits<{
  (e: 'update', todo: Todo): void
  (e: 'view', id: number): void
  (e: 'edit', id: number): void
  (e: 'delete', id: number): void
}>()

const handleUpdate = () => {
  emit('update', {
    ...props.todo,
    done: !props.todo.done
  })
}
</script>