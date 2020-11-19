<template>
    <div class="display-view container">
        <!-- 展示数据和可视化 -->
        <div class="display-info ">
            <el-row :gutter="24" >
                <el-col :md="17" :sm="24" :xs="24">
                    <institution-card class="display-author" :name="name" v-loading="loading"></institution-card>
                    <div
                            class="display-data"
                            v-loading="isLoading"
                            element-loading-text="loading……">
                        <institution-paper :paper="paper" v-loading="loading"></institution-paper>
                    </div>
                </el-col>
                <el-col :md="7" :sm="24" :xs="24" class="right">
                    <div class="display-visual">
                        <p class="up">Bio-entities<el-tooltip effect="dark" content="3 types of bio-entities ranked by frequency mentioned in papers published." placement="top-start">
                            <i class="el-icon-question"></i>
                        </el-tooltip></p>
                        <entityList class="display-visual" :entity="entity" v-loading="loading"></entityList>
                        <wordcloud :wordcloud="wordcloud" v-loading="loading"></wordcloud>
                        <bar class="down" :bar="bar" v-loading="loading"></bar>
                    </div>
                    <co-author :coauthor="coauthor" :flag="flag" v-loading="loading"></co-author>
                </el-col>
            </el-row>
        </div>
    </div>
</template>

<script>
    import AuthorCard from "../components/general/institutionCard";
    import AuthorPaper from "../components/general/paper";
    import CoAuthors from "../components/general/coauthor";
    import Bar from "../components/general/bar";
    import entityList from "../components/general/list";
    import wordcloud from "../components/general/wordcloud";

    import $axios from "../util/axios";

    export default {
        components: {
            "institution-card":AuthorCard,
            "institution-paper": AuthorPaper,
            "co-author": CoAuthors,
            "bar":Bar,
            "entityList": entityList,
            "wordcloud": wordcloud
        },
        data() {
            return {
                name:[],
                d:[],
                card:[],
                paper:[],
                coauthor:[],
                bar:[],
                wordcloud:[],
                entity:[],
                flag:"",
                loading:true
            }
        },
        mounted() {
            this.get_data()
        },
      // computed: {
      //   query() {
      //     // 返回路由中的query，注意用的是$route
      //     return this.$route.params.id
      //   }
      // },
        methods:{
            get_data() {
                $axios.get("/displayInstitution/"+this.$route.params.id).then(response=>{
                    let d=response.data
                   // console.log(d)
                    this.name = this.$route.params.id
                    console.log(this.name)
                    this.paper = d["articles"]
                    this.bar = d["bar"]
                    this.coauthor = d["coauthor"]
                    this.entity = d["wordcloud"]
                    this.wordcloud = d["wordcloud"]
                    this.flag = "ins"
                    this.loading = false
                    console.log(this.paper)
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
        margin-top: 1.8em;
    }

    .down {
        margin-bottom: 1em;
    }

</style>