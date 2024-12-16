
// src/features/todo/stores/todo.ts
import { defineStore } from 'pinia'
import type { Todo, NewTodo, TodoId, UpdateTodo } from '@/features/todo/types'
import { todoApi } from '@/features/todo/api'

interface TodoState {
  todos: Todo[]
  selectedTodoId: TodoId
  editingTodoId: TodoId
  loading: boolean
  error: string | null
}

export const useTodoStore = defineStore('todo', {
  state: (): TodoState => ({
    todos: [],
    selectedTodoId: null,
    editingTodoId: null,
    loading: false,
    error: null
  }),

  getters: {
    selectedTodo: (state): Todo | null => 
      state.selectedTodoId !== null 
        ? state.todos.find(todo => todo.id === state.selectedTodoId) ?? null
        : null,
    editingTodo: (state): Todo | null => 
      state.editingTodoId !== null 
        ? state.todos.find(todo => todo.id === state.editingTodoId) ?? null
        : null
  },

  actions: {
    async fetchTodos() {
      this.loading = true
      try {
        this.todos = await todoApi.fetchTodos()
      } catch (error) {
        this.error = '取得中にエラーが発生しました'
        throw error
      } finally {
        this.loading = false
      }
    },

    async addTodo(newTodo: NewTodo) {
      try {
        const todo = await todoApi.createTodo(newTodo)
        this.todos.push(todo)
      } catch (error) {
        this.error = '追加中にエラーが発生しました'
        throw error
      }
    },

    async updateTodo(id: TodoId, todo: UpdateTodo) {
      if (id === null) return
      try {
        const updatedTodo = await todoApi.updateTodo(id, {
          task: todo.task,
          done: todo.done
        })
        const index = this.todos.findIndex(t => t.id === id)
        if (index !== -1) {
          this.todos[index] = updatedTodo
        }
      } catch (error) {
        this.error = '更新中にエラーが発生しました'
        throw error
      }
    },

    async deleteTodo(id: TodoId) {
      if (id === null) return
      try {
        await todoApi.deleteTodo(id)
        this.todos = this.todos.filter(todo => todo.id !== id)
        
        // 削除後のクリーンアップ
        if (this.selectedTodoId === id) {
          this.selectedTodoId = null
        }
        if (this.editingTodoId === id) {
          this.editingTodoId = null
        }
      } catch (error) {
        this.error = '削除中にエラーが発生しました'
        throw error
      }
    },

    selectTodo(id: TodoId) {
      this.selectedTodoId = id
    },

    startEditing(id: TodoId) {
      this.editingTodoId = id
    },

    stopEditing() {
      this.editingTodoId = null
    }
  }
})