// src/features/contact/api/index.ts

import axios from 'axios'
import type { ContactFormData } from '../types'

const BASE_URL = import.meta.env.VITE_API_BASE_URL
axios.defaults.baseURL = BASE_URL

export const contactApi = {
  async submitContactForm(formData: ContactFormData) {
    const emailRequest = {
      subject: formData.subject,
      textBody: formData.comment,
      htmlBody: formData.htmlBody,
      data: {
        name: formData.name,
        corporate_name: formData.corporate_name,
        email: formData.email,
        phone_number: formData.phone_number,
        zip: formData.zip,
        address: formData.address,
      },
      attachments: formData.attachments,
    };

    const { data } = await axios.post<{ text: string }>('/email/send', emailRequest);
    return data;
  },
}