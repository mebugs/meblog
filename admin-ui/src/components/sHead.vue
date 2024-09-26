<template>
  <div class="head-wrapper">
    <div class="head-left">
      <a-button size="large" class="head-btn" :shape="'circle'" @click="setCollapsed">
        <template #icon>
          <icon-menu-fold v-if="!collapsed" :style="btnSize" />
          <icon-menu-unfold v-else :style="btnSize" />
        </template>
      </a-button>
    </div>
    <div class="head-center">
      <STab />
    </div>
    <div class="head-right">
      <SHeadBtns />
      <SHeadAvatar />
    </div>
  </div>
</template>

<script lang="ts" setup>
import { computed } from 'vue'
import STab from './sTab.vue'
import SHeadBtns from './sHeadBtns.vue'
import SHeadAvatar from './sHeadAvatar.vue'
import { appStore } from '@/store'
// 顶部组件
// 应用配置
const app = appStore()
// 折叠状态
const collapsed = computed(() => {
  return app.collapsed
})
// 设置折叠
const { setCollapsed, getHeadIconSize } = app
// 图标大小
const btnSize = getHeadIconSize
</script>
