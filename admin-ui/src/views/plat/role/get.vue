<template>
  <a-spin :loading="load">
    <a-form size="large" label-align="left" class="form" layout="vertical" :model="formData">
      <a-row :gutter="20">
        <a-col :span="12">
          <a-row :gutter="20">
            <a-col :span="24">
              <a-form-item field="name" :label="$t('role.name')" :rules="[{ required: true, message: $t('rule.required') }]">
                <template #extra>
                  <div>{{ $t('role.name.tips') }}</div>
                </template>
                <span class="formSpan">{{ formData.name }}</span>
              </a-form-item>
            </a-col>
            <a-col :span="24">
              <a-form-item field="permissionIds" :label="$t('role.permissionIds')">
                <template #extra>
                  <div>{{ $t('role.permissionIds.tips') }}</div>
                </template>
                <span class="formSpan">{{ $t('role.permissionIds.place.forGet') }}</span>
              </a-form-item>
            </a-col>
            <a-col :span="24">
              <a-form-item field="remark" :label="$t('role.remark')">
                <span class="formSpan">{{ formData.remark }}</span>
              </a-form-item>
            </a-col>
          </a-row>
        </a-col>
        <a-col :span="12">
          <a-row :gutter="20">
            <a-col :span="24">
              <a-spin :loading="treeLoad" class="treeLe">
                <!-- disableCheckbox: 'id' 保证所有节点都被Disable-->
                <a-tree
                  v-if="list[0]"
                  ref="myTree"
                  v-model:checked-keys="formData.permissionIds"
                  v-model:half-checked-keys="formData.halfPermissionIds"
                  size="large"
                  :field-names="{ key: 'id' }"
                  :data="list"
                  :show-line="true" />
              </a-spin>
            </a-col>
          </a-row>
        </a-col>
      </a-row>
    </a-form>
  </a-spin>
</template>

<script lang="ts" setup>
import { onMounted, ref } from 'vue'
import type { Pop } from '@/utils/hooks/pop'
import { treeNode } from '@/api/comm'
import useLoad from '@/utils/hooks/load'
import { roleInit, roleGet } from '@/api/plat/role'
import { permissionTree } from '@/api/plat/permission'
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
const { load: treeLoad, setLoad: setTreeLoad } = useLoad(false)
// 表单数据初始化
const formData = roleInit()
async function get() {
  setLoad(true)
  try {
    const res = await roleGet(props.pop.itemId)
    formData.value = res.data
    return
  } catch (err) {
    // DoNothing CommonPopUp
  } finally {
    setLoad(false)
  }
}
// 树处理
const list = ref<treeNode[]>([])
// 执行树搜索
function treeFilter() {
  // hasIds
  let hasIds = [...formData.value.permissionIds, ...formData.value.halfPermissionIds]
  // 递归检索
  const loop = (data: treeNode[]) => {
    const result: treeNode[] = []
    data.forEach((item: treeNode) => {
      // 自身满足才处理子集
      if (hasIds.includes(item.id) || item.expand) {
        if (item.children) {
          const filterData = loop(item.children)
          result.push({ ...item, children: filterData })
        } else {
          result.push({ ...item })
        }
      }
    })
    return result
  }
  list.value = loop(list.value)
}
// 权限树查询
async function treeQuery() {
  setTreeLoad(true)
  try {
    // 路由分页
    const res = await permissionTree({})
    list.value = res.data
    treeFilter()
  } catch (e) {
    list.value = []
  } finally {
    setTreeLoad(false)
  }
}
// 页面渲染
onMounted(() => {
  // Nothing
  get()
  // 初始化树查询
  treeQuery()
})
</script>
