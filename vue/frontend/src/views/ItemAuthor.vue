<template>
    <div class="display-view container">
      <div  class="display-info ">
      <el-row :gutter="24">
<!--        <el-col :md="24" :lg="4"></el-col>-->
        <el-col :md="15" :sm="24" :xs="24">

      <el-card
          shadow="hover"
          class="display-query">
        <!-- 展示检索词 -->
        <p>what you search:</p>
        <p id="query">{{ query }}</p>
      </el-card>
<!--        <div v-for="(author,index) in authors" :key="index" is="author-card" v-bind:authors="authors">-->
            <author-card :author="author" v-loading="loading"></author-card>
<!--        </div>-->
        </el-col>
        <div class="display-visual">
        <el-col :md="9" :sm="24" :xs="24" class="right">
          <p class="title">Popular Authors
          </p>
          <topAuthors></topAuthors>

        </el-col>
        </div>
      </el-row>

      </div>
    </div>
    
</template>

<script>
    import AuthorCard from "../components/general/AuthorCardNoBox";
    import topAuthors from "@/components/general/topAuthors";
    import $axios from "@/util/axios";
    export default {
        name: "ItemAuthor",
      components: {
        'author-card': AuthorCard,
        "topAuthors":topAuthors

      },
        data() {
            return {
              loading:true,
              author:[],
              d:[]
            }
        },
      created() {
        this.get_data()

      },
      watch: {
        '$route': 'get_data'
      },

      methods:{


        get_data() {
          // console.log(this.d)
          this.loading = true
          $axios.get("/displayAuthorList/"+ this.$route.params.name).then(response=>{
            this.loading = false
            let d=response.data
            this.author=d
            // this.loading = false
            console.log("id is:",  this.$route.params.name)
            console.log("itemauthors", this.author)
          })
        },
        },


      computed: {
        query() {
          // 返回路由中的query，注意用的是$route
          return this.$route.params.name
        }
      },


    }
</script>

<style scoped>
.container {
  max-width: 1500px;
  width: 100%;
  margin: 0 auto;
  /*padding: .5em;*/
}
.right{
  margin: 0 auto;
  /*width: 100%;*/
}
.display-author{
  width: 100%;
}
.display-view {
  width: 100%;
  /* max-width: 970px; */
  margin: 0 auto;
}

.display-query {
  margin-bottom: 1.5em;
}
.display-query #query {
  /* text-transform: uppercase; */
  font-weight: bold;
  font-size: 1.5em;
}

.display-info {
  min-height: 30em;
  margin:auto;
}

.display-visual .title {
  text-transform: capitalize;
  margin-bottom: 1em;
  /* color: #409EFF; */
}
.display-query{
  width: 100%;
  margin-bottom: 1em;
  margin-top: 0;
}
.display-query #query {
  /* text-transform: uppercase; */
  font-weight: bold;
  font-size: 1.5em;
}
    .display-view {
        width: 80%;
        /* max-width: 970px; */
      /*margin-left: auto;*/
      /*margin-right: auto;*/
      margin-top: 1em;
      /*margin-left: 18em;*/
    }
    element.style {
      margin-top: 1em;
    }
.el-pagination {
  text-align: center;
  padding: 3rem 5rem;
}

img {
  display: inline;
  margin: 0 auto;
  height: 100%;
  width: 100%;
  /*border-radius: 50%;*/
}

div.link:hover{
  background-color:#e9e9e9;
}
.display-visual .title {
  text-transform: capitalize;
  margin-bottom: 1em;
  font-weight: bold;
  /* color: #409EFF; */
}

</style>