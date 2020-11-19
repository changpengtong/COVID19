<template>
    <div class="display-view container">
        <!-- 展示数据和可视化 -->
        <div class="display-info ">
            <el-row :gutter="24" >

                <el-col :md="17" :sm="24" :xs="24">
<!--                    author info-->
                    <author-card class="display-author" v-loading="loading" :card="card"></author-card>
                    <!-- 展示数据 -->
                    <div
                            v-loading="isLoading"
                            element-loading-text="loading……">
                        <author-paper :paper="paper" v-loading="loading"></author-paper>
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
    import AuthorCard from "../components/general/AuthorCard";
    import AuthorPaper from "../components/general/paper";
    import CoAuthors from "../components/general/coauthor";
    import entityList from "../components/general/list";
    // import wordcloud from "../components/Author/AuthorWordcloud";
    import bar from "../components/general/bar";
    import wordcloud from "../components/general/wordcloud";
    import $axios from "../util/axios";
 //   import $route from "../util/router"
    //import DisplayProfile from "../components/display_info/DisplayProfile";
    export default {
        name: "Author",
        components: {
            "author-card":AuthorCard,
            "author-paper": AuthorPaper,
            "co-author": CoAuthors,
            "bar":bar,
            'entityList' : entityList,
            'wordcloud':wordcloud
        },
        data() {
            return {
                paper:[],
                card:[],
                coauthor:[],
                entity:[],
                bar:[],
                wordcloud:[],
                flag:"",
                loading:true
            }
        },
        mounted() {
            this.get_data()
        },
        methods: {
            get_data() {
              //  $axios.get("/paperDetail/"+this.$route.params.id).then(response=>{
                $axios.get("/displayAuthor/"+this.$route.params.id).then(response=>{
                    let d=response.data
                   // console.log(d)
                    this.paper = d["articles"]
                    this.bar = d["bar"]
                    this.coauthor = d["coauthor"]
                    this.entity = d["wordcloud"]
                    this.wordcloud = d["wordcloud"]
                    this.card = d["card"]
                  //  this.paper = d
                    this.flag = "aut"
                    this.loading = false
                   // console.log(this.paper)
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