import appStore from './app'
import tabStore from './tab'
import userStore from './user'
import { createPinia } from 'pinia'

const pinia = createPinia()

export { appStore, tabStore, userStore }
export default pinia
