<script lang="tsx">
import { onUnmounted, defineComponent, computed, ref, resolveComponent, h } from 'vue'
import { useI18n } from 'vue-i18n'
import { appStore, userStore } from '@/store'
import { useRouter, RouteRecordRaw } from 'vue-router'
import { listenerRouteChange, removeRouteListener } from '@/utils/routeListener'
// 动态生成菜单
export default defineComponent({
  setup() {
    // 应用配置
    const app = appStore()
    // 折叠状态
    const collapsed = computed(() => {
      return app.collapsed
    })
    // 用户数据
    const user = userStore()
    const { initMenusTree } = user
    initMenusTree()
    const userTree = computed(() => {
      return user.menusTree
    })
    // 国际化
    const { t } = useI18n()
    // 选中和打开的菜单
    const openKeys = ref<string[]>([])
    const selectedKey = ref<string[]>([])
    // 路由
    const router = useRouter()
    // 跳转
    const goto = (item: RouteRecordRaw) => {
      // TODO Open external link

      // Trigger router change
      router.push({
        name: item.name
      })
    }
    // 生成子路由
    const renderSubMenu = () => {
      function travel(_route: RouteRecordRaw[], nodes = []) {
        if (_route) {
          _route.forEach((element) => {
            // 加载图标
            const icon = element?.meta?.icon ? h(resolveComponent(element?.meta?.icon)) : ''
            const node =
              element?.children && element?.children.length !== 0 ? (
                <a-sub-menu
                  key={element?.name}
                  v-slots={{
                    icon: () => icon,
                    title: () => t(element?.meta?.locale || '')
                  }}>
                  {travel(element?.children)}
                </a-sub-menu>
              ) : (
                <a-menu-item v-ripples key={element?.name} v-slots={{ icon: () => icon }} onClick={() => goto(element)}>
                  {t(element?.meta?.locale || '')}
                </a-menu-item>
              )
            nodes.push(node as never)
          })
        }
        return nodes
      }
      return travel(userTree.value)
    }
    // 绑定Key
    const findMenuOpenKeys = (target: string) => {
      const result: string[] = []
      let isFind = false
      const backtrack = (item: RouteRecordRaw, keys: string[]) => {
        if (item.name === target) {
          isFind = true
          result.push(...keys)
          return
        }
        if (item.children?.length) {
          item.children.forEach((el) => {
            backtrack(el, [...keys, el.name as string])
          })
        }
      }
      userTree.value.forEach((el: RouteRecordRaw) => {
        if (isFind) return // Performance optimization
        backtrack(el, [el.name as string])
      })
      return result
    }
    // 路由跳转事件
    listenerRouteChange((newRoute) => {
      // const { requiresAuth, activeMenu, hideInMenu } = newRoute.meta;
      const { hideInMenu } = newRoute.meta
      // if (requiresAuth && (!hideInMenu || activeMenu)) {
      if (!hideInMenu) {
        const menuOpenKeys = findMenuOpenKeys(newRoute.name as string)
        const keySet = new Set([...menuOpenKeys, ...openKeys.value])
        // 打开的Key
        openKeys.value = [...keySet]
        // 最后一个是被选
        selectedKey.value = [menuOpenKeys[menuOpenKeys.length - 1]]
      }
    }, true)
    // 注销监听
    onUnmounted(() => {
      removeRouteListener()
    })
    return () => (
      <a-menu
        mode={'vertical'}
        v-model:collapsed={collapsed.value}
        auto-open={true}
        selected-keys={selectedKey.value}
        auto-open-selected={true}
        collapsed-width={60}
        level-indent={0}
        style='height: 100%;width:100%;'>
        {renderSubMenu()}
      </a-menu>
    )
  }
})
</script>
