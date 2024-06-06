import { defineConfig } from 'vite'
import path from "path"
import { svelte } from '@sveltejs/vite-plugin-svelte'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [svelte()],
  build: {
    manifest: true,
    outDir: path.resolve(__dirname, "../priv/assets")
  }
})
