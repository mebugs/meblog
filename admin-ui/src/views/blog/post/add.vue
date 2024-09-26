<template>
  <a-form size="large" label-align="left" class="form" layout="vertical" :model="formData" @submit="submit">
    <a-row :gutter="20">
      <a-col :span="24">
        <a-form-item field="title" :label="$t('post.title')" :rules="[{ required: true, message: $t('rule.required') }]">
          <a-input v-model="formData.title" :max-length="64" allow-clear show-word-limit :placeholder="$t('post.title.place')" />
        </a-form-item>
      </a-col>
      <a-col :span="24">
        <a-form-item field="summary" :label="$t('post.summary')" :rules="[{ required: true, message: $t('rule.summary') }]">
          <a-textarea
            v-model="formData.summary"
            :max-length="256"
            allow-clear
            show-word-limit
            :placeholder="$t('post.summary.place')" />
        </a-form-item>
      </a-col>
      <a-col :span="24" style="margin: 20px 0">
        <div id="vditorAdd"></div>
      </a-col>
      <a-col :span="24">
        <a-form-item field="tags" :label="$t('post.tags')">
          <a-select
            v-model="formData.tags"
            :options="pop.dictList.roles"
            multiple
            allow-clear
            allow-search
            :placeholder="$t('rule.select')" />
        </a-form-item>
      </a-col>
      <a-col :span="12">
        <a-form-item field="group" :label="$t('post.group')" :rules="[{ required: true, message: $t('rule.required') }]">
          <a-select
            v-model="formData.group"
            :options="pop.dictList.postGroupList"
            allow-clear
            allow-search
            :placeholder="$t('rule.select')" />
        </a-form-item>
      </a-col>
      <a-col :span="12">
        <a-form-item field="status" :label="$t('post.status')" :rules="[{ required: true, message: $t('rule.required') }]">
          <a-select
            v-model="formData.status"
            :options="pop.dictList.postStatus"
            allow-clear
            allow-search
            :placeholder="$t('rule.select')" />
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
import '~/vditor/src/assets/less/index.less'
import { onMounted, ref } from 'vue'
import type { Pop } from '@/utils/hooks/pop'
import useLoad from '@/utils/hooks/load'
import vditor from '@/utils/hooks/vditor'
import { postInit, postAdd } from '@/api/blog/post'
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
// 0 编辑 1 正文预览 2 确认发布
const step = ref(0)
const { vd, md, getNew, toPreview, getResponse } = vditor(step, 'vditorAdd')

// 加载中变量
const { load, setLoad } = useLoad(false)
// 表单数据初始化
const formData = postInit()
// const postAdd = ref<FormInstance>();
// 提交数据
const submit = async ({ errors, values }: { errors: any; values: any }) => {
  if (load.value) return
  if (!errors) {
    setLoad(true)
    try {
      // const res = await postAdd.value?.validate();
      await postAdd(values)
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
  getNew()
})
</script>
