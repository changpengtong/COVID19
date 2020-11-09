<template>
    <div>
    <el-table
            :data="tempList"
            style="width: 100%"
            class="my-table"
            :default-sort="{prop:'year',order:'descending'}"
            @sort-change="sort_change"
            :header-cell-style="{
                'background-color': '#F5F5F5',
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
                        {{item}}
                    </span>
                    <span v-if="index==='journal'" style="font-style: italic; ont-size: small">
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
            @size-change="handleSizeChange"
            @current-change="handleCurrentChange" :current-page="currentPage"
            :page-sizes="[5,10,20,50,100]"
            :page-size="pageSize" layout="total, sizes, prev, pager, next, jumper"
            :total="tableData.length">
    </el-pagination>
    </div>
</template>


<script>
    import $axios from "../../util/axios"
    export default {
        data() {
            return {
                total:0,
                currentPage:1,
                pageSize:10,
                tempList:[],
                tableData:[]
            }
        },
        mounted() {
            this.get_data()
            this.currentChangePage(this.tableData,1)
        },

        methods: {
            get_data() {
                let d=[]
                let t=[]
                $axios.get("/BioentityArticles").then(response=>{
                    d=response.data
                    console.log(d)
                     for(let i=0;i<d.length;i++) {
                         let article = {}
                         article["title"] = d[i]["ArticleTitle"]
                        article["author"] = d[i]["Authors"]
                         article["journal"] = d[i]["Journal_Title"]
                         article["url"] = ""
                         if(d[i]["PubYear"]!=null)
                            t.push({"article":article,"year":d[i]["PubYear"]})
                     }

                    this.tableData = t
                    console.log(this.tableData)
                    this.currentChangePage(this.tableData,1)
                })
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
        }
    }
</script>

<style scoped>
    .tab {
        padding: 0;
        position: relative;
        margin: 0 0;
        /*border: 1px solid #eeeeee;*/

    }
    .el-table__row>td {
        border:none;
    }

    .el-table::before {
        height: 0px
    }
    .el-table{
        border: 1px solid #eeeeee;
    }

    .el-table::before {
        height: 0px
    }

    .el-pagination{
        text-align: center;
        padding: 3rem 5rem;
    }
</style>
