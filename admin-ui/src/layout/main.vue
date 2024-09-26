<template>
  <a-layout class="layout" :class="{ collapsed: collapsed }">
    <!-- 左侧导航栏 -->
    <a-layout-sider hide-trigger collapsible :width="220" :collapsed-width="60" :collapsed="collapsed">
      <div class="logo-wrapper">
        <img :src="logo[collapsed ? 'collapsed' : theme]" />
      </div>
      <div class="menu-wrapper">
        <SMenu />
      </div>
    </a-layout-sider>
    <a-layout class="layout-content">
      <SHead />
      <a-layout-content>
        <Full />
      </a-layout-content>
    </a-layout>
  </a-layout>
</template>

<script lang="ts" setup>
import { computed } from 'vue'
import SHead from '@/components/sHead.vue'
import SMenu from '@/components/sMenu.vue'
import Full from './full.vue' // 全屏路由模式
import { appStore } from '@/store'
// 应用配置
const app = appStore()
// 主题状态
const theme = computed(() => {
  return app.theme
})
// 折叠状态
const collapsed = computed(() => {
  return app.collapsed
})
// LOGO数据
const logo: any = {
  light: '/static/img/logo.png',
  dark: '/static/img/logo-dark.png',
  collapsed: '/static/img/icon.png'
}
</script>
