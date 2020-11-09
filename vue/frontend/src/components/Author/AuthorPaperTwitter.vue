<template>
    <el-tabs :tab-position="tabPostion" stretch="true" class="tab">
        <el-tab-pane label="Covid-19 Papers">
            <div>
                <el-table
                        :data="tempList"
                        style="width: 100%"
                        class="my-table"
                        :default-sort="{prop:'year',order:'descending'}"
                        :header-cell-style="{
                            // 'background-color': '#F5F5F5'
                            'padding-left': '1rem',
                        }"
                >
                    <el-table-column
                            prop="article"
                            label=TITLE
                            sortable
                            align="left"
                            width="600"
                    >
                        <template slot-scope="scope">
                            <div v-for="(item,index) in scope.row.article" :key="index"  style="margin-left: 1rem">
                                <span v-if="index==='title'" style="font-size: large">
                                    <a href="https://www.google.com/">
                                    {{item}}
                                    </a>
                                </span>
                                <span v-if="index==='author'" style="color: darkgray;font-size: small">
                                    {{item}}
                                </span>
                                <span v-if="index==='journal'" style="color: darkgray;font-size: small">
                                    {{item}}
                                </span>
                            </div>
                        </template>
                    </el-table-column>

                    <el-table-column
                            prop="year"
                            label="YEAR"
                            sortable
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
        </el-tab-pane>
        <el-tab-pane label="ClinicalTrials">
            <div>
                <el-table
                        :data="tempClist"
                        style="width: 100%"
                        :default-sort="{prop:'year',order:'descending'}"
                        :header-cell-style="{
                            // 'background-color': '#F5F5F5'
                            'padding-left': '1rem'
                        }"
                >
                    <el-table-column
                            prop="article"
                            label=TITLE
                            sortable
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
                                <span v-if="index==='author'" style="color: darkgray;font-size: small">
                                    {{item}}
                                </span>
                                <span v-if="index==='journal'" style="color: darkgray;font-size: small">
                                    {{item}}
                                </span>
                            </div>
                        </template>
                    </el-table-column>

                    <el-table-column
                            prop="year"
                            label="YEAR"
                            sortable
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
        <el-tab-pane label="Most Discussed">
            <div>
                <el-table
                        :data="tempClist"
                        style="width: 100%"
                        :default-sort="{prop:'year',order:'descending'}"
                        :header-cell-style="{
                            // 'background-color': '#F5F5F5'
                            'padding-left': '1rem'
                        }"
                >
                    <el-table-column
                            prop="article"
                            label=TITLE
                            sortable
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
                                <span v-if="index==='author'" style="color: darkgray;font-size: small">
                                    {{item}}
                                </span>
                                <span v-if="index==='journal'" style="color: darkgray;font-size: small">
                                    {{item}}
                                </span>
                            </div>
                        </template>
                    </el-table-column>

                    <el-table-column
                            prop="year"
                            label="YEAR"
                            sortable
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
    var request = new XMLHttpRequest()

    request.open('GET', 'https://api.altmetric.com/v1/citations/1w?num_results=100', true)
    request.onload = function () {
        // Begin accessing JSON data here
        var data = JSON.parse(this.response)["results"]

        if (request.status >= 200 && request.status < 400) {
            Object.keys(data).forEach(function (key){
                console.log(key);
                console.log(data[key]['title']);
            });
        } else {
            console.log('error')
        }
    }

    request.send()





    // export default {
    //     data() {
    //         return {
    //             tabPostion: 'top',
    //             total:0,
    //             currentPage:1,
    //             pageSize:10,
    //             total1:0,
    //             currentPage1:1,
    //             pageSize1:10,
    //             tempList:[],
    //             tempClist:[],
    //             tableData: [{
    //                 article: {
    //                     title: "A survey of current trends in computational drug repositioning",
    //                     author: "J Li, S Zheng, B Chen, AJ Butte, SJ Swamidass, Z Lu",
    //                     journal: "Elife 8, e48847",
    //                     url:""
    //                 },
    //                 cited: 269,
    //                 year: "2016"
    //             }, {
    //                 article: {
    //                     title: "Chem2Bio2RDF: a semantic framework for linking and data mining chemogenomic and systems chemical biology data",
    //                     author: "J Li, S Zheng, B Chen, AJ Butte, SJ Swamidass, Z Lu",
    //                     journal: "Elife 8, e48847",
    //                     url:""
    //                 },
    //                 cited: 210,
    //                 year: "2010"
    //             }, {
    //                 article: {
    //                     title: "Gaining insight into off-target mediated effects of drug candidates with a comprehensive systems chemical biology analysis",
    //                     author: "J Li, S Zheng, B Chen, AJ Butte, SJ Swamidass, Z Lu",
    //                     journal: "Elife 8, e48847",
    //                     url:""
    //                 },
    //                 cited: 167,
    //                 year: "2009"
    //             }, {
    //                 article: {
    //                     title: "PubChem as a source of polypharmacology",
    //                     author: "J Li, S Zheng, B Chen, AJ Butte, SJ Swamidass, Z Lu",
    //                     journal: "Elife 8, e48847",
    //
    //                 },
    //                 cited: 147,
    //                 year: 2009
    //             },
    //                 {
    //                     article: {
    //                         title: "Chem2Bio2RDF: a semantic framework for linking and data mining chemogenomic and systems chemical biology data",
    //                         author: "J Li, S Zheng, B Chen, AJ Butte, SJ Swamidass, Z Lu",
    //                         journal: "Elife 8, e48847",
    //                         url:""
    //                     },
    //                     cited: 210,
    //                     year: "2010"
    //                 },
    //                 {
    //                     article: {
    //                         title: "Chem2Bio2RDF: a semantic framework for linking and data mining chemogenomic and systems chemical biology data",
    //                         author: "J Li, S Zheng, B Chen, AJ Butte, SJ Swamidass, Z Lu",
    //                         journal: "Elife 8, e48847",
    //                         url:""
    //                     },
    //                     cited: 210,
    //                     year: "2010"
    //                 },
    //                 {
    //                     article: {
    //                         title: "Chem2Bio2RDF: a semantic framework for linking and data mining chemogenomic and systems chemical biology data",
    //                         author: "J Li, S Zheng, B Chen, AJ Butte, SJ Swamidass, Z Lu",
    //                         journal: "Elife 8, e48847",
    //                         url:""
    //                     },
    //                     cited: 210,
    //                     year: "2010"
    //                 },
    //                 {
    //                     article: {
    //                         title: "Chem2Bio2RDF: a semantic framework for linking and data mining chemogenomic and systems chemical biology data",
    //                         author: "J Li, S Zheng, B Chen, AJ Butte, SJ Swamidass, Z Lu",
    //                         journal: "Elife 8, e48847",
    //                         url:""
    //                     },
    //                     cited: 210,
    //                     year: "2010"
    //                 },
    //                 {
    //                     article: {
    //                         title: "Chem2Bio2RDF: a semantic framework for linking and data mining chemogenomic and systems chemical biology data",
    //                         author: "J Li, S Zheng, B Chen, AJ Butte, SJ Swamidass, Z Lu",
    //                         journal: "Elife 8, e48847",
    //                         url:""
    //                     },
    //                     cited: 210,
    //                     year: "2010"
    //                 },
    //                 {
    //                     article: {
    //                         title: "Chem2Bio2RDF: a semantic framework for linking and data mining chemogenomic and systems chemical biology data",
    //                         author: "J Li, S Zheng, B Chen, AJ Butte, SJ Swamidass, Z Lu",
    //                         journal: "Elife 8, e48847",
    //                         url:""
    //                     },
    //                     cited: 210,
    //                     year: "2010"
    //                 },
    //                 {
    //                     article: {
    //                         title: "Chem2Bio2RDF: a semantic framework for linking and data mining chemogenomic and systems chemical biology data",
    //                         author: "J Li, S Zheng, B Chen, AJ Butte, SJ Swamidass, Z Lu",
    //                         journal: "Elife 8, e48847",
    //                         url:""
    //                     },
    //                     cited: 210,
    //                     year: "2010"
    //                 }],
    //             cList:[{
    //                 article: {
    //                     title: "Monte Carlo Calculations for Alcohols and Their Mixtures with Alkanes. Transferable Potentials for Phase Equilibria. 5. United-Atom Description of Primary, Secondary, and Tertiary Alcohols",
    //                     author: "Bin Chen, Jeffrey J. Potoff, and J. Ilja Siepmann",
    //                     journal: "The Journal of Physical Chemistry B 105 (15), 3093-3104",
    //                     url:""
    //                 },
    //                 year: "2001"
    //             },{
    //                 article: {
    //                     title: "Monte Carlo Calculations for Alcohols and Their Mixtures with Alkanes. Transferable Potentials for Phase Equilibria. 5. United-Atom Description of Primary, Secondary, and Tertiary Alcohols",
    //                     author: "Bin Chen, Jeffrey J. Potoff, and J. Ilja Siepmann",
    //                     journal: "The Journal of Physical Chemistry B 105 (15), 3093-3104",
    //                     url:""
    //                 },
    //                 year: "2001"
    //             },{
    //                 article: {
    //                     title: "Potentials for Phase Equilibria. 5. United-Atom Description of Primary, Secondary, and Tertiary Alcohols",
    //                     author: "Bin Chen, Jeffrey J. Potoff, and J. Ilja Siepmann",
    //                     journal: "The Journal of Physical Chemistry B 105 (15), 3093-3104",
    //                     url:""
    //                 },
    //                 year: "2001"
    //             },{
    //                 article: {
    //                     title: "Monte Carlo Calculations for Alcohols and Their Mixtures with Alkanes",
    //                     author: "Jeffrey J. Potoff, and J. Ilja Siepmann",
    //                     journal: "The Journal of Physical Chemistry B 105 (15), 3093-3104",
    //                     url:""
    //                 },
    //                 year: "2002"
    //             },{
    //                 article: {
    //                     title: " 5. United-Atom Description of Primary, Secondary, and Tertiary Alcohols",
    //                     author: "Bin Chen, Jeffrey J. Potoff, and J. Ilja Siepmann",
    //                     journal: "The Journal of Physical Chemistry B 105 (15), 3093-3104",
    //                     url:""
    //                 },
    //                 year: "2006"
    //             },{
    //                 article: {
    //                     title: "Transferable Potentials for Phase Equilibria. 5. United-Atom Description of Primary, Secondary, and Tertiary Alcohols",
    //                     author: "Bin Chen, Jeffrey J. Potoff, and J. Ilja Siepmann",
    //                     journal: "The Journal of Physical Chemistry B 105 (15), 3093-3104",
    //                     url:""
    //                 },
    //                 year: "2009"
    //             }]
    //         }
    //
    //     },
    //     mounted() {
    //         this.currentChangePage(this.tableData,1)
    //         this.currentChangePage1(this.cList,1)
    //     },
    //     methods: {
    //         objectSpanMethod({rowIndex,columnIndex}) {
    //             if (columnIndex !== 0) {
    //                 if (rowIndex % 2 === 0) {
    //                     return {
    //                         rowspan: 2,
    //                         colspan: 1
    //                     };
    //                 } else {
    //                     return {
    //                         rowspan: 0,
    //                         colspan: 0
    //                     };
    //                 }
    //             }
    //         },
    //         handleSizeChange: function (pageSize) {
    //             this.pageSize = pageSize
    //             this.handleCurrentChange(this.currentPage)
    //         },
    //         handleCurrentChange:function (currentPage) {
    //             this.currentPage = currentPage
    //             this.currentChangePage(this.tableData,currentPage)
    //         },
    //         handleSizeChange1: function (pageSize) {
    //             this.pageSize1 = pageSize
    //             this.handleCurrentChange1(this.currentPage)
    //         },
    //         handleCurrentChange1:function (currentPage) {
    //             this.currentPage1 = currentPage
    //             this.currentChangePage1(this.cList,currentPage)
    //         },
    //         currentChangePage(list,currentPage) {
    //             let from = (currentPage-1)*this.pageSize
    //             let to = currentPage*this.pageSize
    //             this.tempList = []
    //             for(;from<to;from++) {
    //                 if(list[from]) {
    //                     this.tempList.push(list[from]);
    //                 }
    //             }
    //         },
    //         currentChangePage1(list,currentPage) {
    //             let from = (currentPage-1)*this.pageSize1
    //             let to = currentPage*this.pageSize1
    //             this.tempClist = []
    //             for(;from<to;from++) {
    //                 if(list[from]) {
    //                     this.tempClist.push(list[from]);
    //                 }
    //             }
    //         },
    //
    //     }
    // }
</script>

<style scoped>
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
