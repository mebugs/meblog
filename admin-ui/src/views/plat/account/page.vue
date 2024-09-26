<template>
  <a-form size="large" :model="query" label-align="left" layout="vertical">
    <a-row :gutter="20">
      <a-col :span="8">
        <a-form-item field="code" :label="$t('account.account')">
          <a-input v-model="query.account" :max-length="16" allow-clear show-word-limit :placeholder="$t('account.account.sc')" />
        </a-form-item>
      </a-col>
      <a-col :span="8">
        <a-form-item field="code" :label="$t('account.name')">
          <a-input v-model="query.name" :max-length="16" allow-clear show-word-limit :placeholder="$t('account.name.sc')" />
        </a-form-item>
      </a-col>
      <a-col :span="8">
        <a-form-item field="type" :label="$t('account.deptId')">
          <a-tree-select
            v-model="query.deptId"
            :allow-search="true"
            :allow-clear="true"
            :data="pop.dictList.deptTree"
            :filter-tree-node="filterDept"
            :placeholder="$t('button.all')" />
        </a-form-item>
      </a-col>
      <a-col :span="12">
        <a-space>
          <a-tooltip :content="$t('button.add')" :mini="true">
            <a-button
              v-permission="'PlatAccountAdd'"
              size="large"
              type="primary"
              status="danger"
              @click="pop.open('add', 0, $t('account.add'), $t('account.add.sub'), {}, search)">
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
            <a-button v-permission="'PlatAccountView'" size="large" type="primary" @click="search">
              <template #icon>
                <icon-search />
              </template>
            </a-button>
          </a-tooltip>
          <a-tooltip :content="$t('button.reset')" :mini="true">
            <a-button v-permission="'PlatAccountView'" size="large" @click="resetQuery">
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
    <template #isLeader="{ record }">{{ dictMap.deptAccount[record.isLeader] }}</template>
    <template #permissionType="{ record }">{{ dictMap.accountPermission[record.permissionType] }}</template>
    <template #status="{ record }">
      <a-tag :color="typeTag[record.status]">{{ dictMap.accountStatus[record.status] }}</a-tag>
    </template>
    <template #roleNames="{ record }">{{ record.roleNames.join('/') }}</template>
    <template #operations="{ record }">
      <a-space>
        <a-tooltip :content="$t('button.get')" :mini="true">
          <a-button
            v-permission="'PlatAccountView'"
            type="text"
            @click="pop.open('get', record.id, $t('account.get'), record.code, {}, search)">
            <template #icon> <icon-eye /> </template>
          </a-button>
        </a-tooltip>
        <a-tooltip :content="$t('button.edit')" :mini="true">
          <a-button
            v-permission="'PlatAccountEdit'"
            type="text"
            :disabled="record.mark === '1'"
            @click="pop.open('edit', record.id, $t('account.edit'), record.code, {}, search)">
            <template #icon> <icon-edit /> </template>
          </a-button>
        </a-tooltip>
        <a-tooltip :content="$t('account.resetPwd')" :mini="true">
          <a-button v-permission="'PlatAccountReset'" type="text" :disabled="record.mark === '1'" @click="openReset(record)">
            <template #icon> <icon-sync /> </template>
          </a-button>
        </a-tooltip>
        <a-tooltip :content="$t('button.delete')" :mini="true">
          <a-button v-permission="'PlatAccountDel'" type="text" :disabled="record.mark === '1'" @click="openDelete(record)">
            <template #icon> <icon-delete /> </template>
          </a-button>
        </a-tooltip>
      </a-space>
    </template>
  </a-table>
  <!-- 刪除确认-->
  <a-modal v-model:visible="delItem.delConfirm" :width="400" :title="$t('title.delete')" @before-ok="deleting">
    <div>{{ $t('account.del.tips') }}</div>
  </a-modal>
  <!-- 重置确认-->
  <a-modal v-model:visible="resetItem.resetConfirm" :width="400" :title="$t('account.resetPwd')" @before-ok="reseting">
    <div>{{ $t('account.resetPwd.tips') }}</div>
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
import { accountPage, accountDel, accountReset } from '@/api/plat/account'
import { deptTree } from '@/api/plat/dept'
import { roleList } from '@/api/plat/role'
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
  return { account: '', name: '', deptId: '' }
}
// 查询对象
const query = ref(initQuery())
// 状态标签
const typeTag: any = { '0': 'green', '1': 'orange' }
// 表格表头和数据指定
const columns = computed(() => [
  { title: t('account.account'), dataIndex: 'account' },
  { title: t('account.name'), dataIndex: 'name' },
  { title: t('account.deptName'), dataIndex: 'deptName', ellipsis: true, tooltip: true },
  { title: t('account.isLeader'), dataIndex: 'isLeader', slotName: 'isLeader' },
  { title: t('account.permissionType'), dataIndex: 'permissionType', slotName: 'permissionType' },
  { title: t('account.status'), dataIndex: 'status', slotName: 'status' },
  { title: t('account.roleNames'), dataIndex: 'roleNames', slotName: 'roleNames' },
  { title: t('base.oper'), slotName: 'operations', width: 180 }
])
// 列表对象
const list = ref([])
// 分页检索
async function pageQuery() {
  if (load.value) return
  setLoad(true)
  try {
    const res = await accountPage({ ...query.value, ...page })
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
const dictList = ref({ accountStatus: [], deptAccount: [], accountPermission: [], deptTree: [], roles: [] })
const dictMap = ref({ accountStatus: {} as any, deptAccount: {} as any, accountPermission: {} as any })
// 字段初始化
async function dictInit() {
  // 指定字典Key
  await dictRead({ groupKeys: ['accountStatus', 'deptAccount', 'accountPermission'] }).then((r) => {
    dictList.value.accountStatus = r.data.list.accountStatus
    dictList.value.deptAccount = r.data.list.deptAccount
    dictList.value.accountPermission = r.data.list.accountPermission
    dictMap.value = r.data.map
    props.pop.dictMap = dictMap
  })
  await deptTree({}).then((dr) => {
    dictList.value.deptTree = dr.data
  })
  await roleList().then((rr) => {
    dictList.value.roles = rr.data
  })
  props.pop.dictList = dictList
}
// 检索部门
function filterDept(searchValue: string, nodeData: any) {
  return nodeData.title.toLowerCase().indexOf(searchValue.toLowerCase()) > -1
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
    await accountDel(delItem.delId)
  } catch (err) {
    return false
  } finally {
    // Nothing
    pageQuery()
  }
}
// 重置对象
const resetItem = reactive({
  resetConfirm: false,
  resetId: 0
})
// 打开重置
function openReset(item: any) {
  resetItem.resetId = item.id
  resetItem.resetConfirm = true
}
// 确认重置
async function reseting() {
  try {
    await accountReset(resetItem.resetId)
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
