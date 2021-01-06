<template>
  <div class="home-search-pane">
    <search-box
        class="search-box"
        :type="type"
        :data="data"
        :key="data"
    ></search-box>

    <div class="radios">
      <el-radio-group v-model="type" @change="handleClick">
        <el-radio :label="2">INSTITUTION</el-radio>
        <el-radio :label="3">BIOENTITY</el-radio>
        <el-radio :label="4">AUTHOR</el-radio>
      </el-radio-group>
    </div>
  </div>
</template>

<script scoped>
import SearchBox from "../common/SearchBox.vue";
import $axios from "../../util/axios";
import sample from "../../data/a.json";

export default {
  components: {
    "search-box": SearchBox,
  },
  data() {
    return {
      type: 0,
      data: [],
      sample: sample,
      dataEntities: [],
      dataAuthors: [],
      dataInstitutions: [],
    };
  },

  methods: {
    handleClick() {
      console.log("clicked");
      // this.data = this.sample;
      // console.log(sample);
      switch (this.type) {
        case 2:
          this.data = this.dataInstitutions;
          console.log(this.dataInstitutions);
          break;

        case 3:
          this.data = this.dataEntities;
          console.log(this.dataEntities);
          break;

        case 4:
          this.data = this.dataAuthors;
          console.log(this.dataAuthors);
          break;
      }
    },
    fetchAPI(api) {
      console.log("in fetch");
      $axios
          .get("/" + api)
          .then((response) => {
            let temp = response.data;
            switch (api) {
              case "institutionNames":
                this.dataInstitutions = temp;
                console.log(this.dataInstitutions);
                break;

              case "bioentityNames":
                this.dataEntities = temp;
                break;

              case "authorNames":
                this.dataAuthors = temp;
                break;
            }
          })
          .catch((error) => console.log(error));
    },
  },
  mounted: function () {
    this.fetchAPI("institutionNames");
    this.fetchAPI("bioentityNames");
    this.fetchAPI("authorNames");
  },
};
</script>

<style>
/*.home-search-pane {*/
/*  display: block;*/
/*  width: 100%;*/
/*  padding-bottom: -100px;*/
/*  padding-top: 15px;*/
/*}*/

/*.home-search-pane .search-box {*/
/*  padding-bottom: 10px;*/
/*}*/

/*.home-search-pane .radios .el-radio-group {*/
/*  padding-top: 13px;*/
/*  padding-left: 15%;*/
/*}*/
</style>
