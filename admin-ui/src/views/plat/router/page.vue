<template>
  <a-form size="large" :model="query" label-align="left" layout="vertical">
    <a-row :gutter="20">
      <a-col :span="8">
        <a-form-item field="name" :label="$t('router.name')">
          <a-input v-model="query.name" :max-length="64" allow-clear show-word-limit :placeholder="$t('router.name.sc')" />
        </a-form-item>
      </a-col>
      <a-col :span="8">
        <a-form-item field="url" :label="$t('router.url')">
          <a-input v-model="query.url" :max-length="64" allow-clear show-word-limit :placeholder="$t('router.url.sc')" />
        </a-form-item>
      </a-col>
      <a-col :span="8">
        <a-form-item field="serviceCode" :label="$t('plat.serviceCode')">
          <a-select
            v-model="query.serviceCode"
            :options="dictList.serviceCode"
            allow-clear
            allow-search
            :placeholder="$t('button.all')" />
        </a-form-item>
      </a-col>
      <a-col :span="12">
        <a-space>
          <a-tooltip :content="$t('button.add')" :mini="true">
            <a-button
              v-permission="'PlatRouterAdd'"
              size="large"
              type="primary"
              status="danger"
              @click="pop.open('add', 0, $t('router.add'), $t('router.add.sub'), {}, search)">
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
            <a-button v-permission="'PlatRouterView'" size="large" type="primary" @click="search">
              <template #icon>
                <icon-search />
              </template>
            </a-button>
          </a-tooltip>
          <a-tooltip :content="$t('button.reset')" :mini="true">
            <a-button v-permission="'PlatRouterView'" size="large" @click="resetQuery">
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
    <template #serviceCode="{ record }">
      {{ dictMap.serviceCode[record.serviceCode] }}
    </template>
    <template #type="{ record }">
      <a-tag :color="flagTag[record.type]">{{ dictMap.routerType[record.type] }}</a-tag>
    </template>
    <template #logInDb="{ record }">
      <a-tag :color="flagTag[record.logInDb]">{{ dictMap.openStatus[record.logInDb] }}</a-tag>
    </template>
    <template #reqLogPrint="{ record }">
      <a-tag :color="flagTag[record.reqLogPrint]">{{ dictMap.openStatus[record.reqLogPrint] }}</a-tag>
    </template>
    <template #resLogPrint="{ record }">
      <a-tag :color="flagTag[record.resLogPrint]">{{ dictMap.openStatus[record.resLogPrint] }}</a-tag>
    </template>
    <template #operations="{ record }">
      <a-space>
        <a-tooltip :content="$t('button.get')" :mini="true">
          <a-button
            v-permission="'PlatRouterView'"
            type="text"
            @click="pop.open('get', record.id, $t('router.get'), record.code, {}, search)">
            <template #icon> <icon-eye /> </template>
          </a-button>
        </a-tooltip>
        <a-tooltip :content="$t('button.edit')" :mini="true">
          <a-button
            v-permission="'PlatRouterEdit'"
            type="text"
            @click="pop.open('edit', record.id, $t('router.edit'), record.code, {}, search)">
            <template #icon> <icon-edit /> </template>
          </a-button>
        </a-tooltip>
        <a-tooltip :content="$t('button.delete')" :mini="true">
          <a-button v-permission="'PlatRouterDel'" type="text" :disabled="record.mark === '1'" @click="openDelete(record)">
            <template #icon> <icon-delete /> </template>
          </a-button>
        </a-tooltip>
      </a-space>
    </template>
  </a-table>
  <!-- 刪除确认-->
  <a-modal v-model:visible="delItem.delConfirm" :width="400" :title="$t('title.delete')" @before-ok="deleting">
    <div>{{ $t('router.del.tip') }}</div>
  </a-modal>
</template>

<script lang="ts" setup>
import { ref, reactive, onMounted, watch, computed } from 'vue'
import { useI18n } from 'vue-i18n'
import type { Pop } from '@/utils/hooks/pop'
import useLocale from '@/utils/hooks/locale'
import useLoad from '@/utils/hooks/load'
import usePage from '@/utils/hooks/page'
import { dictRead } from '@/api/plat/dict'
import { routerPage, routerDel } from '@/api/plat/router'
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
  return { name: '', url: '', serviceCode: '', type: '' }
}
// 查询对象
const query = ref(initQuery())
// 状态标签
const flagTag: any = { '0': 'green', '1': 'orange', '2': 'red' }
// 表格表头和数据指定
const columns = computed(() => [
  { title: t('plat.serviceCode'), dataIndex: 'serviceCode', slotName: 'serviceCode', width: 150 },
  { title: t('router.name'), dataIndex: 'name', width: 150 },
  { title: t('router.type'), dataIndex: 'type', slotName: 'type', width: 150 },
  { title: t('router.url'), dataIndex: 'url', ellipsis: true, tooltip: true },
  { title: t('router.logInDb'), dataIndex: 'logInDb', slotName: 'logInDb' },
  { title: t('router.reqLogPrint'), dataIndex: 'reqLogPrint', slotName: 'reqLogPrint' },
  { title: t('router.resLogPrint'), dataIndex: 'resLogPrint', slotName: 'resLogPrint' },
  { title: t('base.oper'), slotName: 'operations', width: 140 }
])
// 列表对象
const list = ref([])
// 分页检索
async function pageQuery() {
  if (load.value) return
  setLoad(true)
  try {
    const res = await routerPage({ ...query.value, ...page })
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
const dictList = ref({ serviceCode: [], routerType: [], openStatus: [] })
const dictMap = ref({ serviceCode: {} as any, routerType: {} as any, openStatus: {} as any })
// 字段初始化
async function dictInit() {
  // 指定字典Key
  await dictRead({ groupKeys: ['serviceCode', 'routerType', 'openStatus'] }).then((r) => {
    dictList.value = r.data.list
    dictMap.value = r.data.map
    props.pop.dictList = dictList
    props.pop.dictMap = dictMap
  })
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
    await routerDel(delItem.delId)
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
