<template>
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
export default {
  name: "SearchBox",
  props: {
    type: Number,
    data: Array,
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
          if(!names.has(row.name)) {
            filterList.push(row);
            names.add(row.name);
          }

        }
      });
      return filterList;
    },
    querySearch(queryString, callback) {
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
    },

    search(item) {
      switch (this.type) {
        case 0:
          this.$alert(
            "<strong>Please select the category before searching.</strong>",
            {
              dangerouslyUseHTMLString: true,
            }
          );
          break;

      case 2:
        this.$router.push({
          path: "/institution/" + item.id,
        });
        break;
      case 3:
        this.$router.push({
          path: "/bioentity/" + item.id,
        });
        break;
      case 4:
        this.$router.push({
          path: "/authorList/" + item.name,
        });
        break;
      }
    },

    handleSelect(item) {
      this.search(item);
    },

    handleButtonSearch() {
      this.search(this.input);
    },
  },

};
</script>