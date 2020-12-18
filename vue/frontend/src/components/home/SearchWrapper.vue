<template>
  <div class="home-search-pane">
    <search-box class="search-box" :type="type" :key="type"></search-box>

    <div class="radios">
      <el-radio-group v-model="type" @change="handleClick">
<!--        <el-radio :label=1>PAPER</el-radio>-->
        <el-radio :label=2>INSTITUTION</el-radio>
        <el-radio :label=3>BIOENTITY</el-radio>
        <el-radio :label=4>AUTHOR</el-radio>
      </el-radio-group>
    </div>
  </div>
</template>

<script>
import SearchBox from "../common/SearchBox.vue";

//  import SearchBox3 from "./../common/SearchBox3.vue"
export default {
  components: {
    "search-box": SearchBox,
  },
  data() {
    return {
      type: 0,
    };
  },
  methods: {
    handleClick() {
      console.log("clicked");
      console.log(this.type);
    },
  },
};
</script>

<style>
.home-search-pane {
  display: block;
  width: 100%;
  padding-bottom: -100px;
  padding-top: 15px;
}
.home-search-pane .search-box {
  padding-bottom:10px;

}

.home-search-pane .radios .el-radio-group{
  padding-left: 10px;
}
</style>
