<template>
  <a-form size="large" label-align="left" class="form" layout="vertical" :model="formData" @submit="submit">
    <a-row :gutter="20">
      <a-col :span="12">
        <a-form-item field="level" :label="$t('permission.level')">
          <span class="formSpan">{{ pop.dictMap.permissionLevel[formData.level] }}</span>
        </a-form-item>
      </a-col>
      <a-col :span="12">
        <a-form-item field="static" :label="$t('permission.static')">
          <template #extra>
            <div>{{ $t('permission.static.tips') }}</div>
          </template>
          <a-select v-model="formData.static" :options="pop.dictList.openStatus" :placeholder="$t('rule.select')" />
        </a-form-item>
      </a-col>
      <a-col :span="12">
        <a-form-item field="name" :label="$t('permission.name')">
          <a-input v-model="formData.name" :max-length="32" allow-clear show-word-limit :placeholder="$t('permission.name.place')" />
        </a-form-item>
      </a-col>
      <a-col :span="12">
        <a-form-item field="alias" :label="$t('permission.alias')">
          <a-input v-model="formData.alias" :max-length="32" allow-clear show-word-limit :placeholder="$t('permission.alias.place')" />
        </a-form-item>
      </a-col>
      <a-col :span="24">
        <a-divider orientation="left">
          {{ $t('permission.routers') }}
          <a-tooltip :content="$t('permission.routers.check')" :mini="true">
            <a-button type="primary" size="small" @click="open">
              <template #icon>
                <icon-edit />
              </template>
            </a-button>
          </a-tooltip>
        </a-divider>
        <routers :list="formData.routers" :remove="true" :dict-map="pop.dictMap" :remove-func="removeQuick" />
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
              {{ $t('button.cancel') }}</a-button
            >
          </a-space>
        </div>
      </a-col>
    </a-row>
  </a-form>
  <a-modal v-model:visible="check" :hide-title="true" :width="900" :on-before-ok="done" :on-before-cancell="cancel">
    <router-check v-if="check" :do-select="doSelect" :data="selectData" />
  </a-modal>
</template>

<script lang="ts" setup>
import { onMounted } from 'vue'
import type { Pop } from '@/utils/hooks/pop'
import { CheckList } from '@/utils/hooks/check'
import useLoad from '@/utils/hooks/load'
import { permissionInit, permissionAdd } from '@/api/plat/permission'
import Routers from './routers.vue'
import RouterCheck from './routerCheck.vue'
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
const formData = permissionInit()
// 等级初始化
let level = Number(props.pop.useObj.level) + 1
formData.value.level = level.toString()
formData.value.pid = props.pop.useObj.id
// 选择框初始化
const { check, selectData, open, cancel, done, doSelect, removeQuick } = CheckList(formData, 'routerIds', 'routers')
// const permissionAdd = ref<FormInstance>();
// 提交数据
const submit = async ({ errors, values }: { errors: any; values: any }) => {
  if (load.value) return
  if (!errors) {
    setLoad(true)
    try {
      // const res = await permissionAdd.value?.validate();
      await permissionAdd(values)
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
  setLoad(false)
})
</script>
