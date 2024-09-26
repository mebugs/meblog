import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import ArcoVue from '@arco-design/web-vue'
import ArcoVueIcon from '@arco-design/web-vue/es/icon'
import '~/@arco-design/web-vue/dist/arco.css'
import '@/assets/style/base.less'
import i18n from './locale' // 国际化注册
import router from './router' // 路由注册
import store from './store' // 全局状态管理注册
import directive from './directive' // 全局指令注册
import '@/api/interceptor' // 全局API拦截器

const app = createApp(App)
app.use(ArcoVue)
app.use(ArcoVueIcon)
app.use(i18n)
app.use(router)
app.use(store)
app.use(directive)
app.mount('#app')
