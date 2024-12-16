import type { Forklift } from '@/features/stocklist/types';
interface ForkliftState {
    forklifts: Forklift[];
    currentForklift: Forklift | null;
    loading: boolean;
    error: string | null;
}
export declare const useForkliftStore: import("pinia").StoreDefinition<"forklift", ForkliftState, {}, {
    clearCurrentForklift(): void;
    fetchForkliftsByType(type: string): Promise<void>;
    fetchForkliftByDetails(enginetype: string, model: string, serial: string): Promise<void>;
}>;
export {};
