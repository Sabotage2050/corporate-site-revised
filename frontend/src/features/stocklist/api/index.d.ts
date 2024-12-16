import type { Forklift } from '../types';
export declare const forkliftApi: {
    fetchForkliftsByType(type: string): Promise<Forklift[]>;
    getForkliftByDetails(enginetype: string, model: string, serial: string): Promise<Forklift>;
};
