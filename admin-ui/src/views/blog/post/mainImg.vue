<template>
  <label class="upImg" for="upMainImg">
    <img :src="formData.img" v-if="formData.img" />
    <p for="upImg" v-else>
      <icon-upload />
      {{ $t('button.upload') }}
    </p>
  </label>
  <input id="upMainImg" accept="image/gif, image/jpeg, image/png, image/jpg" type="file" style="display: none" @change="chooesMain" />
</template>

<script lang="ts" setup>
import useImgs from '@/utils/hooks/imgs'
// 入参读取
const props = defineProps({
  formData: {
    type: Object,
    required: true,
    default: () => {
      return {}
    }
  }
})
// 主图加载器
const { imgObj, initImgQuick, chooesImg } = useImgs()
initImgQuick('main')
const chooesMain = async (e: Event) => {
  await chooesImg(e)
    .then(() => {
      props.formData.img = imgObj.value.baseUrl
      props.formData.imgs = imgObj.value.smallBaseUrl
      console.log(props.formData)
    })
    .catch((er) => {
      console.log(er)
    })
}
</script>
