<!--<template>-->
<!--  <div class="topnav">-->
<!--    <div class="_container">-->
<!--      <div class="framework-width">-->
<!--        <div id="icon">-->
<!--          <router-link to="/" tag="div">-->
<!--            <a href="#">COVID-19 PORTAL</a>-->
<!--          </router-link>-->
<!--        </div>-->

<!--      </div>-->
<!--    </div>-->
<!--  </div>-->
<!--</template>-->
<template>
  <div class="topnav">
    <div class="_container">
      <div class="framework-width">
        <div id="icon">
          <a href=" ">COVID-19 PORTAL</a >
        </div>
      </div>
    </div>
  </div>
</template>

<script >
// import SearchBox from "./../common/SearchBox.vue";
// import SearchWrapper from "../components/home/SearchWrapper";
export default {
  components: { },
  component: {
    // "search-box": SearchBox,
    // "search-wrapper": SearchWrapper,
  },
  computed: {
    ifSearchBox() {
      return this.$route.name == "displayInfo";
    },
    routeName() {
      return this.$route.name;
    },
  },
};
</script>

<style>
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
  height: 80px;
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
  margin-top:-15px;
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

.topnav .search-box {
  float: right;
  display: inline-block;
  /* height: 25px; */
  line-height: 40px;
  width: 85%;
}
</style>