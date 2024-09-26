<template>
  <a-spin :loading="load">
    <a-form size="large" label-align="left" class="form" layout="vertical" :model="formData">
      <a-row :gutter="20">
        <a-col :span="12">
          <a-form-item field="account" :label="$t('account.account')">
            <span class="formSpan">{{ formData.account }}</span>
          </a-form-item>
        </a-col>
        <a-col :span="12">
          <a-form-item field="permissionType" :label="$t('account.permissionType')">
            <span class="formSpan">{{ pop.dictMap.accountPermission[formData.permissionType] }}</span>
          </a-form-item>
        </a-col>

        <a-col :span="12">
          <a-form-item field="deptId" :label="$t('account.deptId')">
            <span class="formSpan">{{ formData.deptName }}</span>
          </a-form-item>
        </a-col>
        <a-col :span="12">
          <a-form-item field="isLeader" :label="$t('account.isLeader')">
            <span class="formSpan">{{ pop.dictMap.deptAccount[formData.isLeader] }}</span>
          </a-form-item>
        </a-col>
        <a-col :span="12">
          <a-form-item field="status" :label="$t('account.status')">
            <span class="formSpan">{{ pop.dictMap.accountStatus[formData.status] }}</span>
          </a-form-item>
        </a-col>
        <a-col :span="12">
          <a-form-item field="name" :label="$t('account.name')">
            <span class="formSpan">{{ formData.name }}</span>
          </a-form-item>
        </a-col>
        <a-col :span="24">
          <a-form-item field="roleIds" :label="$t('account.roleIds')">
            <span class="formSpan">
              <a-tag v-if="formData.roleNames" v-for="item in formData.roleNames">{{ item }}</a-tag>
            </span>
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
import { accountInit, accountGet } from '@/api/plat/account'
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
const formData = accountInit()
async function get() {
  setLoad(true)
  try {
    const res = await accountGet(props.pop.itemId)
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
