import { ref } from 'vue'
import { storeToRefs } from 'pinia'
import { useTodoStore } from '@/stores/todo'
import type { Todo, TodoId } from '@/features/todo/types'

export function useTodoList() {
  const store = useTodoStore()
  const { todos, selectedTodo, editingTodo, loading, error } = storeToRefs(store)
  const newTodoTask = ref('')

  const actions = {
    async fetchTodos() {
      await store.fetchTodos()
    },

    async add() {
      if (!newTodoTask.value.trim()) return
      await store.addTodo({
        task: newTodoTask.value,
        done: false
      })
      newTodoTask.value = ''
    },

    async update(todo: Todo) {
      await store.updateTodo(todo.id, todo)
    },

    async delete(id: TodoId) {
      if (!confirm('本当に削除しますか？')) return
      await store.deleteTodo(id)
    },

    viewDetails(id: TodoId | null) {
      store.selectTodo(id)
    },

    startEditing(id: TodoId) {
      store.startEditing(id)
    },

    stopEditing() {
      store.stopEditing()
    },

    async saveEdit(todo: Todo) {
      if (!todo.task.trim()) return
      await store.updateTodo(todo.id, todo)
      store.stopEditing()
    }
  }

  return {
    // State
    todos,
    selectedTodo,
    editingTodo,
    loading,
    error,
    newTodoTask,
    // Actions
    actions
  }
}