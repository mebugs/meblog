<template>
  <a-form size="large" label-align="left" class="form" layout="vertical" :model="formData" @submit="submit">
    <a-row :gutter="20">
      <a-col :span="24">
        <a-form-item field="sourceShow" :label="$t('category.sourceId')" :rules="[{ required: true, message: $t('rule.required') }]">
          <label class="upIcon" for="upMainImg">
            <img :src="formData.sourceShow" v-if="formData.sourceShow" />
            <p v-else>
              <icon-upload />
              {{ $t('button.upload') }}
            </p>
          </label>
          <input
            id="upMainImg"
            accept="image/gif, image/jpeg, image/png, image/jpg"
            type="file"
            style="display: none"
            @change="chooesMain" />
        </a-form-item>
      </a-col>
      <a-col :span="12">
        <a-form-item field="title" :label="$t('category.title')" :rules="[{ required: true, message: $t('rule.required') }]">
          <a-input v-model="formData.title" :max-length="32" allow-clear show-word-limit :placeholder="$t('category.title.place')" />
        </a-form-item>
      </a-col>
      <a-col :span="12">
        <a-form-item field="url" :label="$t('category.url')" :rules="[{ required: true, message: $t('rule.required') }]">
          <a-input v-model="formData.url" :max-length="32" allow-clear show-word-limit :placeholder="$t('category.url.place')" />
        </a-form-item>
      </a-col>
      <a-col :span="24">
        <a-form-item field="summary" :label="$t('category.summary')" :rules="[{ required: true, message: $t('rule.required') }]">
          <a-input
            v-model="formData.summary"
            :max-length="64"
            allow-clear
            show-word-limit
            :placeholder="$t('category.summary.place')" />
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
import { categoryInit, categoryAdd } from '@/api/blog/category'
import useImgs from '@/utils/hooks/imgs'
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
const formData = categoryInit()
// const categoryAdd = ref<FormInstance>();
// 提交数据
const submit = async ({ errors, values }: { errors: any; values: any }) => {
  if (load.value) return
  if (!errors) {
    setLoad(true)
    try {
      // const res = await categoryAdd.value?.validate();
      await categoryAdd(values)
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
// 主图加载器
const { imgObj, initImgQuick, chooesImg } = useImgs()
initImgQuick('icon')
const chooesMain = async (e: Event) => {
  await chooesImg(e)
    .then(() => {
      formData.value.source = imgObj.value.baseUrls
      formData.value.sourceShow = imgObj.value.baseUrls[0]
    })
    .catch((er) => {
      console.log(er)
    })
}
// 页面渲染
onMounted(() => {
  // Nothing
  setLoad(false)
})
</script>
