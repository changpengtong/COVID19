<template>
    <div class="home">
        <div class="home-pane">
            <div class="home-search-wrapper">
                <div class="_container">
                    <search-wrapper class="home-search-pane"></search-wrapper>
                </div>
            </div>
        </div>
        <!-- 词云 -->
        <!-- 响应式布局 -->
        <div class="home-pane">
            <div class="home-wordcloud-pane _container">
                <h1 class="slogan">all about Covid-19</h1>
                <el-row>
                    <el-col :md="24" :lg="8">
                        <h2 class="title">technology</h2>
                        <div id="tech"></div>
                    </el-col>
                    <el-col :md="24" :lg="8">
                        <h2 class="title">Covid19</h2>
                        <div id="field"></div>
                    </el-col>
                    <el-col :md="24" :lg="8">
                        <h2 class="title">organization</h2>
                        <div id="ins"></div>
                    </el-col>
                </el-row>
            </div>
        </div>

        <div class="home-pane">
            <div class="home-intro _container">
                <h1 class="slogan">big data platform for sci-tech knowledge</h1>
                <el-row>
                    <el-col :md="24" :lg="12">
                        <img src="./../assets/img/home/pull-file.png" alt="files">
                    </el-col>
                    <el-col :md="24" :lg="12">
                        <div class="item">
                            <h2>Intelligent retrieval in vertical domain is significant.</h2>
                            <p>In the era of artificial intelligence with big data, knowledge base-based intelligent retrieval in vertical domain is a significant development. We benefit from intelligent retrieval to efficiently manage large data, display massive fragmented knowledge structurally, and extract valuable information from it for trending analysis and prediction.</p>
                        </div>
                        <div class="item">
                            <h2>Many retrieval platforms emerge</h2>
                            <p>In the field of Sci-tech Literature Retrieval, many retrieval platforms emerge, such as Google Scholar, Scopus, Semantic Scholar, AMiner and so on. However, the above platforms lack the multi-source data fusion and diffusion of innovations in the industry, while, with the development of artificial intelligence, empowering the traditional industry with artificial intelligence become a hot spot in the research field.</p>
                        </div>
                    </el-col>
                </el-row>
            </div>
        </div>

        <div class="home-pane">
            <div class="home-what-do _container">
                <h1 class="slogan">what <span class="_font-blue">COVID-19 Portal</span> do?</h1>
                <el-row>
                    <el-col :md="24" :lg="8">
                        <div class="item">
                            <img src="./../assets/img/home/pull-file.png" alt="files">
                            <p>Researching the field of artificial intelligence, we pulled over 3 million records of multi-source data such as papers, news and patents.</p>
                        </div>
                    </el-col>
                    <el-col :md="24" :lg="8">
                        <div class="item">
                            <img src="./../assets/img/home/pull-file.png" alt="files">
                            <p>We designed a reusable vertical domain retrieval framework named VIIA-RIIR, based on which the intelligent retrieval and visualization platform Helenus.ai was constructed</p>
                        </div>
                    </el-col>
                    <el-col :md="24" :lg="8">
                        <div class="item">
                            <img src="./../assets/img/home/pull-file.png" alt="files">
                            <p>The Covid-19 Portal could support the intelligent retrieval and visualization at the million level of multi-source data</p>
                        </div>
                    </el-col>
                </el-row>
            </div>
        </div>

        <div class="home-pane">
            <div class="home-what-u-get _container">
                <h1 class="slogan">what <span class="_font-blue">covid-19 portal</span> can help you?</h1>
                <h2>Based on the information visualization of Covid-19 Portal, you can analyze development status, technology ecology, organization layout, industry application and national competition concerning Covid-19.</h2>
            </div>
        </div>

        <!-- <div class="home-pane home-pane-carousel">
          <div class="home-carousel-inner">
          </div>
        </div> -->
    </div>
</template>

<script>
    import Vue from 'vue'
    import SearchWrapper from "./../components/home/SearchWrapper.vue"
     import wordcloud from "./../graphs/wordcloud"

    export default {
        name: 'home',
        components: {
            "search-wrapper": SearchWrapper,
        },
        data() {
            return {
                // 检索建议
                search_suggestion: ["svm","baidu","finance"]
            }
        },


        ////////词云///////////////////////////
        mounted() {
            this.$axios.get("/wordcloud")
            .then(res=> {
                let data = res.data
                // let text=[]
                // let size=[]
                let map=[]
                console.log(data)
                for(let i=0;i<data.length;i++) {
                    var d = {}
                    d['name'] = data[i]['text']
                    d['value'] = data[i]['size']
                    //   console.log(d)
                    map[i] = d
                }
                    let subclass = Vue.extend(wordcloud);
                    new subclass({
                        propsData:{
                            items:map
                        }
                    }).$mount("#"+"field")
              //  }
            })
            .catch(err=> {
                console.log("ERROR"+err);
            })
        },
        methods: {
        }
    }
</script>

<style scoped>
    h1 {
        text-transform: capitalize;
        font-size: 2em;
        font-weight: 400;
        letter-spacing: 3px;
    }

    h2 {
        text-transform: capitalize;
        font-size: 1.5em;
        font-weight: 400;
        letter-spacing: 3px;
    }

    p {
        margin-top: .5em;
        margin-bottom: .5em;
    }

    .center {
        margin: 0 auto;
    }

    .home {
        margin: 0 auto;
        text-align: center;
    }

    .home-pane {
        min-height: 200px;
        padding: 2em 0;
        margin: 0;
    }
    /* 奇偶选择器 */
    .home-pane:nth-child(even) {
        background-color: #F5F5F5;
    }

    .home-wordcloud-pane {
        min-height: 200px;
        padding: 10px 5px;
        /* 最大宽度是当前窗口的65% 设置margin 保证自适应和两边留白 不影响其他元素宽度*/
        margin: 0 auto;
        /* max-width: calc(65%); */
    }

    .home-search-wrapper {
        /* background: url('../assets/img/white-background.jpg') center center no-repeat; */
        /* background: url('../assets/img/gray-background.png'); */
        background: url('../assets/img/home/blue-background.png');
        /* 100%是按照图片大小铺满 cover是根据屏幕大小铺满（可能会截取掉一部分） */
        background-size: cover;
        /* 是否跟随鼠标一起滑动 */
        background-attachment:fixed;
        /* 是否重复铺满  */
        /* background-repeat:no-repeat; */
        width: 100%;
        background-position:center;
        height: calc(100vh - 72px);
        /* 父元素设置relative 子元素设置absolute 实现相对定位*/
        position: relative;

    }
    .home-search-wrapper .home-search-pane {
        /* 相对定位 */
        position: absolute;
        top: 50%;
        left: 50%;
        transform: translate(-50%,-50%);
        padding: 1em;
    }

    .home-intro,
    .home-what-do,
    .home-what-u-get {
        /* margin: 0 5em; */
    }

    .home-intro .item {
        text-align: left;
    }

    .home-intro img {
        width: 13em;
    }

    .home-what-do .item {
        margin: .5em 2em;
    }

    .home-what-do .item img {
        /* height: 100%; */
        margin: 0 auto 1em auto;
        width: 13em;
    }

    .home-pane-carousel {
        min-height: 400px;
        /* background: url('../assets/img/white-background.jpg') center center no-repeat; */
        background: url('./../assets/img/blackMachine-background.jpg');
        /* 100%是按照图片大小铺满 cover是根据屏幕大小铺满（可能会截取掉一部分） */
        background-size: cover;
        /* 是否跟随鼠标一起滑动 */
        /* background-attachment:fixed; */
        /* 是否重复铺满  */
        /* background-repeat:no-repeat; */
        width:calc(100%);
        background-position:center;
        height: calc(50%);
        /* 父元素设置relative 子元素设置absolute 实现相对定位*/
        /* position: relative; */
    }
</style>