import { mergeConfig } from 'vite';
import baseConfig from './vite.config.ts';

export default mergeConfig(
  {
    mode: 'production',
  },
  baseConfig
);
