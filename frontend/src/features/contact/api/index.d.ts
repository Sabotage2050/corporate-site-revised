import type { ContactFormData } from '../types';
export declare const contactApi: {
    submitContactForm(formData: ContactFormData): Promise<{
        text: string;
    }>;
};
