<template>
  <a-input-search
    v-model="treeSearch"
    :placeholder="$t('dept.name.sc')"
    size="large"
    allow-clear
    style="margin: 4px 0 10px 0"
    @press-enter="treeFilter"
    @search="treeFilter"
    @clear="treeFilter" />
  <a-spin :loading="load" class="treeLe">
    <a-tree
      v-if="list[0]"
      ref="myTree"
      size="large"
      v-model:selected-keys="treeSelect"
      :data="list"
      :show-line="true"
      @select="onSelect" />
  </a-spin>
</template>

<script lang="ts" setup>
import { onMounted, ref, nextTick } from 'vue'
import useLoad from '@/utils/hooks/load'
import { treeNode } from '@/api/comm'
import type { Pop } from '@/utils/hooks/pop'
import { deptTree } from '@/api/plat/dept'
// 树Dom
const myTree = ref<any>()
// 入参读取
const props = defineProps({
  pop: {
    type: Object,
    required: true,
    default: () => {
      return {} as Pop
    }
  }
})
// 加载中变量
const { load, setLoad } = useLoad(false)
// 树检索
const treeSearch = ref('')
// 列表对象
const origList = ref<treeNode[]>([])
const list = ref<treeNode[]>([])
// 执行树搜索
function treeFilter() {
  // 无检索
  if (treeSearch.value === '') {
    list.value = origList.value
  }
  // 递归检索
  const loop = (data: treeNode[]) => {
    const result: treeNode[] = []
    data.forEach((item: treeNode) => {
      if (item.title.toLowerCase().indexOf(treeSearch.value.toLowerCase()) > -1) {
        result.push({ ...item })
      } else if (item.children) {
        const filterData = loop(item.children)
        if (filterData.length) {
          result.push({
            ...item,
            children: filterData
          })
        }
      }
    })
    return result
  }
  list.value = loop(origList.value)
}
// 树选择对象 默认选择ROOT
const treeSelect = ref<any>(['1'])
// 刷新选择
async function refreshSelect(keyObj: any) {
  props.pop.itemId = keyObj.id
}
// 树选择
function onSelect(keys: any[], event: any) {
  treeSelect.value = keys
  const { node } = event
  refreshSelect(node)
}
// 如果选择了不存在值，则重定向选择
function syncSelect() {
  nextTick(() => {
    if (myTree.value && myTree.value.getSelectedNodes().length == 0) {
      getRootNode()
    }
  })
}
// 权限树查询
async function treeQuery(init: boolean) {
  setLoad(true)
  try {
    // 路由分页
    const res = await deptTree({})
    origList.value = res.data
    list.value = res.data
    // 树过滤保留
    treeFilter()
    if (init) {
      // 初始化select
      getRootNode()
    }
    syncSelect()
  } catch (e) {
    // 清空数据
    origList.value = []
    list.value = []
  } finally {
    setLoad(false)
  }
}
// 取根节点
function getRootNode() {
  if (origList.value.length > 0) {
    treeSelect.value = ['1']
    const root = origList.value[0]
    return refreshSelect({ id: root.id, title: root.title })
  }
}
// 绑定树刷新函数
props.pop.treeRefresh = treeQuery
function init() {
  // 初始化树查询
  treeQuery(true)
}
// 页面渲染
onMounted(() => {
  init()
})
</script>
