export declare function useForkliftDetail(): {
    currentForklift: import("vue").Ref<import("../../features/stocklist/types").Forklift | null, import("../../features/stocklist/types").Forklift | null>;
    loading: import("vue").Ref<boolean, boolean>;
    error: import("vue").Ref<string | null, string | null>;
    actions: {
        fetchForkliftByDetails(enginetype: string, model: string, serial: string): Promise<void>;
        clearCurrentForklift(): void;
    };
};
