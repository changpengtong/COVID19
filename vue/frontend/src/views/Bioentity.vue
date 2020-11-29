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
<!--                    <word-cloud :wordcloud="wordcloud" v-loading="loading"></word-cloud>-->
                    <trend-linefold :linefold="linefold" v-loading="loading"></trend-linefold>
                    <br/>
                    <co-author :coauthor="coauthor" :flag="flag" v-loading="loading"></co-author>
                </el-col>
            </el-row>
        </div>
    </div>
</template>

<script>
    import EntityCard from "../components/general/BioentityCard";
    import entityPapers from "../components/general/paper";
    import CoAuthors from "../components/general/coauthor";
    import EntityLinefold from "../components/general/linefold";
    import institutionList from "../components/general/institutionList";
    // import wordcloud from "../components/general/wordcloud";
    import $axios from "./../util/axios"
    export default {
        components: {
            "entity-card":EntityCard,
            "author-paper": entityPapers,
            "co-author": CoAuthors,
            "institution-list": institutionList,
            "trend-linefold":EntityLinefold,
            // "word-cloud": wordcloud,
        },
        data() {
          return {
              paper:[],
              linefold:[],
              coauthor:[],
              institution:[],
              // wordcloud:[],
              flag:"",
              name:"",
            type:[],
              loading:true
          }
        },
      created() {
        this.get_data()

      },
      watch: {
        '$route': 'get_data'
      },
        // mounted() {
        //     this.get_data()
        // },
        methods:{
            get_data() {
              this.loading = true
              $axios.get("/displayBioentity/"+this.$route.params.id).then(response=>{
              //  $axios.get("/A01Paper/"+this.$route.params.id).then(response=>{
                this.loading = false
                let d=response.data
                    console.log("d is :", d)
                    this.name = this.$route.params.id
                  // this.type=d["type"]
                    this.paper = d["articles"]
                    this.linefold = d["bar"]
                    this.coauthor = d["coauthor"]
                    this.institution = d["institution"]
                    // this.wordcloud = d["wordcloud"]
                  //  this.paper = d
                    this.flag = "bio"
                    // this.loading = false
                  console.log("type is :")
                   console.log(this.type)
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
        /*padding: .5em;*/
    }
    .right{
        margin: 0 auto;
        /*width: 100%;*/
    }
    .display-author{
        width: 100%;
    }
    .display-view {
        width: 100%;
        /* max-width: 970px; */
        margin: 0 auto;
    }

    .display-query {
        margin-bottom: 1.5em;
    }
    .display-query #query {
        /* text-transform: uppercase; */
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
        /* color: #409EFF; */
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