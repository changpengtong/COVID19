// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'

import 'url-search-params-polyfill';
// 引入element-ui，注意放置位置
import ElementUI from 'element-ui';
import lang from 'element-ui/lib/locale/lang/en'
import locale from 'element-ui/lib/locale'
import 'element-ui/lib/theme-chalk/index.css';
import App from './App'
import router from './util/router'
// 引入 axios
// import axios from 'axios'
import axios from "./util/axios"

// 引入echarts
import echarts from 'echarts'
import "./assets/css/main.css"
import "./assets/css/icon-type.css";
// 引入bus
import bus from "./util/bus"
import search from "./util/search"



locale.use(lang)

// import "./util/eChart"


Vue.config.prouctionTip = false

Vue.use(ElementUI)
Vue.component('v-chart',echarts)
Vue.prototype.$search=search
Vue.prototype.$echarts=echarts
Vue.prototype.$bus=bus
Vue.prototype.$axios=axios


/* eslint-disable no-new */
/**new Vue({
  el: 'app',
  router,
  components: { App },
  template: '<App/>'
})**/
new Vue({
  router,
  render: h => h(App)
}).$mount('#app')

