import { DirectiveBinding } from 'vue'
import { userStore } from '@/store'

// 权限检查指令 替换为字符串指令
function checkPermission(el: HTMLElement, binding: DirectiveBinding) {
  const { value } = binding
  if (!value) {
    return
  }
  const user = userStore()
  const { permissions } = user
  // 检查用户是否具备这个元素的访问权限
  const hasPermission = permissions.includes(value)
  if (!hasPermission && el.parentNode) {
    el.parentNode.removeChild(el)
  }
}

// 执行时机，初始化或更新时
export default {
  mounted(el: HTMLElement, binding: DirectiveBinding) {
    checkPermission(el, binding)
  },
  updated(el: HTMLElement, binding: DirectiveBinding) {
    checkPermission(el, binding)
  }
}
