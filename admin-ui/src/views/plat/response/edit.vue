<template>
  <a-form size="large" label-align="left" class="form" layout="vertical" :model="formData" @submit="submit">
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
          <a-input v-model="formData.remark" :max-length="64" allow-clear show-word-limit :placeholder="$t('response.remark.place')" />
        </a-form-item>
      </a-col>
      <a-col :span="12">
        <a-form-item field="zhCn" :label="$t('response.zhCn')" :rules="[{ required: true, message: $t('rule.required') }]">
          <a-input v-model="formData.zhCn" :max-length="64" allow-clear show-word-limit :placeholder="$t('response.zhCn.place')" />
        </a-form-item>
      </a-col>
      <a-col :span="12">
        <a-form-item field="enUs" :label="$t('response.enUs')" :rules="[{ required: true, message: $t('rule.required') }]">
          <a-input v-model="formData.enUs" :max-length="64" allow-clear show-word-limit :placeholder="$t('response.enUs.place')" />
        </a-form-item>
      </a-col>
      <a-col :span="24">
        <a-divider />
        <div class="doBtn">
          <a-space>
            <a-button size="large" type="primary" html-type="submit" :loading="load">
              <template #icon>
                <icon-check />
              </template>
              {{ $t('button.submit') }}
            </a-button>
            <a-button size="large" @click="pop.close()">
              <template #icon>
                <icon-close />
              </template>
              {{ $t('button.cancel') }}</a-button
            >
          </a-space>
        </div>
      </a-col>
    </a-row>
  </a-form>
</template>

<script lang="ts" setup>
import { onMounted } from 'vue'
import type { Pop } from '@/utils/hooks/pop'
import useLoad from '@/utils/hooks/load'
import { responseInit, responseGet, responseEdit } from '@/api/plat/response'
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
// const responseAdd = ref<FormInstance>();
// 提交数据
const submit = async ({ errors, values }: { errors: any; values: any }) => {
  if (load.value) return
  if (!errors) {
    setLoad(true)
    try {
      // const res = await responseAdd.value?.validate();
      await responseEdit(values)
      // Pop Close & Back
      props.pop.close()
      props.pop.callBack()
    } catch (err) {
      // DoNothing
    } finally {
      setLoad(false)
    }
  }
}
// 页面渲染
onMounted(() => {
  // Nothing
  get()
})
</script>
