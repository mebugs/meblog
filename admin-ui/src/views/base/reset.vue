<template>
  <a-layout class="layout reset">
    <div class="l_login">
      <div class="l_title">
        <h1>{{ $t('base.reset') }}</h1>
        <p>{{ $t('base.reset.tips') }}</p>
      </div>
      <a-form
        size="large"
        :label-col-props="{ span: 0 }"
        label-align="left"
        layout="vertical"
        class="form"
        :model="formData"
        @submit="submit">
        <a-row :gutter="20">
          <a-col :span="24">
            <a-form-item field="pwd" :rules="[{ required: true, message: $t('rule.required') }]">
              <a-input-password v-model="formData.pwd" :defaultVisibility="false" allow-clear :placeholder="$t('base.reset.pwd')">
                <template #prefix>
                  <icon-lock :size="20" />
                </template>
              </a-input-password>
            </a-form-item>
          </a-col>
          <a-col :span="24">
            <a-form-item field="pwdCfm" :rules="[{ required: true, message: $t('rule.required') }]">
              <a-input-password v-model="formData.pwdCfm" :defaultVisibility="false" allow-clear :placeholder="$t('base.reset.pwdCfm')">
                <template #prefix>
                  <icon-lock :size="20" />
                </template>
              </a-input-password>
            </a-form-item>
          </a-col>
          <a-col :span="24" style="text-align: center; margin-top: 20px">
            <a-button style="width: 200px; height: 40px; border-radius: 20px" type="primary" html-type="submit" :loading="load">
              <template #icon>
                <icon-thunderbolt />
              </template>
              {{ $t('button.reset') }}
            </a-button>
          </a-col>
        </a-row>
      </a-form>
    </div>
  </a-layout>
</template>
<script setup lang="ts">
import { ref } from 'vue'
import useLoad from '@/utils/hooks/load'
import { authReset } from '@/api/base/auth'
import { useRouter } from 'vue-router'
import { clearToken } from '@/utils/hooks/token'
// 加载中变量
const { load, setLoad } = useLoad(false)
// 登录表单
const formData = ref({ pwd: '', pwdCfm: '' })
// 路由对象
const router = useRouter()
// 提交登陆
const submit = async ({ errors, values }: { errors: any; values: any }) => {
  if (load.value) return
  if (!errors) {
    setLoad(true)
    try {
      // const res = await routerAdd.value?.validate();
      await authReset(values)
      clearToken() // 清理Token
      // 主动登出
      setTimeout(() => {
        router.push('/center/index')
      }, 500)
    } catch (err) {
      // DoNothing
    } finally {
      setLoad(false)
    }
  }
}
</script>

<style scoped lang="less">
.reset {
  overflow: hidden;
  position: relative;
  flex-direction: row;
  align-items: center;
  justify-content: center;
  * {
    transition: all 0.6s ease-in-out;
  }
  .l_fix {
    position: absolute;
    top: 0;
    right: 15px;
    width: 155px;
    display: flex;
    align-items: center;
    z-index: 2;
  }
  .l_login {
    margin-top: 20px;
    padding: 25px;
    width: 450px;
    height: 400px;
    background: var(--color-bg-1);
    border-radius: 10px;
    .l_title {
      width: 300px;
      height: 80px;
      margin: 20px auto;
      text-align: center;
      margin-bottom: 40px;
      position: relative;
      line-height: 30px;
    }
    .form * {
      transition: unset;
    }
    .arco-input-wrapper {
      padding: 5px 15px 5px 18px;
      border-radius: 23px;
    }
  }
}
</style>
