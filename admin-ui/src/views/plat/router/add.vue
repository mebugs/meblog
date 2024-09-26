<template>
  <a-form size="large" label-align="left" class="form" layout="vertical" :model="formData" @submit="submit">
    <a-row :gutter="20">
      <a-col :span="12">
        <a-form-item field="type" :label="$t('router.type')">
          <template #extra>
            <div>{{ $t('router.type.tips') }}</div>
          </template>
          <a-select v-model="formData.type" :options="pop.dictList.routerType" :placeholder="$t('rule.select')" />
        </a-form-item>
      </a-col>
      <a-col :span="12">
        <a-form-item field="serviceCode" :label="$t('plat.serviceCode')" :rules="[{ required: true, message: $t('rule.required') }]">
          <a-select
            v-model="formData.serviceCode"
            :options="pop.dictList.serviceCode"
            allow-clear
            allow-search
            :placeholder="$t('rule.select')" />
        </a-form-item>
      </a-col>
      <a-col :span="12">
        <a-form-item
          field="url"
          :label="$t('router.url')"
          :rules="[
            { required: true, message: $t('rule.required') },
            {
              match: /^(\/)[a-z0-9\/]*$/,
              message: $t('router.url.format')
            }
          ]">
          <a-input v-model="formData.url" :max-length="64" allow-clear show-word-limit :placeholder="$t('router.url.place')" />
        </a-form-item>
      </a-col>
      <a-col :span="12">
        <a-form-item field="name" :label="$t('router.name')" :rules="[{ required: true, message: $t('rule.required') }]">
          <a-input v-model="formData.name" :max-length="32" allow-clear show-word-limit :placeholder="$t('router.name.place')" />
        </a-form-item>
      </a-col>
      <a-col :span="8">
        <a-form-item field="reqLogPrint" :label="$t('router.reqLogPrint')">
          <template #extra>
            <div>{{ $t('router.reqLogPrint.tips') }}</div>
          </template>
          <a-select v-model="formData.reqLogPrint" :options="pop.dictList.openStatus" :placeholder="$t('rule.select')" />
        </a-form-item>
      </a-col>
      <a-col :span="8">
        <a-form-item field="resLogPrint" :label="$t('router.resLogPrint')">
          <template #extra>
            <div>{{ $t('router.resLogPrint.tips') }}</div>
          </template>
          <a-select v-model="formData.resLogPrint" :options="pop.dictList.openStatus" :placeholder="$t('rule.select')" />
        </a-form-item>
      </a-col>
      <a-col :span="8">
        <a-form-item field="logInDb" :label="$t('router.logInDb')">
          <template #extra>
            <div>{{ $t('router.logInDb.tips') }}</div>
          </template>
          <a-select v-model="formData.logInDb" :options="pop.dictList.openStatus" :placeholder="$t('rule.select')" />
        </a-form-item>
      </a-col>
      <a-col :span="24" class="formGroup">
        <a-row :gutter="20">
          <a-col :span="24">
            <a-form-item field="reqLogSecure" :label="$t('router.reqLogSecure')">
              <template #extra>
                <div>{{ $t('router.reqLogSecure.tips') }}</div>
              </template>
              <a-textarea
                v-model="formData.reqLogSecure"
                :max-length="256"
                allow-clear
                show-word-limit
                :placeholder="$t('router.reqLogSecure.place')" />
            </a-form-item>
          </a-col>
          <a-col :span="24">
            <a-form-item field="resLogSecure" :label="$t('router.resLogSecure')">
              <template #extra>
                <div>{{ $t('router.resLogSecure.tips') }}</div>
              </template>
              <a-textarea
                v-model="formData.resLogSecure"
                :max-length="256"
                allow-clear
                show-word-limit
                :placeholder="$t('router.resLogSecure.place')" />
            </a-form-item>
          </a-col>
        </a-row>
      </a-col>
      <a-col :span="24">
        <a-form-item field="remark" :label="$t('router.remark')">
          <a-textarea
            v-model="formData.remark"
            :max-length="128"
            allow-clear
            show-word-limit
            :placeholder="$t('router.remark.place')" />
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
import { routerInit, routerAdd } from '@/api/plat/router'
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
// const routerAdd = ref<FormInstance>();
// 提交数据
const submit = async ({ errors, values }: { errors: any; values: any }) => {
  if (load.value) return
  if (!errors) {
    setLoad(true)
    try {
      // const res = await routerAdd.value?.validate();
      await routerAdd(values)
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
  setLoad(false)
})
</script>
