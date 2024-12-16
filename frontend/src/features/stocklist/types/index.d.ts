export interface Forklift {
    enginetype: string;
    maker: string;
    model: string;
    serialNo: string;
    height: number;
    ct: string;
    attachment: string;
    year: number;
    hourMeter: number;
    application: string;
    fob: number;
    createdAt: string;
    updatedAt: string;
}
export interface ForkliftDisplay extends Forklift {
    topImage?: string;
}
export type ForkliftType = 'battery' | 'diesel' | 'gasoline';
export interface StockTopItem {
    image: string;
    path: string;
    alt: string;
    label: string;
}
