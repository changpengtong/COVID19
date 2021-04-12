<template>
  <div class="table">
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
  </div>
</template>

<script>
    export default {
        props:['institution'],
        data() {
            return {
                tableData:[],
                d:[],
            }
        },
        mounted() {
        },
        watch: {
            institution:function(newData) {
                this.d = newData
                this.get_data()
            }
        },
        methods: {
            get_data() {
              let t=[]
              if (this.d!=null){
                for(let i=0;i<this.d.length;i++) {
                  if (this.d[i]["Affiliation"]!=null) {
                    this.d[i]["Affiliation"]=this.d[i]["Affiliation"].replace(/^,+/,"").replace(/,+$/,"")
                    t.push({
                      'ins': {
                        'name': this.d[i]["Affiliation"],
                        'url': "/COVID19/#/institution/" + this.d[i]["id"]
                      },
                      'num': this.d[i]['NumberOfPapers'],
                    })
                  }
                }
              }

              this.tableData = t
            }
        }
    }
</script>

<style>
    .entity-name {
        color: black;
    }

    .scroll{
      overflow-y: scroll;
      height: 300px;
    }
    .scroll::-webkit-scrollbar {
      -webkit-appearance: none;
      width: 7px;
    }

    .scroll::-webkit-scrollbar-thumb {
      border-radius: 4px;
      background-color: rgba(0, 0, 0, .5);
      box-shadow: 0 0 1px rgba(255, 255, 255, .5);
    }
</style>