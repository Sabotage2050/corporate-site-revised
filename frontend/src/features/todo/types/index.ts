export interface Todo {
  id: number
  task: string
  done: boolean
}

export interface NewTodo {
  task: string
  done: boolean
}

export interface UpdateTodo {
  task: string
  done: boolean
}
// nullを許容するTodoId型
export type TodoId = number | null

// API関連のレスポンス型
export interface ApiResponse<T> {
  data: T
  error?: string
}

// エラー型
export interface ApiError {
  message: string
  code?: string
}