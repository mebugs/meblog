<template>
  <a-form size="large" :model="query" label-align="left" layout="vertical">
    <a-row :gutter="20">
      <a-col :span="8">
        <a-form-item field="name" :label="$t('role.name')">
          <a-input v-model="query.name" :max-length="16" allow-clear show-word-limit :placeholder="$t('role.name.sc')" />
        </a-form-item>
      </a-col>
      <a-col :span="16"> </a-col>
      <a-col :span="12">
        <a-space>
          <a-tooltip :content="$t('button.add')" :mini="true">
            <a-button
              v-permission="'PlatRoleAdd'"
              size="large"
              type="primary"
              status="danger"
              @click="pop.open('add', 0, $t('role.add'), $t('role.add.sub'), {}, search)">
              <template #icon>
                <icon-plus />
              </template>
            </a-button>
          </a-tooltip>
          <!-- <a-divider direction="vertical" /> -->
        </a-space>
      </a-col>
      <a-col :span="12" class="doBtn">
        <a-space>
          <a-tooltip :content="$t('button.search')" :mini="true">
            <a-button v-permission="'PlatRoleView'" size="large" type="primary" @click="search">
              <template #icon>
                <icon-search />
              </template>
            </a-button>
          </a-tooltip>
          <a-tooltip :content="$t('button.reset')" :mini="true">
            <a-button v-permission="'PlatRoleView'" size="large" @click="resetQuery">
              <template #icon>
                <icon-refresh />
              </template>
            </a-button>
          </a-tooltip>
        </a-space>
      </a-col>
    </a-row>
  </a-form>
  <a-divider />
  <!--表格，吸顶和滚动条不可同时使用 -->
  <a-table
    :scrollbar="false"
    :sticky-header="true"
    :row-key="'id'"
    :loading="load"
    :pagination="page"
    :columns="columns"
    :data="list"
    @page-change="changePage">
    <template #operations="{ record }">
      <a-space>
        <a-tooltip :content="$t('button.get')" :mini="true">
          <a-button
            v-permission="'PlatRoleView'"
            type="text"
            @click="pop.open('get', record.id, $t('role.get'), record.code, {}, search)">
            <template #icon> <icon-eye /> </template>
          </a-button>
        </a-tooltip>
        <a-tooltip :content="$t('button.edit')" :mini="true">
          <a-button
            v-permission="'PlatRoleEdit'"
            type="text"
            @click="pop.open('edit', record.id, $t('role.edit'), record.code, {}, search)">
            <template #icon> <icon-edit /> </template>
          </a-button>
        </a-tooltip>
        <a-tooltip :content="$t('button.delete')" :mini="true">
          <a-button v-permission="'PlatRoleDel'" type="text" :disabled="record.mark === '1'" @click="openDelete(record)">
            <template #icon> <icon-delete /> </template>
          </a-button>
        </a-tooltip>
      </a-space>
    </template>
  </a-table>
  <!-- 刪除确认-->
  <a-modal v-model:visible="delItem.delConfirm" :width="400" :title="$t('title.delete')" @before-ok="deleting">
    <div>{{ $t('role.del.tips') }}</div>
  </a-modal>
</template>

<script lang="ts" setup>
import { ref, reactive, onMounted, watch, computed } from 'vue'
import { useI18n } from 'vue-i18n'
import type { Pop } from '@/utils/hooks/pop'
import useLocale from '@/utils/hooks/locale'
import useLoad from '@/utils/hooks/load'
import usePage from '@/utils/hooks/page'
import { rolePage, roleDel } from '@/api/plat/role'
// 入参读取
defineProps({
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
// 当前语言
const { currentLocale } = useLocale()
const { t } = useI18n()
// 分页
const { page, setQuery, search, changePage, resetPage } = usePage()
// 初始化查询对象
const initQuery = () => {
  return { name: '' }
}
// 查询对象
const query = ref(initQuery())
// 表格表头和数据指定
const columns = computed(() => [
  { title: t('role.name'), dataIndex: 'name', ellipsis: true, tooltip: true },
  { title: t('role.remark'), dataIndex: 'remark', ellipsis: true, tooltip: true },
  { title: t('base.oper'), slotName: 'operations', width: 140 }
])
// 列表对象
const list = ref([])
// 分页检索
async function pageQuery() {
  if (load.value) return
  setLoad(true)
  try {
    const res = await rolePage({ ...query.value, ...page })
    list.value = res.data.list
    page.total = res.data.total
  } catch (e) {
    // 清空数据
    list.value = []
    page.current = 1
    page.total = 0
  } finally {
    setLoad(false)
  }
}
// 初始化分页
setQuery(pageQuery)
function init() {
  // 初始化搜索
  pageQuery()
}
// 重置查询
function resetQuery() {
  query.value = initQuery()
  list.value = []
  // 重置分页
  resetPage()
  init()
}
// 删除对象
const delItem = reactive({
  delConfirm: false,
  delId: 0
})
// 打开删除
function openDelete(item: any) {
  delItem.delId = item.id
  delItem.delConfirm = true
}
// 确认删除
async function deleting() {
  try {
    await roleDel(delItem.delId)
  } catch (err) {
    return false
  } finally {
    // Nothing
    pageQuery()
  }
}
// 页面渲染
onMounted(() => {
  init()
})
// 语言监听
watch(currentLocale, (n, o) => {
  if (n !== o) {
    // Nothing
  }
})
</script>
