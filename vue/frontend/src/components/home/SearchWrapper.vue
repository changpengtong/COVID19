<template>
  <div >

    <div class="home-search-pane">
      <search-box   v-if="activeName=='1'" class="center search-box" ></search-box>
      <search-box2   v-else-if="activeName=='2'" class="center search-box2" ></search-box2>
      <search-box3   v-else-if="activeName=='3'" class="center search-box3" ></search-box3>
      <search-box4  v-else-if="activeName=='4'" class="center search-box4" ></search-box4>
      <div class="center search-box" v-else>
        <el-input
            v-model="query"
            @keyup.enter.native="alert"
            :placeholder="placeholder"
        >
          <el-button slot="append" icon="el-icon-search" @click="alert"></el-button>
        </el-input>
      </div>
    </div>
    <div class="radios">
    <el-radio-group v-model="activeName" @change="handleClick">
      <el-radio :label="1">PAPER</el-radio>
      <el-radio :label="2">INSTITUTION</el-radio>
      <el-radio :label="3">BIOENTITY</el-radio>
      <el-radio :label="4">AUTHOR</el-radio>

    </el-radio-group>
    </div>
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
        "search-box4": SearchBox4,
      },
      data() {
        return {
          search_suggestion1: ["CoronaVirus", "2020", "pneumonia"],
          search_suggestion2: ["The University of Texas", "Michigan State University"],
          search_suggestion3: ["COVID-19", "ACE2"],
          search_suggestion4: ["Jason S. McLellan"],
          // activeName: 'first'
          query: "",
          activeName:"",

          placeholder: "Whatever comes to your mind"
        }
      },
      methods: {
        handleClick(tab, event) {
          console.log(tab, event);
        },
        alert() {
          this.$alert('<strong>Please Select the Category Before Searching.</strong>',  {
            dangerouslyUseHTMLString: true
          });
        },
      }
    }
</script>

<style scoped>
    .home-search-pane {
        display: block;
        width: 85%;
        /*padding-bottom: -100px;*/
      margin-bottom: 1.2em;
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


    .el-radio-group{
      margin-top: 15px ;
    }
</style>
