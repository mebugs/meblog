import { ref } from 'vue'

export default function useLoad(initValue = false) {
  const load = ref(initValue)
  const setLoad = (value: boolean) => {
    load.value = value
  }
  const toggle = () => {
    load.value = !load.value
  }
  return {
    load,
    setLoad,
    toggle
  }
}
