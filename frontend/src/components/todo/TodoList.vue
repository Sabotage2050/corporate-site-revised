// src/features/todo/components/TodoList.vue
<template>
  <div class="bg-white rounded-lg shadow-sm">
    <div class="p-6">
      <h1 class="text-2xl font-bold text-gray-900 mb-6">Todo List</h1>
      
      <TodoForm 
        v-model="newTodoTask"
        :loading="loading"
        @submit="actions.add"
      />

      <!-- ローディング状態 -->
      <div v-if="loading" class="text-center py-8">
        <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-primary-600 mx-auto"></div>
      </div>

      <!-- エラー表示 -->
      <div v-else-if="error" class="bg-red-50 text-red-600 p-4 rounded-lg mb-4">
        {{ error }}
      </div>

      <!-- Todoリスト -->
      <TransitionGroup name="list" tag="ul" class="divide-y divide-gray-200">
        <TodoItem
          v-for="todo in todos"
          :key="todo.id"
          :todo="todo"
          @update="actions.update"
          @view="actions.viewDetails"
          @edit="actions.startEditing"
          @delete="actions.delete"
        />
      </TransitionGroup>

      <!-- 空の状態 -->
      <div 
        v-if="!loading && !error && todos.length === 0" 
        class="text-center py-8 text-gray-500"
      >
        タスクがありません。新しいタスクを追加してください。
      </div>
    </div>

    <TodoModal
      :selected-todo="selectedTodo ?? null"
      :editing-todo="editingTodo ?? null"
      @close="actions.viewDetails(null)"
      @stop-editing="actions.stopEditing"
      @save="actions.saveEdit"
    />
  </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue'
import { useTodoList } from '@/hooks/todo/useTodoList'
import TodoForm from './TodoForm.vue'
import TodoItem from './TodoItem.vue'
import TodoModal from './TodoModal.vue'

const {
  todos,
  newTodoTask,
  selectedTodo,
  editingTodo,
  loading,
  error,
  actions
} = useTodoList()

onMounted(async () => {
  await actions.fetchTodos()
})
</script>