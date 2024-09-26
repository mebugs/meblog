<template>
  <!--表格，子页无需吸顶 -->
  <a-table
    :scrollbar="false"
    :sticky-header="false"
    :row-key="'id'"
    :columns="RouterSelectColumns"
    :data="list"
    :pagination="false"
    size="medium"
    style="margin-top: 10px">
    <template #serviceCode="{ record }"> {{ dictMap.serviceCode[record.serviceCode] }} </template>
    <template #operations="{ record }" v-if="remove">
      <a-space>
        <a-tooltip :content="$t('button.delete')" :mini="true">
          <a-button type="text" size="small" @click="remove(record.id)">
            <template #icon> <icon-delete /> </template>
          </a-button>
        </a-tooltip>
      </a-space>
    </template>
  </a-table>
</template>

<script lang="ts" setup>
import { computed } from 'vue'
import { useI18n } from 'vue-i18n'
const { t } = useI18n()
// 入参读取
const props = defineProps({
  list: {
    type: Object,
    required: true,
    default: () => {
      return []
    }
  },
  dictMap: {
    type: Object,
    required: true,
    default: () => {
      return {}
    }
  },
  remove: {
    type: Boolean,
    default: () => {
      return false
    }
  },
  removeFunc: {
    type: Function,
    default: (id: number) => {
      return id
    }
  }
})
// 移除数据
function remove(id: number) {
  props.removeFunc(id)
}

// 路由列表对象
const RouterSelectColumns = computed(() => {
  let c: any[] = [
    { title: t('plat.serviceCode'), dataIndex: 'serviceCode', slotName: 'serviceCode' },
    { title: t('router.name'), dataIndex: 'name' },
    { title: t('router.url'), dataIndex: 'url' }
  ]
  if (props.remove) {
    c.push({ title: t('base.oper'), slotName: 'operations', width: 80 })
  }
  return c
})
</script>
