import { defineConfig } from 'vite'
import preact from '@preact/preset-vite'
import viteCompression from 'vite-plugin-compression';
import eslint from 'vite-plugin-eslint'
import zlib from 'zlib';
import { VitePWA } from 'vite-plugin-pwa'


// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    eslint(),
    preact(),
    VitePWA({
      manifest: {
        name: 'SVLZ',
        short_name: 'SVLZ',
        id: '/',
        start_url: '/?from=homescreen',
        display: 'standalone',
        orientation: 'portrait-primary',
        lang: 'en',
        dir: 'ltr',
        display_override: [
          'standalone',
          'minimal-ui',
          'window-controls-overlay'
        ],
        prefer_related_applications: false,
      }
    }),
    viteCompression({
      algorithm: 'brotliCompress',
      filter: /\.(js|mjs|json|txt|css|html|svg|woff2)$/i,
      compressionOptions: {
        params: {
					[zlib.constants.BROTLI_PARAM_QUALITY]: zlib.constants.BROTLI_MAX_QUALITY,
          [zlib.constants.BROTLI_PARAM_MODE]: zlib.constants.BROTLI_MODE_TEXT,
				},
      }
    })
  ],
})
