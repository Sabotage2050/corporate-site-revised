// src/features/contact/constants/index.ts

import type { FormField } from '../types'

export const CONTACT_FORM_FIELDS: FormField[] = [
  {
    id: 'comment',
    label: 'detail',
    type: 'textarea',
    rows: 3,
    cols: 50
  },
  {
    id: 'name',
    label: 'Name',
    type: 'text',
    required: true
  },
  {
    id: 'corporate_name',
    label: 'Corporate Name / Organization Name',
    type: 'text'
  },
  {
    id: 'email',
    label: 'E-mail Address',
    type: 'text',
    required: true
  },
  {
    id: 'phone_number',
    label: 'Phone Number',
    type: 'text'
  },
  {
    id: 'zip',
    label: 'Postal Code',
    type: 'text',
    size: 15
  },
  {
    id: 'address',
    label: 'Country, City, Street Address,\nName of Apartment Building, etc.',
    type: 'text',
    size: 50
  }
]