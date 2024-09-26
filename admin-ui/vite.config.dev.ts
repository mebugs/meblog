import { mergeConfig } from 'vite';
import baseConfig from './vite.config.ts';

export default mergeConfig(
  {
    mode: 'development',
    server: {
      port: 5000,
      open: true,
      proxy: {
        '/api': {
          target: 'http://127.0.0.1:8000',
          changeOrigin: true,
          rewrite: (path: string) => path.replace(/^\/api/, ''),
        },
      },
    },
    plugins: [],
  },
  baseConfig
);
