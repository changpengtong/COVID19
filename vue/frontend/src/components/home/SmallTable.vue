<template>
  <el-card class="small-table-card">
    <div class="small-table" >
      <div class="title">Top 10 {{tableName}}</div>
      <el-table ref="singleTable" :data="dataList"
                stripe height="370" highlight-current-row fit
                @current-change="handleCurrentChange"
                @row-click="getDetailPage"
                style="width: 100%"
                max-height="325">
        <el-table-column
            type="index"
            width="35">
        </el-table-column>

        <el-table-column
            property="name"
            label="Name"
        >
        </el-table-column>

        <el-table-column
            property="paper"
            label="Paper"
            width="60">
        </el-table-column>
      </el-table>
    </div>
  </el-card>

</template>

<script>
export default {
  name: "SmallTable",
  props: {
    tableName : String,
    dataList : Array
  },
  methods: {
    // 发送检索请求
    getDetailPage(row) {
      this.$router.push("path").catch(error => {
        if (error.name != "NavigationDuplicated") {
          throw error;
        }
      });
      this.$router.push({
        path: "/COVID19/#/bioentity/" + row.name,

      });
    }
  }
};
</script>

<style >

.small-table-card{
  text-align: center;
  height: 345px;
}

.small-table .title {
  font-size: 21px;
  margin-top: -10px;
  font-weight: bold;
}


.small-table-card .el-table td, .el-table th.is-leaf {
  border-bottom: 0;
}
</style>