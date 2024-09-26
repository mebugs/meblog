<template>
  <a-form size="large" label-align="left" class="form" layout="vertical" :model="formData" @submit="submit">
    <a-row :gutter="20">
      <a-col :span="12">
        <a-form-item field="groupKey" :label="$t('dict.groupKey')" :rules="[{ required: true, message: $t('rule.required') }]">
          <a-select
            v-model="formData.groupKey"
            :options="pop.dictList.dictGroup"
            allow-clear
            allow-search
            @change="groupKeyChange"
            :placeholder="$t('rule.select')" />
        </a-form-item>
      </a-col>
      <a-col :span="12">
        <a-form-item field="remark" :label="$t('dict.remark')">
          <a-input v-model="formData.remark" :max-length="64" allow-clear show-word-limit :placeholder="$t('dict.remark.place')" />
        </a-form-item>
      </a-col>
      <a-col :span="12">
        <a-form-item field="val" :label="$t('dict.val')" :rules="[{ required: true, message: $t('rule.required') }]">
          <template #extra>
            <div>{{ $t('dict.val.tips') + (nextVal != -1 ? nextVal : $t('dict.val.nextWait')) }}</div>
          </template>
          <a-input v-model="formData.val" :max-length="32" allow-clear show-word-limit :placeholder="$t('dict.val.place')" />
        </a-form-item>
      </a-col>
      <a-col :span="12">
        <a-form-item field="choose" :label="$t('dict.choose')">
          <template #extra>
            <div>{{ $t('dict.choose.tips') }}</div>
          </template>
          <a-select v-model="formData.choose" :options="pop.dictList.openStatus" :placeholder="$t('rule.select')" />
        </a-form-item>
      </a-col>
      <a-col :span="12">
        <a-form-item field="label" :label="$t('dict.label')" :rules="[{ required: true, message: $t('rule.required') }]">
          <a-input v-model="formData.label" :max-length="64" allow-clear show-word-limit :placeholder="$t('dict.label.place')" />
        </a-form-item>
      </a-col>
      <a-col :span="12">
        <a-form-item field="labelEn" :label="$t('dict.labelEn')" :rules="[{ required: true, message: $t('rule.required') }]">
          <a-input v-model="formData.labelEn" :max-length="64" allow-clear show-word-limit :placeholder="$t('dict.labelEn.place')" />
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
import { onMounted, ref } from 'vue'
import type { Pop } from '@/utils/hooks/pop'
import useLoad from '@/utils/hooks/load'
import { dictInit, dictNextVal, dictAdd } from '@/api/plat/dict'
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
// const dictAdd = ref<FormInstance>();
// 提交数据
const submit = async ({ errors, values }: { errors: any; values: any }) => {
  if (load.value) return
  if (!errors) {
    setLoad(true)
    try {
      // const res = await dictAdd.value?.validate();
      await dictAdd(values)
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
// 推荐值
const nextVal = ref(-1)
async function groupKeyChange(val: string) {
  // 选择的groupKey产生的变更
  if (val) {
    setLoad(true)
    try {
      const res = await dictNextVal(val)
      nextVal.value = res.data
      return
    } catch (err) {
      // DoNothing CommonPopUp
    } finally {
      setLoad(false)
    }
  }
  nextVal.value = -1
}
// 页面渲染
onMounted(() => {
  // Nothing
  setLoad(false)
})
</script>
