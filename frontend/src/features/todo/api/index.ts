// src/features/todo/api/index.ts
import axios from 'axios'
import type { Todo, NewTodo, UpdateTodo, TodoId } from '../types'

const BASE_URL = import.meta.env.VITE_API_BASE_URL
axios.defaults.baseURL = BASE_URL

export const todoApi = {
  async fetchTodos() {
    const { data } = await axios.get<Todo[]>('/todos')
    return data
  },

  async getTodo(id: TodoId) {
    if (id === null) throw new Error('Invalid ID')
    const { data } = await axios.get<Todo>(`/todos/${id}`)
    return data
  },

  async createTodo(todo: NewTodo) {
    const { data } = await axios.post<Todo>('/todos', todo)
    return data
  },

  async updateTodo(id: TodoId, todo: UpdateTodo) {
    if (id === null) throw new Error('Invalid ID')
    const { data } = await axios.put<Todo>(`/todos/${id}`, todo)
    return data
  },

  async deleteTodo(id: TodoId) {
    if (id === null) throw new Error('Invalid ID')
    await axios.delete(`/todos/${id}`)
  }
}