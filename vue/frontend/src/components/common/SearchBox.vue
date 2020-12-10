<template>
    <!-- 检索框 -->
    <div class="search-box">
      <el-input
          v-model="query"
          @keyup.enter.native="search"
          :placeholder="placeholder"
          :fetch-suggestions="querySearchAsync"
          @select="handleSelect"
      >
        <el-button slot="append" icon="el-icon-search" @click="search"></el-button>
      </el-input>
    </div>
</template>
<script>

    export default {
        name: "SearchBox",
        data: function (){
            return {
                query:'',
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
                this.$router.push({
                    path: "/displayInfo",
                    query: {
                      // keyword: this.query
                      query: this.query

                    }

                });


            }
        }
    }
</script>

<style>
    .search-box {
        /* display: inline-block; */
    }
</style>