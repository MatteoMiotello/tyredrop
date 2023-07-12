import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';
import viteSvgr from "vite-plugin-svgr";
import {esbuildCommonjs} from "@originjs/vite-plugin-commonjs";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
      react(),
      viteSvgr(),
  ],
    optimizeDeps: {
        esbuildOptions: {
            plugins: [esbuildCommonjs(['react-moment'])],
        },
    },
});
