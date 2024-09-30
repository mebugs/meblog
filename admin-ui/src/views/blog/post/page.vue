<template>
  <a-form size="large" :model="query" label-align="left" layout="vertical">
    <a-row :gutter="20">
      <a-col :span="8">
        <a-form-item field="title" :label="$t('post.title')">
          <a-input v-model="query.title" allow-clear show-word-limit :placeholder="$t('post.title.sc')" />
        </a-form-item>
      </a-col>
      <a-col :span="8">
        <a-form-item field="status" :label="$t('post.status')">
          <a-select v-model="query.status" :options="dictList.postStatus" allow-clear allow-search :placeholder="$t('button.all')" />
        </a-form-item>
      </a-col>
      <a-col :span="8">
        <a-form-item field="group" :label="$t('post.group')">
          <a-select v-model="query.group" :options="dictList.postGroupList" allow-clear allow-search :placeholder="$t('button.all')" />
        </a-form-item>
      </a-col>
      <a-col :span="12">
        <a-space>
          <a-tooltip :content="$t('button.add')" :mini="true">
            <a-button
              v-permission="''"
              size="large"
              type="primary"
              status="danger"
              @click="pop.open('add', 0, $t('post.add'), $t('post.add.sub'), {}, search)">
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
    <template #img="{ record }">{{ record.img }}</template>
    <template #status="{ record }">
      <a-tag :color="typeTag[record.status]">{{ dictMap.postStatus[record.status] }}</a-tag>
    </template>
    <template #operations="{ record }">
      <a-space>
        <a-tooltip :content="$t('button.edit')" :mini="true">
          <a-button
            v-permission="'PlatpostEdit'"
            type="text"
            :disabled="record.mark === '1'"
            @click="pop.open('edit', record.id, $t('post.edit'), record.code, {}, search)">
            <template #icon> <icon-edit /> </template>
          </a-button>
        </a-tooltip>
        <a-tooltip :content="$t('button.delete')" :mini="true">
          <a-button v-permission="'PlatpostDel'" type="text" :disabled="record.mark === '1'" @click="openDelete(record)">
            <template #icon> <icon-delete /> </template>
          </a-button>
        </a-tooltip>
      </a-space>
    </template>
  </a-table>
  <!-- 刪除确认-->
  <a-modal v-model:visible="delItem.delConfirm" :width="400" :title="$t('title.delete')" @before-ok="deleting">
    <div>{{ $t('post.del.tips') }}</div>
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
import { postPage, postDel } from '@/api/blog/post'
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
  return { title: '', status: '', group: '' }
}
// 查询对象
const query = ref(initQuery())
// 状态标签
const typeTag: any = { '0': 'green', '1': 'orange' }
// 表格表头和数据指定
const columns = computed(() => [
  { title: t('post.img'), dataIndex: 'img' },
  { title: t('post.title'), dataIndex: 'title' },
  { title: t('post.group'), dataIndex: 'groupName' },
  { title: t('post.status'), dataIndex: 'status', slotName: 'status' },
  { title: t('base.oper'), slotName: 'operations', width: 180 }
])
// 列表对象
const list = ref([])
// 分页检索
async function pageQuery() {
  if (load.value) return
  setLoad(true)
  try {
    const res = await postPage({ ...query.value, ...page })
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
const dictList = ref({ postStatus: [], postGroupList: [] })
const dictMap = ref({ postStatus: {} as any })
// 字段初始化
async function dictInit() {
  // 指定字典Key
  await dictRead({ groupKeys: ['postStatus'] }).then((r) => {
    dictList.value.postStatus = r.data.list.postStatus
    dictMap.value.postStatus = r.data.map.postStatus

    props.pop.dictMap = dictMap
  })
  // await deptTree({}).then((dr) => {
  //   dictList.value.deptTree = dr.data
  // })
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
    await postDel(delItem.delId)
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
