<template>
  <a-form size="large" label-align="left" class="form" layout="vertical" :model="formData" @submit="submit">
    <a-row :gutter="20">
      <a-col :span="12">
        <a-form-item field="name" :label="$t('dept.name')">
          <a-input v-model="formData.name" :max-length="16" allow-clear show-word-limit :placeholder="$t('dept.name.place')" />
        </a-form-item>
      </a-col>
      <a-col :span="12">
        <a-form-item field="permissionType" :label="$t('dept.permissionType')">
          <template #extra>
            <div>{{ $t('dept.permissionType.tips') }}</div>
          </template>
          <a-select v-model="formData.permissionType" :options="pop.dictList.deptPermission" :placeholder="$t('rule.select')" />
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
            <a-button size="large" @click="pop.treeClose()">
              <template #icon>
                <icon-close />
              </template>
              {{ $t('button.cancel') }}
            </a-button>
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
import { deptInit, deptGet, deptEdit } from '@/api/plat/dept'
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
const formData = deptInit()
async function get() {
  setLoad(true)
  try {
    const res = await deptGet(props.pop.itemId)
    formData.value = res.data
    return
  } catch (err) {
    // DoNothing CommonPopUp
  } finally {
    setLoad(false)
  }
} // const deptAdd = ref<FormInstance>();
// 提交数据
const submit = async ({ errors, values }: { errors: any; values: any }) => {
  if (load.value) return
  if (!errors) {
    setLoad(true)
    try {
      // const res = await deptAdd.value?.validate();
      await deptEdit(values)
      // Pop Close & Back 刷新树
      props.pop.treeClose()
      props.pop.treeRefresh(false)
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
