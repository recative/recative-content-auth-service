/// <reference types="vitest" />

import { defineConfig } from "vitest/config";

export default defineConfig({
  test: {
    deps: {
      fallbackCJS: true,
    },
    include: ["test/*.{test,spec}.{js,mjs,cjs,ts,mts,cts,jsx,tsx}"],
    globalSetup: ["./src/globalSetup/index.ts"],
  },
});
