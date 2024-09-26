<template>
  <div class="stab" :class="{ stabCheck: routeCheck, stabMounted: mounteDone }" @click="go()">
    <div class="stabi" :class="{ stabinit: initDone }">
      <icon-loading class="stabLoad" :size="20" />
      <span class="stabGo">
        <icon-dashboard :class="{ stabShow: routeCheck && item.isHome }" :size="24" />
        <icon-compass :class="{ stabShow: routeCheck && !item.isHome }" :size="24" />
        <icon-pause-circle :class="{ stabShow: !routeCheck }" :size="20" />
      </span>
    </div>
    <div class="stabp">
      {{ $t(item.title) }}
    </div>
    <div class="stabx" v-if="!item.isHome">
      <a-button size="mini" type="text" shape="circle" @click.stop="removeTab()">
        <icon-close :size="18" />
      </a-button>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { PropType, ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import type { TagProps } from '@/store/tab'
import { tabStore } from '@/store'
import { computed } from 'vue'
const props = defineProps({
  item: {
    type: Object as PropType<TagProps>,
    default() {
      return {}
    }
  },
  index: {
    type: Number,
    default: 0
  }
})
const route = useRoute()
const routeCheck = computed(() => {
  return route.name === props.item.name
})
// 图标初始化结束
const initDone = ref(false)
// 元素初始化结束
const mounteDone = ref(false)
function init() {
  setTimeout(() => {
    mounteDone.value = true
  }, 100)
  setTimeout(() => {
    initDone.value = true
  }, 1200)
}
const router = useRouter()
const tab = tabStore()
const tagList = computed(() => {
  return tab.getTabList
})
// 移除标签
function removeTab() {
  mounteDone.value = false
  setTimeout(() => {
    tab.removeTabList(props.item.name, props.index)
    // 关闭的当前才进行跳转
    if (routeCheck.value) {
      router.push({ name: tagList.value[props.index - 1].name })
    }
  }, 350)
}
// 跳转
function go() {
  router.push({ name: props.item.name })
}
// 初次加载
onMounted(() => {
  init()
})
</script>
