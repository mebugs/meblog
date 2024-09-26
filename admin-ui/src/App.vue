<template>
  <a-config-provider :locale="locale">
    <router-view />
  </a-config-provider>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import enUS from '@arco-design/web-vue/es/locale/lang/en-us' // arco语言包
import zhCN from '@arco-design/web-vue/es/locale/lang/zh-cn' // arco语言包
import useLocale from '@/utils/hooks/locale'
const { currentLocale } = useLocale()
// 语言包绑定
const locale = computed(() => {
  switch (currentLocale.value) {
    case 'zh-CN':
      return zhCN
    case 'en-US':
      return enUS
    default:
      return enUS
  }
})
</script>

<style scoped></style>
