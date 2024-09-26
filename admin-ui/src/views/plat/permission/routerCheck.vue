<template>
  <a-form size="medium" :model="query" label-align="left" layout="vertical">
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
      <a-col :span="24" class="doBtn">
        <a-space>
          <a-tooltip :content="$t('button.search')" :mini="true">
            <a-button size="medium" type="primary" @click="search">
              <template #icon>
                <icon-search />
              </template>
            </a-button>
          </a-tooltip>
          <a-tooltip :content="$t('button.reset')" :mini="true">
            <a-button size="medium" @click="resetQuery">
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
    size="medium"
    v-model:selectedKeys="selectedKeys"
    :scrollbar="false"
    :sticky-header="true"
    :row-key="'id'"
    :loading="load"
    :pagination="page"
    :columns="columns"
    :data="list"
    :row-selection="{ type: 'checkbox' }"
    style="height: 410px"
    @page-change="changePage"
    @select="selectKeys">
    <template #serviceCode="{ record }">
      {{ dictMap.serviceCode[record.serviceCode] }}
    </template>
  </a-table>
</template>

<script lang="ts" setup>
import { ref, onMounted, watch, computed } from 'vue'
import { useI18n } from 'vue-i18n'
import useLocale from '@/utils/hooks/locale'
import useLoad from '@/utils/hooks/load'
import usePage from '@/utils/hooks/page'
import { dictRead } from '@/api/plat/dict'
import { routerPage } from '@/api/plat/router'
// 入参读取
const props = defineProps({
  doSelect: {
    type: Function,
    required: true
  },
  data: {
    type: Object,
    required: true,
    default: () => {
      return {}
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
  return { name: '', url: '', serviceCode: '', type: '0' }
}
// 查询对象
const query = ref(initQuery())
// 表格表头和数据指定
const columns = computed(() => [
  { title: t('plat.serviceCode'), dataIndex: 'serviceCode', slotName: 'serviceCode', width: 150 },
  { title: t('router.name'), dataIndex: 'name', width: 150 },
  { title: t('router.type'), dataIndex: 'type', slotName: 'type', width: 150 },
  { title: t('router.url'), dataIndex: 'url', ellipsis: true, tooltip: true }
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
const dictList = ref({ serviceCode: [], routerType: [] })
const dictMap = ref({ serviceCode: {} as any, routerType: {} as any })
// 字段初始化
async function dictInit() {
  // 指定字典Key
  await dictRead({ groupKeys: ['serviceCode', 'routerType'] }).then((r) => {
    dictList.value = r.data.list
    dictMap.value = r.data.map
  })
}
// 选择对象
const selectedKeys = ref<any>([])
const selectedItems = ref<any>([])
// 去掉响应式
const { data } = props
// 初始化赋值
selectedItems.value = data.selectedItems
selectedKeys.value = data.selectedKeys
// 选择事件
function selectKeys(keys: (string | number)[], thisKey: string | number, item: any) {
  if (keys.includes(thisKey)) {
    selectedItems.value.push(item)
  } else {
    selectedItems.value = selectedItems.value.filter((i: any) => {
      return i.id !== thisKey
    })
  }
  props.doSelect(selectedItems.value, keys)
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
