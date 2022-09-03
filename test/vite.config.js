import { defineConfig } from "vite";

// https://vitejs.dev/config/
export default defineConfig({
  build: {
    outDir: "dist/static",
  },
  server: {
    host: true,
    port: 53172,
  },
  optimizeDeps: {
    exclude: ["photon-lib-js"],
  },
});
