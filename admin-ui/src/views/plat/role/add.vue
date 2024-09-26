<template>
  <a-form size="large" label-align="left" class="form" layout="vertical" :model="formData" @submit="submit">
    <a-row :gutter="20">
      <a-col :span="12">
        <a-row :gutter="20">
          <a-col :span="24">
            <a-form-item field="name" :label="$t('role.name')" :rules="[{ required: true, message: $t('rule.required') }]">
              <template #extra>
                <div>{{ $t('role.name.tips') }}</div>
              </template>
              <a-input v-model="formData.name" :max-length="16" allow-clear show-word-limit :placeholder="$t('role.name.place')" />
            </a-form-item>
          </a-col>
          <a-col :span="24">
            <a-form-item field="permissionIds" :label="$t('role.permissionIds')">
              <template #extra>
                <div>{{ $t('role.permissionIds.tips') }}</div>
              </template>
              <span class="formSpan">{{ $t('role.permissionIds.place') }}</span>
            </a-form-item>
          </a-col>
          <a-col :span="24">
            <a-form-item field="remark" :label="$t('role.remark')">
              <a-input v-model="formData.remark" :max-length="64" allow-clear show-word-limit :placeholder="$t('role.remark.place')" />
            </a-form-item>
          </a-col>
        </a-row>
      </a-col>
      <a-col :span="12">
        <a-row :gutter="20">
          <a-col :span="24">
            <a-spin :loading="treeLoad" class="treeLe">
              <a-tree
                v-if="list[0]"
                ref="myTree"
                v-model:checked-keys="formData.permissionIds"
                v-model:half-checked-keys="formData.halfPermissionIds"
                size="large"
                :field-names="{ key: 'id' }"
                :checkable="true"
                :data="list"
                :show-line="true" />
            </a-spin>
          </a-col>
        </a-row>
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
              {{ $t('button.cancel') }}
            </a-button>
          </a-space>
        </div>
      </a-col>
    </a-row>
  </a-form>
</template>

<script lang="ts" setup>
import { onMounted, ref } from 'vue'
import type { Pop } from '@/utils/hooks/pop'
import { treeNode } from '@/api/comm'
import useLoad from '@/utils/hooks/load'
import { roleInit, roleAdd } from '@/api/plat/role'
import { permissionTree } from '@/api/plat/permission'
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
const { load: treeLoad, setLoad: setTreeLoad } = useLoad(false)
// 表单数据初始化
const formData = roleInit()
// const roleAdd = ref<FormInstance>();
// 提交数据
const submit = async ({ errors, values }: { errors: any; values: any }) => {
  if (load.value) return
  if (!errors) {
    setLoad(true)
    try {
      // const res = await roleAdd.value?.validate();
      await roleAdd(values)
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
// 树处理
const list = ref<treeNode[]>([])
// 权限树查询
async function treeQuery() {
  setTreeLoad(true)
  try {
    // 路由分页
    const res = await permissionTree({})
    list.value = res.data
    // 未选部分清理
  } catch (e) {
    list.value = []
  } finally {
    setTreeLoad(false)
  }
}
// 页面渲染
onMounted(() => {
  // Nothing
  setLoad(false)
  setTreeLoad(false)
  // 初始化树查询
  treeQuery()
})
</script>
