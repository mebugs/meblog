<template>
  <div id="vditor" style="margin-top: 60px"></div>
  <button @click="get">下一步</button>
  <button @click="send">提交</button>
</template>

<script setup lang="ts">
import '~/vditor/src/assets/less/index.less'
import { onMounted, ref } from 'vue'
import vditor from '@/utils/hooks/vditor'

// 0 编辑 1 正文预览 2 确认发布
const step = ref(0)
const { vd, md, getNew, toPreview, getResponse } = vditor(step, 'vditor')
function get() {
  toPreview()
}

function send() {
  getResponse()
  console.log(md.md)
  console.log(md.html)
  console.log(md.outline)
}

// 界面渲染后
onMounted(() => {
  getNew()
  console.log(vd)
})
</script>
<style scoped></style>
