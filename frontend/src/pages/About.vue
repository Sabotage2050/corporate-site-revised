<template>
  <div class="flex flex-col bg-gray-300 text-center">
    <div class="order-1">
      <ul class="flex">
        <li>
          <RouterLink to="/" class="text-blue-800 font-serif">Top</RouterLink>
        </li>
        <li>&nbsp;&gt;&nbsp;</li>
        <li class="font-serif">About us</li>
      </ul>
      <img
        src="@/assets/about/sanichi_about_us_logo.png"
        alt="Sanichi About Us"
        class="w-[500px] h-[60px]"
      />
    </div>

    <div class="order-2 mt-20 bg-gray-200 m-auto">
      <table class="border bg-gray-300">
        <tbody>
          <tr v-for="info in COMPANY_INFO" :key="info.label">
            <td class="border border-black font-serif">{{ info.label }}:</td>
            <td class="border border-black font-serif whitespace-pre-line">
              {{ info.value }}
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <div class="order-3 mt-20 mb-20 m-auto">
      <iframe
        src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3281.0506555165593!2d135.49153871552656!3d34.67867098044029!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x6000e7024c61ab53%3A0x4f0065c1d7e037ea!2z44CSNTUwLTAwMTMg5aSn6Ziq5bqc5aSn6Ziq5biC6KW_5Yy65paw55S677yR5LiB55uu77yT77yU4oiS77yS77yRIOOCqOODvOODgOOCpOODk-ODqw!5e0!3m2!1sja!2sjp!4v1661669622189!5m2!1sja!2sjp"
        width="450"
        height="450"
        loading="lazy"
        referrerpolicy="no-referrer-when-downgrade"
        class="border-0"
      ></iframe>
    </div>
  </div>
</template>

<script setup lang="ts">
import { RouterLink } from 'vue-router'
import { COMPANY_INFO } from '@/features/about/constants'
</script>