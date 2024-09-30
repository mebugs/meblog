import { ref } from 'vue'
import { Message } from '@arco-design/web-vue'
import i18n from '@/locale'
export type img = {
  baseUrl: string
  smallBaseUrl: Array<string> // 缩略图
  smallSize: Array<Number>
  sizeScale: Number // 高/宽比例 1:1
  chooseDown: boolean // 选择处理完成
}
export default function useImgs() {
  // 默认无需压缩，不限比例
  const imgObj = ref<img>({
    baseUrl: '',
    smallBaseUrl: [],
    smallSize: [1024],
    sizeScale: 0,
    chooseDown: false
  })
  const initImgQuick = (type: string) => {
    if (type == 'main') {
      initImg([1024, 512, 128], 0.56)
    }
    if (type == 'icon') {
      initImg([128], 1)
    }
  }
  const initImg = (smallSize: Array<Number>, sizeScale: Number) => {
    imgObj.value.smallSize = smallSize
    imgObj.value.sizeScale = sizeScale
  }
  const chooesImg = (e: Event) => {
    return new Promise(async (resolve, reject) => {
      if (e.target) {
        const files = (e.target as HTMLInputElement).files
        if (files && files[0]) {
          let file = files[0]
          if (window.FileReader && file) {
            if (window.FileReader && file) {
              const url = await readFile(file).catch((e) => {
                reject(e)
                return
              })
              imgObj.value.smallBaseUrl = []
              for (var i in imgObj.value.smallSize) {
                await resizeImg(url as string, file.type, Number(i)).catch((e) => {
                  reject(e)
                  return
                })
              }
              resolve(0)
              return
            }
          }
        }
      }
      Message.error(i18n.global.t('upimg.notallow'))
      reject(1)
      return
    })
  }
  const readFile = (file: File) => {
    return new Promise((resolve, reject) => {
      var reader = new FileReader()
      reader.readAsDataURL(file)
      //监听文件读取结束后事件
      reader.onloadend = function (ex: ProgressEvent<FileReader>) {
        // 原图
        if (ex.target) {
          resolve(ex.target.result)
          return
        }
      }
      reader.onerror = function () {
        Message.error(i18n.global.t('upimg.filereadfail'))
        reject(2) // 读取失败
        return
      }
    })
  }
  const resizeImg = (url: string | ArrayBuffer | null, type: string, index: number) => {
    return new Promise((resolve, reject) => {
      // 创建一个 Image 对象
      var image = new Image()
      // 绑定 load 事件处理器，加载完成后执行
      image.onload = function () {
        // 验证比例
        if (index === 0 && imgObj.value.sizeScale != 0) {
          let sco = (image.height / image.width).toFixed(2)
          if (sco !== imgObj.value.sizeScale.toFixed(2)) {
            Message.error(i18n.global.t('upimg.filesizefail'))
            reject(3)
            return // 比例不正确
          }
        }
        // 获取 canvas DOM 对象
        var canvas = document.createElement('canvas')
        var thisSize = imgObj.value.smallSize[index] as number
        //生成首页展示图
        if (image.width > thisSize) {
          // 宽度等比例缩放 *=
          image.height *= thisSize / image.width
          image.width = thisSize
        }
        // 获取 canvas的 2d 画布对象,
        var ctx = canvas.getContext('2d')
        if (ctx) {
          // canvas清屏，并设置为上面宽高
          ctx.clearRect(0, 0, canvas.width, canvas.height)
          // 重置canvas宽高
          canvas.width = image.width
          canvas.height = image.height
          // 将图像绘制到canvas上
          ctx.drawImage(image, 0, 0, image.width, image.height)
          // !!! 注意，image 没有加入到 dom之中
          var blob = canvas.toDataURL(type)
          if (index == 0) {
            imgObj.value.baseUrl = blob
          } else {
            imgObj.value.smallBaseUrl.push(blob)
          }
          resolve(0)
          return
        }
      }
      image.src = url as string
    })
  }
  return {
    imgObj,
    chooesImg,
    initImgQuick
  }
}
