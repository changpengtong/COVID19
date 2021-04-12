<template>
    <div class="display-view container">
        <div class="display-info ">
            <el-row :gutter="24" >

                <el-col :md="17" :sm="24" :xs="24">
                    <entity-card class="display-author" v-loading="loading" :name="name"></entity-card>
                    <div
                            class="display-data"
                            v-loading="isLoading"
                            element-loading-text="loading……">
                        <author-paper :paper="paper" v-loading="loading"></author-paper>
                    </div>
                </el-col>
                <el-col :md="7" :sm="24" :xs="24" class="right">
                    <div class="display-visual">
                        <institution-list class="up" :institution="institution" v-loading="loading"></institution-list>
                    </div>
                    <bioentity-graph :bioentity_data="bioentity_data" v-loading="loading"></bioentity-graph>
                    <trend-linefold :linefold="linefold" v-loading="loading"></trend-linefold>
                    <br/>
                    <co-author :coauthor="coauthor" :flag="flag" v-loading="loading"></co-author>
                </el-col>
            </el-row>
        </div>
    </div>
</template>

<script>
    import bioData from "../data/bioentity-data.json";
    import EntityCard from "../components/general/BioentityCard";
    import entityPapers from "../components/general/paper";
    import CoAuthors from "../components/general/coauthor";
    import EntityLinefold from "../components/general/linefold";
    import institutionList from "../components/general/institutionList";
    import BioentityGraph from "../components/general/BioentityGraph";
    import $axios from "./../util/axios"
    export default {
        components: {
            "entity-card": EntityCard,
            "author-paper": entityPapers,
            "co-author": CoAuthors,
            "institution-list": institutionList,
            "trend-linefold": EntityLinefold,
            "bioentity-graph": BioentityGraph,
        },
        data() {
          return {
              paper:[],
              linefold:[],
              coauthor:[],
              institution:[],
              flag:"",
              name:"",
              type:[],
              bioentity_data: [],
              loading:true,
          }
        },
      created() {
        this.get_data()

      },
      watch: {
        '$route': 'get_data'
      },
        methods:{
            get_data() {
              this.loading = true
              $axios.get("/displayBioentity/"+this.$route.params.id).then(response=>{
                this.loading = false
                let d=response.data
                    this.name = d["name"]
                    this.paper = [d["articles"], d["clinicals"]]
                    this.linefold = d["bar"]
                    this.coauthor = d["coauthor"]
                    this.institution = d["institution"]
                    this.bioentity_data = bioData[this.$route.params.id]
                    this.flag = "bio"
                })
            }
        }


    }

</script>

<style>

    .container {
        max-width: 1200px;
        width: 100%;
        margin: 0 auto;
    }
    .right{
        margin: 0 auto;
    }
    .display-author{
        width: 100%;
    }
    .display-view {
        width: 100%;
        margin: 0 auto;
    }

    .display-query {
        margin-bottom: 1.5em;
    }
    .display-query #query {
        font-weight: bold;
        font-size: 1.5em;
    }

    .display-info {
        min-height: 30em;
        margin:auto;
    }


    .display-data {
        border-right: 1px solid rgba(150,150,150,0.2);
    }

    .display-visual .title {
        text-transform: capitalize;
        margin-bottom: 1em;
    }

    .el-divider--horizontal {
        margin:0 0;
    }

    .up {
        margin-top: 2em;
    }

    .down {
        margin-bottom: 1em;
    }

</style>