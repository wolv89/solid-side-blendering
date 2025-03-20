import { defineConfig } from 'vite';
import solidPlugin from 'vite-plugin-solid';

export default defineConfig({
  plugins: [solidPlugin()],
  server: {
    port: 8300,
    allowedHosts: [
      'blewb.build'
    ]
  },
  build: {
    target: 'esnext',
    outDir: '../dist',
    emptyOutDir: true
  },
});
