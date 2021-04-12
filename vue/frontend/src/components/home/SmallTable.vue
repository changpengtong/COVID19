<template>
  <el-card class="small-table-card">
    <div class="small-table">
      <div class="title">Popular {{ tableName }}</div>
      <el-table
        ref="singleTable"
        :data="dataList"
        stripe
        highlight-current-row
        fit
        @current-change="handleCurrentChange"
        @row-click="getDetailPage"
        :show-header="false"
      >
      <el-table-column type="index" width="30"> </el-table-column>
      <el-table-column width="200">
        <template slot-scope="scope">
          <button class="item-button">{{ scope.row.name }}</button>
        </template>
      </el-table-column>
        <!-- <el-table-column property="name" label="Name"><button></button></el-table-column> -->
      </el-table>

    </div>
  </el-card>
</template>

<script>
export default {
  name: "SmallTable",
  props: {
    tableName: String,
    dataList: Array,
    type: Number,
  },
  methods: {
    getDetailPage(row) {
      switch (this.type) {
        case 1:
          this.$router.push({
            path: "/institution/" + row.id,
          });
          break;
        case 2:
          this.$router.push({
            path: "/bioentity/" + row.id,
          });
          break;
        case 3:
          this.$router.push({
            path: "/author/" + row.aid,
          });
          break;
      }
    },
  },
};
</script>

<style >
.el-table {
  height: 10%;
  max-height: 10%;
}

.small-table-card {
  text-align: center;
  height: 30%;
  width: 100%;
  border: hidden;
  margin-right: 0px;
}

.small-table .title {
  font-size: 21px;
  margin-top: 0px;
  font-weight: bold;
}

.small-table-card .el-table td,
.el-table th.is-leaf {
  border-bottom: 0;
}
.small-table-card .small-table .el-table {
  padding-top: 10px;
}

.small-table-card .small-table .el-table__row {
  font-size: 12px;
  height: 30px;
}

.small-table-card
  .small-table
  .el-table--striped
  .el-table__body
  tr.el-table__row--striped
  td {
}

.small-table-card .small-table .el-table td,
.el-table th {
  padding: 0;
  border: 0;
}

.small-table-card .small-table .title {
  font-weight: bold;
  font-family: "Slabo 27px", serif;
}

.item-button {
  border: none;
  background-color: transparent;
  outline: none; 
  width: 100%;
  text-align: left;
}


</style>