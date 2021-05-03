<template xmlns:img="http://www.w3.org/1999/html">
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
                  <el-scrollbar wrap-style="max-height:200px">
                  <div class="info-item">
                      <a :href="author.url">
                      <span id="name" class="info-item-title">{{author.name}}</span>
                      </a>
                    </div>
                    <!-- affilliation-->
                    <div class="info-item" v-if="author">
                            <a :href="author.affiliationUrl">
                           <span class="institution" v-if="author.affiliation">{{author.affiliation}}</span>
                            </a>
                    </div>
                    <div class="info-item" v-if="author.email">
                        <span class="info-item-title"></span><span>{{author.location}}</span>
                    </div>
                    <div class="info-item" v-if="author.email">
                        <span class="info-item-title">Email:</span><span>{{author.email}}</span>
                    </div>
                 </el-scrollbar>
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
export default {
    props:['author'],
    components: {
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
        let t=[]
        for(let i=0;i<this.d.length;i++) {
            t.push({
              "email": this.d[i]["Email"],
              "name" : this.d[i]["Name"],
              "affiliation" : this.d[i]["Affiliation"],
              "aid" :this.d[i]["aid"],
              "affiliationUrl":"/#/institution/"+ this.d[i]["institution_id"],
              "url": "/#/author/"+this.d[i]["aid"],
              "location": this.d[i]["Location"]
            })
        }
        this.authors = t
        this.tableData=t
      },

    }

  }
</script>
<style scoped>
    .show-developer {
        margin: 1.5em auto;
    }
    .el-card{
      max-height: 250px;
    }

    .info-card {
        max-width: 50em;
        padding: 1em 1em;
    }
    .center {
        margin: 0 auto;
    }
    .info-card {
    }
    .info-image img {
        object-fit: cover;
        border-radius: 50%;
        height: 80px;
        width: 80px;
        display: block;
        margin: auto;
    }
    .info-text{
        margin: .5em 0;
      font-size: 17px;

    }
    .info-item {
        margin: .5em 0;
    }
    .info-item-title {
        font-weight: normal;
      font-size: 17px;
        text-transform: capitalize;
        margin-right: .5em;
    }
    .institution{
        font-weight: normal;
      font-size: 17px;
      text-transform: capitalize;
        margin-right: .5em;
        text-decoration:underline;
    }
    .words{
        white-space: pre-wrap;
        margin-right: 1rem;
    }
    #name {
        font-size: 17px;
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