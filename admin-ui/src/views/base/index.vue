<template>
  <div class="layout index">
    <div class="wel" :class="{ done: done }">
      <div class="welogo">
        <img class="welbase" :src="logoImg" />
        <img class="welrun" :src="logoImg" />
      </div>
      <div class="welinfo">{{ $t(tips) }}</div>
      <div class="load"><icon-loading :size="36" /></div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import { isLogin } from '@/utils/hooks/token'
import { appStore } from '@/store'
// 应用配置
const app = appStore()
const logoImg = computed(() => {
  let d = app.theme == 'dark' ? '-dark' : ''
  return '/static/img/logo' + d + '.png'
})
// 路由对象
const router = useRouter()

// 默认欢迎语
const tips = ref('base.welcome')
const done = ref(false)
// 读取登陆数据
function readLogin() {
  // 读取登陆信息
  tips.value = 'base.tryReadLogin'
  setTimeout(() => {
    let toUrl = '/center/index'
    // 如果尚未登录提示，跳转登录画面
    if (isLogin()) {
      tips.value = 'base.tryReadLoginHave'
    } else {
      tips.value = 'base.tryReadLoginNull'
      toUrl = '/login'
    }
    setTimeout(() => {
      router.push(toUrl)
    }, 1500)
  }, 1500)
}
// 初始化方法
function init() {
  // 基本数据加载完成
  setTimeout(() => {
    done.value = true
    setTimeout(() => {
      readLogin()
    }, 1500)
  }, 500)
}
// 界面渲染后
onMounted(() => {
  init()
})
</script>
<style scoped lang="less">
.index {
  display: flex;
  flex-wrap: wrap;
  align-items: center;
  height: 100vh;
  width: 100vw;
  justify-content: center;
  color: var(--color-text-1);
  * {
    transition: all 0.4s ease-in-out;
  }
  .i_fix {
    position: absolute;
    top: 0;
    right: 15px;
    width: 155px;
    display: flex;
    align-items: center;
    z-index: 2;
  }
  .wel {
    text-align: center;
    line-height: 30px;
    opacity: 0;
    transform: scale(0.8);
    .welogo {
      position: relative;

      img {
        display: block;
      }
      .welbase {
        opacity: 1;
        animation: fafa 1s linear infinite;
      }
      .welrun {
        position: absolute;
        left: 0;
        top: 0;
        opacity: 0.4;
        transform: scale(1);
        animation: bili 1s linear infinite;
      }
    }
    .welinfo {
      margin: 10px 0 15px 0;
    }
  }
  .wel.done {
    opacity: 1;
    transform: scale(1);
  }
}
</style>
