<template>
    <div class="home-search-pane">
        <el-tabs v-model="activeName" @tab-click="handleClick">
            <el-tab-pane label="PAPER" name="first">
              <search-box class="center search-box" ></search-box>
            </el-tab-pane>
          <el-tab-pane label="INSTITUTION" name="second">
            <search-box2 class="center search-box2" ></search-box2>
          </el-tab-pane>
      <el-tab-pane label="BIOENTITY" name="third">
        <search-box3 class="center search-box3" ></search-box3>
      </el-tab-pane>
          <el-tab-pane label="AUTHOR" name="fourth">
            <search-box4 class="center search-box4" ></search-box4>
          </el-tab-pane>

        </el-tabs>

    </div>

</template>

<script>
    import SearchBox from "./../common/SearchBox.vue"
   import SearchBox2 from "./../common/SearchBox2.vue"
    import SearchBox3 from "./../common/SearchBox3.vue"
    import SearchBox4 from "./../common/SearchBox4.vue"

    //  import SearchBox3 from "./../common/SearchBox3.vue"
    export default {
        components: {
            "search-box": SearchBox,
          "search-box2": SearchBox2,
          "search-box3": SearchBox3,
          "search-box4": SearchBox4

          //  "search-box2": SearchBox2,
          //  "search-box3": SearchBox3
        },
        data() {
            return {
                search_suggestion1: ["CoronaVirus","2020","pneumonia"],
              search_suggestion2: ["The University of Texas", "Michigan State University"],
              search_suggestion3: ["COVID-19", "ACE2"],
              search_suggestion4: ["Jason S. McLellan"],
              activeName: 'first'
            }
        },
        methods: {
            handleClick(tab, event) {
                console.log(tab, event);



            }
        }
    }
</script>

<style scoped>
    .home-search-pane {
        display: block;
        width: 85%;
        padding-bottom: -100px;
    }
    .home-search-pane .slogan{
        padding: 1em;
    }
    .home-search-pane h1 {
        font-size: 3em;
        font-weight: 400;
        letter-spacing: 6px;
        margin-bottom: 0.3em;
        margin-top: 0;
        text-transform: uppercase;
    }
    .home-search-pane h2 {
        font-size: 1.5em;
        font-weight: 400;
        letter-spacing: 3px;
        margin-bottom: 1.5em;
    }

    .home-search-pane .search-box {
        box-shadow: 2px 2px 10px #f5f5f5;
    }

    .home-search-pane .search-box2 {
        box-shadow: 2px 2px 10px #f5f5f5;
    }

    .home-search-pane .search-box3 {
        box-shadow: 2px 2px 10px #f5f5f5;
    }

    .home-search-pane .suggestion {
        font-size: 1.2em;
        margin: 1em;
    }

    .home-search-pane .suggestion .title{
        display: inline-block;
        margin: 0 1em 0 .5em;
    }

    .home-search-pane .suggestion .content div{
        display: inline;
    }

    .home-search-pane .suggestion a {
        display: inline-block;
        text-decoration: none;
        text-transform: capitalize;
        margin: 0 .5em;
    }

    .home-search-pane .suggestion a:hover {
        text-decoration: underline;
    }
</style>
