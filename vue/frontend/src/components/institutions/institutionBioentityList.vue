<template>
        <el-tabs type="card" :tab-position="tabPosition" stretch="true">
            <el-tab-pane label="Disease" style="background-color: #409EFF">
                    <el-table
                            :data="tab1"
                            type="success"
                            style="width: 100%"
                            :show-header="false"
                            :row-class-name="tableRowClassName2"
                            :border="false"
                            :row-style="{height:'40px'}"
                            :cell-style="{padding:'0px'}"
                    >
                        <el-table-column
                                prop="name"
                                label="Co-Institutions"
                                align="center"
                        >
                        </el-table-column>
                        <el-table-column
                                prop="num"
                                label="Paper"
                                align="center"
                        >
                        </el-table-column>
                    </el-table>
                </el-tab-pane>
            <el-tab-pane label="Drug" class="type-drug">
                <el-table
                        :data="tab2"
                        style="width: 100%"
                        :show-header="false"
                        :border="false"
                        :row-class-name="tableRowClassName2"
                        :row-style="{height:'40px'}"
                        :cell-style="{padding:'0px'}"
                >
                    <el-table-column
                            prop="name"
                            label="Co-Institutions"
                            align="center"
                    >
                    </el-table-column>
                    <el-table-column
                            prop="num"
                            label="Paper"
                            align="center"
                    >
                    </el-table-column>
                </el-table>
            </el-tab-pane >
            <el-tab-pane label="Gene" class="type-gene">
                <el-table
                        :data="tab3"
                        style="width: 100%"
                        :show-header="false"
                        :border="false"
                        :row-class-name="tableRowClassName2"
                        :row-style="{height:'40px'}"
                        :cell-style="{padding:'0px'}"
                >
                    <el-table-column
                            prop="name"
                            label="Co-Institutions"
                            align="center"
                    >
                    </el-table-column>
                    <el-table-column
                            prop="num"
                            label="Paper"
                            align="center"
                    >
                    </el-table-column>
                </el-table>
            </el-tab-pane>
        </el-tabs>
</template>

<script>
    import $axios from '../../util/axios'

    export default {
        components: {
        },

        data() {
            return {
                tabPosition:"top",
                tabIndex: 3,
                tab1:[],
                tab2:[],
                tab3:[],
                tableData: [{
                    name: 'Gene'
                }, {
                    name: 'Proteins'
                }, {
                    name: 'Ligands'
                }, {
                    name: 'Molecules'
                },{
                    name: 'Ace2'
                }]
            }
        },
        mounted() {
            this.get_data()
        },
        methods: {
            get_data() {
                $axios.get("/InstitutionWordCloud").then(response=>{
                    let d=response.data
                    console.log(d)
                    let i=1;
                    for(let key in d) {
                        let temp = d[key].slice(0,5)
                        let mention = []
                        for(let i=0;i<5;i++) {
                            mention.push({'name':temp[i]['Mention'],'num':temp[i]['occurences']})
                        }
                        console.log(mention)
                        if(i===1) {
                            this.tab1 = mention
                        } else if(i===2) {
                            this.tab2 = mention
                        } else if(i===3) {
                            this.tab3 = mention
                        }
                        i++;
                    }
                })
            },
            tableRowClassName({row,rowIndex}) {
                console.log(row)
                if(rowIndex%2===1) {
                    return 'white-background'
                } else {
                    return 'gray-background'
                }
            },
            tableRowClassName2({row,rowIndex}) {
                console.log(row)
                if(rowIndex%2===0) {
                    return 'white-background'
                } else {
                    return 'gray-background'
                }
            }
        }
    }
</script>

<style>
    .el-table .gray-background {
        background: #f2f2f2;
    }
    .el-table .white-background {
        background: #ffffff;
    }
    .entity-name {
        /*text-decoration:underline;*/
        color: black;
    }
    .el-tabs el-tabs--left{
        float: left;
        margin-bottom: 0;
        margin-right: 0;
    }
    .el-tabs .el-tabs__header.is-left {
        float: left;
        margin-bottom: 0;
        margin-right: 0;
    }
    .wordcloud {
        max-width:450px;
        max-height:450px;
        /*min-height: 400px;*/
        padding: 0px ;
    }
    .box-card {
        width: 330px;
        margin-top: 2.5rem;
        padding: 0px;
    }
    .el-card__body{
        padding: 0 0 0 0;
    }
</style>