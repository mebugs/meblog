import { reactive } from 'vue'

// 权限树对象
export type treeNode = {
  id: number
  key: string
  title: string
  level: string
  expand: string // 拓展字段
  disabled?: boolean
  children?: treeNode[]
}

export function treeNodeInit() {
  return reactive<treeNode>({
    id: 0,
    key: '',
    title: '',
    level: '',
    expand: '',
    children: []
  })
}
