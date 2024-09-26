import { App } from 'vue';
import permission from './permission';

// 自定义指令注册
export default {
  install(Vue: App) {
    Vue.directive('permission', permission); // 权限指令
  },
};
