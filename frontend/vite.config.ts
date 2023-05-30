import { defineConfig } from 'vite'
import preact from '@preact/preset-vite'
import viteCompression from 'vite-plugin-compression';
import zlib from 'zlib';


// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    preact(),
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
