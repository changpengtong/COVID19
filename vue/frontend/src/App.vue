<template>
  <div id="app">

    <el-container>
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
      <el-footer>
        <vfooter class="override-el-footer footer"></vfooter>
      </el-footer>
    </el-container>
    <back-top></back-top>
  </div>
</template>

<script>
  import TopNav from './components/common/TopNav'
  import Footer from './components/common/Footer'
  // 直接导入外部样式 覆盖element-ui样式
 import './assets/css/override-element-ui.css'
  import backtop from "./components/common/BackTop"
 // import './components/'
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
    /* color: #2c3e50; */
    color: black;
    margin: 0;
  }

  /*
  *  override the el-ui style
  */
  /* .el-container,
  .el-header,
  .el-main,
  .el-footer {
      width: 100%;
      margin: 0 auto;
      padding: 0;
  } */


  .content {
    min-height: 600px;
    /* background-color: #00BFB3; */
  }

  .view {
    min-height: 600px;
  }


  .footer {
    width: 100%;
    margin: 0;
    margin-top: .5em;
    padding: 0;
    min-height: 200px;
  }
  /* 修改局部样式 不影响其他的组件 */
  .override-el-main.el-main {
    padding:0px;
  }
  .override-el-header.el-header {
    padding:0px;
    height:auto !important;
  }
</style>