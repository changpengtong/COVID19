<template>
    <el-table
            :data="tableData"
            :default-sort = "{prop: 'num', order: 'descending'}"
            style="width: 100%"
            height="300"
            :stripe="true"
    >

        <el-table-column
                prop="ins"
                label="Institutions"
        >
            <template slot-scope="scope">
                <a :href="scope.row.ins.url" style="color:#000000">{{scope.row.ins.name}}</a>
            </template>
        </el-table-column>
        <el-table-column
                prop="num"
                label="Papers"
                align="center"
                width="70"
        >
        </el-table-column>
    </el-table>
</template>

<script>

   // import $axios from "../../util/axios";

    export default {
        props:['institution'],
        data() {
            return {
                tableData:[],
                d:[],
            }
            // return {
            //     tableData: [{
            //         name: 'The University of Texas at Austin',
            //     }, {
            //         name: 'The University of Texas at Dallas'
            //     }, {
            //         name: 'The University of California, Los Angeles'
            //     }, {
            //         name: 'Massachusetts Institute of Technology'
            //     }]
            // }
        },
        mounted() {
           // this.get_data()
        },
        watch: {
            institution:function(newData) {
                // console.log(newData)
                this.d = newData
                this.get_data()
            }
        },
        methods: {
            get_data() {
               // let d=[]
                let t=[]
                    for(let i=0;i<this.d.length;i++) {
                      if (this.d[i]["Affiliation"]!=null) {
                        t.push({
                          'ins': {
                            'name': this.d[i]["Affiliation"],
                            'url': "/#/institution/" + this.d[i]["Affiliation"]
                          },
                          'num': this.d[i]['NumberOfPapers'],
                        })
                      }
                    }
                    this.tableData = t
                    console.log(this.tableData)
            }
        }
    }
</script>

<style scoped>
    .entity-name {
        /*text-decoration:underline;*/
        color: black;
    }
</style>