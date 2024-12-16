import type { Todo, TodoId } from '@/features/todo/types';
export declare function useTodoList(): {
    todos: import("vue").Ref<Todo[], Todo[]>;
    selectedTodo: import("vue").ComputedRef<Todo | null>;
    editingTodo: import("vue").ComputedRef<Todo | null>;
    loading: import("vue").Ref<boolean, boolean>;
    error: import("vue").Ref<string | null, string | null>;
    newTodoTask: import("vue").Ref<string, string>;
    actions: {
        fetchTodos(): Promise<void>;
        add(): Promise<void>;
        update(todo: Todo): Promise<void>;
        delete(id: TodoId): Promise<void>;
        viewDetails(id: TodoId | null): void;
        startEditing(id: TodoId): void;
        stopEditing(): void;
        saveEdit(todo: Todo): Promise<void>;
    };
};
