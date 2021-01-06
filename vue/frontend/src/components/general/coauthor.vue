<template>
    <div>
        <!-- head+button -->
        <div class="head">
            <span class="title-left" style="margin-right: 11em" v-if='flag==="aut"'>Co-authors</span>
            <span class="title-left" style="margin-right: 11em" v-else>Authors</span>
            <el-button type="text" @click="dialogTableVisible = true" >
                <span class="title-right" style="text-align: right">View All</span>
            </el-button>
        </div>
        <el-divider></el-divider>
        <!--list-->
        <el-dialog title="Authors" :visible.sync="dialogTableVisible">
            <div>
                <el-table :data="coauthors">
                    <!--                <el-table-column  width="150" class="table-img"> <el-avatar :size="small" src="../../../../../../frontend/src/assets/img/user.png"></el-avatar></el-table-column>-->
                    <el-table-column property="name" label="Name" width="200"></el-table-column>
                    <el-table-column property="institution" label="Affiliation"></el-table-column>
                </el-table>
            </div>
        </el-dialog>

        <div class="info-card" :body-style="{ padding: '0px' }" shadow="hover">

            <div v-for="(author,index) in coauthors" :key="index" class="link">
                <a :href=author.url target="_blank">
<!--                <span @click="redirect(author)">-->
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
<!--                </span>-->
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
                // console.log(newData)
                this.d = newData;
                this.get_data()
            }
        },
        mounted() {
            // this.get_data()
        },
        methods: {
            get_data() {
                let t=[]
                for(let i=0;i<this.d.length;i++) {

                    // Do stuff
    if (typeof this.d[i]["Location"]!='undefined' && typeof this.d[i]["aid"]!='undefined' && typeof this.d[i]["ForeName"]!='undefined' && typeof this.d[i]["LastName"]!='undefined' && typeof this.d[i]["Affiliation"]!='undefined' ) {
    t.push({
    'name': this.d[i]["ForeName"] + " " + this.d[i]["LastName"],
    'institution': this.d[i]["Affiliation"] + " " + this.d[i]["Location"],
    'url': "/#/author/" + this.d[i]["aid"]
  })
}


                }
                this.coauthors = t
                //   console.log(this.coauthors)
            },
            // redirect:function (author) {
            //     this.$router.push({path:author.url})
            // }
        }
    }
</script>

<style scoped>
    .info-card {
        /* margin: 0 auto; */
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
    }
    .info-icon {
        /*margin-left: 27em;*/
        margin-top: 1.5em;
    }
    .custom-icon {
        font-size: 1.2em;
    }
    img {
        display: inline;
        margin: 0 auto;
        height: 100%;
        width: 100%;
        /*border-radius: 50%;*/
    }
    .info-text{
        margin: .5em 0;
    }
    .info-item {
        /*margin: .5em 0;*/

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
        font-size: 2em;
        margin-bottom: .5em;
    }
    div.link:hover{
        background-color:#e9e9e9;
    }
</style>