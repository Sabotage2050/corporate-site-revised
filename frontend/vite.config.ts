/// <reference types="vitest" />
import { defineConfig, loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path'

// https://vitejs.dev/config/
export default defineConfig(({ command, mode }) => {
  // 環境変数をロード（processEnvをtrueに設定）
  const env = loadEnv(mode, process.cwd(), '')

  return {
    plugins: [vue()],
    server: {
      host: true,
      port: 3000
    },
    build: {
      sourcemap: true,
       // ビルドキャッシュを無効化
    cache: false,
      outDir: 'dist',
      rollupOptions: {
        input: {
          main: path.resolve(__dirname, 'index.html')
        }
      },
            // TypeScriptの出力を制御
      typescript: {
        noEmit: true,
      }
    },
    preview: {
      port: 3000
    },
    resolve: {
      alias: {
        '@': path.resolve(__dirname, './src')
      }
    },
    base: './',
    optimizeDeps: {
      include: ['swiper', 'swiper/vue']
    },
    define: {
  'import.meta.env.VITE_API_BASE_URL': JSON.stringify(env.VITE_API_BASE_URL),
  'import.meta.env.VITE_APP_ENV': JSON.stringify(env.VITE_APP_ENV)
}
  }
})