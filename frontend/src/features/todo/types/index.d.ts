export interface Todo {
    id: number;
    task: string;
    done: boolean;
}
export interface NewTodo {
    task: string;
    done: boolean;
}
export interface UpdateTodo {
    task: string;
    done: boolean;
}
export type TodoId = number | null;
export interface ApiResponse<T> {
    data: T;
    error?: string;
}
export interface ApiError {
    message: string;
    code?: string;
}
