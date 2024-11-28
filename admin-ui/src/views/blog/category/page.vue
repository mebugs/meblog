<template>
  <a-form size="large" :model="query" label-align="left" layout="vertical">
    <a-row :gutter="20">
      <a-col :span="8">
        <a-form-item field="title" :label="$t('category.title')">
          <a-input v-model="query.title" :max-length="16" allow-clear show-word-limit :placeholder="$t('category.title.sc')" />
        </a-form-item>
      </a-col>
      <a-col :span="8">
        <a-form-item field="url" :label="$t('category.url')">
          <a-input v-model="query.url" :max-length="16" allow-clear show-word-limit :placeholder="$t('category.url.sc')" />
        </a-form-item>
      </a-col>
      <a-col :span="8"> </a-col>
      <a-col :span="12">
        <a-space>
          <a-tooltip :content="$t('button.add')" :mini="true">
            <a-button
              v-permission="''"
              size="large"
              type="primary"
              status="danger"
              @click="pop.open('add', 0, $t('category.add'), $t('category.add.sub'), {}, search)">
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
            <a-button v-permission="''" size="large" type="primary" @click="search">
              <template #icon>
                <icon-search />
              </template>
            </a-button>
          </a-tooltip>
          <a-tooltip :content="$t('button.reset')" :mini="true">
            <a-button v-permission="''" size="large" @click="resetQuery">
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
    <template #sourceShow="{ record }"> <img class="smallIcon" :src="record.sourceShow" /> </template>
    <template #operations="{ record }">
      <a-space>
        <a-tooltip :content="$t('button.get')" :mini="true">
          <a-button v-permission="''" type="text" @click="pop.open('get', record.id, $t('category.get'), record.title, {}, search)">
            <template #icon> <icon-eye /> </template>
          </a-button>
        </a-tooltip>
        <a-tooltip :content="$t('button.edit')" :mini="true">
          <a-button v-permission="''" type="text" @click="pop.open('edit', record.id, $t('category.edit'), record.title, {}, search)">
            <template #icon> <icon-edit /> </template>
          </a-button>
        </a-tooltip>
        <a-tooltip :content="$t('category.merge')" :mini="true">
          <a-button v-permission="''" type="text" @click="openMerge(record)">
            <template #icon> <icon-branch /> </template>
          </a-button>
        </a-tooltip>
        <a-tooltip :content="$t('button.delete')" :mini="true">
          <a-button v-permission="''" type="text" @click="openDelete(record)">
            <template #icon> <icon-delete /> </template>
          </a-button>
        </a-tooltip>
      </a-space>
    </template>
  </a-table>
  <!-- 刪除确认-->
  <a-modal v-model:visible="delItem.delConfirm" :width="400" :title="$t('title.delete')" @before-ok="deleting">
    <div>{{ $t('category.del.tips') }}</div>
  </a-modal>
  <!-- 刪除合并-->
  <a-modal v-model:visible="mergeItem.mergeConfirm" :width="400" :title="$t('category.merge')" @before-ok="mergeIn">
    <a-form size="large" label-align="left" class="form" :model="mergeItem" layout="vertical">
      <a-col :span="24">
        <a-form-item field="toId" :label="$t('category.toId')">
          <template #extra>
            <div>{{ $t('category.toId.tips') }}</div>
          </template>
          <a-select v-model="mergeItem.toId" :options="cateList" :placeholder="$t('rule.select')" />
        </a-form-item>
      </a-col>
    </a-form>
  </a-modal>
</template>

<script lang="ts" setup>
import { ref, reactive, onMounted, watch, computed } from 'vue'
import { useI18n } from 'vue-i18n'
import type { Pop } from '@/utils/hooks/pop'
import useLocale from '@/utils/hooks/locale'
import useLoad from '@/utils/hooks/load'
import usePage from '@/utils/hooks/page'
import { categoryPage, categoryDel, categoryMerge, categoryList } from '@/api/blog/category'
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
// 当前语言
const { currentLocale } = useLocale()
const { t } = useI18n()
// 分页
const { page, setQuery, search, changePage, resetPage } = usePage()
// 初始化查询对象
const initQuery = () => {
  return { title: '', url: '' }
}
// 查询对象
const query = ref(initQuery())
// 表格表头和数据指定
const columns = computed(() => [
  { title: t('category.sourceId'), dataIndex: 'sourceShow', slotName: 'sourceShow' },
  { title: t('category.title'), dataIndex: 'title' },
  { title: t('category.url'), dataIndex: 'url' },
  { title: t('category.summary'), width: 400, dataIndex: 'summary', ellipsis: true, tooltip: true },
  { title: t('base.oper'), slotName: 'operations', width: 190 }
])
// 列表对象
const list = ref([])
// 分页检索
async function pageQuery() {
  if (load.value) return
  setLoad(true)
  try {
    const res = await categoryPage({ ...query.value, ...page })
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
// 初始化字典对象
const dictList = ref({})
const dictMap = ref({})
// 字段初始化
async function dictInit() {
  props.pop.dictMap = dictMap
  props.pop.dictList = dictList
}
function init() {
  // 初始化后端字典对象
  dictInit()
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
// 合并对象
const mergeItem = reactive({
  mergeConfirm: false,
  id: 0,
  toId: ''
})
// 列表对象
const cateList = ref([])
// 打开合并
async function openMerge(item: any) {
  cateList.value = []
  mergeItem.id = item.id
  mergeItem.mergeConfirm = true
  mergeItem.toId = ''
  await categoryList().then((rr) => {
    cateList.value = rr.data
  })
}
// 确认合并
async function mergeIn() {
  try {
    await categoryMerge({
      id: mergeItem.id,
      toId: Number(mergeItem.toId)
    })
  } catch (err) {
    return false
  } finally {
    // Nothing
    pageQuery()
  }
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
    await categoryDel(delItem.delId)
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
    dictInit()
  }
})
</script>
