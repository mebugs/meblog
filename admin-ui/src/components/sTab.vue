<template>
  <s-tab-item v-ripples v-for="(tag, index) in tagList" :key="tag.name" :index="index" :item="tag" />
</template>

<script lang="ts" setup>
import { computed, onMounted, onUnmounted } from 'vue'
import { listenerRouteChange, removeRouteListener } from '@/utils/routeListener'
import STabItem from './sTabItem.vue'
import { tabStore } from '@/store'
const tab = tabStore()
const tagList = computed(() => {
  return tab.getTabList
})
// 路由切换
listenerRouteChange((route) => {
  if (!route.meta.noAffix && tagList.value && !tagList.value.some((tag) => tag.name === route.name)) {
    tab.addTabList(route)
  }
}, true)
// 标签初始化
onMounted(() => {})

// 注销监听
onUnmounted(() => {
  removeRouteListener()
})
</script>
