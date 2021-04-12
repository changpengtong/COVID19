<template>
  <div id="app">

    <el-container class='override-el-container'>
      <el-header class="override-el-header">
        <div class="top-nav">
          <top-nav></top-nav>
        </div>
      </el-header>
      <el-main class="override-el-main">
        <div class="content">
          <router-view class="view" />
        </div>
      </el-main>
      <el-footer class='m-footer'>
        <vfooter class='footer'></vfooter>
      </el-footer>
    </el-container>
    <back-top></back-top>
  </div>
</template>

<script>
  import TopNav from './components/common/TopNav'
  import Footer from './components/common/Footer'
 import './assets/css/override-element-ui.css'
import backtop from "./components/common/BackTop"
  export default {
    name: 'App',
    components: {
      "top-nav": TopNav,
      "Vfooter": Footer,
      "back-top": backtop
    },
  }

</script>

<style scoped>
  * {
    margin: 0;
    box-sizing: border-box;
  }

  #app {
    font-family: 'Avenir', Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    color: black;
    margin: 0;
    height: 100%;
    min-height:100%;
    padding:0;
  }

  .content {
    min-height: 100%;
  }

  .view {
    min-height: 100%;
  }

  .override-el-main.el-main {
    padding: 0px;
    height: 80%;
  }

  .override-el-header.el-header {
    padding: 0px;
    height: 15% !important;
  }

  .override-el-container {
    height: 100%;
  }
</style>