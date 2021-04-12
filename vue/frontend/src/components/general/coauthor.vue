<template>
    <div>
        <div class="head">
            <span class="title-left" style="margin-right: 11em" v-if='flag==="aut"'>Co-authors</span>
            <span class="title-left" style="margin-right: 11em" v-else>Authors</span>
            <el-button type="text" @click="dialogTableVisible = true" >
                <span class="title-right" style="text-align: right">View All</span>
            </el-button>
        </div>
        <el-divider></el-divider>
        <el-dialog title="Authors" :visible.sync="dialogTableVisible">
            <div>
                <el-table :data="coauthors">
                    <el-table-column property="name" label="Name" width="200"></el-table-column>
                    <el-table-column property="institution" label="Affiliation"></el-table-column>
                </el-table>
            </div>
        </el-dialog>

        <div class="info-card" :body-style="{ padding: '0px' }" shadow="hover">

            <div v-for="(author,index) in coauthors" :key="index" class="link">
                <a :href=author.url target="_blank">
                    <el-row :gutter="20" v-if="index<20">
                        <el-col :xs="6" :sm="2" :md="6" :lg="5">
                            <div class="info-image">
                                <img src="../../../src/assets/img/author.png">
                            </div>
                        </el-col>
                        <el-col :xs="15" :sm="20" :md="16" :lg="17">
                            <div class="info-text">
                                <div class="info-item">
                            <span id="name" class="info-item-title" style="color: black;font-size: smaller">
                                {{author.name}}
                            </span>
                                </div>
                                <div class="info-work">
                                    <span style="color: dimgray;font-size: smaller">{{author.institution}}</span>
                                </div>
                            </div>
                        </el-col>
                        <el-col :xs="2" :sm="2" :md="2" :lg="2">
                            <div class="info-icon">
                                <i class="custom-icon el-icon-arrow-right"></i>
                            </div>
                        </el-col>
                    </el-row>
                </a>
            </div>
        </div>

    </div>
</template>

<script>
    export default {
        props:['coauthor','flag'],
        data() {
            return {
                d:[],
                coauthors:[],
                dialogTableVisible: false,
                dialogFormVisible: false
            }
        },
        watch:{
            coauthor:function (newData) {
                this.d = newData;
                this.get_data()
            }
        },
        mounted() {
        },
        methods: {
            get_data() {
                let t=[]
                for(let i=0;i<this.d.length;i++) {
                    if (typeof this.d[i]["Location"]!='undefined' && typeof this.d[i]["aid"]!='undefined' && typeof this.d[i]["ForeName"]!='undefined' && typeof this.d[i]["LastName"]!='undefined' && typeof this.d[i]["Affiliation"]!='undefined' ) {
                        t.push({
                            'name': this.d[i]["ForeName"] + " " + this.d[i]["LastName"],
                            'institution': this.d[i]["Affiliation"] + " " + this.d[i]["Location"],
                            'url': "/COVID19/#/author/" + this.d[i]["aid"]
                        })
                    }
                }
                this.coauthors = t
            },
    }   
}
</script>

<style scoped>
    .info-card {
        padding: 1em;
    }
    .info-image {
        margin-bottom: .5em;
        margin-top: .7em;
        margin-left: .4em;
        width: 2.5em;
        height: 2.5em;
        position: relative;
        overflow: hidden;
      object-fit: cover;
      border-radius: 50%;
    }
    .info-icon {
        margin-top: 1.5em;
    }
    .custom-icon {
        font-size: 17px;
    }
    img {

        display: inline;
        margin: 0 auto;
        height: 100%;
        width: 100%;
    }
    .info-text{
        margin: .7em 0;
    }
    .info-item {
    }
    .info-work {
        overflow:hidden;
        text-overflow:ellipsis;
        white-space:nowrap;
    }
    .info-item-title {
        font-weight: normal;
        text-transform: capitalize;
    }
    #name {
        font-size: 20px;
        margin-bottom: .5em;
    }
    div.link:hover{
        background-color:#e9e9e9;
    }
    .head{
      white-space: nowrap;

    }
    .head *{
      display: inline;

    }

</style>