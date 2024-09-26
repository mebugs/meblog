<template>
  <!-- 路由引入组件的名称要和路由name保持一致 -->
  <router-view v-slot="{ Component, route }">
    <transition name="fade-transform" mode="out-in" appear>
      <keep-alive :include="cacheList">
        <component :is="Component" :key="route.fullPath" />
      </keep-alive>
    </transition>
  </router-view>
</template>

<script lang="ts" setup>
import { computed } from 'vue'
import { tabStore } from '@/store'

const tabs = tabStore()
// Tab缓存列表
const cacheList = computed(() => tabs.getCacheList)
</script>

<style></style>
