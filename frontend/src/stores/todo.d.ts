import type { Todo, NewTodo, TodoId, UpdateTodo } from '@/features/todo/types';
interface TodoState {
    todos: Todo[];
    selectedTodoId: TodoId;
    editingTodoId: TodoId;
    loading: boolean;
    error: string | null;
}
export declare const useTodoStore: import("pinia").StoreDefinition<"todo", TodoState, {
    selectedTodo: (state: {
        todos: {
            id: number;
            task: string;
            done: boolean;
        }[];
        selectedTodoId: TodoId;
        editingTodoId: TodoId;
        loading: boolean;
        error: string | null;
    } & import("pinia").PiniaCustomStateProperties<TodoState>) => Todo | null;
    editingTodo: (state: {
        todos: {
            id: number;
            task: string;
            done: boolean;
        }[];
        selectedTodoId: TodoId;
        editingTodoId: TodoId;
        loading: boolean;
        error: string | null;
    } & import("pinia").PiniaCustomStateProperties<TodoState>) => Todo | null;
}, {
    fetchTodos(): Promise<void>;
    addTodo(newTodo: NewTodo): Promise<void>;
    updateTodo(id: TodoId, todo: UpdateTodo): Promise<void>;
    deleteTodo(id: TodoId): Promise<void>;
    selectTodo(id: TodoId): void;
    startEditing(id: TodoId): void;
    stopEditing(): void;
}>;
export {};
