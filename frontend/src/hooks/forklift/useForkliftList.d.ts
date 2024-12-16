import type { ForkliftDisplay } from '@/features/stocklist/types';
export declare function useForkliftList(): {
    forklifts: import("vue").ComputedRef<ForkliftDisplay[]>;
    loading: import("vue").Ref<boolean, boolean>;
    error: import("vue").Ref<string | null, string | null>;
    actions: {
        fetchForkliftsByType(type: string): Promise<void>;
    };
};
