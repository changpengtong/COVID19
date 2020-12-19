<template>
  <!-- 检索框 -->
  <div class="search-box">
    <el-input
        class="search-box-auto"
        v-model="input"
        clearable
        @keyup.enter.native="search"
        :fetch-suggestions="querySearch"
        :placeholder="placeholder"
        @select="handleSelect"
    >
      <el-tooltip
          class="item"
          effect="dark"
          slot="suffix"
          content="Search"
          placement="top-start"
      >
        <el-button
            class="el-icon-search"
            @click="search"
            circle
            size="mini"
            plain
        >
        </el-button>
      </el-tooltip>

      <template slot-scope="{ item }">
        <div class="name">{{ item.name }}</div>
      </template>
    </el-input>
  </div>
</template>

<script>
import dataEntities from "../../data/top10Entities.json";
import dataAuthors from "../../data/top10Authors.json";
import dataPapers from "../../data/top10Papers.json";
import $axios from "../../util/axios";

export default {
  name: "SearchBox",
  props: {
    type: Number,
  },
  data() {
    return {
      input: "",
      query: "",
      placeholder: "Whatever comes to your mind",
      dataEntities,
      dataAuthors,
      dataPapers,
      data: [],
      timeout: null,
    };
  },
  methods: {
    createFilter() {
      let filterList = [];
      this.data.forEach((row) => {
        let same = row.name.toLowerCase().startsWith(this.input.toLowerCase());
        if (same) {
          filterList.push(row);
        }
      });
      return filterList;
    },
    querySearch(queryString, callback) {
      console.log(this.data);
      if (this.type === 0) {
        callback([]);
      }
      let list = this.data;
      let results = queryString ? this.createFilter() : list;

      //调用 callback 返回建议列表的数据
      clearTimeout(this.timeout);
      this.timeout = setTimeout(() => {
        callback(results);
      }, 3000 * Math.random());
    },

    // 发送检索请求
    search() {
      switch (this.type) {
        case 0:
          console.log("in switch 0");
          this.$alert(
              "<strong>Please select the category before searching.</strong>",
              {
                dangerouslyUseHTMLString: true,
              }
          );
          break;

        case 1:
          this.$router.push({
            path: "/displayInfo",
            query: {
              // keyword: this.query
              query: this.input,
            },
          });
          break;
        case 2:
          this.$router.replace({
            name: "institution",
            params: { id: this.input },
          });
          break;
        case 3:
          this.$router.push({ name: "bioentity", params: { id: this.input } });
          break;
        case 4:
          this.$router.push({ name: 'itemAuthor', params: { name:this.input }});
          break;
      }
    },
    handleSelect(item) {
      let name = item.name;
      this.search(name);
    },
    fetchAPI() {
      console.log("in fetch");
      $axios
          .get("/institutionNames")
          .then((response) => {
            let temp = response.data;
            this.data = temp;
            console.log(this.data);
          })
          .catch((error) => console.log(error));

    },
  },

  mounted: function () {
    console.log(this.type);

    // let api = "";
    switch (this.type) {
      case 0:
        break;
      case 1:
        // api = "paper";
        this.data = this.dataPapers;
        break;
      case 2:
        // api = "institutionNames";

        this.fetchAPI();

        break;
      case 3:
        // api = "bioentity";
        this.data = this.dataEntities;
        break;
      case 4:
        // api = "author";
        this.data = this.dataAuthors;
        break;
    }
  },
};
</script>

<style >
.search-box .search-box-auto {
  width: 110%;
}

.search-box .el-input__inner {
  border-radius: 20px 20px;
}

.search-box .el-icon-search {
  font-size: 17px;
  margin-right: 10px;
  border: 0px;
}
</style>