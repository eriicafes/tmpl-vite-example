import { svelte } from "@sveltejs/vite-plugin-svelte";
import tailwindcss from "@tailwindcss/vite";
import react from "@vitejs/plugin-react";
import vue from "@vitejs/plugin-vue";
import { defineConfig, Plugin } from "vite";

export default defineConfig({
  plugins: [goDevRefresh(), react(), vue(), svelte(), tailwindcss()],
  build: {
    manifest: true,
    rollupOptions: {
      input: [
        "src/main.css",
        "src/react/index.tsx",
        "src/vue/index.ts",
        "src/svelte/index.ts",
      ],
    },
  },
});

export function goDevRefresh(): Plugin {
  return {
    name: "go-dev-refresh",
    configureServer(server) {
      server.middlewares.use("/__go-dev/refresh", (request, response, next) => {
        if (request.method !== "POST") return next();
        server.ws.send({ type: "full-reload" });
        response.statusCode = 204;
        response.end();
      });
    },
  };
}
