/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}"  // すべてのVueファイルでTailwindを使用可能に
  ],
  theme: {
    extend: {},
  },
  plugins: [],
}
