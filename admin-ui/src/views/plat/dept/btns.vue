<template>
  <a-col :span="24" class="doBtn treeBtn">
    <a-space>
      <a-tooltip :content="$t('button.sync')" :mini="true">
        <a-button v-permission="'PlatDeptView'" size="large" @click="refresh()">
          <template #icon>
            <icon-sync />
          </template>
        </a-button>
      </a-tooltip>
      <!-- 创建子集，3层不可 -->
      <a-tooltip :content="$t('button.add.sub')" :mini="true">
        <a-button
          v-permission="'PlatDeptAdd'"
          size="large"
          type="primary"
          status="danger"
          @click="pop.open('add', data.id, $t('dept.add.sub'), data.name, data, pop.callBack)">
          <template #icon>
            <icon-plus />
          </template>
        </a-button>
      </a-tooltip>
      <!-- 编辑，根节点不可 -->
      <a-tooltip :content="$t('button.edit')" :mini="true">
        <a-button
          v-permission="'PlatDeptEdit'"
          size="large"
          type="primary"
          status="warning"
          @click="pop.open('edit', data.id, $t('dept.edit'), data.name, data, pop.callBack)">
          <template #icon>
            <icon-edit />
          </template>
        </a-button>
      </a-tooltip>
      <!-- 同级排序 -->
      <a-tooltip :content="$t('button.sort')" :mini="true">
        <a-button v-permission="'PlatDeptSort'" size="large" :disabled="data.id == 1" type="primary" @click="confirmSort()">
          <template #icon>
            <icon-sort-ascending />
          </template>
        </a-button>
      </a-tooltip>
      <!-- 转移 -->
      <a-tooltip :content="$t('dept.to')" :mini="true">
        <a-button
          v-permission="'PlatDeptMerge'"
          size="large"
          :disabled="data.id == 1"
          type="primary"
          status="success"
          @click="confirmMerge()">
          <template #icon>
            <icon-branch />
          </template>
        </a-button>
      </a-tooltip>
      <!-- 删除，根节点不可 -->
      <a-tooltip :content="$t('button.delete')" :mini="true">
        <a-button v-permission="'PlatDeptDel'" size="large" :disabled="data.id == 1" @click="confirmDelete()">
          <template #icon>
            <icon-delete />
          </template>
        </a-button>
      </a-tooltip>
    </a-space>
  </a-col>
  <!-- 刪除确认-->
  <a-modal v-model:visible="delItem.delConfirm" :title="$t('title.delete')" @before-ok="deptDelete">
    <div>{{ $t('dept.del.tips') }}</div>
  </a-modal>
  <!-- 排序确认-->
  <a-modal v-model:visible="sortItem.sortConfirm" :width="340" :title="$t('dept.sort')" @before-ok="deptSorting">
    <a-spin :loading="load" class="sortList">
      <a-table
        :columns="columns"
        :data="sortList"
        :draggable="{ type: 'handle', width: 40 }"
        :pagination="false"
        @change="sortChange" />
    </a-spin>
  </a-modal>
  <!-- 迁移确认-->
  <a-modal v-model:visible="mergeItem.mergeConfirm" :width="500" :title="$t('dept.to')" @before-ok="deptMerge">
    <a-spin :loading="load">
      <a-form size="large" label-align="left" class="form" :model="mergeItem" layout="vertical">
        <a-row :gutter="20">
          <a-col :span="24">
            <a-form-item field="toId" :label="$t('dept.toId')" :rules="[{ required: true, message: $t('rule.required') }]">
              <a-tree-select
                v-model="mergeItem.toId"
                :allow-search="true"
                :allow-clear="true"
                :data="mergeItem.deptTree"
                :filter-tree-node="filterDept"
                :placeholder="$t('rule.select')" />
            </a-form-item>
          </a-col>
          <a-col :span="24">
            <a-form-item field="toType" :label="$t('dept.toType')">
              <template #extra>
                <div>{{ $t('dept.toType.tips') }}</div>
              </template>
              <a-select v-model="mergeItem.toType" :options="pop.dictList.deptToType" :placeholder="$t('rule.select')" />
            </a-form-item>
          </a-col>
        </a-row>
      </a-form>
    </a-spin>
  </a-modal>
</template>

<script lang="ts" setup>
import { reactive, ref, computed } from 'vue'
import { deptDel, deptBro, deptSort, deptTo } from '@/api/plat/dept'
import { useI18n } from 'vue-i18n'
import type { Pop } from '@/utils/hooks/pop'
import { deptTree } from '@/api/plat/dept'
import useLoad from '@/utils/hooks/load'
// 加载中变量
const { load, setLoad } = useLoad(false)
const { t } = useI18n()
// 路由列表对象
const columns = computed(() => [{ title: t('dept.name'), dataIndex: 'name' }])
// 入参读取
const props = defineProps({
  pop: {
    type: Object,
    required: true,
    default: () => {
      return {} as Pop
    }
  },
  data: {
    type: Object,
    required: true,
    default: () => {
      return {} as Pop
    }
  }
})
// 检索部门
function filterDept(searchValue: string, nodeData: any) {
  return nodeData.title.toLowerCase().indexOf(searchValue.toLowerCase()) > -1
}
// 部门合并
const mergeItem = reactive({
  mergeConfirm: false,
  id: 0,
  toId: '',
  toType: '0',
  deptTree: []
})
// 获取部门树
async function confirmMerge() {
  mergeItem.mergeConfirm = true
  mergeItem.id = props.data.id
  mergeItem.toId = ''
  mergeItem.toType = '0'
  mergeItem.deptTree = []
  setLoad(true)
  try {
    const res = await deptTree({})
    mergeItem.deptTree = res.data
  } catch (err) {
    // DoNothing CommonPopUp
  } finally {
    setLoad(false)
  }
}
// 部门合并
async function deptMerge() {
  try {
    await deptTo({
      id: mergeItem.id,
      toId: Number(mergeItem.toId),
      toType: mergeItem.toType
    })
    // 刷新树
    props.pop.treeRefresh(false)
    // 选中目标部门

    return true
  } catch (err) {
    return false
  } finally {
    // Nothing
  }
}
// 删除对象
const delItem = reactive({
  delConfirm: false,
  delId: 0
})
// 刷新
function refresh() {
  props.pop.callBack()
  props.pop.treeRefresh(false)
}
// 确认删除
function confirmDelete() {
  delItem.delId = props.data.id
  delItem.delConfirm = true
}
// 权限删除
async function deptDelete() {
  try {
    await deptDel(delItem.delId)
    // 刷新树
    props.pop.treeRefresh(false)
    return true
  } catch (err) {
    return false
  } finally {
    // Nothing
  }
}
// 排序对象
const sortList = ref([])
const sortItem = reactive({
  sortConfirm: false,
  sortId: 0
})
// 查询排序列表
async function getBroList() {
  setLoad(true)
  try {
    const res = await deptBro(sortItem.sortId)
    sortList.value = res.data
  } catch (err) {
    // DoNothing CommonPopUp
  } finally {
    setLoad(false)
  }
}
// 确认排序
function confirmSort() {
  sortItem.sortId = props.data.id
  sortItem.sortConfirm = true
  // 查询排序列表
  getBroList()
}
// 排序更新
function sortChange(data: any) {
  sortList.value = data
}
// 提交排序
async function deptSorting() {
  setLoad(true)
  try {
    const sortObj: any[] = []
    sortList.value.forEach((item: any, i) => {
      sortObj.push({
        id: item.id,
        sort: i
      })
    })
    const res = await deptSort(sortObj)
    if (res.data) {
      // 刷新树
      props.pop.callBack()
      props.pop.treeRefresh(false)
    }
  } catch (err) {
    // DoNothing CommonPopUp
  } finally {
    setLoad(false)
  }
}
</script>
