import Vditor from 'vditor'
import { reactive, ref } from 'vue'

// mardown编辑器对象
export interface markdown {
  md: string | undefined // MD 基础文本
  html: string | undefined // HTML 解析结果
  outline: string | undefined // 大纲 HTML
}

// 编辑器插件初始化
export default function vditor(step: any, nodeId: string) {
  const md: markdown = reactive<markdown>({
    md: '',
    html: '',
    outline: ''
  })
  const vd = ref<Vditor>({} as any)
  // 初始化VD
  const getNew = () => {
    vd.value = new Vditor(nodeId, {
      toolbar,
      lang: 'zh_CN',
      mode: 'ir',
      value: '',
      icon: 'material',
      height: window.innerHeight - 200,
      cache: {
        enable: false
      },
      outline: {
        enable: true,
        position: 'right'
      },
      debugger: true,
      typewriterMode: true,
      placeholder: "❤️你好！米虫！开始编写吧！❤️推荐您Ctrl+'或点击↕按钮打开全屏专注编写↗",
      cdn: '/static/lib/vditor',
      preview: {
        maxWidth: 920,
        actions: [],
        parse(element: HTMLElement) {
          // 延时处理
          // setTimeout(() => {
          //   md.html = element.innerHTML
          //   step.value = 1
          // }, 3000)
        },
        markdown: {
          toc: true,
          mark: true,
          footnotes: true,
          autoSpace: true
        },
        theme: {
          current: 'ant-design',
          path: '/static/lib/vditor'
        },
        hljs: {
          enable: true,
          lineNumber: true,
          defaultLang: '',
          style: 'vs'
        }
      },
      toolbarConfig: {
        pin: true
      },
      counter: {
        enable: true,
        type: 'text'
      },
      tab: '\t',
      hint: {
        emoji: emoji
      }
      // upload: {
      //     accept: 'image/*,.mp3, .wav, .rar',
      //     token: 'test',
      //     url: '/api/upload/editor',
      //     linkToImgUrl: '/api/upload/fetch',
      //     filename(name) {
      //         return name.replace(/[^(a-zA-Z0-9\u4e00-\u9fa5\.)]/g, '').replace(/[\?\\/:|<>\*\[\]\(\)\$%\{\}@~]/g, '').replace('/\\s/g', '')
      //     },
      // },
    })
  }
  // 模拟预览
  const toPreview = () => {
    var elements = document.querySelectorAll('#' + nodeId + ' .vditor-toolbar div button')
    var lastElement = elements[elements.length - 1]
    lastElement.dispatchEvent(new MouseEvent('click'))
  }
  // 获取提交结果
  const getResponse = () => {
    // 读取预览HTML
    md.outline = document.querySelector('#' + nodeId + ' .vditor-outline')?.innerHTML
    md.html = document.querySelector('#' + nodeId + ' .vditor-preview')?.innerHTML
    md.md = vd.value.getValue()
    if (md.html) {
      // 替换处理
      // Copy方法替换
      // 特殊符号替换
      md.html = md.html.replaceAll(
        "this.previousElementSibling.select();document.execCommand('copy');this.setAttribute('aria-label', '已复制')",
        'copyCode(this)'
      )
      md.html = md.html.replaceAll('<p><img src="', '<p class="mimg"><img src="')
    }
  }
  return {
    vd,
    md,
    getNew,
    toPreview,
    getResponse
  }
}

// 编辑器工具栏
const toolbar = [
  'emoji',
  'headings',
  'bold',
  'italic',
  'strike',
  'link',
  '|',
  'list',
  'ordered-list',
  'check',
  'outdent',
  'indent',
  '|',
  'quote',
  'line',
  'code',
  'inline-code',
  'insert-before',
  'insert-after',
  '|',
  'upload',
  'table',
  '|',
  'undo',
  'redo',
  '|',
  'fullscreen',
  'preview'
  // 'edit-mode',
  // 'content-theme',
  // 'code-theme',
  // 'export',
  // {
  //     name: 'more',
  //     toolbar: [
  //         'info',
  //         'help',
  //     ],
  // }
]

// 表情对象
const emoji = {
  '100': '💯',
  anchor: '⚓️',
  anger: '💢',
  boom: '💥',
  dart: '🎯',
  fire: '🔥',
  gem: '💎',
  construction: '🚧',
  crossed_swords: '⚔️',
  crown: '👑',
  art: '🎨',
  key: '🔑',
  book: '📖',
  heavy_check_mark: '✔️',
  heavy_multiplication_x: '✖️',
  pushpin: '📌',
  bell: '🔔',
  heart: '❤️',
  broken_heart: '💔',
  bulb: '💡',
  alarm_clock: '⏰',
  balance_scale: '⚖️',
  basketball: '🏀',
  blossom: '🌼',
  bomb: '💣',
  closed_umbrella: '🌂',
  lemon: '🍋',
  cheese: '🧀',
  cherries: '🍒',
  beers: '🍻',
  banana: '🍌',
  apple: '🍎',
  reminder_ribbon: '🎗',
  ring: '💍',
  trophy: '🏆',
  underage: '🔞',
  zap: '⚡️',
  zzz: '💤',
  airplane: '✈️',
  rocket: '🚀',
  car: '🚗',
  bus: '🚌',
  sunny: '☀️',
  sparkles: '✨',
  star: '⭐️',
  calendar: '📆',
  camera: '📷',
  computer: '💻',
  shit: '💩',
  christmas_tree: '🎄',
  ferris_wheel: '🎡',
  bee: '🐝',
  cactus: '🌵',
  fallen_leaf: '🍂',
  strawberry: '🍓',
  balloon: '🎈',
  '+1': '👍',
  '-1': '👎',
  crossed_fingers: '🤞',
  fist: '✊',
  point_right: '👉',
  raised_hand_with_fingers_splayed: '🖐',
  vulcan_salute: '🖖',
  boy: '👦',
  girl: '👧',
  bear: '🐻',
  cat: '🐱',
  cow: '🐮',
  dog: '🐶',
  fox_face: '🦊',
  lion: '🦁',
  mouse: '🐭',
  panda_face: '🐼',
  pig: '🐷',
  rabbit: '🐰',
  frog: '🐸',
  skull: '💀',
  kissing_heart: '😘',
  smile: '😄',
  smirk: '😏',
  wink: '😉',
  heart_eyes: '😍',
  laughing: '😆',
  yum: '😋',
  angry: '😠',
  unamused: '😒',
  worried: '😟',
  cry: '😢',
  face_with_head_bandage: '🤕',
  persevere: '😣',
  triumph: '😤',
  confused: '😕',
  confounded: '😖',
  expressionless: '😑',
  frowning_face: '☹️',
  zipper_mouth_face: '🤐',
  face_with_thermometer: '🤒',
  thinking: '🤔',
  astonished: '😲',
  dizzy_face: '😵',
  cold_sweat: '😰',
  fearful: '😨',
  flushed: '😳',
  grimacing: '😬',
  stuck_out_tongue: '😛',
  sleeping: '😴',
  sneezing_face: '🤧',
  sob: '😭',
  tired_face: '😫',
  weary: '😩',
  hushed: '😯',
  joy: '😂'
}
