<template>
  <a-spin :loading="load">
    <a-form size="large" label-align="left" class="form" layout="vertical" :model="formData">
      <a-row :gutter="20">
        <a-col :span="12">
          <a-form-item field="serviceCode" :label="$t('plat.serviceCode')">
            <span class="formSpan">{{ pop.dictMap.serviceCode[formData.serviceCode] }}</span>
          </a-form-item>
        </a-col>
        <a-col :span="12">
          <a-form-item field="type" :label="$t('response.type')">
            <span class="formSpan">{{ pop.dictMap.responseType[formData.type] }}</span>
          </a-form-item>
        </a-col>
        <a-col :span="12">
          <a-form-item field="code" :label="$t('response.code')">
            <span class="formSpan">{{ formData.code }}</span>
          </a-form-item>
        </a-col>
        <a-col :span="12">
          <a-form-item field="remark" :label="$t('response.remark')">
            <span class="formSpan">{{ formData.remark }}</span>
          </a-form-item>
        </a-col>
        <a-col :span="12">
          <a-form-item field="zhCn" :label="$t('response.zhCn')">
            <span class="formSpan">{{ formData.zhCn }}</span>
          </a-form-item>
        </a-col>
        <a-col :span="12">
          <a-form-item field="enUs" :label="$t('response.enUs')">
            <span class="formSpan">{{ formData.enUs }}</span>
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
import { responseInit, responseGet } from '@/api/plat/response'
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
const formData = responseInit()
async function get() {
  setLoad(true)
  try {
    const res = await responseGet(props.pop.itemId)
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
