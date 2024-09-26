<template>
  <a-col :span="24" class="doBtn treeBtn">
    <a-space>
      <a-tooltip :content="$t('button.sync')" :mini="true">
        <a-button v-permission="'PlatPermissionView'" size="large" @click="refresh()">
          <template #icon>
            <icon-sync />
          </template>
        </a-button>
      </a-tooltip>
      <!-- 创建子集，3层不可 -->
      <a-tooltip :content="$t('button.add.sub')" :mini="true">
        <a-button
          v-permission="'PlatPermissionAdd'"
          size="large"
          :disabled="!data.level || data.level == '3'"
          type="primary"
          status="danger"
          @click="pop.open('add', data.id, $t('permission.add.sub'), data.name, data, pop.callBack)">
          <template #icon>
            <icon-plus />
          </template>
        </a-button>
      </a-tooltip>
      <!-- 编辑，根节点不可 -->
      <a-tooltip :content="$t('button.edit')" :mini="true">
        <a-button
          v-permission="'PlatPermissionEdit'"
          size="large"
          :disabled="!data.level || data.id == 1"
          type="primary"
          status="warning"
          @click="pop.open('edit', data.id, $t('permission.edit'), data.name, data, pop.callBack)">
          <template #icon>
            <icon-edit />
          </template>
        </a-button>
      </a-tooltip>
      <!-- 同级排序 -->
      <a-tooltip :content="$t('button.sort')" :mini="true">
        <a-button
          v-permission="'PlatPermissionSort'"
          size="large"
          :disabled="!data.level || data.id == 1"
          type="primary"
          @click="confirmSort()">
          <template #icon>
            <icon-sort-ascending />
          </template>
        </a-button>
      </a-tooltip>
      <!-- 删除，根节点不可 -->
      <a-tooltip :content="$t('button.delete')" :mini="true">
        <a-button v-permission="'PlatPermissionDel'" size="large" :disabled="!data.level || data.id == 1" @click="confirmDelete()">
          <template #icon>
            <icon-delete />
          </template>
        </a-button>
      </a-tooltip>
    </a-space>
  </a-col>
  <!-- 刪除确认-->
  <a-modal v-model:visible="delItem.delConfirm" :title="$t('title.delete')" @before-ok="permissionDelete">
    <div>{{ $t('permission.del.tips') }}</div>
  </a-modal>
  <!-- 排序确认-->
  <a-modal v-model:visible="sortItem.sortConfirm" :width="340" :title="$t('permission.sort')" @before-ok="permissionSorting">
    <a-spin :loading="load" class="sortList">
      <a-table
        :columns="columns"
        :data="sortList"
        :draggable="{ type: 'handle', width: 40 }"
        :pagination="false"
        @change="sortChange" />
    </a-spin>
  </a-modal>
</template>

<script lang="ts" setup>
import { reactive, ref, computed } from 'vue'
import { permissionDel, permissionBro, permissionSort } from '@/api/plat/permission'
import { useI18n } from 'vue-i18n'
import type { Pop } from '@/utils/hooks/pop'
import useLoad from '@/utils/hooks/load'
// 加载中变量
const { load, setLoad } = useLoad(false)
const { t } = useI18n()
// 路由列表对象
const columns = computed(() => [{ title: t('permission.name'), dataIndex: 'name' }])
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
async function permissionDelete() {
  try {
    await permissionDel(delItem.delId)
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
    const res = await permissionBro(sortItem.sortId)
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
async function permissionSorting() {
  setLoad(true)
  try {
    const sortObj: any[] = []
    sortList.value.forEach((item: any, i) => {
      sortObj.push({
        id: item.id,
        sort: i
      })
    })
    const res = await permissionSort(sortObj)
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
