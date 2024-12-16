import type { Todo } from '@/features/todo/types';
type __VLS_Props = {
    selectedTodo: Todo | null;
    editingTodo: Todo | null;
};
declare const _default: import("vue").DefineComponent<__VLS_Props, {}, {}, {}, {}, import("vue").ComponentOptionsMixin, import("vue").ComponentOptionsMixin, {} & {
    close: () => any;
    "stop-editing": () => any;
    save: (todo: Todo) => any;
}, string, import("vue").PublicProps, Readonly<__VLS_Props> & Readonly<{
    onClose?: (() => any) | undefined;
    "onStop-editing"?: (() => any) | undefined;
    onSave?: ((todo: Todo) => any) | undefined;
}>, {}, {}, {}, {}, string, import("vue").ComponentProvideOptions, false, {}, any>;
export default _default;
