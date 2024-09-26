<template>
  <a-spin :loading="load">
    <a-form size="large" label-align="left" class="form" layout="vertical" :model="formData">
      <a-row :gutter="24">
        <a-col :span="12">
          <a-form-item field="type" :label="$t('router.type')">
            <template #extra>
              <div>{{ $t('router.type.tips') }}</div>
            </template>
            <span class="formSpan">{{ pop.dictMap.routerType[formData.type] }}</span>
          </a-form-item>
        </a-col>
        <a-col :span="12">
          <a-form-item field="serviceCode" :label="$t('plat.serviceCode')">
            <span class="formSpan">{{ pop.dictMap.serviceCode[formData.serviceCode] }}</span>
          </a-form-item>
        </a-col>
        <a-col :span="12">
          <a-form-item field="url" :label="$t('router.url')">
            <span class="formSpan">{{ formData.url }}</span>
          </a-form-item>
        </a-col>
        <a-col :span="12">
          <a-form-item field="name" :label="$t('router.name')">
            <span class="formSpan">{{ formData.name }}</span>
          </a-form-item>
        </a-col>
        <a-col :span="8">
          <a-form-item field="reqLogPrint" :label="$t('router.reqLogPrint')">
            <template #extra>
              <div>{{ $t('router.reqLogPrint.tips') }}</div>
            </template>
            <span class="formSpan">{{ pop.dictMap.openStatus[formData.reqLogPrint] }}</span>
          </a-form-item>
        </a-col>
        <a-col :span="8">
          <a-form-item field="resLogPrint" :label="$t('router.resLogPrint')">
            <template #extra>
              <div>{{ $t('router.resLogPrint.tips') }}</div>
            </template>
            <span class="formSpan">{{ pop.dictMap.openStatus[formData.resLogPrint] }}</span>
          </a-form-item>
        </a-col>
        <a-col :span="8">
          <a-form-item field="logInDb" :label="$t('router.logInDb')">
            <template #extra>
              <div>{{ $t('router.logInDb.tips') }}</div>
            </template>
            <span class="formSpan">{{ pop.dictMap.openStatus[formData.logInDb] }}</span>
          </a-form-item>
        </a-col>
        <a-col :span="24" class="formGroup">
          <a-row :gutter="20">
            <a-col :span="24">
              <a-form-item field="reqLogSecure" :label="$t('router.reqLogSecure')">
                <template #extra>
                  <div>{{ $t('router.reqLogSecure.tips') }}</div>
                </template>
                <span class="formSpan">{{ formData.reqLogSecure }}</span>
              </a-form-item>
            </a-col>
            <a-col :span="24">
              <a-form-item field="resLogSecure" :label="$t('router.resLogSecure')">
                <template #extra>
                  <div>{{ $t('router.resLogSecure.tips') }}</div>
                </template>
                <span class="formSpan">{{ formData.resLogSecure }}</span>
              </a-form-item>
            </a-col>
          </a-row>
        </a-col>
        <a-col :span="24">
          <a-form-item field="remark" :label="$t('router.remark')">
            <span class="formSpan">{{ formData.remark }}</span>
          </a-form-item>
        </a-col>
      </a-row>
    </a-form>
  </a-spin>
</template>

<script lang="ts" setup>
import { onMounted } from 'vue'
import type { Pop } from '@/utils/hooks/pop'
import useLoad from '@/utils/hooks/load'
import { routerInit, routerGet } from '@/api/plat/router'
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
const formData = routerInit()
async function get() {
  setLoad(true)
  try {
    const res = await routerGet(props.pop.itemId)
    formData.value = res.data
    return
  } catch (err) {
    // DoNothing CommonPopUp
  } finally {
    setLoad(false)
  }
}
// 页面渲染
onMounted(() => {
  // Nothing
  get()
})
</script>
