<template>
    <div class="display-view container">
        <div class="display-info ">
            <el-row :gutter="24" >

                <el-col :md="17" :sm="24" :xs="24">
                    <author-card class="display-author"  :card="card" v-loading="loading"></author-card>
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
    import bar from "../components/general/bar";
    import $axios from "../util/axios";
    export default {
        name: "Author",
        components: {
            "author-card":AuthorCard,
            "author-paper": AuthorPaper,
            "co-author": CoAuthors,
            "bar":bar,
            'entityList' : entityList
        },
        data() {
            return {
                paper:[],
                card:[],
                coauthor:[],
                entity:[],
                bar:[],
                flag:"",
                loading:true
            }
        },
        mounted() {
            this.get_data()
        },
      created() {
        this.get_data()

      },
      watch: {
        '$route': 'get_data'
      },
        methods: {
          get_data() {
            this.loading = true
            $axios.get("/displayAuthor/" + this.$route.params.id).then(response => {
              this.loading = false
              let d = response.data
              this.paper = [d["articles"], d["clinicals"]]
              this.bar = d["bar"]
              this.coauthor = d["coauthor"]
              this.entity = d["wordcloud"]
              this.card=d["card"]
              this.flag = "aut"
            })
          }
        },

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



    .display-visual .title {
        text-transform: capitalize;
        margin-bottom: 1em;
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