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

export default {
  components: {
    "search-box": SearchBox,
  },
  data() {
    return {
      type: 0,
      data: [],
      dataEntities: [],
      dataAuthors: [],
      dataInstitutions: [],
      test: [],
    };
  },

  methods: {
    handleClick() {
      switch (this.type) {
        case 2:
          this.data = this.dataInstitutions;
          break;

        case 3:
          this.data = this.dataEntities;
          break;

        case 4:
          this.data = this.dataAuthors;
          break;
      }
    },
    fetchAPI(api) {
      $axios 
        .get("/" + api)
        .then((response) => {
          let temp = response.data;
          switch (api) {
            case "institutionNames":
              this.dataInstitutions = temp;
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
