import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

import { library } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

import { faNewspaper } from '@fortawesome/free-regular-svg-icons'
library.add(faNewspaper)

import { faThumbsUp } from '@fortawesome/free-regular-svg-icons'
library.add(faThumbsUp)

const app = createApp(App)
    .component('font-awesome-icon', FontAwesomeIcon);

app.use(router)

app.mount('#app')
