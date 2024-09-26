<template>
  <a-spin :loading="load">
    <a-page-header :title="$t('dept.get')" :subtitle="formData.name" :show-back="false" />
    <btns :pop="pop" :data="formData" />
    <a-form size="large" label-align="left" class="form" layout="vertical" :model="formData">
      <a-row :gutter="20">
        <a-col :span="12">
          <a-form-item field="name" :label="$t('dept.name')">
            <span class="formSpan">{{ formData.name }}</span>
          </a-form-item>
        </a-col>
        <a-col :span="12">
          <a-form-item field="permissionType" :label="$t('dept.permissionType')">
            <template #extra>
              <div>{{ $t('dept.permissionType.tips') }}</div>
            </template>
            <span class="formSpan">{{ dictMap.deptPermission[formData.permissionType] }}</span>
          </a-form-item>
        </a-col>
        <a-col :span="24">
          <a-row :gutter="20">
            <a-col :span="12">
              <a-divider orientation="left">{{ $t('dept.accounts.isLeader') }}</a-divider>
              <a-table
                :scrollbar="false"
                :sticky-header="false"
                :row-key="'id'"
                :columns="columns"
                :data="formData.accounts[0]"
                :pagination="false"
                size="medium"
                style="margin-top: 10px">
              </a-table>
            </a-col>
            <a-col :span="12">
              <a-divider orientation="left">{{ $t('dept.accounts.unLeader') }}</a-divider>
              <a-table
                :scrollbar="false"
                :sticky-header="false"
                :row-key="'id'"
                :columns="columns"
                :data="formData.accounts[1]"
                :pagination="false"
                size="medium"
                style="margin-top: 10px">
              </a-table>
            </a-col>
          </a-row>
        </a-col>
      </a-row>
    </a-form>
  </a-spin>
</template>

<script lang="ts" setup>
import { onMounted, watch, ref, computed } from 'vue'
import type { Pop } from '@/utils/hooks/pop'
import useLocale from '@/utils/hooks/locale'
import useLoad from '@/utils/hooks/load'
import { dictRead } from '@/api/plat/dict'
import { deptInit, deptGet } from '@/api/plat/dept'
import Btns from './btns.vue'
import { useI18n } from 'vue-i18n'
// 当前语言
const { currentLocale } = useLocale()
const { t } = useI18n()
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
// 表单数据初始化
const formData = deptInit()
async function get() {
  setLoad(true)
  try {
    const res = await deptGet(props.pop.itemId)
    formData.value = res.data
    return
  } catch (err) {
    // DoNothing CommonPopUp
  } finally {
    setLoad(false)
  }
}
props.pop.callBack = get
// 初始化字典对象
const dictList = ref({ deptPermission: [], deptToType: [] })
const dictMap = ref({ deptPermission: {} as any, deptToType: {} as any })
// 字段初始化
async function dictInit() {
  // 指定字典Key
  await dictRead({ groupKeys: ['deptPermission', 'deptToType'] }).then((r) => {
    dictList.value = r.data.list
    dictMap.value = r.data.map
    props.pop.dictList = dictList
    props.pop.dictMap = dictMap
  })
}
// 账号表格
const columns = computed(() => [
  { title: t('account.account'), dataIndex: 'account' },
  { title: t('account.name'), dataIndex: 'name' }
])
// 业务初始化
function init() {
  // 初始化后端字典对象
  dictInit()
}
// 页面渲染
onMounted(() => {
  init()
  get()
})
// 语言监听
watch(currentLocale, (n, o) => {
  if (n !== o) {
    init()
  }
})
watch(
  () => props.pop.itemId, // 监听defineProps的特殊处理方式
  (n, o) => {
    if (n !== o) {
      get()
    }
  }
)
</script>
