import type { Todo, NewTodo, UpdateTodo, TodoId } from '../types';
export declare const todoApi: {
    fetchTodos(): Promise<Todo[]>;
    getTodo(id: TodoId): Promise<Todo>;
    createTodo(todo: NewTodo): Promise<Todo>;
    updateTodo(id: TodoId, todo: UpdateTodo): Promise<Todo>;
    deleteTodo(id: TodoId): Promise<void>;
};
