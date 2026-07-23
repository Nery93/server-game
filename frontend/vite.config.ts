import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    proxy: {
      // Forwards /game* requests to the Go backend during development,
      // so the browser never has to deal with cross-origin requests.
      '/game': 'http://localhost:8080',
    },
  },
})
