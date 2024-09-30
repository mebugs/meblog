<template>
  <div class="container" :class="{ sPop: pop.pop }">
    <a-card class="sGroup sList" :bordered="false">
      <page :pop="pop" />
    </a-card>
    <a-card class="sGroup sItem" :bordered="false" :loading="false">
      <a-page-header :title="pop.header" :subtitle="pop.subHeader" @back="pop.close()" />
      <add v-if="pop.add" :pop="pop" />
      <edit v-if="pop.edit" :pop="pop" />
    </a-card>
  </div>
</template>

<script lang="ts" setup>
import { onUnmounted } from 'vue'
import usePop from '@/utils/hooks/pop'
const pop = usePop()
onUnmounted(() => {
  pop.close()
})
</script>
<script lang="ts">
export default {
  name: 'BlogComment'
}
</script>
