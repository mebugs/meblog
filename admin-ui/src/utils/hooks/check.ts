import { ref } from 'vue'
// 选择工具类
// 适用于弹窗列表选择
export function CheckList(formData: any, idKey: string, valKey: string) {
  // 路由添加层
  const check = ref(false)
  // 路由选择对象
  const selectData: any = ref({
    selectedKeys: [],
    selectedItems: []
  })
  // 初始化选择对象
  function init() {
    selectData.value = {
      selectedKeys: [],
      selectedItems: []
    }
    if (formData.value[valKey] && formData.value[valKey].length > 0) {
      // 重新赋值，消除响应性
      formData.value[valKey].forEach((i: any) => {
        selectData.value.selectedItems.push(i)
        selectData.value.selectedKeys.push(i.id)
      })
    }
  }
  // 打开选择画面
  function open() {
    init()
    check.value = true
  }
  // 取消选择画面
  function cancel() {
    init()
    check.value = false
  }
  // 确认选择
  function done() {
    // 与内部选择组件进行双向同步
    formData.value[valKey] = selectData.value.selectedItems
    formData.value[idKey] = selectData.value.selectedKeys
    return true
  }
  // 子组件选择(传入方法)
  function doSelect(items: any[], keys: any[]) {
    selectData.value.selectedItems = items
    selectData.value.selectedKeys = keys
  }

  // 快捷移除外层选择
  function removeQuick(id: number) {
    formData.value[valKey] = formData.value[valKey].filter((i: any) => {
      return i.id !== id
    })
    formData.value[idKey] = formData.value[idKey].filter((i: any) => {
      return i !== id
    })
  }
  return {
    check,
    selectData,
    open,
    cancel,
    done,
    doSelect,
    removeQuick
  }
}
