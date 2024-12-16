import { Todo } from '@/features/todo/types';
type __VLS_Props = {
    todo: Todo;
};
declare const _default: import("vue").DefineComponent<__VLS_Props, {}, {}, {}, {}, import("vue").ComponentOptionsMixin, import("vue").ComponentOptionsMixin, {} & {
    update: (todo: Todo) => any;
    view: (id: number) => any;
    delete: (id: number) => any;
    edit: (id: number) => any;
}, string, import("vue").PublicProps, Readonly<__VLS_Props> & Readonly<{
    onUpdate?: ((todo: Todo) => any) | undefined;
    onView?: ((id: number) => any) | undefined;
    onDelete?: ((id: number) => any) | undefined;
    onEdit?: ((id: number) => any) | undefined;
}>, {}, {}, {}, {}, string, import("vue").ComponentProvideOptions, false, {}, HTMLLIElement>;
export default _default;
