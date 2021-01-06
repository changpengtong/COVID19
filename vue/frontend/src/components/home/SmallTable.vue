<template>
  <el-card class="small-table-card">
    <div class="small-table">
      <div class="title">Popular {{ tableName }}</div>
      <el-table
        ref="singleTable"
        :data="dataList"
        stripe
        height="200"
        highlight-current-row
        fit
        @current-change="handleCurrentChange"
        @row-click="getDetailPage"
        max-height="325"
        :show-header="false"
      >
        <el-table-column type="index" width="35"> </el-table-column>
        <el-table-column property="name" label="Name"> </el-table-column>
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
    // 发送检索请求
    getDetailPage(row) {
      switch (this.type) {
        case 1:
          this.$router.push({
            path: "/institution/" + row.name,
          });
          break;
        case 2:
          this.$router.push({
            path: "/bioentity/" + row.name,
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
.small-table-card {
  text-align: center;
  height: 220px;
  width: 265px;
  border: hidden;
  margin-right: 0px;
}

.small-table .title {
  font-size: 21px;
  margin-top: -10px;
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
  font-size: 13px;
  height: 30px;
}

.small-table-card
  .small-table
  .el-table--striped
  .el-table__body
  tr.el-table__row--striped
  td {
  background-color: #edf4ff;
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


</style>