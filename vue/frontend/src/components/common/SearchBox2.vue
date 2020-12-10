<template>
  <!-- 检索框 -->
  <div class="search-box">
    <el-input
        v-model="query"
        @keyup.enter.native="search"
        :placeholder="placeholder"
        prepend="query"
        >
      <el-button slot="append" icon="el-icon-search" @click="search"></el-button>
    </el-input>
    <!-- <el-button type="primary" @click="search_query">submit</el-button> -->
  </div>
</template>
<script>
export default {
  name: "SearchBox2",
  props: ['input'],
  data() {
    return {
      query:"",
      placeholder: "Whatever comes to your mind",
    }
  },
  methods: {
    // 发送检索请求
    search() {

      this.$router.push("path").catch(error => {
        if (error.name != "NavigationDuplicated") {
          throw error;
        }
      });
      this.$router.push(
          { name: 'institution', params: { id: this.query}})

    },
  },

}
</script>

<style>
.search-box {
  /* display: inline-block; */
}
</style>