<template>
  <div class="head-btn-item">
    <!-- 待办提示 -->
    <a-tooltip :content="$t('base.notification')">
      <a-button size="large" class="head-btn" :shape="'circle'">
        <template #icon>
          <icon-notification :style="btnSize" />
        </template>
      </a-button>
    </a-tooltip>
  </div>
</template>

<script lang="ts" setup>
import { appStore } from '@/store'
// 顶部按钮区域组件
// 应用配置
const app = appStore()
const { initTheme, getHeadIconSize } = app
// 初始化模式
initTheme()
// 图标大小
const btnSize = getHeadIconSize
</script>
