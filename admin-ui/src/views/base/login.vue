<template>
  <a-layout class="layout login">
    <div class="l_bg" :class="{ l_bg_dark: theme == 'dark' }" />
    <div class="l_fix head-wrapper"></div>
    <a-layout class="l_pan">
      <div class="l_mod">
        <img class="lm_sd" src="/static/img/login-left-s-dark.png" />
        <img class="lm_sd" :class="{ lm_hide: theme == 'dark' }" src="/static/img/login-left-s.png" />
        <img class="lm_bt" :class="{ lm_bdark: theme == 'dark' }" src="/static/img/login-left-btn.png" />
      </div>
    </a-layout>
    <a-layout class="l_pan">
      <div class="l_login">
        <div class="l_logo">
          <img src="/static/img/logo-dark.png" />
          <img :class="{ ll_hide: theme == 'dark' }" src="/static/img/logo.png" />
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
              <a-form-item field="acc" :rules="[{ required: true, message: $t('rule.required') }]">
                <a-input v-model="formData.acc" allow-clear :placeholder="$t('base.login.acc')">
                  <template #prefix>
                    <icon-user :size="20" />
                  </template>
                </a-input>
              </a-form-item>
            </a-col>
            <a-col :span="24">
              <a-form-item field="pwd" :rules="[{ required: true, message: $t('rule.required') }]">
                <a-input-password v-model="formData.pwd" :defaultVisibility="false" allow-clear :placeholder="$t('base.login.pwd')">
                  <template #prefix>
                    <icon-lock :size="20" />
                  </template>
                </a-input-password>
              </a-form-item>
            </a-col>
            <a-col :span="24">
              <a-form-item>
                <a-checkbox v-model="rem"> {{ $t('base.login.rem') }} </a-checkbox>
              </a-form-item>
            </a-col>
            <a-col :span="24" style="text-align: center">
              <a-button style="width: 200px; height: 40px; border-radius: 20px" type="primary" html-type="submit" :loading="load">
                <template #icon>
                  <icon-thunderbolt />
                </template>
                {{ $t('button.login') }}
              </a-button>
            </a-col>
          </a-row>
        </a-form>
      </div>
    </a-layout>
  </a-layout>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import useLoad from '@/utils/hooks/load'
import { appStore } from '@/store'
import { authLogin } from '@/api/base/auth'
import { setToken, getAcc, setAcc, clearAcc } from '@/utils/hooks/token'
import { useRouter } from 'vue-router'
// 应用配置
const app = appStore()
const theme = computed(() => {
  return app.theme
})
// 加载中变量
const { load, setLoad } = useLoad(false)
// 登录表单
const formData = ref({ acc: '', pwd: '' })
// 获取记忆
const remData = getAcc()
// 记住密码
const rem = ref(false)
if (remData) {
  formData.value = remData
  rem.value = true
}
// 路由对象
const router = useRouter()
// 提交登陆
const submit = async ({ errors, values }: { errors: any; values: any }) => {
  if (load.value) return
  if (!errors) {
    setLoad(true)
    try {
      // const res = await routerAdd.value?.validate();
      const res = await authLogin(values)
      if (res != null && res.data != null) {
        let lRes = res.data
        let toUrl = '/center/index'
        if (lRes.re) {
          toUrl = '/reset' // 前往密码重置
        }
        setToken(lRes.tk)
        // 记住账密
        if (rem.value) {
          setAcc(values)
        } else {
          clearAcc()
        }
        setTimeout(() => {
          router.push(toUrl)
        }, 500)
      }
    } catch (err) {
      // DoNothing
    } finally {
      setLoad(false)
    }
  }
}
</script>

<style scoped lang="less">
.login {
  overflow: hidden;
  position: relative;
  flex-direction: row;
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
  .l_bg {
    position: absolute;
    height: 300%;
    width: 150%;
    top: -150%;
    right: 45%;
    background-image: linear-gradient(-90deg, #47b9e3 0, #164792 100%);
    border-radius: 50%;
  }
  .l_bg_dark {
    transform: rotate(180deg);
  }
  .l_pan {
    height: 100%;
    width: 50%;
    justify-content: center;
    align-items: center;
    z-index: 1;
    .l_mod {
      width: 500px;
      height: 600px;
      position: relative;
      .lm_sd {
        position: absolute;
        z-index: 2;
        top: -5%;
        transform: scale(0.9);
        animation: sdd 5s linear infinite;
        opacity: 1;
      }
      .lm_hide {
        opacity: 0;
      }
      .lm_bt {
        position: absolute;
        bottom: 20px;
        z-index: 1;
        animation: btt 5s linear infinite;
      }
      .lm_bdark {
        filter: brightness(80%) contrast(80%);
      }
    }
    .l_login {
      margin-top: 40px;
      padding: 25px;
      width: 450px;
      height: 450px;
      background: var(--color-bg-1);
      border-radius: 10px;
      .l_logo {
        width: 300px;
        height: 80px;
        margin: 20px auto;
        margin-bottom: 40px;
        position: relative;
        img {
          position: absolute;
          opacity: 1;
        }
        .ll_hide {
          opacity: 0;
        }
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
}

@keyframes btt {
  50% {
    transform: scale(0.85);
    opacity: 1;
  }
}
@keyframes sdd {
  50% {
    top: -10%;
    transform: scale(1);
  }
}
</style>
