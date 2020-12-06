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
        <div v-for="(card,index) in authors" :key="index" is="author-card" v-bind:card="card">
<!--            <author-card :card="card"></author-card>-->
        </div>
        </el-col>
        <div class="display-visual">
        <el-col :md="9" :sm="24" :xs="24" class="right">
          <p class="title">Top 10 Authors
          </p>
          <topAuthors></topAuthors>

        </el-col>
        </div>
      </el-row>
      <div class="pager">
        <el-pagination
            @current-change="handleCurrentChange" :current-page="currentPage"
            :page-size="pageSize" layout="prev, pager, next"
            :total="authors.length">
        </el-pagination>
      </div>
      </div>
    </div>
    
</template>

<script>
    import AuthorCard from "../components/general/AuthorCardNoBox";
    import topAuthors from "@/components/general/topAuthors";
    export default {
        name: "ItemAuthor",
      components: {
        'author-card': AuthorCard,
        "topAuthors":topAuthors

      },
        data() {
            return {
              currentPage:1,
              pageSize:10,
              tempList:[],
                authors:[{},{},{},{},{},{},{},{},{},{},{
                    'ForeName':'a',
                    'LastName':'b',
                    'lab':'COVID Lab',
                    'url':'google.com',
                    'institution':'msu',
                    'mail':'12345@qq.com',
                    'paper':123,
                    'quotes':14,
                }],
            }
        },
      mounted() {
      },
      watch: {
        paper:function (newData) {
          //console.log(newData)
          this.d = newData
          this.get_data()
          this.currentChangePage(this.tableData,1)
          this.currentChangePage1(this.cList,1)
        }
      },
      methods:{
          handleCurrentChange:function (currentPage) {
          this.currentPage = currentPage
          this.currentChangePage(this.tableData,currentPage)

        },
        currentChangePage(list,currentPage) {
          let from = (currentPage-1)*this.pageSize
          let to = currentPage*this.pageSize
          this.tempList = []
          for(;from<to;from++) {
            if(list[from]) {
              this.tempList.push(list[from]);
            }
          }
        },
        get_data() {
          // console.log(this.d)
          let t=[]
          for(let i=0;i<this.d.length;i++) {
            let article = {}
            article["title"] = this.d[i]["ArticleTitle"]
            article["author"] = this.d[i]["Authors"]
            article["journal"] = this.d[i]["Journal_Title"]
            article["url"] = this.d[i]["url"]
            // article["tweet"]=this.d[i]["tweet"]
            if (!this.d[i]["tweet"]) {
              // data attribute doesn't exist
              t.push({"article":article,"year":this.d[i]["PubYear"],"tweet":"0"})
            }
            else{
              t.push({"article":article,"year":this.d[i]["PubYear"],"tweet":this.d[i]["tweet"]})
            }
            // else{
            //   t.push({"article":article,"year":this.d[i]["PubYear"],"tweet":0})
            // }
            // console.log(t)

          }
          this.tableData = t
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