<template>
  <div class="container treeE" :class="{ sPop: pop.pop }">
    <div class="treeL">
      <div class="sGroup sBg" :bordered="false">
        <tree :pop="pop" />
      </div>
    </div>
    <div class="treeR">
      <a-card class="sGroup sList" :bordered="false">
        <get :pop="pop" />
      </a-card>
      <a-card class="sGroup sItem" :bordered="false" :loading="false">
        <a-page-header :title="pop.header" :subtitle="pop.subHeader" @back="pop.treeClose()" />
        <add v-if="pop.add" :pop="pop" />
        <edit v-if="pop.edit" :pop="pop" />
      </a-card>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { onUnmounted } from 'vue'
import usePop from '@/utils/hooks/pop'
import Tree from './tree.vue'
import Get from './get.vue'
import Add from './add.vue'
import Edit from './edit.vue'

// 树选择节点
const pop = usePop()
pop.itemId = 1 // 默认选择ROOT节点
onUnmounted(() => {
  pop.close()
})
</script>
<script lang="ts">
export default {
  name: 'PlatDept'
}
</script>
