<template>
    <div class="display-view container">
      <div  class="display-info ">
      <el-row :gutter="24">
        <el-col :md="15" :sm="24" :xs="24">

      <el-card
          shadow="hover"
          class="display-query">
        <p>what you search:</p>
        <p id="query">{{ query }}</p>
      </el-card>
            <author-card :author="author" v-loading="loading"></author-card>
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
          this.loading = true
          $axios.get("/displayAuthorList/"+ this.$route.params.name).then(response=>{
            this.loading = false
            let d=response.data
            this.author=d
          })
        },
        },


      computed: {
        query() {
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
}
.right{
  margin: 0 auto;
}
.display-author{
  width: 100%;
}
.display-view {
  width: 100%;
  margin: 0 auto;
}

.display-query {
  margin-bottom: 1.5em;
}
.display-query #query {
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
}
.display-query{
  width: 100%;
  margin-bottom: 1em;
  margin-top: 0;
}
.display-query #query {
  font-weight: bold;
  font-size: 1.5em;
}
    .display-view {
        width: 80%;
      margin-top: 1em;
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
}

div.link:hover{
  background-color:#e9e9e9;
}
.display-visual .title {
  text-transform: capitalize;
  margin-bottom: 1em;
  font-weight: bold;
}
</style>