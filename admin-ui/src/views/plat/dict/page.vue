<template>
  <a-form size="large" :model="query" label-align="left" layout="vertical">
    <a-row :gutter="20">
      <a-col :span="24">
        <a-row :gutter="20">
          <a-col :span="8">
            <a-form-item field="groupKey" :label="$t('dict.groupKey')">
              <a-select
                v-model="query.groupKey"
                :options="dictList.dictGroup"
                allow-clear
                allow-search
                :placeholder="$t('button.all')" />
            </a-form-item>
          </a-col>
          <a-col :span="8">
            <a-form-item field="choose" :label="$t('dict.choose')">
              <a-select
                v-model="query.choose"
                :options="dictList.openStatus"
                allow-clear
                allow-search
                :placeholder="$t('button.all')" />
            </a-form-item>
          </a-col>
        </a-row>
      </a-col>
      <a-col :span="12">
        <a-space>
          <a-tooltip :content="$t('button.add')" :mini="true">
            <a-button
              v-permission="'PlatDictAdd'"
              size="large"
              type="primary"
              status="danger"
              @click="pop.open('add', 0, $t('dict.add'), $t('dict.add.sub'), {}, search)">
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
            <a-button v-permission="'PlatDictView'" size="large" type="primary" @click="search">
              <template #icon>
                <icon-search />
              </template>
            </a-button>
          </a-tooltip>
          <a-tooltip :content="$t('button.reset')" :mini="true">
            <a-button v-permission="'PlatDictView'" size="large" @click="resetQuery">
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
    <template #groupKey="{ record }">
      {{ dictMap.dictGroup[record.groupKey] }}
    </template>
    <template #choose="{ record }">
      <a-tag :color="chooseTag[record.choose]">{{ dictMap.openStatus[record.choose] }}</a-tag>
    </template>
    <template #operations="{ record }">
      <a-space>
        <a-tooltip :content="$t('button.get')" :mini="true">
          <a-button
            v-permission="'PlatDictView'"
            type="text"
            @click="pop.open('get', record.id, $t('dict.get'), record.groupKey, {}, search)">
            <template #icon> <icon-eye /> </template>
          </a-button>
        </a-tooltip>
        <a-tooltip :content="$t('button.edit')" :mini="true">
          <a-button
            v-permission="'PlatDictEdit'"
            type="text"
            @click="pop.open('edit', record.id, $t('dict.edit'), record.groupKey, {}, search)">
            <template #icon> <icon-edit /> </template>
          </a-button>
        </a-tooltip>
        <a-tooltip :content="$t('button.sort')" :mini="true">
          <a-button v-permission="'PlatDictSort'" type="text" @click="openSort(record.groupKey)">
            <template #icon> <icon-sort-ascending /> </template>
          </a-button>
        </a-tooltip>
        <a-tooltip :content="$t('button.lock')" :mini="true">
          <a-button v-permission="'PlatDictDel'" type="text" :disabled="record.mark === '1'" @click="openDelete(record)">
            <template #icon> <icon-lock /> </template>
          </a-button>
        </a-tooltip>
      </a-space>
    </template>
  </a-table>
  <!-- 排序确认-->
  <a-modal v-model:visible="sortPop.pop" :width="360" :title="sortPop.groupKey + $t('dict.sort')" @before-ok="sorting">
    <a-spin :loading="load" class="sortList">
      <a-table
        :columns="sortColumns"
        :data="sortList"
        :draggable="{ type: 'handle', width: 40 }"
        :pagination="false"
        :show-header="false"
        @change="sortChange" />
    </a-spin>
  </a-modal>
  <!-- 刪除确认-->
  <a-modal v-model:visible="delItem.delConfirm" :width="400" :title="$t('title.lock')" @before-ok="deleting">
    <div>{{ $t('dict.del.tip') }}</div>
  </a-modal>
</template>

<script lang="ts" setup>
import { ref, reactive, onMounted, watch, computed } from 'vue'
import { useI18n } from 'vue-i18n'
import type { Pop } from '@/utils/hooks/pop'
import useLocale from '@/utils/hooks/locale'
import useLoad from '@/utils/hooks/load'
import usePage from '@/utils/hooks/page'
import { dictPage, dictBro, dictSort, dictDel, dictRead } from '@/api/plat/dict'
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
  return { groupKey: '', choose: '' }
}
// 状态标签
const chooseTag: any = { '0': 'green', '1': 'orange', '2': 'red' }
// 查询对象
const query = ref(initQuery())
// 表格表头和数据指定
const columns = computed(() => [
  { title: t('dict.groupKey'), dataIndex: 'groupKey', slotName: 'groupKey' },
  { title: t('dict.val'), dataIndex: 'val' },
  { title: t('dict.label'), dataIndex: 'label', width: 150, ellipsis: true, tooltip: true },
  { title: t('dict.labelEn'), dataIndex: 'labelEn', width: 150, ellipsis: true, tooltip: true },
  { title: t('dict.choose'), dataIndex: 'choose', slotName: 'choose' },
  { title: t('dict.remark'), dataIndex: 'remark', ellipsis: true, tooltip: true },
  { title: t('base.oper'), slotName: 'operations', width: 180 }
])
// 列表对象
const list = ref([])
// 分页检索
async function pageQuery() {
  if (load.value) return
  setLoad(true)
  try {
    const res = await dictPage({ ...query.value, ...page })
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
const dictList = ref({ dictGroup: [], openStatus: [] })
const dictMap = ref({ dictGroup: {} as any, openStatus: {} as any })
// 临时对象处理

// 字段初始化
async function dictInit() {
  // 指定字典Key
  await dictRead({ groupKeys: ['dictGroup', 'openStatus'] }).then((r) => {
    dictList.value = r.data.list
    dictMap.value = r.data.map
    props.pop.dictList = dictList
    props.pop.dictMap = dictMap
  })
}
function init() {
  // 字典初始化
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
// 排序对象
const sortList = ref([])
const sortPop = reactive({
  pop: false,
  groupKey: ''
})
// 排序列表对象
const sortColumns = computed(() => [{ title: t('dict.label'), dataIndex: 'name', ellipsis: true, tooltip: true }])
// 打开Sort
function openSort(groupKey: string) {
  sortPop.pop = true
  sortPop.groupKey = groupKey
  getSortList()
}
// 查询同级
async function getSortList() {
  setLoad(true)
  try {
    const res = await dictBro(sortPop.groupKey)
    sortList.value = res.data
  } catch (err) {
    // DoNothing
  } finally {
    setLoad(false)
  }
}
// 排序提交
async function sorting() {
  setLoad(true)
  try {
    const sortObj: any[] = []
    sortList.value.forEach((item: any, i) => {
      sortObj.push({
        id: item.id,
        sort: i
      })
    })
    await dictSort(sortObj)
  } catch (err) {
    // DoNothing
    return false
  } finally {
    setLoad(false)
    pageQuery()
  }
}
// 排序更新
function sortChange(data: any) {
  sortList.value = data
}
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
    await dictDel(delItem.delId)
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
