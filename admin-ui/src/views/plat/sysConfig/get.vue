<template>
  <a-spin :loading="load">
    <a-form size="large" label-align="left" class="form" layout="vertical" :model="formData">
      <a-row :gutter="20">
        <a-col :span="24">
          <a-row :gutter="20">
            <a-col :span="12">
              <a-form-item field="loginSwitch" :label="$t('sysConfig.loginSwitch')">
                <template #extra>
                  <div>{{ $t('sysConfig.loginSwitch.tips') }}</div>
                </template>
                <span class="formSpan">{{ dictMap.openStatus[formData.loginSwitch] }}</span>
              </a-form-item>
            </a-col>
          </a-row>
        </a-col>
        <a-col :span="24" class="formGroup" v-if="formData.loginSwitch == '0'">
          <a-row :gutter="20">
            <a-col :span="12">
              <a-form-item field="loginNum" :label="$t('sysConfig.loginNum')">
                <span class="formSpan">{{ formData.loginNum }}</span>
              </a-form-item>
            </a-col>
          </a-row>
        </a-col>
        <a-col :span="24">
          <a-row :gutter="20">
            <a-col :span="12">
              <a-form-item field="loginFailSwitch" :label="$t('sysConfig.loginFailSwitch')">
                <template #extra>
                  <div>{{ $t('sysConfig.loginFailSwitch.tips') }}</div>
                </template>
                <span class="formSpan">{{ dictMap.openStatus[formData.loginFailSwitch] }}</span>
              </a-form-item>
            </a-col>
          </a-row>
        </a-col>
        <a-col :span="24" class="formGroup" v-if="formData.loginFailSwitch == '0'">
          <a-row :gutter="20">
            <a-col :span="12">
              <a-form-item field="loginFailNum" :label="$t('sysConfig.loginFailNum')">
                <span class="formSpan">{{ formData.loginFailNum }}</span>
              </a-form-item>
            </a-col>
            <a-col :span="12">
              <a-form-item field="loginFailUnit" :label="$t('sysConfig.loginFailUnit')">
                <span class="formSpan"> {{ dictMap.timeUnit[formData.loginFailUnit] }} </span>
              </a-form-item>
            </a-col>
            <a-col :span="12">
              <a-form-item field="loginFailLockNum" :label="$t('sysConfig.loginFailLockNum')">
                <span class="formSpan">{{ formData.loginFailLockNum }}</span>
              </a-form-item>
            </a-col>
            <a-col :span="12">
              <a-form-item field="loginFailLockUnit" :label="$t('sysConfig.loginFailLockUnit')">
                <span class="formSpan">{{ dictMap.timeUnit[formData.loginFailLockUnit] }}</span>
              </a-form-item>
            </a-col>
            <a-col :span="12">
              <a-form-item field="loginFailTryNum" :label="$t('sysConfig.loginFailTryNum')">
                <span class="formSpan">{{ formData.loginFailTryNum }}</span>
              </a-form-item>
            </a-col>
          </a-row>
        </a-col>
        <a-col :span="24">
          <a-row :gutter="20">
            <a-col :span="12">
              <a-form-item field="logoutSwitch" :label="$t('sysConfig.logoutSwitch')">
                <template #extra>
                  <div>{{ $t('sysConfig.logoutSwitch.tips') }}</div>
                </template>
                <span class="formSpan">{{ dictMap.openStatus[formData.logoutSwitch] }}</span>
              </a-form-item>
            </a-col>
          </a-row>
        </a-col>
        <a-col :span="24" class="formGroup" v-if="formData.logoutSwitch == '0'">
          <a-row :gutter="20">
            <a-col :span="12">
              <a-form-item field="logoutNum" :label="$t('sysConfig.logoutNum')">
                <span class="formSpan">{{ formData.logoutNum }}</span>
              </a-form-item>
            </a-col>
            <a-col :span="12">
              <a-form-item field="logoutUnit" :label="$t('sysConfig.logoutUnit')">
                <span class="formSpan">{{ dictMap.timeUnit[formData.logoutUnit] }}</span>
              </a-form-item>
            </a-col>
          </a-row>
        </a-col>
        <a-col :span="24">
          <a-divider />
          <div class="doBtn">
            <a-space>
              <a-button
                v-permission="'PlatConfigEdit'"
                size="large"
                type="primary"
                :loading="load"
                @click="pop.open('edit', 1, $t('sysConfig.edit'), $t('sysConfig.system'), {}, get)">
                <template #icon>
                  <icon-edit />
                </template>
                {{ $t('button.edit') }}
              </a-button>
              <a-button v-permission="'PlatConfigView'" size="large" :loading="load" @click="get()">
                <template #icon>
                  <icon-sync />
                </template>
                {{ $t('button.sync') }}</a-button
              >
            </a-space>
          </div>
        </a-col>
      </a-row>
    </a-form>
  </a-spin>
</template>

<script lang="ts" setup>
import { onMounted, watch, ref } from 'vue'
import type { Pop } from '@/utils/hooks/pop'
import useLocale from '@/utils/hooks/locale'
import useLoad from '@/utils/hooks/load'
import { dictRead } from '@/api/plat/dict'
import { sysConfigInit, sysConfigGet } from '@/api/plat/sysConfig'
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
// 当前语言
const { currentLocale } = useLocale()
// 加载中变量
const { load, setLoad } = useLoad(false)
// 表单数据初始化
const formData = sysConfigInit()
async function get() {
  setLoad(true)
  try {
    const res = await sysConfigGet()
    formData.value = res.data
    return
  } catch (err) {
    // DoNothing CommonPopUp
  } finally {
    setLoad(false)
  }
}
// 初始化字典对象
const dictList = ref({ timeUnit: [], openStatus: [] })
const dictMap = ref({ timeUnit: {} as any, openStatus: {} as any })
// 字段初始化
async function dictInit() {
  // 指定字典Key
  await dictRead({ groupKeys: ['timeUnit', 'openStatus'] }).then((r) => {
    dictList.value = r.data.list
    dictMap.value = r.data.map
    props.pop.dictList = dictList
    props.pop.dictMap = dictMap
  })
}
// 初始化
function init() {
  dictInit()
  get()
}
// 页面渲染
onMounted(() => {
  init()
})
// 语言监听
watch(currentLocale, (n, o) => {
  if (n !== o) {
    dictInit()
  }
})
</script>
