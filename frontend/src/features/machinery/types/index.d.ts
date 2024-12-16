export interface MachineImage {
    src: string;
    alt: string;
    model?: string;
}
export interface MachineCategory {
    title: string;
    images: MachineImage[];
}
