import Vue from 'vue'

import 'url-search-params-polyfill';
import ElementUI from 'element-ui';
import lang from 'element-ui/lib/locale/lang/en'
import locale from 'element-ui/lib/locale'
import 'element-ui/lib/theme-chalk/index.css';
import App from './App'
import router from './util/router'
import axios from "./util/axios"
import echarts from 'echarts'
import "./assets/css/main.css"
import "./assets/css/icon-type.css";
import bus from "./util/bus"
import search from "./util/search"
import cytoscape from 'cytoscape';

locale.use(lang)

Vue.config.prouctionTip = false

Vue.use(ElementUI)
Vue.component('v-chart',echarts)
Vue.prototype.$search = search
Vue.prototype.$echarts = echarts
Vue.prototype.$bus = bus
Vue.prototype.$axios = axios
Vue.prototype.$cytoscape = cytoscape;

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')

