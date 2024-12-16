<!-- src/pages/Contact.vue -->
<template>
  <div class="bg-gray-300">
    <form @submit.prevent="handleSubmit" class="flex flex-col justify-center m-10">
      <h2 class="text-2xl text-black font-serif text-center mb-5">
        CONTACT FORM
      </h2>

      <div 
        v-for="field in CONTACT_FORM_FIELDS" 
        :key="field.id"
        class="md:flex md:items-center justify-center mb-5"
      >
        <label :for="field.id" class="mx-5 font-serif whitespace-pre-line">{{ field.label }}</label>
        
        <template v-if="field.type === 'textarea'">
          <textarea
            :id="field.id"
            :value="String(formData[field.id as keyof ContactFormData] || '')"
            @input="updateField(field.id, ($event.target as HTMLTextAreaElement).value)"
            :rows="field.rows"
            :cols="field.cols"
            class="bg-gray-300 border-2 border-black focus:outline-none focus:bg-white font-serif"
          />
        </template>
        
        <template v-else>
          <input
            :id="field.id"
            :value="String(formData[field.id as keyof ContactFormData] || '')"
            @input="updateField(field.id, ($event.target as HTMLInputElement).value)"
            :size="field.size"
            :required="field.required"
            class="bg-gray-300 border-2 border-black focus:outline-none focus:bg-white font-serif"
          />
        </template>

        <span 
          v-if="field.required && errors[field.id]"
          class="mx-5 text-red-600 font-serif"
        >
          Required field
        </span>
      </div>

      <div class="flex justify-center mt-10 mx-10 mb-10">
        <button 
          type="submit" 
          class="text-2xl font-serif"
          :disabled="loading"
        >
          Submit
        </button>
      </div>
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { CONTACT_FORM_FIELDS } from '@/features/contact/constants'
import type { ContactFormData } from '@/features/contact/types'
import { contactApi } from '@/features/contact/api'

const router = useRouter()
const loading = ref(false)
const errors = ref<Record<string, string>>({})

const formData = reactive<ContactFormData>({
  comment: '',
  name: '',
  corporate_name: '',
  email: '',
  phone_number: '',
  zip: '',
  address: '',
  subject: '新しいコンタクト',
  textBody: '',
  htmlBody: '',
  data: {},
  attachments: []
})

const updateField = (fieldId: string, value: string) => {
  if (fieldId in formData) {
    (formData as any)[fieldId] = value
  }
}

const handleSubmit = async () => {
  loading.value = true
  errors.value = {}

  try {
    // フォームのバリデーション
    let hasError = false
    CONTACT_FORM_FIELDS.forEach(field => {
      if (field.required && !formData[field.id as keyof ContactFormData]) {
        errors.value[field.id] = 'Required field'
        hasError = true
      }
    })

    if (hasError) {
      return
    }


    // フォームデータの送信
    await contactApi.submitContactForm(formData)

    // 成功時は結果ページへリダイレクト
    router.push('/contact/result')
  } catch (error) {
    console.error('Error submitting form:', error)
    errors.value.submit = 'Failed to submit form. Please try again.'
  } finally {
    loading.value = false
  }
}
</script>