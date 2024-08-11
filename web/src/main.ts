import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from '@/App.vue'
import router from '@/router'
import '@/assets/main.css'
import 'virtual:uno.css'

const app = createApp(App)
const store = createPinia()

app.use(store)
app.use(router)

app.mount('#app')
