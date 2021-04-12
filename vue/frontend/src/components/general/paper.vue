<template>
    <el-tabs :tab-position="tabPostion" stretch="true" class="tab">
        <el-tab-pane label="Covid-19 Papers">
            <div>
                <el-table
                        :data="tempList"
                        style="width: 100%"
                        class="my-table"
                        :default-sort="{prop:'year',order:'ascending'}"
                        @sort-change="sort_change"
                        :header-cell-style="{
                            'padding-left': '1rem',
                        }"
                >
                    <el-table-column
                            prop="article"
                            label=TITLE
                            sortable="custom"
                            align="left"
                            width="600"
                    >
                        <template slot-scope="scope">
                            <div v-for="(item,index) in scope.row.article" :key="index"  style="margin-left: 1rem">
                                <span v-if="index==='title'" style="font-size: 17px">
                                    <a :href="scope.row.article.url">
                                    {{item}}
                                    </a>
                                </span>
                                <span v-if="index==='author'" style="font-size: small">
                                    {{item}}
                                </span>
                                <span v-if="index==='journal'" style="font-style: italic; font-size: small">
                                    {{item}}
                                </span>
                            </div>
                        </template>
                    </el-table-column>

                    <el-table-column
                            prop="year"
                            label="YEAR"
                            sortable="custom"
                            align="center"
                            width="100"
                    >
                    </el-table-column>
                  <el-table-column
                      prop="tweet"
                      label="TWEETS"
                      sortable="custom"
                      align="center"
                      width="130"
                  >
                  </el-table-column>
                </el-table>
                <el-pagination
                        @size-change="handleSizeChange"
                        @current-change="handleCurrentChange" :current-page="currentPage"
                        :page-sizes="[5,10,15,20,50,100]"
                        :page-size="pageSize" layout="total, sizes, prev, pager, next, jumper"
                        :total="tableData.length">
                </el-pagination>
            </div>
        </el-tab-pane>
        <el-tab-pane label="Clinical Trials">
            <div>
                <el-table
                        :data="tempClist"
                        style="width: 100%"
                        class="my-table"
                        :default-sort="{prop:'year',order:'ascending'}"
                        @sort-change="sort_change"
                        :header-cell-style="{
                            'padding-left': '1rem',
                        }"
                >
                    <el-table-column
                            prop="article"
                            label=TITLE
                            sortable="custom"
                            align="left"
                            width="600"
                    >
                        <template slot-scope="scope">
                            <div v-for="(item,index) in scope.row.article" :key="index"  style="margin-left: 1rem">
                                <span v-if="index==='title'" style="font-size: 17px">
                                    <a :href="scope.row.article.url">
                                    {{item}}
                                    </a>
                                </span>
                                <span v-if="index==='researchers'" style="font-size: small">
                                    {{item}}
                                </span>
                                <span v-if="index==='sponsors'" style="font-style: italic; font-size: small">
                                    {{item}}
                                </span>
                            </div>
                        </template>
                    </el-table-column>

                    <el-table-column
                            prop="year"
                            label="YEAR"
                            sortable="custom"
                            align="center"
                            width="100"
                    >
                    </el-table-column>
                </el-table>
                <el-pagination
                        @size-change="handleSizeChange1"
                        @current-change="handleCurrentChange1" :current-page="currentPage1"
                        :page-sizes="[5,10,15,20,50,100]"
                        :page-size="pageSize" layout="total, sizes, prev, pager, next, jumper"
                        :total="clinicalData.length">
                </el-pagination>
            </div>
        </el-tab-pane>
    </el-tabs>

</template>

<script>
    export default {
        components: {},
        props:['paper'],
        data() {
            return {
                d:[],
                tabPostion: 'top',
                total:0,
                currentPage:1,
                pageSize:15,
                total1:0,
                currentPage1:1,
                pageSize1:15,
                tempList:[],
                tempClist:[],
                tableData: [],
                clinicalData:[],
            }

        },
        mounted() {
        },
        watch: {
            paper:function (newData) {
                this.d = newData
                this.get_data()
                this.currentChangePage(this.tableData,1)
                this.currentChangePage1(this.clinicalData,1)
            }
        },
        methods: {
            get_data() {
                let t=[]
                for(let i=0;i<this.d[0].length;i++) {
                    let article = {}
                    article["title"] = this.d[0][i]["ArticleTitle"]
                    article["author"] = this.d[0][i]["Authors"]
                    article["journal"] = this.d[0][i]["Journal_Title"]
                    article["url"] = this.d[0][i]["url"]

                  if (!this.d[0][i]["tweet"]) {
                    t.push({"article":article,"year":this.d[0][i]["PubYear"],"tweet":"0"})
                  }
                  else{
                    this.d[0][i]["tweet"]=String(this.d[0][i]["tweet"]).split('.')[0]
                    t.push({"article":article,"year":this.d[0][i]["PubYear"],"tweet":this.d[0][i]["tweet"]})
                  }
                }
                this.tableData = t

                let t1 = []
                if (this.d[1] != null) {
                    for (let i = 0; i < this.d[1].length; i++) {
                        let clinical = {}
                        clinical['title'] = this.d[1][i]['title']
                        clinical['url'] = this.d[1][i]['url']
                        clinical['sponsors'] = this.d[1][i]['sponsors']
                        clinical['researchers'] = this.d[1][i]['researchers']

                        t1.push({"article": clinical, "year": this.d[1][i]['year']})
                    }
                }
                this.clinicalData = t1
            },
            sort_change(column) {
                this.currentPage = 1; 
                if (column.prop === "year") {
                    this.proptype = column.prop; 
                    if (column.order === "descending") {
                        this.tableData.sort(this.my_desc_sort);
                        this.currentChangePage(this.tableData,1)
                    } else if (column.order === "ascending") {
                        this.tableData.sort(this.my_asc_sort);
                        this.currentChangePage(this.tableData,1)
                    }
                }
              if (column.prop === "tweet") {
                this.proptype = column.prop; 
                if (column.order === "descending") {
                  this.tableData.sort(this.my_desc_sort);
                  this.currentChangePage(this.tableData,1)
                } else if (column.order === "ascending") {
                  this.tableData.sort(this.my_asc_sort);
                  this.currentChangePage(this.tableData,1)
                }
              }
                if(column.prop === "article") {
                    this.proptype = column.prop; 
                    if (column.order === "descending") {
                        this.tableData.sort(this.my_desc_sort);
                        this.currentChangePage(this.tableData,1)
                    } else if (column.order === "ascending") {
                        this.tableData.sort(this.my_asc_sort);
                        this.currentChangePage(this.tableData,1)
                    }
                }
            },

            my_desc_sort(a, b) {
                return b[this.proptype] - a[this.proptype]; 
            },
            my_asc_sort(a, b) {
                return a[this.proptype] - b[this.proptype];
            },
            objectSpanMethod({rowIndex,columnIndex}) {
                if (columnIndex !== 0) {
                    if (rowIndex % 2 === 0) {
                        return {
                            rowspan: 2,
                            colspan: 1
                        };
                    } else {
                        return {
                            rowspan: 0,
                            colspan: 0
                        };
                    }
                }
            },
            handleSizeChange: function (pageSize) {
                this.pageSize = pageSize
                this.handleCurrentChange(this.currentPage)
            },
            handleCurrentChange:function (currentPage) {
                this.currentPage = currentPage
                this.currentChangePage(this.tableData,currentPage)
            },
            handleSizeChange1: function (pageSize) {
                this.pageSize1 = pageSize
                this.handleCurrentChange1(this.currentPage)
            },
            handleCurrentChange1:function (currentPage) {
                this.currentPage1 = currentPage
                this.currentChangePage1(this.clinicalData,currentPage)
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
            currentChangePage1(list,currentPage) {
                let from = (currentPage-1)*this.pageSize1
                let to = currentPage*this.pageSize1
                this.tempClist = []
                for(;from<to;from++) {
                    if(list[from]) {
                        this.tempClist.push(list[from]);
                    }
                }
            },

        }
    }
</script>

<style>
    .tab {
        padding: 0;
        position: relative;
        margin: 0 0;
    }
    .el-table__row>td {
    }
    .el-table{
        border: 1px solid #eeeeee;
    }

    .el-table::before {
        height: 0px
    }
    .my-tab {
    }
    .el-pagination{
        text-align: center;
        padding: 3rem 5rem;
    }

</style>
