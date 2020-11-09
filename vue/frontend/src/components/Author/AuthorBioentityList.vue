<template>
    <el-tabs type="card" :tab-position="tabPosition" stretch="true" >
        <el-tab-pane id="tab1" class="tab" label="Disease" >
            <el-table
                    :data="tab1"
                    style="width: 100%"
                    :show-header="false"
                    :stripe="true"
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
                        label="Co-Institutions"
                        align="center"
                >
                </el-table-column>
            </el-table>
        </el-tab-pane>
        <el-tab-pane id="tab2" class="tab" label="Drug">
            <el-table
                    :data="tab2"
                    style="width: 100%"
                    :show-header="false"
                    :stripe="true"
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
                        label="Co-Institutions"
                        align="center"
                >
                </el-table-column>
            </el-table>
        </el-tab-pane>
        <el-tab-pane id="tab3" class="tab" label="Gene">
            <el-table
                    :data="tab3"
                    style="width: 100%"
                    :show-header="false"
                    :stripe="true"
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
                        label="Co-Institutions"
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

        data() {
            return {
                tabPosition:"top",
                tab1:[],
                tab2:[],
                tab3:[],
            }
        },
        mounted() {
            this.get_data()
        },
        methods: {
            get_data() {
                $axios.get("/AuthorWordCloudType").then(response=>{
                    let d=response.data
                    // console.log(d)
                    let i=1;
                    for(let key in d) {
                        let temp=d[key]
                        if (d[key]!=null){
                            temp = d[key].slice(0,5)
                        }

                        let mention = []
                        for(let i=0;i<5;i++) {
                            if (temp!=null){
                                mention.push({'name':temp[i]['Mention'],'num':temp[i]['occurences']})
                            }

                        }
                        console.log(mention)
                        if(i===1) {
                            this.tab1 = mention
                        } else if(i===2) {
                            this.tab2 = mention
                        } else if(i===3) {
                            this.tab3 = mention
                        } else if(i===4) {
                                this.tab4 = mention
                        } else if(i===5) {
                            this.tab5 = mention
                            console.log('mention5')
                        }
                        i++;
                    }
                })
            },
        }

    }
</script>

<style scoped>
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
    /deep/ .el-tabs__item.is-active{
        color: white;
        background-color: #409EFF;
    }
    /deep/ .el-tabs__item {
        font-size: 0.8em;
        letter-spacing: 0.1em;
    }
</style>