<template>
  <a-form size="large" label-align="left" class="form" layout="vertical" :model="formData" @submit="submit">
    <a-row :gutter="20">
      <a-col :span="12">
        <a-form-item field="account" :label="$t('account.account')" :rules="[{ required: true, message: $t('rule.required') }]">
          <template #extra>
            <div>{{ $t('account.account.tips') }}</div>
          </template>
          <a-input v-model="formData.account" :max-length="16" allow-clear show-word-limit :placeholder="$t('account.account.place')" />
        </a-form-item>
      </a-col>
      <a-col :span="12">
        <a-form-item
          field="permissionType"
          :label="$t('account.permissionType')"
          :rules="[{ required: true, message: $t('rule.required') }]">
          <template #extra>
            <div>{{ $t('account.permissionType.tips') }}</div>
          </template>
          <a-select
            v-model="formData.permissionType"
            :options="pop.dictList.accountPermission"
            allow-clear
            allow-search
            :placeholder="$t('rule.select')" />
        </a-form-item>
      </a-col>
      <a-col :span="12">
        <a-form-item field="deptId" :label="$t('account.deptId')" :rules="[{ required: true, message: $t('rule.required') }]">
          <a-tree-select
            v-model="formData.deptId"
            :allow-search="true"
            :allow-clear="true"
            :data="pop.dictList.deptTree"
            :filter-tree-node="filterDept"
            :placeholder="$t('rule.select')" />
        </a-form-item>
      </a-col>
      <a-col :span="12">
        <a-form-item field="isLeader" :label="$t('account.isLeader')" :rules="[{ required: true, message: $t('rule.required') }]">
          <a-select
            v-model="formData.isLeader"
            :options="pop.dictList.deptAccount"
            allow-clear
            allow-search
            :placeholder="$t('rule.select')" />
        </a-form-item>
      </a-col>
      <a-col :span="12">
        <a-form-item field="status" :label="$t('account.status')" :rules="[{ required: true, message: $t('rule.required') }]">
          <a-select
            v-model="formData.status"
            :options="pop.dictList.accountStatus"
            allow-clear
            allow-search
            :placeholder="$t('rule.select')" />
        </a-form-item>
      </a-col>
      <a-col :span="12">
        <a-form-item field="name" :label="$t('account.name')">
          <a-input v-model="formData.name" :max-length="16" allow-clear show-word-limit :placeholder="$t('account.name.place')" />
        </a-form-item>
      </a-col>
      <a-col :span="24">
        <a-form-item field="roleIds" :label="$t('account.roleIds')">
          <a-select
            v-model="formData.roleIds"
            :options="pop.dictList.roles"
            multiple
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
import { onMounted } from 'vue'
import type { Pop } from '@/utils/hooks/pop'
import useLoad from '@/utils/hooks/load'
import { accountInit, accountAdd } from '@/api/plat/account'
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
// 检索部门
function filterDept(searchValue: string, nodeData: any) {
  return nodeData.title.toLowerCase().indexOf(searchValue.toLowerCase()) > -1
}
// 加载中变量
const { load, setLoad } = useLoad(false)
// 表单数据初始化
const formData = accountInit()
// const accountAdd = ref<FormInstance>();
// 提交数据
const submit = async ({ errors, values }: { errors: any; values: any }) => {
  if (load.value) return
  if (!errors) {
    setLoad(true)
    try {
      // const res = await accountAdd.value?.validate();
      await accountAdd(values)
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
