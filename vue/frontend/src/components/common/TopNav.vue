
<template>
  <div class="topnav">
    <div class="_container">
      <div class="framework-width">
        <el-row :gutter="20">
          <el-col :span="15">
        <div id="icon">
          <a href="/COVID19/">COVID-19 PORTAL</a >
        </div>
          </el-col>
          <el-col :span="9">
        <transition name="el-zoom-in-top">
          <search-wrapper id="search-wrapper" v-if="ifSearchBox"></search-wrapper>
        </transition>
          </el-col>
        </el-row>
      </div>
    </div>
  </div>
</template>

<script scoped>
// import SearchBox from "./../common/SearchBox.vue";
import SearchWrapper from "../home/SearchWrapper";
export default {
  components: {
    // "search-box": SearchBox,
    "search-wrapper": SearchWrapper,
  },
  computed: {
    ifSearchBox() {
      console.log("page: "+this.$route.name);
      return this.$route.name !== 'home';
    },
    routeName() {
      return this.$route.name;
    },
  },
};
</script>

<style>
/*<!--search box wrapper-->*/
.home-search-pane {
  display: block;
  width: 80%;
  padding-bottom: -100px;
  padding-top: 15px;
}

.home-search-pane .search-box {
  padding-bottom: 10px;
}

.home-search-pane .radios .el-radio-group {
  padding-top: 13px;
  padding-left: 15%;
}
.search-box .search-box-auto {
  position: center;
  width: 30em;
}
/*serach box*/
.search-box .el-col-6 {
  min-width: 100px;
  padding-right: 15px;
}

.search-box .el-input__inner {
  border-radius: 20px 20px;
}

.search-box .el-icon-search {
  margin-top: 4px;
  font-size: 17px;
  margin-right: 10px;
  border: 0px;
}

* {
  box-sizing: border-box;
}

/*
    导航栏样式
     */
.topnav {
  overflow: hidden;
  /* background-color: #e9e9e9; */
  background-color: #ffffff;
  /* box-shadow: 0px 1px 2px #aaaaaa; */
  /* position: fixed; */
  border-bottom: 1px solid #e1e1e1;
  top: 0;
  height: 150px;
  width: 100%;
  padding-top: 15px;
}

/*
     * ---------------------------------------------
     *  暂时的网站标识
     */
#icon {
  float: left;
  display: block;
  line-height: 80px;
  font-weight: bold;
  margin-top:15px;
  margin-left: 15px;
}
#icon a {
  font-size: 35px;
}

.topnav a,
.topnav .link {
  float: right;
  display: inline-block;
  color: #409eff;
  font-weight: bold;
  text-decoration: none;
  text-align: center;
  padding: 20px 16px;
  font-size: 1em;
  line-height: 40px;
  border-bottom: 4px solid transparent;
}

.topnav .link:hover {
  /* border-bottom: 4px solid #409EFF; */
  border-bottom-color: #409eff;
}

/*.topnav .search-wrapperr{*/
/*  float: right;*/
/*  display: inline-block;*/
/*  !* height: 25px; *!*/
/*  line-height: 40px;*/
/*  width: 55%;*/
/*  border: #503C0E;*/
/*}*/
</style>