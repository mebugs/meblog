<template>
  <div class="head-btn-item">
    <a-dropdown trigger="click">
      <a-avatar :size="36" class="head-btn" :style="{ cursor: 'pointer' }">
        <img alt="avatar" :src="'/static/img/icon.png'" />
      </a-avatar>
      <template #content>
        <a-doption>
          <a-space @click="confirm = true">
            <icon-export />
            <span>
              {{ $t('base.logout') }}
            </span>
          </a-space>
        </a-doption>
      </template>
    </a-dropdown>
  </div>
  <!-- 刪除确认-->
  <a-modal v-model:visible="confirm" :width="400" :title="$t('base.logout')" @before-ok="confirmLogout">
    <div>{{ $t('base.logout.tips') }}</div>
  </a-modal>
</template>
<script lang="ts" setup>
import { ref } from 'vue'
import { userStore } from '@/store'
const confirm = ref(false)
async function confirmLogout() {
  // 登出处理
  const user = userStore()
  await user.logout()
}
</script>
