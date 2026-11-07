import { defineConfig } from 'orval'

export default defineConfig({
  server: {
    input: '../openapi/server.openapi.yaml',
    output: {
      target: './src/apiclient/gen',
      mode: 'tags-split',
      client: 'axios-functions',
      mock: false,
    },
  },
})
