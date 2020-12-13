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
              if (this.d!=null){
                for(let i=0;i<this.d.length;i++) {
                  if (this.d[i]["Affiliation"]!=null && this.d[i]["Location"]!=null) {
                    this.d[i]["Affiliation"]=this.d[i]["Affiliation"].replace(/^,+/,"").replace(/,+$/,"")
                    t.push({
                      'ins': {
                        'name': this.d[i]["Affiliation"]+", "+this.d[i]["Location"],
                        'url': "/COVID19/#/institution/" + this.d[i]["Affiliation"]
                      },
                      'num': this.d[i]['NumberOfPapers'],
                    })
                  }
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