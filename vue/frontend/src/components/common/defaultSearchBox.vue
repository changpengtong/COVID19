<template>
    <!-- 检索框 -->
    <div class="search-box">
      <el-input
          v-model="query"
          @keyup.enter.native="alert"
          :placeholder="placeholder"
      >
        <el-button slot="append" icon="el-icon-search" @click="alert"></el-button>
      </el-input>
    </div>
</template>
<script>

    export default {
        name: "defaultSearchBox",
        data: function (){
            return {
                query:'',
                placeholder: "Whatever comes to your mind",

            }
        },
        methods: {
          alert() {
            this.$alert('<strong>Please select the category before searching.</strong>',  {
              dangerouslyUseHTMLString: true
            });
          },

        }
    }
</script>

<style>
    .search-box {
        /* display: inline-block; */
    }
</style>