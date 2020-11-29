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
                            // 'background-color': '#F5F5F5'
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
                                <span v-if="index==='title'" style="font-size: large">
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
                      width="120"
                  >
                  </el-table-column>
                </el-table>
                <el-pagination
                        @size-change="handleSizeChange"
                        @current-change="handleCurrentChange" :current-page="currentPage"
                        :page-sizes="[5,10,20,50,100]"
                        :page-size="pageSize" layout="total, sizes, prev, pager, next, jumper"
                        :total="tableData.length">
                </el-pagination>
            </div>
        </el-tab-pane>
        <el-tab-pane label="ClinicalTrials">
            <div>
                <el-table
                        :data="tempClist"
                        style="width: 100%"
                        :default-sort="{prop:'year',order:'ascending'}"
                        @sort-change="sort_change"
                        :header-cell-style="{
                            // 'background-color': '#F5F5F5'
                            'padding-left': '1rem'
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
                            <div v-for="(item,index) in scope.row.article" :key="index" style="margin-left: 1rem">
                                <span v-if="index==='title'" style="font-size: large">
                                    <a href="https://www.google.com/">
                                    {{item}}
                                    </a>
                                </span>
                                <span v-if="index==='author'" style="font-size: small">
                                    <text-expand-shrink :text="item" :key="index"></text-expand-shrink>
                                    <!--                                    {{item}}-->
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
                    >
                    </el-table-column>

                </el-table>
                <el-pagination
                        @size-change="handleSizeChange1"
                        @current-change="handleCurrentChange1" :current-page="currentPage1"
                        :page-sizes="[5,10,20,50,100]"
                        :page-size="pageSize1" layout="total, sizes, prev, pager, next, jumper"
                        :total="cList.length"
                        style="text-align: center;"
                >
                </el-pagination>
            </div>
        </el-tab-pane>

    </el-tabs>

</template>

<script>
    import textExpandShrink from "./../textExpandShrink";
    export default {
        components: {
            "text-expand-shrink": textExpandShrink
        },
        props:['paper'],
        data() {
            return {
                d:[],
                tabPostion: 'top',
                total:0,
                currentPage:1,
                pageSize:10,
                total1:0,
                currentPage1:1,
                pageSize1:10,
                tempList:[],
                //    dataList:[],
                tempClist:[],
                tableData: [],
                cList:[{
                    article: {
                        title: "Monte Carlo Calculations for Alcohols and Their Mixtures with Alkanes. Transferable Potentials for Phase Equilibria. 5. United-Atom Description of Primary, Secondary, and Tertiary Alcohols",
                        author: "Bin Chen, Jeffrey J. Potoff, and J. Ilja Siepmann",
                        journal: "The Journal of Physical Chemistry B 105 (15), 3093-3104",
                        url:""
                    },
                    year: 12
                },{
                    article: {
                        title: "Monte Carlo Calculations for Alcohols and Their Mixtures with Alkanes. Transferable Potentials for Phase Equilibria. 5. United-Atom Description of Primary, Secondary, and Tertiary Alcohols",
                        author: "Bin Chen, Jeffrey J. Potoff, and J. Ilja Siepmann",
                        journal: "The Journal of Physical Chemistry B 105 (15), 3093-3104",
                        url:""
                    },
                    year: 13
                },{
                    article: {
                        title: "Potentials for Phase Equilibria. 5. United-Atom Description of Primary, Secondary, and Tertiary Alcohols",
                        author: "Bin Chen, Jeffrey J. Potoff, and J. Ilja Siepmann",
                        journal: "The Journal of Physical Chemistry B 105 (15), 3093-3104",
                        url:""
                    },
                    year: 14
                },{
                    article: {
                        title: "Monte Carlo Calculations for Alcohols and Their Mixtures with Alkanes",
                        author: "Jeffrey J. Potoff, and J. Ilja Siepmann",
                        journal: "The Journal of Physical Chemistry B 105 (15), 3093-3104",
                        url:""
                    },
                    year: 15
                },{
                    article: {
                        title: " 5. United-Atom Description of Primary, Secondary, and Tertiary Alcohols",
                        author: "Bin Chen, Jeffrey J. Potoff, and J. Ilja Siepmann",
                        journal: "The Journal of Physical Chemistry B 105 (15), 3093-3104",
                        url:""
                    },
                    year: 16
                },{
                    article: {
                        title: "Transferable Potentials for Phase Equilibria. 5. United-Atom Description of Primary, Secondary, and Tertiary Alcohols",
                        author: "Bin Chen, Jeffrey J. Potoff, and J. Ilja Siepmann",
                        journal: "The Journal of Physical Chemistry B 105 (15), 3093-3104",
                        url:""
                    },
                    year: 17
                }]
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
        methods: {
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
            sort_change(column) {
                this.currentPage = 1; // 排序后返回第一页
                if (column.prop === "year") {
                    this.proptype = column.prop; // 将键名prop赋值给变量proptype
                    if (column.order === "descending") {
                        this.tableData.sort(this.my_desc_sort);
                        this.currentChangePage(this.tableData,1)
                    } else if (column.order === "ascending") {
                        this.tableData.sort(this.my_asc_sort);
                        this.currentChangePage(this.tableData,1)
                    }
                }
              if (column.prop === "tweet") {
                this.proptype = column.prop; // 将键名prop赋值给变量proptype
                if (column.order === "descending") {
                  this.tableData.sort(this.my_desc_sort);
                  this.currentChangePage(this.tableData,1)
                } else if (column.order === "ascending") {
                  this.tableData.sort(this.my_asc_sort);
                  this.currentChangePage(this.tableData,1)
                }
              }
                if(column.prop === "article") {
                    this.proptype = column.prop; //
                    if (column.order === "descending") {
                        this.tableData.sort(this.my_desc_sort);
                        this.currentChangePage(this.tableData,1)
                    } else if (column.order === "ascending") {
                        this.tableData.sort(this.my_asc_sort);
                        this.currentChangePage(this.tableData,1)
                    }
                }
            },
            //排序方法
            my_desc_sort(a, b) {
                return b[this.proptype] - a[this.proptype]; // a["time"] 等价于 a.time
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
                this.currentChangePage1(this.cList,currentPage)
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
        /*border: 1px solid #eeeeee;*/

    }
    .el-table__row>td {
        /*margin-left: 10em;*/
    }
    .el-table{
        border: 1px solid #eeeeee;
    }

    .el-table::before {
        height: 0px
    }
    .my-tab {
        /*margin-top: -1.5em;*/
    }
    .el-pagination{
        text-align: center;
        padding: 3rem 5rem;
    }

</style>
