<template>
  <a-spin :loading="load">
    <a-page-header :title="$t('permission.get')" :subtitle="formData.name" :show-back="false" />
    <btns :pop="pop" :data="formData" />
    <a-form size="large" label-align="left" class="form" layout="vertical" :model="formData">
      <a-row :gutter="20">
        <a-col :span="12">
          <a-form-item field="level" :label="$t('permission.level')">
            <span class="formSpan">{{ dictMap.permissionLevel[formData.level] }}</span>
          </a-form-item>
        </a-col>
        <a-col :span="12">
          <a-form-item field="static" :label="$t('permission.static')">
            <template #extra>
              <div>{{ $t('permission.static.tips') }}</div>
            </template>
            <span class="formSpan">{{ dictMap.openStatus[formData.static] }}</span>
          </a-form-item>
        </a-col>
        <a-col :span="12">
          <a-form-item field="name" :label="$t('permission.name')">
            <span class="formSpan">{{ formData.name }}</span>
          </a-form-item>
        </a-col>
        <a-col :span="12">
          <a-form-item field="alias" :label="$t('permission.alias')">
            <span class="formSpan">{{ formData.alias }}</span>
          </a-form-item>
        </a-col>
        <a-col :span="24">
          <a-divider orientation="left">{{ $t('permission.routers') }}</a-divider>
          <routers :list="formData.routers" :remove="false" :dict-map="dictMap" />
        </a-col>
      </a-row>
    </a-form>
  </a-spin>
</template>

<script lang="ts" setup>
import { onMounted, watch, ref } from 'vue'
import type { Pop } from '@/utils/hooks/pop'
import useLocale from '@/utils/hooks/locale'
import useLoad from '@/utils/hooks/load'
import { dictRead } from '@/api/plat/dict'
import { permissionInit, permissionGet } from '@/api/plat/permission'
import Routers from './routers.vue'
import Btns from './btns.vue'
// 当前语言
const { currentLocale } = useLocale()
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
const formData = permissionInit()
async function get() {
  setLoad(true)
  try {
    const res = await permissionGet(props.pop.itemId)
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
const dictList = ref({ permissionLevel: [], serviceCode: [], openStatus: [] })
const dictMap = ref({ permissionLevel: {} as any, serviceCode: {} as any, openStatus: {} as any })
// 字段初始化
async function dictInit() {
  // 指定字典Key
  await dictRead({ groupKeys: ['permissionLevel', 'serviceCode', 'openStatus'] }).then((r) => {
    dictList.value = r.data.list
    dictMap.value = r.data.map
    props.pop.dictList = dictList
    props.pop.dictMap = dictMap
  })
}
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
