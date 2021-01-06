<template>
  <!-- 检索框 -->
  <div class="search-box">

    <el-autocomplete
      class="search-box-auto"
      v-model="input"
      clearable
      @keyup.enter.native="handleButtonSearch"
      :fetch-suggestions="querySearch"
      :placeholder="placeholder"
      @select="handleSelect"
      @change="handleInputChange"
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
          @click="handleButtonSearch"
          circle
          size="mini"
          plain
        >
        </el-button>
      </el-tooltip>

      <template slot-scope="{ item }">
        <div class="name">{{ item.name }}</div>
      </template>
    </el-autocomplete>
  </div>
</template>

<script>
// import $axios from "../../util/axios";

export default {
  name: "SearchBox",
  props: {
    type: Number,
    data: Array,
    // method: {setInput: Function},
  },
  data() {
    return {
      input: "",
      placeholder: "Whatever comes to your mind",
      timeout: null,
    };
  },
  methods: {
    createFilter() {
      let filterList = [];
      let names = new Set();
      this.data.forEach((row) => {
        let same = row.name.toLowerCase().startsWith(this.input.toLowerCase());
        if (same) {
          row.name = row.name.replace(/^,+/, "").replace(/,+$/, "");
          if(!names.has(row.name)) {
            filterList.push(row);
            names.add(row.name);
          }

        }
      });
      return filterList;
    },
    querySearch(queryString, callback) {
      console.log(this.data);
      if (this.type === 0) {
        this.$notify({
          type: 'warning',
          title: 'Please select a type (Institution/Bioentity/Author).',
          message: 'Type selector is under search box.',
          offset: 100,
          position: 'top-left',
        });
      }

      let list = [];
      let results = queryString ? this.createFilter() : list;

      callback(results);

      //调用 callback 返回建议列表的数据
      // clearTimeout(this.timeout);
      // this.timeout = setTimeout(() => {
      //   callback(results);
      // }, 3000 * Math.random());
    },

    // 发送检索请求
    search(name) {
      switch (this.type) {
        case 0:
          this.$alert(
            "<strong>Please select the category before searching.</strong>",
            {
              dangerouslyUseHTMLString: true,
            }
          );
          break;

    case 1:
      this.$router.push({
        path: "/displayInfo/" + name,
      });
      break;
    case 2:
      this.$router.push({
        path: "/institution/" + name,
      });
      break;
    case 3:
      this.$router.push({
        path: "/bioentity/" + name,
      });
      break;
    case 4:
      this.$router.push({
        path: "/authorList/" + name,
      });
      break;
    }
    },
    handleSelect(item) {
      this.search(item.name);
    },
    handleButtonSearch() {
      this.search(this.input);
    },
    // fetchAPI(api) {
    //   console.log("in fetch");
    //   $axios
    //     .get("/" + api)
    //     .then((response) => {
    //       let temp = response.data;
    //       this.data = temp;
    //       console.log(this.data);
    //     })
    //     .catch((error) => console.log(error));
    // },
  },

};
</script>

<style >
/*.search-box .search-box-auto {*/
/*  position: center;*/
/*  width: 770px;*/
/*}*/

/*.search-box .el-col-6 {*/
/*  min-width: 100px;*/
/*  padding-right: 100px;*/
/*}*/

/*.search-box .el-input__inner {*/
/*  border-radius: 20px 20px;*/
/*}*/

/*.search-box .el-icon-search {*/
/*  margin-top: 4px;*/
/*  font-size: 17px;*/
/*  margin-right: 10px;*/
/*  border: 0px;*/
/*}*/
</style>