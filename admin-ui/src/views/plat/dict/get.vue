<template>
  <a-spin :loading="load">
    <a-form size="large" label-align="left" class="form" layout="vertical" :model="formData">
      <a-row :gutter="20">
        <a-col :span="12">
          <a-form-item field="groupKey" :label="$t('dict.groupKey')">
            <span class="formSpan">{{ pop.dictMap.dictGroup[formData.groupKey] }}</span>
          </a-form-item>
        </a-col>
        <a-col :span="12">
          <a-form-item field="remark" :label="$t('dict.remark')">
            <span class="formSpan">{{ formData.remark }}</span>
          </a-form-item>
        </a-col>
        <a-col :span="12">
          <a-form-item field="val" :label="$t('dict.val')">
            <span class="formSpan">{{ formData.val }}</span>
          </a-form-item>
        </a-col>
        <a-col :span="12">
          <a-form-item field="choose" :label="$t('dict.choose')">
            <template #extra>
              <div>{{ $t('dict.choose.tips') }}</div>
            </template>
            <span class="formSpan">{{ pop.dictMap.openStatus[formData.choose] }}</span>
          </a-form-item>
        </a-col>
        <a-col :span="12">
          <a-form-item field="label" :label="$t('dict.label')">
            <span class="formSpan">{{ formData.label }}</span>
          </a-form-item>
        </a-col>
        <a-col :span="12">
          <a-form-item field="labelEn" :label="$t('dict.labelEn')">
            <span class="formSpan">{{ formData.labelEn }}</span>
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
import { dictInit, dictGet } from '@/api/plat/dict'
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
const formData = dictInit()
async function get() {
  setLoad(true)
  try {
    const res = await dictGet(props.pop.itemId)
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
