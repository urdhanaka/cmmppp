import { svelte } from '@sveltejs/vite-plugin-svelte'
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { defineConfig } from 'vite'

// https://vite.dev/config/
export default defineConfig({
  plugins: [svelte()],
  resolve: {
    alias: {
      $lib: fileURLToPath(new URL('./src/lib', import.meta.url))
    }
  }
})
