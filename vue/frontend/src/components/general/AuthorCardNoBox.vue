<template xmlns:img="http://www.w3.org/1999/html">
    <!-- 开发者名片 -->
  <div>
      <div v-for="(author,index) in tempList" :key="index">
        <el-card class="info-card" :body-style="{ padding: '0px' }" shadow="hover">
        <el-row :gutter="24">
            <el-col :md="24" :lg="6">
                <div class="info-image" >
                    <img src="../../assets/img/author.png">
                </div>
            </el-col>
            <el-col :md="24" :lg="16">
                <div class="info-text">
                    <div class="info-item">
                      <a :href="author.url">
                      <span id="name" class="info-item-title">{{author.name}}</span>
                      </a>
                    </div>
                    <!-- affilliation-->
                    <div class="info-item" v-if="author">
<!--                        <router-link :to=affiliation.institution.url >-->
                            <a :href="author.affiliationUrl">
                           <span class="institution" v-if="author.affiliation">{{author.affiliation}}</span>
<!--                        </router-link>-->
                            </a>
                      <span v-if="author.location">{{author.location}}, </span>

                    </div>
                    <!-- mail -->
                    <div class="info-item" v-if="author.email">
                        <span class="info-item-title">Email:</span><span>{{author.email}}</span>
                    </div>
                </div>
            </el-col>
        </el-row>
        </el-card>
      </div>
    <div class="pager">
      <el-pagination
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange" :current-page="currentPage"
          :page-sizes="[5,10,15,20,50,100]"
          :page-size="pageSize" layout="total, sizes, prev, pager, next, jumper"
          :total="authors.length">
      </el-pagination>
    </div>
  </div>
</template>

<script>
// import http from "../../../util/axios";


export default {
    props:['author'],
    components: {
      // "info-card": CardHelper
    },
    data() {
      return {
        currentPage:1,
        pageSize:10,
        tempList:[],
        authors:[],
        d:[]
      }
    },
    watch:{
      author:function (newData) {
        console.log("card,newdata",newData)
        this.d = newData
        this.get_data()
        this.currentChangePage(this.tableData,1)

        // this.currentChangePage(this.tableData,1)
      }
    },
    methods:{
      handleSizeChange: function (pageSize) {
        this.pageSize = pageSize
        this.handleCurrentChange(this.currentPage)
      },
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
        console.log("get data")
        let t=[]
        for(let i=0;i<this.d.length;i++) {
          this.d[i]["Affiliation"]=this.d[i]["Affiliation"].replace(/^,+/,"").replace(/,+$/,"")
            t.push({
              "email": this.d[i]["Email"],
              "name" : this.d[i]["Name"],
              "affiliation" : this.d[i]["Affiliation"]+",",
              "aid" :"",
              "affiliationUrl":"/COVID19/#/institution/"+ this.d[i]["Affiliation"],
              "url": "/#/author/"+this.d[i]["Name"]+"/institution/"+ this.d[i]["Affiliation"],
              "location":this.d[i]["Location"]
            })
        }
        this.authors = t
        this.tableData=t
        console.log("authorcard"),
          console.log(this.authors)
      },

    }

  }
</script>
<style scoped>

    .show-developer {
        margin: 1.5em auto;
    }
    .el-card{
      max-height: 200px;
    }

    .info-card {
        max-width: 50em;
         /*width: 50em;*/
        /* display: inline-block; */
        /*margin: 2em auto;*/
        padding: 1em 1em;
    }
    .center {
        margin: 0 auto;
    }
    .info-card {
         /*margin: 0 auto;*/
        /*padding: 1em;*/
    }
    .info-image img {
        object-fit: cover;
        border-radius: 50%;
        height: 80px;
        width: 80px;
        display: block;
        margin: auto;
        /*margin-top: 1em;*/

    }
    .info-text{
        margin: .5em 0;
      font-size: 1.2rem;

    }
    .info-item {
        margin: .5em 0;
      /*font-size: 1rem;*/

    }
    .info-item-title {
        font-weight: normal;
      font-size: 1.2rem;
        text-transform: capitalize;
        margin-right: .5em;
    }
    .institution{
        font-weight: normal;
      font-size: 1.2rem;
      text-transform: capitalize;
        margin-right: .5em;
        text-decoration:underline;
    }
    .words{
        white-space: pre-wrap;
        margin-right: 1rem;
    }
    #name {
        font-size: 1.2rem;
        margin-bottom: .5em;
        font-weight: 500;
    }
    .info-descript {
        color: rgba(0,0,0,.5);
        margin-bottom: .5em;
    }
    .el-pagination{
      text-align: center;
      padding: 3rem 5rem;
    }
</style>