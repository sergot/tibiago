import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import Buefy from 'buefy';
import 'buefy/dist/buefy.css';

import App from './App.vue'
import router from './router'

const app = createApp(App)

app.use(createPinia())
app.use(router)

// @ts-ignore
app.use(Buefy)

app.mount('#app')
