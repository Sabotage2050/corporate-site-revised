export interface ContactFormData {
    comment: string;
    name: string;
    corporate_name: string;
    email: string;
    phone_number: string;
    zip: string;
    address: string;
    subject: string;
    textBody: string;
    htmlBody?: string;
    data: {
        [key: string]: string;
    };
    attachments?: EmailAttachment[];
}
export interface FormField {
    id: string;
    label: string;
    type: 'text' | 'textarea';
    required?: boolean;
    size?: number;
    rows?: number;
    cols?: number;
}
export interface EmailAttachment {
    filename: string;
    content: string;
    mimeType: string;
}
